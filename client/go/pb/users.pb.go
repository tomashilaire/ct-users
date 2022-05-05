// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Created              string   `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string   `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *User) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

type DeleteUserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SignInRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignInResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{5}
}

func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (m *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(m, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *SignInResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignUpRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,4,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{6}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SignUpRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *SignUpRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type SignUpResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{7}
}

func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(m, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *SignUpResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthenticateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{8}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

type AuthenticateResponse struct {
	AuthId               string   `protobuf:"bytes,1,opt,name=authId,proto3" json:"authId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{9}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetAuthId() string {
	if m != nil {
		return m.AuthId
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*GetUserRequest)(nil), "pb.GetUserRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "pb.ListUsersRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "pb.DeleteUserResponse")
	proto.RegisterType((*SignInRequest)(nil), "pb.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "pb.SignInResponse")
	proto.RegisterType((*SignUpRequest)(nil), "pb.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "pb.SignUpResponse")
	proto.RegisterType((*AuthenticateRequest)(nil), "pb.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "pb.AuthenticateResponse")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcb, 0x6e, 0xda, 0x40,
	0x14, 0x95, 0x8d, 0x71, 0xe9, 0xa5, 0x75, 0xdb, 0x29, 0x6d, 0x2c, 0x94, 0x05, 0xb2, 0xb2, 0x60,
	0x65, 0x24, 0xf2, 0x05, 0x44, 0x48, 0x11, 0x52, 0x16, 0x11, 0x11, 0x9b, 0xec, 0x6c, 0x7c, 0x13,
	0x46, 0x81, 0xf1, 0xc4, 0x33, 0x56, 0x94, 0xef, 0xc8, 0xb7, 0xe4, 0xff, 0xa2, 0x79, 0x81, 0x79,
	0x64, 0x95, 0xdd, 0xdc, 0x73, 0x1f, 0xe7, 0xdc, 0x73, 0x35, 0xd0, 0xad, 0x05, 0x56, 0x22, 0xe5,
	0x55, 0x29, 0x4b, 0xe2, 0xf3, 0x3c, 0x91, 0x10, 0x2c, 0x04, 0x56, 0x24, 0x02, 0x9f, 0x16, 0xb1,
	0x37, 0xf0, 0x86, 0xdf, 0xe7, 0x3e, 0x2d, 0x08, 0x81, 0x80, 0x65, 0x1b, 0x8c, 0x7d, 0x8d, 0xe8,
	0x37, 0xe9, 0x41, 0x1b, 0x37, 0x19, 0x5d, 0xc7, 0x2d, 0x0d, 0x9a, 0x80, 0xc4, 0xf0, 0x6d, 0x59,
	0x61, 0x26, 0xb1, 0x88, 0x03, 0x8d, 0xbb, 0x50, 0x65, 0x6a, 0x5e, 0xe8, 0x4c, 0xdb, 0x64, 0x6c,
	0x98, 0x0c, 0x20, 0xba, 0x46, 0xa9, 0x88, 0xe7, 0xf8, 0x5c, 0xa3, 0x90, 0x87, 0xfc, 0x09, 0x81,
	0xdf, 0x37, 0x54, 0xe8, 0x12, 0x61, 0x6b, 0x92, 0x0b, 0x20, 0x53, 0x5c, 0xa3, 0x44, 0xd3, 0x28,
	0x78, 0xc9, 0x04, 0x1e, 0x75, 0x4e, 0xe0, 0xe7, 0x1d, 0x7d, 0x64, 0x33, 0xe6, 0x46, 0x6f, 0x65,
	0x7b, 0x4d, 0xd9, 0x7d, 0xe8, 0xf0, 0x4c, 0x88, 0x97, 0xb2, 0x2a, 0xec, 0x92, 0xdb, 0x38, 0x99,
	0x42, 0xe4, 0x46, 0x58, 0x92, 0x73, 0x08, 0x94, 0x73, 0x7a, 0x44, 0x77, 0xdc, 0x49, 0x79, 0x9e,
	0x6a, 0x11, 0x1a, 0x55, 0x0c, 0xb2, 0x7c, 0x42, 0x66, 0x07, 0x99, 0x20, 0x79, 0xf3, 0x8c, 0x92,
	0x05, 0x77, 0x4a, 0x9c, 0xa9, 0xde, 0x29, 0x53, 0xfd, 0xcf, 0xd4, 0xb5, 0xf6, 0xd5, 0x91, 0x21,
	0xfc, 0x5a, 0x96, 0xec, 0x81, 0x56, 0x9b, 0x5b, 0x57, 0x62, 0x8c, 0x3f, 0x84, 0x15, 0x9f, 0x7c,
	0xe5, 0x68, 0xdd, 0xd7, 0x6f, 0xb7, 0x9b, 0x12, 0xf5, 0x85, 0xdd, 0xfe, 0xc1, 0xdf, 0x49, 0x2d,
	0x57, 0xc8, 0x24, 0x5d, 0x66, 0x12, 0xdd, 0x85, 0x52, 0xe8, 0xed, 0xc3, 0x96, 0xe2, 0x3f, 0x84,
	0x59, 0x2d, 0x57, 0x33, 0x77, 0x27, 0x1b, 0x8d, 0xdf, 0x3d, 0x88, 0x1a, 0x0d, 0xb4, 0x64, 0x64,
	0x04, 0xa1, 0xd1, 0x47, 0xfe, 0x28, 0x25, 0x7b, 0x06, 0xf6, 0x49, 0x13, 0xb2, 0xb3, 0x6d, 0xc3,
	0x8c, 0xed, 0x1a, 0xb6, 0xb7, 0xdf, 0x35, 0x34, 0x6e, 0x39, 0x81, 0x1f, 0x4d, 0x91, 0xe4, 0x4c,
	0xd5, 0x9c, 0xd8, 0xa6, 0x1f, 0x1f, 0x27, 0xcc, 0x88, 0xab, 0xf0, 0x3e, 0x48, 0x47, 0x3c, 0xcf,
	0x43, 0xfd, 0x91, 0x2e, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xe6, 0x8f, 0x43, 0x57, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/pb.Authentication/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/pb.Authentication/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/pb.Authentication/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
}

// UnimplementedAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (*UnimplementedAuthenticationServer) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthenticationServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAuthenticationServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Authentication_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Authentication_SignIn_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _Authentication_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
