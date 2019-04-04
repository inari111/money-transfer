// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_query.proto

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

type UserGetRequest struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetRequest) Reset()         { *m = UserGetRequest{} }
func (m *UserGetRequest) String() string { return proto.CompactTextString(m) }
func (*UserGetRequest) ProtoMessage()    {}
func (*UserGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_query_72f4fd42b85e5263, []int{0}
}
func (m *UserGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetRequest.Unmarshal(m, b)
}
func (m *UserGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetRequest.Marshal(b, m, deterministic)
}
func (dst *UserGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetRequest.Merge(dst, src)
}
func (m *UserGetRequest) XXX_Size() int {
	return xxx_messageInfo_UserGetRequest.Size(m)
}
func (m *UserGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetRequest proto.InternalMessageInfo

func (m *UserGetRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UserGetResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetResponse) Reset()         { *m = UserGetResponse{} }
func (m *UserGetResponse) String() string { return proto.CompactTextString(m) }
func (*UserGetResponse) ProtoMessage()    {}
func (*UserGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_query_72f4fd42b85e5263, []int{1}
}
func (m *UserGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetResponse.Unmarshal(m, b)
}
func (m *UserGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetResponse.Marshal(b, m, deterministic)
}
func (dst *UserGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetResponse.Merge(dst, src)
}
func (m *UserGetResponse) XXX_Size() int {
	return xxx_messageInfo_UserGetResponse.Size(m)
}
func (m *UserGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetResponse proto.InternalMessageInfo

func (m *UserGetResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserGetRequest)(nil), "user.UserGetRequest")
	proto.RegisterType((*UserGetResponse)(nil), "user.UserGetResponse")
}

func init() { proto.RegisterFile("user_query.proto", fileDescriptor_user_query_72f4fd42b85e5263) }

var fileDescriptor_user_query_72f4fd42b85e5263 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89,
	0x48, 0x71, 0x81, 0x48, 0x88, 0x88, 0x92, 0x26, 0x17, 0x5f, 0x68, 0x71, 0x6a, 0x91, 0x7b, 0x6a,
	0x49, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x38, 0x17, 0x3b, 0x58, 0x5f, 0x66, 0x8a,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x1b, 0x88, 0xeb, 0x99, 0xa2, 0x64, 0xc8, 0xc5, 0x0f,
	0x57, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc7, 0x05, 0x36, 0x11, 0xac, 0x90, 0xdb,
	0x88, 0x4b, 0x0f, 0x6c, 0x30, 0x48, 0x51, 0x10, 0x58, 0xdc, 0xc8, 0x9e, 0x8b, 0x13, 0xc4, 0x0b,
	0x04, 0x39, 0x41, 0xc8, 0x88, 0x8b, 0xd9, 0x3d, 0xb5, 0x44, 0x48, 0x04, 0xa1, 0x0a, 0x61, 0xab,
	0x94, 0x28, 0x9a, 0x28, 0xc4, 0x02, 0x27, 0x96, 0x28, 0xa6, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x5b,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xd5, 0x3c, 0x45, 0xd1, 0x00, 0x00, 0x00,
}
