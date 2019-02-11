// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_command.proto

package pb

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

type UserRegisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterRequest) Reset()         { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()    {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_command_a025a86c7e8a5698, []int{0}
}
func (m *UserRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterRequest.Unmarshal(m, b)
}
func (m *UserRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterRequest.Marshal(b, m, deterministic)
}
func (dst *UserRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterRequest.Merge(dst, src)
}
func (m *UserRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegisterRequest.Size(m)
}
func (m *UserRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterRequest proto.InternalMessageInfo

type UserRegisterResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterResponse) Reset()         { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()    {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_command_a025a86c7e8a5698, []int{1}
}
func (m *UserRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResponse.Unmarshal(m, b)
}
func (m *UserRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResponse.Marshal(b, m, deterministic)
}
func (dst *UserRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResponse.Merge(dst, src)
}
func (m *UserRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResponse.Size(m)
}
func (m *UserRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResponse proto.InternalMessageInfo

func (m *UserRegisterResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserUpdateProfileRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserUpdateProfileRequest) Reset()         { *m = UserUpdateProfileRequest{} }
func (m *UserUpdateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UserUpdateProfileRequest) ProtoMessage()    {}
func (*UserUpdateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_command_a025a86c7e8a5698, []int{2}
}
func (m *UserUpdateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateProfileRequest.Unmarshal(m, b)
}
func (m *UserUpdateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateProfileRequest.Marshal(b, m, deterministic)
}
func (dst *UserUpdateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateProfileRequest.Merge(dst, src)
}
func (m *UserUpdateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UserUpdateProfileRequest.Size(m)
}
func (m *UserUpdateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateProfileRequest proto.InternalMessageInfo

func (m *UserUpdateProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserUpdateProfileResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserUpdateProfileResponse) Reset()         { *m = UserUpdateProfileResponse{} }
func (m *UserUpdateProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UserUpdateProfileResponse) ProtoMessage()    {}
func (*UserUpdateProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_command_a025a86c7e8a5698, []int{3}
}
func (m *UserUpdateProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateProfileResponse.Unmarshal(m, b)
}
func (m *UserUpdateProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateProfileResponse.Marshal(b, m, deterministic)
}
func (dst *UserUpdateProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateProfileResponse.Merge(dst, src)
}
func (m *UserUpdateProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UserUpdateProfileResponse.Size(m)
}
func (m *UserUpdateProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateProfileResponse proto.InternalMessageInfo

func (m *UserUpdateProfileResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserRegisterRequest)(nil), "user.UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "user.UserRegisterResponse")
	proto.RegisterType((*UserUpdateProfileRequest)(nil), "user.UserUpdateProfileRequest")
	proto.RegisterType((*UserUpdateProfileResponse)(nil), "user.UserUpdateProfileResponse")
}

func init() { proto.RegisterFile("user_command.proto", fileDescriptor_user_command_a025a86c7e8a5698) }

var fileDescriptor_user_command_a025a86c7e8a5698 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2d, 0x4e, 0x2d,
	0x8a, 0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x89, 0x49, 0x71, 0x81, 0x48, 0x88, 0x88, 0x92, 0x28, 0x97, 0x70, 0x68, 0x71, 0x6a, 0x51,
	0x50, 0x6a, 0x7a, 0x66, 0x71, 0x09, 0x88, 0x2e, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x32, 0xe3, 0x12,
	0x41, 0x15, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe3, 0x02, 0x1b, 0x21, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xa5, 0x07, 0x36, 0x09, 0xac, 0x12, 0x2c, 0xae, 0xa4, 0xc7, 0x25,
	0x01, 0xe2, 0x85, 0x16, 0xa4, 0x24, 0x96, 0xa4, 0x06, 0x14, 0xe5, 0xa7, 0x65, 0xe6, 0xa4, 0x42,
	0xcd, 0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x05, 0xeb, 0xe5, 0x0c, 0x02, 0xb3, 0x95,
	0xac, 0xb9, 0x24, 0xb1, 0xa8, 0x27, 0xce, 0x32, 0xa3, 0x45, 0x8c, 0x5c, 0xdc, 0x20, 0xae, 0x33,
	0xc4, 0x8f, 0x42, 0x8e, 0x5c, 0x1c, 0x30, 0x07, 0x0b, 0x49, 0x22, 0xa9, 0x46, 0xf5, 0x9b, 0x94,
	0x14, 0x36, 0x29, 0xa8, 0x95, 0x01, 0x5c, 0xbc, 0x28, 0x6e, 0x11, 0x92, 0x43, 0x28, 0xc6, 0xe6,
	0x29, 0x29, 0x79, 0x9c, 0xf2, 0x10, 0x13, 0x9d, 0x58, 0xa2, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0,
	0xa1, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x3c, 0xd3, 0x71, 0x95, 0x01, 0x00, 0x00,
}