// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flagTableWrapper.proto

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

type FlagTableWrapper struct {
	Body                 map[string]int64 `protobuf:"bytes,1,rep,name=Body,proto3" json:"Body,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FlagTableWrapper) Reset()         { *m = FlagTableWrapper{} }
func (m *FlagTableWrapper) String() string { return proto.CompactTextString(m) }
func (*FlagTableWrapper) ProtoMessage()    {}
func (*FlagTableWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_145876e9aa9c0e93, []int{0}
}

func (m *FlagTableWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagTableWrapper.Unmarshal(m, b)
}
func (m *FlagTableWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagTableWrapper.Marshal(b, m, deterministic)
}
func (m *FlagTableWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagTableWrapper.Merge(m, src)
}
func (m *FlagTableWrapper) XXX_Size() int {
	return xxx_messageInfo_FlagTableWrapper.Size(m)
}
func (m *FlagTableWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagTableWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_FlagTableWrapper proto.InternalMessageInfo

func (m *FlagTableWrapper) GetBody() map[string]int64 {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*FlagTableWrapper)(nil), "poset.FlagTableWrapper")
	proto.RegisterMapType((map[string]int64)(nil), "poset.FlagTableWrapper.BodyEntry")
}

func init() { proto.RegisterFile("flagTableWrapper.proto", fileDescriptor_145876e9aa9c0e93) }

var fileDescriptor_145876e9aa9c0e93 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcb, 0x49, 0x4c,
	0x0f, 0x49, 0x4c, 0xca, 0x49, 0x0d, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xc8, 0x2f, 0x4e, 0x2d, 0x51, 0x6a, 0x62, 0xe4, 0x12, 0x70, 0x43,
	0x53, 0x21, 0x64, 0xca, 0xc5, 0xe2, 0x94, 0x9f, 0x52, 0x29, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d,
	0xa4, 0xa8, 0x07, 0x56, 0xaa, 0x87, 0xae, 0x4c, 0x0f, 0xa4, 0xc6, 0x35, 0xaf, 0xa4, 0xa8, 0x32,
	0x08, 0xac, 0x5c, 0xca, 0x9c, 0x8b, 0x13, 0x2e, 0x24, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x0a, 0x32,
	0x82, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95,
	0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x93, 0xd8, 0xc0, 0x4e,
	0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x73, 0x14, 0xa6, 0x81, 0xac, 0x00, 0x00, 0x00,
}
