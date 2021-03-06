// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/group.proto

package core // import "github.com/duckchat/duckchat-gateway/proto/core"

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

type GroupJoinPermissionType int32

const (
	GroupJoinPermissionType_GroupJoinPermissionPublic GroupJoinPermissionType = 0
	GroupJoinPermissionType_GroupJoinPermissionMember GroupJoinPermissionType = 1
	GroupJoinPermissionType_GroupJoinPermissionAdmin  GroupJoinPermissionType = 2
)

var GroupJoinPermissionType_name = map[int32]string{
	0: "GroupJoinPermissionPublic",
	1: "GroupJoinPermissionMember",
	2: "GroupJoinPermissionAdmin",
}
var GroupJoinPermissionType_value = map[string]int32{
	"GroupJoinPermissionPublic": 0,
	"GroupJoinPermissionMember": 1,
	"GroupJoinPermissionAdmin":  2,
}

func (x GroupJoinPermissionType) String() string {
	return proto.EnumName(GroupJoinPermissionType_name, int32(x))
}
func (GroupJoinPermissionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_group_b8ae31a59ac9fc67, []int{0}
}

type GroupMemberType int32

const (
	GroupMemberType_GroupMemberGuest  GroupMemberType = 0
	GroupMemberType_GroupMemberNormal GroupMemberType = 1
	GroupMemberType_GroupMemberAdmin  GroupMemberType = 2
	GroupMemberType_GroupMemberOwner  GroupMemberType = 3
)

var GroupMemberType_name = map[int32]string{
	0: "GroupMemberGuest",
	1: "GroupMemberNormal",
	2: "GroupMemberAdmin",
	3: "GroupMemberOwner",
}
var GroupMemberType_value = map[string]int32{
	"GroupMemberGuest":  0,
	"GroupMemberNormal": 1,
	"GroupMemberAdmin":  2,
	"GroupMemberOwner":  3,
}

func (x GroupMemberType) String() string {
	return proto.EnumName(GroupMemberType_name, int32(x))
}
func (GroupMemberType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_group_b8ae31a59ac9fc67, []int{1}
}

type GroupDescriptionType int32

const (
	GroupDescriptionType_GroupDescriptionText     GroupDescriptionType = 0
	GroupDescriptionType_GroupDescriptionMarkdown GroupDescriptionType = 1
)

var GroupDescriptionType_name = map[int32]string{
	0: "GroupDescriptionText",
	1: "GroupDescriptionMarkdown",
}
var GroupDescriptionType_value = map[string]int32{
	"GroupDescriptionText":     0,
	"GroupDescriptionMarkdown": 1,
}

func (x GroupDescriptionType) String() string {
	return proto.EnumName(GroupDescriptionType_name, int32(x))
}
func (GroupDescriptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_group_b8ae31a59ac9fc67, []int{2}
}

