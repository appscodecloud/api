// Code generated by protoc-gen-go.
// source: authentication.proto
// DO NOT EDIT!

/*
Package authentication is a generated protocol buffer package.

It is generated from these files:
	authentication.proto
	conduit.proto

It has these top-level messages:
	TokenRequest
	TokenResponse
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
	ConduitRequest
	ConduitWhoAmIResponse
	ConduitUsersResponse
	ConduitUser
*/
package authentication

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Next Id 4
type TokenRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TokenResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *TokenResponse) Reset()                    { *m = TokenResponse{} }
func (m *TokenResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()               {}
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TokenResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type LoginRequest struct {
	Namespace  string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Secret     string `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	IssueToken bool   `protobuf:"varint,4,opt,name=issue_token,json=issueToken" json:"issue_token,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type LoginResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Token  string         `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type LogoutRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type LogoutResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogoutResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenRequest)(nil), "authentication.TokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "authentication.TokenResponse")
	proto.RegisterType((*LoginRequest)(nil), "authentication.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "authentication.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "authentication.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "authentication.LogoutResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Authentication service

type AuthenticationClient interface {
	// This rpc is used to check a valid user from other applications.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	// appctl used this to validates the user token with phabricator.
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/authentication.Authentication/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := grpc.Invoke(ctx, "/authentication.Authentication/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/authentication.Authentication/Token", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	// This rpc is used to check a valid user from other applications.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	// appctl used this to validates the user token with phabricator.
	Token(context.Context, *TokenRequest) (*TokenResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Authentication/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authentication_Logout_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Authentication_Token_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x4d, 0x79, 0x0f, 0x02, 0x17, 0xe8, 0x62, 0x42, 0x08, 0x69, 0x80, 0xf7, 0x32, 0x26, 0x06,
	0x5d, 0xb4, 0x8a, 0x0b, 0x8d, 0x3b, 0x75, 0xab, 0x9b, 0xea, 0xde, 0xd4, 0x32, 0xc1, 0x46, 0x9c,
	0xa9, 0x9d, 0xa9, 0x89, 0x4b, 0xfd, 0x05, 0x3f, 0xcd, 0x5f, 0xe0, 0x43, 0x9c, 0xde, 0x19, 0x05,
	0x44, 0x54, 0xe2, 0x86, 0xe4, 0x9e, 0x73, 0xb9, 0xe7, 0xdc, 0x3b, 0xa7, 0xd0, 0x8a, 0x72, 0x75,
	0xcd, 0xb8, 0x4a, 0xe2, 0x48, 0x25, 0x82, 0xfb, 0x69, 0x26, 0x94, 0x20, 0xee, 0x22, 0xea, 0x75,
	0xc7, 0x42, 0x8c, 0x27, 0x2c, 0x88, 0xd2, 0x24, 0x88, 0x38, 0x17, 0x0a, 0x61, 0x69, 0xba, 0xbd,
	0x76, 0x01, 0x8f, 0xd4, 0x43, 0xca, 0x64, 0x80, 0xbf, 0x06, 0xa7, 0xc7, 0xd0, 0xb8, 0x10, 0x37,
	0x8c, 0x87, 0xec, 0x2e, 0x67, 0x52, 0x91, 0x2e, 0xd4, 0x78, 0x74, 0xcb, 0x64, 0x1a, 0xc5, 0xac,
	0xe3, 0xfc, 0x77, 0x06, 0xb5, 0x70, 0x06, 0x90, 0x16, 0x94, 0x55, 0xd1, 0xdd, 0x29, 0x21, 0x63,
	0x0a, 0xba, 0x0f, 0x4d, 0x3b, 0x43, 0xa6, 0x5a, 0x91, 0x91, 0x4d, 0xa8, 0x48, 0x2d, 0x9f, 0x4b,
	0x9c, 0x50, 0x1f, 0xba, 0xbe, 0x51, 0xf6, 0xcf, 0x11, 0x0d, 0x2d, 0x4b, 0x1f, 0x1d, 0x68, 0x9c,
	0x8a, 0x71, 0xf2, 0x43, 0x75, 0x0f, 0xaa, 0xb9, 0x64, 0x59, 0x01, 0x58, 0x03, 0xef, 0x35, 0x69,
	0x6b, 0x49, 0x16, 0x67, 0x4c, 0x75, 0xfe, 0x20, 0x63, 0x2b, 0xf2, 0x0f, 0xea, 0x89, 0x94, 0x39,
	0xbb, 0x34, 0xbe, 0xff, 0x6a, 0xb2, 0x1a, 0x02, 0x42, 0xe8, 0x99, 0x9e, 0x41, 0xd3, 0x5a, 0x58,
	0xcf, 0xfc, 0x8a, 0x5b, 0x9c, 0xe0, 0x38, 0x91, 0xab, 0xdf, 0x1c, 0xf4, 0x00, 0xdc, 0xb7, 0x21,
	0xeb, 0x99, 0x1a, 0x4e, 0x4b, 0xe0, 0x1e, 0x2d, 0xe4, 0x82, 0x70, 0x28, 0xe3, 0x82, 0xa4, 0xeb,
	0x7f, 0xc8, 0xd1, 0xfc, 0xe9, 0xbd, 0xde, 0x0a, 0xd6, 0x18, 0xa0, 0x83, 0xa7, 0x97, 0xe9, 0x73,
	0x89, 0xd2, 0x9e, 0xce, 0x57, 0x2a, 0x63, 0x31, 0xb2, 0x41, 0xd3, 0xff, 0x09, 0xee, 0x77, 0xfc,
	0xdd, 0x60, 0x52, 0xb4, 0x1f, 0x3a, 0xdb, 0x24, 0x83, 0x8a, 0x31, 0x4f, 0x3e, 0x1b, 0x39, 0xbb,
	0x8c, 0xd7, 0x5f, 0x45, 0x5b, 0xc9, 0x2d, 0x94, 0xdc, 0xa0, 0xfd, 0x2f, 0x24, 0x75, 0x7f, 0xa1,
	0xa9, 0x77, 0xc4, 0xd7, 0x5c, 0xde, 0x71, 0x3e, 0xdc, 0xcb, 0x3b, 0x2e, 0xc4, 0xf6, 0xfb, 0x1d,
	0xf1, 0x75, 0xb4, 0xde, 0x55, 0x05, 0x3f, 0x9e, 0xbd, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0x12, 0xea, 0x4b, 0x9a, 0x03, 0x00, 0x00,
}
