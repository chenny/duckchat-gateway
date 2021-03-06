// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/im_cts.updatePointer.proto

package site // import "github.com/duckchat/duckchat-gateway/proto/site"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// action: im.cts.updatePointer
type ImCtsUpdatePointerRequest struct {
	U2Pointer            int64            `protobuf:"varint,1,opt,name=u2Pointer,proto3" json:"u2Pointer,omitempty"`
	GroupsPointer        map[string]int64 `protobuf:"bytes,2,rep,name=groupsPointer,proto3" json:"groupsPointer,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ImCtsUpdatePointerRequest) Reset()         { *m = ImCtsUpdatePointerRequest{} }
func (m *ImCtsUpdatePointerRequest) String() string { return proto.CompactTextString(m) }
func (*ImCtsUpdatePointerRequest) ProtoMessage()    {}
func (*ImCtsUpdatePointerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_cts_updatePointer_b5bc9945c7557863, []int{0}
}
func (m *ImCtsUpdatePointerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImCtsUpdatePointerRequest.Unmarshal(m, b)
}
func (m *ImCtsUpdatePointerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImCtsUpdatePointerRequest.Marshal(b, m, deterministic)
}
func (dst *ImCtsUpdatePointerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImCtsUpdatePointerRequest.Merge(dst, src)
}
func (m *ImCtsUpdatePointerRequest) XXX_Size() int {
	return xxx_messageInfo_ImCtsUpdatePointerRequest.Size(m)
}
func (m *ImCtsUpdatePointerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImCtsUpdatePointerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImCtsUpdatePointerRequest proto.InternalMessageInfo

func (m *ImCtsUpdatePointerRequest) GetU2Pointer() int64 {
	if m != nil {
		return m.U2Pointer
	}
	return 0
}

func (m *ImCtsUpdatePointerRequest) GetGroupsPointer() map[string]int64 {
	if m != nil {
		return m.GroupsPointer
	}
	return nil
}

func init() {
	proto.RegisterType((*ImCtsUpdatePointerRequest)(nil), "site.ImCtsUpdatePointerRequest")
	proto.RegisterMapType((map[string]int64)(nil), "site.ImCtsUpdatePointerRequest.GroupsPointerEntry")
}

func init() {
	proto.RegisterFile("site/im_cts.updatePointer.proto", fileDescriptor_im_cts_updatePointer_b5bc9945c7557863)
}

var fileDescriptor_im_cts_updatePointer_b5bc9945c7557863 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xce, 0x2c, 0x49,
	0xd5, 0xcf, 0xcc, 0x8d, 0x4f, 0x2e, 0x29, 0xd6, 0x2b, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x0d, 0xc8,
	0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x29, 0x50,
	0xba, 0xca, 0xc8, 0x25, 0xe9, 0x99, 0xeb, 0x5c, 0x52, 0x1c, 0x8a, 0xac, 0x24, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x48, 0x86, 0x8b, 0xb3, 0xd4, 0x08, 0x2a, 0x26, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x1c, 0x84, 0x10, 0x10, 0x8a, 0xe0, 0xe2, 0x4d, 0x2f, 0xca, 0x2f, 0x2d, 0x28, 0x86, 0xa9,
	0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x32, 0xd2, 0x03, 0x99, 0xac, 0x87, 0xd3, 0x54, 0x3d, 0x77,
	0x64, 0x4d, 0xae, 0x79, 0x25, 0x45, 0x95, 0x41, 0xa8, 0x06, 0x49, 0x39, 0x70, 0x09, 0x61, 0x2a,
	0x12, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x04, 0xbb, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe1,
	0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x02, 0xbb, 0x0d, 0xc2, 0xb1, 0x62, 0xb2, 0x60,
	0x74, 0x8a, 0xe0, 0x12, 0x4e, 0xce, 0xcf, 0xd5, 0xab, 0x4a, 0xcc, 0xa9, 0x84, 0xf8, 0x17, 0xec,
	0xa8, 0x28, 0xfd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x94, 0xd2,
	0xe4, 0xec, 0xe4, 0x8c, 0xc4, 0x12, 0x38, 0x43, 0x37, 0x3d, 0xb1, 0x24, 0xb5, 0x3c, 0xb1, 0x52,
	0x1f, 0xac, 0x41, 0x1f, 0xa4, 0xe1, 0x14, 0x13, 0x7f, 0x54, 0x62, 0x4e, 0x65, 0x4c, 0x00, 0x48,
	0x24, 0x26, 0x38, 0xb3, 0x24, 0x35, 0x89, 0x0d, 0x2c, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0x02, 0xba, 0x75, 0x61, 0x01, 0x00, 0x00,
}