type GroupDescription struct {
	Type                 GroupDescriptionType `protobuf:"varint,1,opt,name=type,proto3,enum=core.GroupDescriptionType" json:"type,omitempty"`
	Body                 string               `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GroupDescription) Reset()         { *m = GroupDescription{} }
func (m *GroupDescription) String() string { return proto.CompactTextString(m) }
func (*GroupDescription) ProtoMessage()    {}
func (*GroupDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_b8ae31a59ac9fc67, []int{0}
}
func (m *GroupDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupDescription.Unmarshal(m, b)
}
func (m *GroupDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupDescription.Marshal(b, m, deterministic)
}
func (dst *GroupDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupDescription.Merge(dst, src)
}
func (m *GroupDescription) XXX_Size() int {
	return xxx_messageInfo_GroupDescription.Size(m)
}
func (m *GroupDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupDescription.DiscardUnknown(m)
}

var xxx_messageInfo_GroupDescription proto.InternalMessageInfo

func (m *GroupDescription) GetType() GroupDescriptionType {
	if m != nil {
		return m.Type
	}
	return GroupDescriptionType_GroupDescriptionText
}

func (m *GroupDescription) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type PublicGroupProfile struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NameInLatin          string                  `protobuf:"bytes,3,opt,name=nameInLatin,proto3" json:"nameInLatin,omitempty"`
	Avatar               string                  `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Description          *GroupDescription       `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PermissionJoin       GroupJoinPermissionType `protobuf:"varint,6,opt,name=permissionJoin,proto3,enum=core.GroupJoinPermissionType" json:"permissionJoin,omitempty"`
	CanGuestReadMessage  bool                    `protobuf:"varint,7,opt,name=canGuestReadMessage,proto3" json:"canGuestReadMessage,omitempty"`
	TimeCreate           int64                   `protobuf:"varint,8,opt,name=timeCreate,proto3" json:"timeCreate,omitempty"`
	Owner                *PublicUserProfile      `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Admins               []*PublicUserProfile    `protobuf:"bytes,10,rep,name=admins,proto3" json:"admins,omitempty"`
	Speakers             []*PublicUserProfile    `protobuf:"bytes,11,rep,name=speakers,proto3" json:"speakers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PublicGroupProfile) Reset()         { *m = PublicGroupProfile{} }
func (m *PublicGroupProfile) String() string { return proto.CompactTextString(m) }
func (*PublicGroupProfile) ProtoMessage()    {}
func (*PublicGroupProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_b8ae31a59ac9fc67, []int{1}
}
func (m *PublicGroupProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicGroupProfile.Unmarshal(m, b)
}
func (m *PublicGroupProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicGroupProfile.Marshal(b, m, deterministic)
}
func (dst *PublicGroupProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicGroupProfile.Merge(dst, src)
}
func (m *PublicGroupProfile) XXX_Size() int {
	return xxx_messageInfo_PublicGroupProfile.Size(m)
}
func (m *PublicGroupProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicGroupProfile.DiscardUnknown(m)
}

var xxx_messageInfo_PublicGroupProfile proto.InternalMessageInfo

func (m *PublicGroupProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PublicGroupProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PublicGroupProfile) GetNameInLatin() string {
	if m != nil {
		return m.NameInLatin
	}
	return ""
}

func (m *PublicGroupProfile) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *PublicGroupProfile) GetDescription() *GroupDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *PublicGroupProfile) GetPermissionJoin() GroupJoinPermissionType {
	if m != nil {
		return m.PermissionJoin
	}
	return GroupJoinPermissionType_GroupJoinPermissionPublic
}

func (m *PublicGroupProfile) GetCanGuestReadMessage() bool {
	if m != nil {
		return m.CanGuestReadMessage
	}
	return false
}

func (m *PublicGroupProfile) GetTimeCreate() int64 {
	if m != nil {
		return m.TimeCreate
	}
	return 0
}

func (m *PublicGroupProfile) GetOwner() *PublicUserProfile {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PublicGroupProfile) GetAdmins() []*PublicUserProfile {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *PublicGroupProfile) GetSpeakers() []*PublicUserProfile {
	if m != nil {
		return m.Speakers
	}
	return nil
}

// TODO:
// remove this message
type AllGroupProfile struct {
	Profile              *PublicGroupProfile  `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Owner                *PublicUserProfile   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Admins               []*PublicUserProfile `protobuf:"bytes,3,rep,name=admins,proto3" json:"admins,omitempty"`
	Speakers             []*PublicUserProfile `protobuf:"bytes,4,rep,name=speakers,proto3" json:"speakers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AllGroupProfile) Reset()         { *m = AllGroupProfile{} }
func (m *AllGroupProfile) String() string { return proto.CompactTextString(m) }
func (*AllGroupProfile) ProtoMessage()    {}
func (*AllGroupProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_group_b8ae31a59ac9fc67, []int{2}
}
func (m *AllGroupProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllGroupProfile.Unmarshal(m, b)
}
func (m *AllGroupProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllGroupProfile.Marshal(b, m, deterministic)
}
func (dst *AllGroupProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllGroupProfile.Merge(dst, src)
}
func (m *AllGroupProfile) XXX_Size() int {
	return xxx_messageInfo_AllGroupProfile.Size(m)
}
func (m *AllGroupProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_AllGroupProfile.DiscardUnknown(m)
}

var xxx_messageInfo_AllGroupProfile proto.InternalMessageInfo

func (m *AllGroupProfile) GetProfile() *PublicGroupProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *AllGroupProfile) GetOwner() *PublicUserProfile {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AllGroupProfile) GetAdmins() []*PublicUserProfile {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *AllGroupProfile) GetSpeakers() []*PublicUserProfile {
	if m != nil {
		return m.Speakers
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupDescription)(nil), "core.GroupDescription")
	proto.RegisterType((*PublicGroupProfile)(nil), "core.PublicGroupProfile")
	proto.RegisterType((*AllGroupProfile)(nil), "core.AllGroupProfile")
	proto.RegisterEnum("core.GroupJoinPermissionType", GroupJoinPermissionType_name, GroupJoinPermissionType_value)
	proto.RegisterEnum("core.GroupMemberType", GroupMemberType_name, GroupMemberType_value)
	proto.RegisterEnum("core.GroupDescriptionType", GroupDescriptionType_name, GroupDescriptionType_value)
}

func init() { proto.RegisterFile("core/group.proto", fileDescriptor_group_b8ae31a59ac9fc67) }

var fileDescriptor_group_b8ae31a59ac9fc67 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x6e, 0xd3, 0x3c,
	0x14, 0x5f, 0x92, 0xae, 0xdb, 0x4e, 0xa5, 0x36, 0x9f, 0xb7, 0x6f, 0x33, 0x13, 0x43, 0x51, 0xaf,
	0xaa, 0x4a, 0x4b, 0x50, 0x77, 0xc3, 0xed, 0x18, 0x68, 0x02, 0xd1, 0x51, 0x45, 0x80, 0xd0, 0xb4,
	0x1b, 0x37, 0x39, 0x74, 0xa6, 0x49, 0x1c, 0x39, 0x09, 0x25, 0x3c, 0x22, 0x8f, 0xb0, 0xa7, 0x41,
	0x76, 0xda, 0x2e, 0x54, 0xd9, 0xc4, 0xae, 0x7a, 0xfa, 0xfb, 0x73, 0x7c, 0x7e, 0xc7, 0x75, 0xc1,
	0x0e, 0x84, 0x44, 0x6f, 0x26, 0x45, 0x91, 0xba, 0xa9, 0x14, 0xb9, 0x20, 0x2d, 0x85, 0x1c, 0xf7,
	0x34, 0x5e, 0x64, 0x28, 0x2b, 0xb8, 0xff, 0x05, 0xec, 0x4b, 0xa5, 0x7a, 0x83, 0x59, 0x20, 0x79,
	0x9a, 0x73, 0x91, 0x10, 0x17, 0x5a, 0x79, 0x99, 0x22, 0x35, 0x1c, 0x63, 0xd0, 0x1d, 0x1d, 0xbb,
	0xca, 0xe3, 0x6e, 0xaa, 0x3e, 0x95, 0x29, 0xfa, 0x5a, 0x47, 0x08, 0xb4, 0xa6, 0x22, 0x2c, 0xa9,
	0xe9, 0x18, 0x83, 0x3d, 0x5f, 0xd7, 0xfd, 0x3b, 0x0b, 0xc8, 0xa4, 0x98, 0x46, 0x3c, 0xd0, 0xc6,
	0x89, 0x14, 0xdf, 0x78, 0x84, 0xa4, 0x0b, 0x26, 0x0f, 0x75, 0xe3, 0x3d, 0xdf, 0xe4, 0xa1, 0xb2,
	0x26, 0x2c, 0xc6, 0x95, 0x55, 0xd5, 0xc4, 0x81, 0x8e, 0xfa, 0x7c, 0x97, 0x7c, 0x60, 0x39, 0x4f,
	0xa8, 0xa5, 0xa9, 0x3a, 0x44, 0x0e, 0xa1, 0xcd, 0x7e, 0xb0, 0x9c, 0x49, 0xda, 0xd2, 0xe4, 0xf2,
	0x1b, 0x79, 0x05, 0x9d, 0xf0, 0x7e, 0x42, 0xba, 0xed, 0x18, 0x83, 0xce, 0xe8, 0xb0, 0x79, 0x7e,
	0xbf, 0x2e, 0x25, 0x6f, 0xa1, 0x9b, 0xa2, 0x8c, 0x79, 0x96, 0x71, 0x91, 0xbc, 0x17, 0x3c, 0xa1,
	0x6d, 0x1d, 0xfe, 0xa4, 0x66, 0x56, 0xf0, 0x64, 0x2d, 0xd2, 0xf9, 0x37, 0x4c, 0xe4, 0x25, 0xec,
	0x07, 0x2c, 0xb9, 0x2c, 0x30, 0xcb, 0x7d, 0x64, 0xe1, 0x18, 0xb3, 0x8c, 0xcd, 0x90, 0xee, 0x38,
	0xc6, 0x60, 0xd7, 0x6f, 0xa2, 0xc8, 0x0b, 0x80, 0x9c, 0xc7, 0x78, 0x21, 0x91, 0xe5, 0x48, 0x77,
	0x1d, 0x63, 0x60, 0xf9, 0x35, 0x84, 0x9c, 0xc2, 0xb6, 0x58, 0x24, 0x28, 0xe9, 0x9e, 0x0e, 0x73,
	0x54, 0xcd, 0x53, 0x6d, 0xf6, 0x73, 0x86, 0x72, 0xb9, 0x58, 0xbf, 0x52, 0x11, 0x0f, 0xda, 0x2c,
	0x8c, 0x79, 0x92, 0x51, 0x70, 0xac, 0xc7, 0xf4, 0x4b, 0x19, 0x39, 0x83, 0xdd, 0x2c, 0x45, 0x36,
	0x47, 0x99, 0xd1, 0xce, 0xe3, 0x96, 0xb5, 0xb0, 0x7f, 0x67, 0x40, 0xef, 0x3c, 0x8a, 0xfe, 0xba,
	0xd9, 0x11, 0xec, 0xa4, 0x55, 0xa9, 0xaf, 0xb7, 0x33, 0xa2, 0xf5, 0x3e, 0x75, 0xa9, 0xbf, 0x12,
	0xde, 0x87, 0x33, 0x9f, 0x18, 0xce, 0x7a, 0x7a, 0xb8, 0xd6, 0x3f, 0x86, 0x1b, 0x16, 0x70, 0xf4,
	0xc0, 0x75, 0x93, 0x13, 0x78, 0xd6, 0x40, 0x55, 0xcd, 0xec, 0xad, 0x07, 0xe8, 0x31, 0xc6, 0x53,
	0x94, 0xb6, 0x41, 0x9e, 0x03, 0x6d, 0xa0, 0xcf, 0xd5, 0xa8, 0xb6, 0x39, 0xfc, 0x0e, 0x3d, 0xcd,
	0x56, 0x72, 0x7d, 0xdc, 0xc1, 0xf2, 0x6d, 0x56, 0x90, 0xfe, 0xe9, 0xd8, 0x5b, 0xe4, 0x7f, 0xf8,
	0xaf, 0x86, 0x5e, 0x09, 0x19, 0xb3, 0xc8, 0x36, 0x36, 0xc4, 0xcb, 0xae, 0x1b, 0xe8, 0x47, 0xb5,
	0x46, 0xdb, 0x1a, 0x5e, 0xc1, 0x41, 0xd3, 0x73, 0x26, 0xb4, 0x01, 0xc7, 0x9f, 0xea, 0xd0, 0xd5,
	0xec, 0x35, 0x66, 0xcc, 0xe4, 0x3c, 0x14, 0x8b, 0xc4, 0x36, 0x5e, 0x7f, 0x85, 0xfd, 0x40, 0xc4,
	0xee, 0x2f, 0x16, 0x95, 0xd5, 0xdf, 0x8a, 0xde, 0xf2, 0xb5, 0x37, 0xe3, 0xf9, 0x6d, 0x31, 0x75,
	0x03, 0x11, 0x7b, 0x61, 0x11, 0xcc, 0x83, 0x5b, 0x96, 0xaf, 0x8b, 0xd3, 0x19, 0xcb, 0x71, 0xc1,
	0x4a, 0x4f, 0x1b, 0x3c, 0x65, 0xf8, 0x6d, 0xf6, 0xae, 0x59, 0x54, 0xde, 0x4c, 0x14, 0x72, 0x73,
	0x21, 0x24, 0x4e, 0xdb, 0x9a, 0x3d, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xa5, 0x41, 0x50,
	0xd0, 0x04, 0x00, 0x00,
}
