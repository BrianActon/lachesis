// Code generated by protoc-gen-go. DO NOT EDIT.
// source: block.proto

package poset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockBody struct {
	Index                int64    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	RoundReceived        int64    `protobuf:"varint,2,opt,name=RoundReceived,proto3" json:"RoundReceived,omitempty"`
	StateHash            []byte   `protobuf:"bytes,3,opt,name=StateHash,proto3" json:"StateHash,omitempty"`
	FrameHash            []byte   `protobuf:"bytes,4,opt,name=FrameHash,proto3" json:"FrameHash,omitempty"`
	Transactions         [][]byte `protobuf:"bytes,5,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockBody) Reset()         { *m = BlockBody{} }
func (m *BlockBody) String() string { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()    {}
func (*BlockBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}

func (m *BlockBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockBody.Unmarshal(m, b)
}
func (m *BlockBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockBody.Marshal(b, m, deterministic)
}
func (m *BlockBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockBody.Merge(m, src)
}
func (m *BlockBody) XXX_Size() int {
	return xxx_messageInfo_BlockBody.Size(m)
}
func (m *BlockBody) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockBody.DiscardUnknown(m)
}

var xxx_messageInfo_BlockBody proto.InternalMessageInfo

func (m *BlockBody) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BlockBody) GetRoundReceived() int64 {
	if m != nil {
		return m.RoundReceived
	}
	return 0
}

func (m *BlockBody) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *BlockBody) GetFrameHash() []byte {
	if m != nil {
		return m.FrameHash
	}
	return nil
}

func (m *BlockBody) GetTransactions() [][]byte {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type WireBlockSignature struct {
	Index                int64    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Signature            string   `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WireBlockSignature) Reset()         { *m = WireBlockSignature{} }
func (m *WireBlockSignature) String() string { return proto.CompactTextString(m) }
func (*WireBlockSignature) ProtoMessage()    {}
func (*WireBlockSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{1}
}

func (m *WireBlockSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WireBlockSignature.Unmarshal(m, b)
}
func (m *WireBlockSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WireBlockSignature.Marshal(b, m, deterministic)
}
func (m *WireBlockSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WireBlockSignature.Merge(m, src)
}
func (m *WireBlockSignature) XXX_Size() int {
	return xxx_messageInfo_WireBlockSignature.Size(m)
}
func (m *WireBlockSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_WireBlockSignature.DiscardUnknown(m)
}

var xxx_messageInfo_WireBlockSignature proto.InternalMessageInfo

func (m *WireBlockSignature) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WireBlockSignature) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type Block struct {
	Body                 *BlockBody        `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
	Signatures           map[string]string `protobuf:"bytes,2,rep,name=Signatures,proto3" json:"Signatures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hash                 []byte            `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Hex                  string            `protobuf:"bytes,4,opt,name=hex,proto3" json:"hex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{2}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetBody() *BlockBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Block) GetSignatures() map[string]string {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Block) GetHex() string {
	if m != nil {
		return m.Hex
	}
	return ""
}

func init() {
	proto.RegisterType((*BlockBody)(nil), "poset.BlockBody")
	proto.RegisterType((*WireBlockSignature)(nil), "poset.WireBlockSignature")
	proto.RegisterType((*Block)(nil), "poset.Block")
	proto.RegisterMapType((map[string]string)(nil), "poset.Block.SignaturesEntry")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x3d, 0x6b, 0xc3, 0x30,
	0x14, 0x44, 0xfe, 0x28, 0xf8, 0xd9, 0xa5, 0xe1, 0xd1, 0xc1, 0x94, 0x0c, 0xc6, 0x64, 0xf0, 0xe4,
	0x21, 0x5d, 0x4a, 0x69, 0x97, 0x40, 0x4b, 0xba, 0x2a, 0x85, 0xce, 0x8a, 0x2d, 0x6a, 0x93, 0x54,
	0x0a, 0xb2, 0x1c, 0x92, 0x5f, 0xd4, 0xbf, 0xd3, 0x9f, 0x54, 0x24, 0x11, 0x3b, 0x29, 0x74, 0x7b,
	0xba, 0xbb, 0xf7, 0xb8, 0x3b, 0x41, 0xbc, 0xde, 0xca, 0x6a, 0x53, 0xee, 0x94, 0xd4, 0x12, 0xc3,
	0x9d, 0xec, 0xb8, 0xce, 0xbf, 0x09, 0x44, 0x0b, 0x03, 0x2f, 0x64, 0x7d, 0xc4, 0x5b, 0x08, 0xdf,
	0x44, 0xcd, 0x0f, 0x29, 0xc9, 0x48, 0xe1, 0x53, 0xf7, 0xc0, 0x19, 0x5c, 0x53, 0xd9, 0x8b, 0x9a,
	0xf2, 0x8a, 0xb7, 0x7b, 0x5e, 0xa7, 0x9e, 0x65, 0x2f, 0x41, 0x9c, 0x42, 0xb4, 0xd2, 0x4c, 0xf3,
	0x25, 0xeb, 0x9a, 0xd4, 0xcf, 0x48, 0x91, 0xd0, 0x11, 0x30, 0xec, 0xab, 0x62, 0x5f, 0x8e, 0x0d,
	0x1c, 0x3b, 0x00, 0x98, 0x43, 0xf2, 0xae, 0x98, 0xe8, 0x58, 0xa5, 0x5b, 0x29, 0xba, 0x34, 0xcc,
	0xfc, 0x22, 0xa1, 0x17, 0x58, 0xbe, 0x04, 0xfc, 0x68, 0x15, 0xb7, 0x66, 0x57, 0xed, 0xa7, 0x60,
	0xba, 0x57, 0xfc, 0x1f, 0xc7, 0xc6, 0xcb, 0x49, 0x62, 0xdd, 0x46, 0x74, 0x04, 0xf2, 0x1f, 0x02,
	0xa1, 0x3d, 0x83, 0x33, 0x08, 0x4c, 0x6e, 0xbb, 0x1c, 0xcf, 0x27, 0xa5, 0xed, 0xa4, 0x1c, 0xfa,
	0xa0, 0x96, 0xc5, 0x27, 0x80, 0x61, 0xb9, 0x4b, 0xbd, 0xcc, 0x2f, 0xe2, 0xf9, 0xf4, 0x5c, 0x5b,
	0x8e, 0xf4, 0x8b, 0xd0, 0xea, 0x48, 0xcf, 0xf4, 0x88, 0x10, 0x34, 0x63, 0x25, 0x76, 0xc6, 0x09,
	0xf8, 0x0d, 0x3f, 0xd8, 0x1e, 0x22, 0x6a, 0xc6, 0xbb, 0x67, 0xb8, 0xf9, 0x73, 0xc4, 0x88, 0x36,
	0xdc, 0x79, 0x8b, 0xa8, 0x19, 0x4d, 0xd8, 0x3d, 0xdb, 0xf6, 0xa7, 0x48, 0xee, 0xf1, 0xe8, 0x3d,
	0x90, 0xf5, 0x95, 0xfd, 0xd4, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x7d, 0x1a, 0xa7,
	0xe3, 0x01, 0x00, 0x00,
}
