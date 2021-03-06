package proxy

import (
	"fmt"
	"os"

	"time"

	"github.com/andrecronje/lachesis/src/crypto"
	"github.com/andrecronje/lachesis/src/poset"
	bproxy "github.com/andrecronje/lachesis/src/proxy/lachesis"
	"github.com/sirupsen/logrus"
)

type State struct {
	stateHash []byte
	snapshots map[int][]byte
	logger    *logrus.Logger
}

func (a *State) CommitBlock(block poset.Block) ([]byte, error) {
	a.logger.WithField("block", block).Debug("CommitBlock")
	err := a.writeBlock(block)
	if err != nil {
		return nil, err
	}
	return a.stateHash, nil
}

func (a *State) GetSnapshot(blockIndex int) ([]byte, error) {
	a.logger.WithField("block", blockIndex).Debug("GetSnapshot")

	snapshot, ok := a.snapshots[blockIndex]
	if !ok {
		return nil, fmt.Errorf("Snapshot %d not found", blockIndex)
	}

	return snapshot, nil
}

func (a *State) Restore(snapshot []byte) ([]byte, error) {
	//XXX do something smart here
	a.stateHash = snapshot
	return a.stateHash, nil
}

func (a *State) writeBlock(block poset.Block) error {
	file, err := a.getFile()
	if err != nil {
		a.logger.Error(err)
		return err
	}
	defer file.Close()

	// write some text to file
	//and update state hash
	hash := a.stateHash
	for _, tx := range block.Transactions() {
		_, err = file.WriteString(fmt.Sprintf("%s\n", string(tx)))
		if err != nil {
			a.logger.Error(err)
			return err
		}
		hash = crypto.SimpleHashFromTwoHashes(hash, crypto.SHA256(tx))
	}

	err = file.Sync()
	if err != nil {
		a.logger.Error(err)
		return err
	}

	a.stateHash = hash

	//XXX do something smart here
	a.snapshots[block.Index()] = hash

	return nil
}

func (a *State) writeMessage(tx []byte) {
	file, err := a.getFile()
	if err != nil {
		a.logger.Error(err)
		return
	}
	defer file.Close()

	// write some text to file
	_, err = file.WriteString(fmt.Sprintf("%s\n", string(tx)))
	if err != nil {
		a.logger.Error(err)
	}
	err = file.Sync()
	if err != nil {
		a.logger.Error(err)
	}
}

func (a *State) getFile() (*os.File, error) {
	path := "messages.txt"
	return os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
}

//------------------------------------------------------

type DummySocketClient struct {
	state         *State
	lachesisProxy *bproxy.SocketLachesisProxy
	logger        *logrus.Logger
}

func NewDummySocketClient(clientAddr string, nodeAddr string, logger *logrus.Logger) (*DummySocketClient, error) {

	lachesisProxy, err := bproxy.NewSocketLachesisProxy(nodeAddr, clientAddr, 1*time.Second, logger)
	if err != nil {
		return nil, err
	}

	state := State{
		stateHash: []byte{},
		snapshots: make(map[int][]byte),
		logger:    logger,
	}
	state.writeMessage([]byte(clientAddr))

	client := &DummySocketClient{
		state:         &state,
		lachesisProxy: lachesisProxy,
		logger:        logger,
	}

	go client.Run()

	return client, nil
}

func (c *DummySocketClient) Run() {
	for {
		select {
		case commit := <-c.lachesisProxy.CommitCh():
			c.logger.Debug("CommitBlock")
			stateHash, err := c.state.CommitBlock(commit.Block)
			commit.Respond(stateHash, err)
		case snapshotRequest := <-c.lachesisProxy.SnapshotRequestCh():
			c.logger.Debug("GetSnapshot")
			snapshot, err := c.state.GetSnapshot(snapshotRequest.BlockIndex)
			snapshotRequest.Respond(snapshot, err)
		case restoreRequest := <-c.lachesisProxy.RestoreCh():
			c.logger.Debug("Restore")
			stateHash, err := c.state.Restore(restoreRequest.Snapshot)
			restoreRequest.Respond(stateHash, err)
		}
	}
}

func (c *DummySocketClient) SubmitTx(tx []byte) error {
	return c.lachesisProxy.SubmitTx(tx)
}
