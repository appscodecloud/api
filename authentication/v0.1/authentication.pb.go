// Code generated by protoc-gen-go.
// source: authentication.proto
// DO NOT EDIT!

/*
Package authentication is a generated protocol buffer package.

It is generated from these files:
	authentication.proto

It has these top-level messages:
	ValidateRequest
	ValidateResponse
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
type ValidateRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Secret    string `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
}

func (m *ValidateRequest) Reset()                    { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()               {}
func (*ValidateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ValidateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ValidateResponse) Reset()                    { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string            { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()               {}
func (*ValidateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidateRequest)(nil), "authentication.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "authentication.ValidateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Authentication service

type AuthenticationClient interface {
	// This rpc is used to check a valid user from other applications.
	User(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	// appctl used this to validates the user token with phabricator.
	Token(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) User(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := grpc.Invoke(ctx, "/authentication.Authentication/User", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Token(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := grpc.Invoke(ctx, "/authentication.Authentication/Token", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	// This rpc is used to check a valid user from other applications.
	User(context.Context, *ValidateRequest) (*ValidateResponse, error)
	// appctl used this to validates the user token with phabricator.
	Token(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthenticationServer).User(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Authentication_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthenticationServer).Token(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "User",
			Handler:    _Authentication_User_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Authentication_Token_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x90, 0x3b, 0x4e, 0xc4, 0x30,
	0x10, 0x86, 0x95, 0x05, 0x56, 0xec, 0x20, 0x85, 0xc5, 0x5a, 0x50, 0x88, 0x10, 0x1b, 0xa5, 0x42,
	0x14, 0x36, 0x84, 0x8e, 0x8e, 0x86, 0x03, 0xf0, 0xea, 0x4d, 0x76, 0xb4, 0x04, 0x82, 0x6d, 0xe2,
	0x09, 0x12, 0x2d, 0x57, 0xe0, 0x68, 0x5c, 0x81, 0x8a, 0x53, 0xe0, 0x47, 0x81, 0x36, 0x12, 0x1d,
	0x8d, 0x25, 0x7f, 0xfa, 0x35, 0xdf, 0xfc, 0x03, 0x33, 0xd9, 0xd3, 0x03, 0x2a, 0x6a, 0x6a, 0x49,
	0x8d, 0x56, 0xdc, 0x74, 0x9a, 0x34, 0x4b, 0x57, 0x69, 0x7e, 0xb0, 0xd4, 0x7a, 0xd9, 0xa2, 0x90,
	0xa6, 0x11, 0x52, 0x29, 0x4d, 0x01, 0xdb, 0x98, 0xce, 0xf7, 0x3c, 0x5e, 0xd0, 0x9b, 0x41, 0x2b,
	0xc2, 0x1b, 0x79, 0x79, 0x09, 0xdb, 0x77, 0xb2, 0x6d, 0x16, 0x92, 0xf0, 0x0a, 0x5f, 0x7a, 0xb4,
	0xc4, 0x76, 0x60, 0xa2, 0xe4, 0x33, 0x5a, 0x23, 0x6b, 0xcc, 0x92, 0x22, 0x39, 0x9a, 0xb0, 0x29,
	0x6c, 0xf6, 0x16, 0x3b, 0x8f, 0xb3, 0x51, 0x20, 0x29, 0x8c, 0x2d, 0xd6, 0x1d, 0x52, 0xb6, 0xe6,
	0xff, 0x65, 0x05, 0xd3, 0xdf, 0x39, 0xd6, 0x38, 0x31, 0xb2, 0x43, 0x97, 0x71, 0x5b, 0xf4, 0x36,
	0x4c, 0xd9, 0xaa, 0x52, 0x1e, 0x17, 0xe0, 0xd7, 0x81, 0x56, 0xdf, 0x09, 0xa4, 0x17, 0x2b, 0x25,
	0xd8, 0x23, 0xac, 0xdf, 0x3a, 0x11, 0x9b, 0xf3, 0x41, 0xe7, 0xc1, 0x92, 0x79, 0xf1, 0x77, 0x20,
	0xda, 0xcb, 0xf9, 0xfb, 0xe7, 0xd7, 0xc7, 0x68, 0xbf, 0x9c, 0xc5, 0x8b, 0xb8, 0xb4, 0x12, 0xaf,
	0x27, 0xfc, 0x54, 0xf8, 0x26, 0xe7, 0xc9, 0x31, 0x6b, 0x61, 0xe3, 0x46, 0x3f, 0xa1, 0xfa, 0x0f,
	0x59, 0x11, 0x64, 0x79, 0xb9, 0x3b, 0x94, 0x91, 0x37, 0x38, 0xdb, 0xfd, 0x38, 0xdc, 0xfb, 0xec,
	0x27, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x5f, 0x3c, 0x75, 0xcd, 0x01, 0x00, 0x00,
}
