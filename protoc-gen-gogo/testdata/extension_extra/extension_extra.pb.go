// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extension_extra/extension_extra.proto

package extension_extra

import (
	fmt "fmt"
	proto "github.com/yangshengBE/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExtraMessage struct {
	Width                *int32   `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraMessage) Reset()         { *m = ExtraMessage{} }
func (m *ExtraMessage) String() string { return proto.CompactTextString(m) }
func (*ExtraMessage) ProtoMessage()    {}
func (*ExtraMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_fce75f5a63502cd5, []int{0}
}
func (m *ExtraMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraMessage.Unmarshal(m, b)
}
func (m *ExtraMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraMessage.Marshal(b, m, deterministic)
}
func (m *ExtraMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraMessage.Merge(m, src)
}
func (m *ExtraMessage) XXX_Size() int {
	return xxx_messageInfo_ExtraMessage.Size(m)
}
func (m *ExtraMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraMessage proto.InternalMessageInfo

func (m *ExtraMessage) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func init() {
	proto.RegisterType((*ExtraMessage)(nil), "extension_extra.ExtraMessage")
}

func init() {
	proto.RegisterFile("extension_extra/extension_extra.proto", fileDescriptor_fce75f5a63502cd5)
}

var fileDescriptor_fce75f5a63502cd5 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xad, 0x28, 0x49,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x8b, 0x4f, 0xad, 0x28, 0x29, 0x4a, 0xd4, 0x47, 0xe3, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa3, 0x09, 0x2b, 0xa9, 0x70, 0xf1, 0xb8, 0x82, 0x18, 0xbe,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x22, 0x5c, 0xac, 0xe5, 0x99, 0x29, 0x25, 0x19, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x93, 0x73, 0x94, 0x63, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe, 0x3e, 0xd8, 0xbc, 0xa4, 0xd2,
	0x34, 0x08, 0x23, 0x59, 0x37, 0x3d, 0x35, 0x4f, 0x17, 0x2c, 0x51, 0x92, 0x5a, 0x5c, 0x92, 0x92,
	0x58, 0x82, 0xe1, 0x02, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x01, 0x1c, 0x4a, 0xa3, 0x00,
	0x00, 0x00,
}
