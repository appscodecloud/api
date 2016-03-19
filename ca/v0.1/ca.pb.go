// Code generated by protoc-gen-go.
// source: ca.proto
// DO NOT EDIT!

/*
Package ca is a generated protocol buffer package.

It is generated from these files:
	ca.proto

It has these top-level messages:
	CertificateCreateRequest
	CertificateCreateResponse
*/
package ca

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

// Use specific requests for protos
type CertificateCreateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Csr  string `protobuf:"bytes,2,opt,name=csr" json:"csr,omitempty"`
}

func (m *CertificateCreateRequest) Reset()                    { *m = CertificateCreateRequest{} }
func (m *CertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateRequest) ProtoMessage()               {}
func (*CertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CertificateCreateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string         `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	Cert   []byte         `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	Key    []byte         `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Root   []byte         `protobuf:"bytes,5,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *CertificateCreateResponse) Reset()                    { *m = CertificateCreateResponse{} }
func (m *CertificateCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateResponse) ProtoMessage()               {}
func (*CertificateCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateCreateRequest)(nil), "ca.CertificateCreateRequest")
	proto.RegisterType((*CertificateCreateResponse)(nil), "ca.CertificateCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for AuthorityCertificate service

type AuthorityCertificateClient interface {
	Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error)
}

type authorityCertificateClient struct {
	cc *grpc.ClientConn
}

func NewAuthorityCertificateClient(cc *grpc.ClientConn) AuthorityCertificateClient {
	return &authorityCertificateClient{cc}
}

func (c *authorityCertificateClient) Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error) {
	out := new(CertificateCreateResponse)
	err := grpc.Invoke(ctx, "/ca.AuthorityCertificate/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthorityCertificate service

type AuthorityCertificateServer interface {
	Create(context.Context, *CertificateCreateRequest) (*CertificateCreateResponse, error)
}

func RegisterAuthorityCertificateServer(s *grpc.Server, srv AuthorityCertificateServer) {
	s.RegisterService(&_AuthorityCertificate_serviceDesc, srv)
}

func _AuthorityCertificate_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthorityCertificateServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _AuthorityCertificate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.AuthorityCertificate",
	HandlerType: (*AuthorityCertificateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthorityCertificate_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for ClientCertificate service

type ClientCertificateClient interface {
	Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error)
}

type clientCertificateClient struct {
	cc *grpc.ClientConn
}

func NewClientCertificateClient(cc *grpc.ClientConn) ClientCertificateClient {
	return &clientCertificateClient{cc}
}

func (c *clientCertificateClient) Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error) {
	out := new(CertificateCreateResponse)
	err := grpc.Invoke(ctx, "/ca.ClientCertificate/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClientCertificate service

type ClientCertificateServer interface {
	Create(context.Context, *CertificateCreateRequest) (*CertificateCreateResponse, error)
}

func RegisterClientCertificateServer(s *grpc.Server, srv ClientCertificateServer) {
	s.RegisterService(&_ClientCertificate_serviceDesc, srv)
}

func _ClientCertificate_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClientCertificateServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ClientCertificate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.ClientCertificate",
	HandlerType: (*ClientCertificateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientCertificate_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x91, 0x41, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0x95, 0xb6, 0x7f, 0xf5, 0x63, 0x10, 0xa2, 0x56, 0x55, 0xb9, 0x51, 0x10, 0x28, 0x0b,
	0x84, 0x58, 0x24, 0x28, 0xec, 0x58, 0x81, 0x72, 0x83, 0x70, 0x02, 0xe3, 0xba, 0xad, 0x45, 0x89,
	0x83, 0x3d, 0x41, 0x64, 0xc3, 0x82, 0x1b, 0x20, 0x8e, 0xc6, 0x15, 0x38, 0x08, 0xf6, 0xb8, 0x48,
	0x59, 0xd0, 0x1d, 0x1b, 0xeb, 0xe9, 0x79, 0xf4, 0xbe, 0xe7, 0x31, 0xf9, 0x2f, 0x78, 0xd6, 0x18,
	0x0d, 0x9a, 0x0e, 0x04, 0x8f, 0x93, 0x95, 0xd6, 0xab, 0x8d, 0xcc, 0x79, 0xa3, 0x72, 0x5e, 0xd7,
	0x1a, 0x38, 0x28, 0x5d, 0xdb, 0x30, 0x11, 0xcf, 0xbc, 0xbd, 0x80, 0xae, 0x91, 0x36, 0xc7, 0x33,
	0xf8, 0xe9, 0x0d, 0x61, 0xa5, 0x34, 0xa0, 0x96, 0x4a, 0x70, 0x90, 0xa5, 0x91, 0xee, 0xac, 0xe4,
	0x53, 0x2b, 0x2d, 0x50, 0x4a, 0x46, 0x35, 0x7f, 0x94, 0x2c, 0x3a, 0x8d, 0xce, 0xf7, 0x2a, 0xd4,
	0xf4, 0x88, 0x0c, 0x85, 0x35, 0x6c, 0x80, 0x96, 0x97, 0xe9, 0x7b, 0x44, 0xe6, 0xbf, 0x44, 0xd8,
	0xc6, 0xc1, 0x25, 0x3d, 0x23, 0x63, 0xeb, 0x9a, 0xb4, 0x16, 0x53, 0xf6, 0x8b, 0xc3, 0x2c, 0x94,
	0xc8, 0xee, 0xd0, 0xad, 0xb6, 0xb7, 0x9e, 0xd5, 0xac, 0xd5, 0x62, 0x1b, 0x8c, 0xda, 0x7b, 0xc2,
	0x05, 0xb3, 0xa1, 0xf3, 0x0e, 0x2a, 0xd4, 0x9e, 0xff, 0x20, 0x3b, 0x36, 0x42, 0xcb, 0x4b, 0x3f,
	0x65, 0xb4, 0x06, 0xf6, 0x2f, 0x4c, 0x79, 0x5d, 0xbc, 0x92, 0xe9, 0x6d, 0x0b, 0x6b, 0x6d, 0x14,
	0x74, 0xbd, 0x6e, 0x74, 0x49, 0xc6, 0xa1, 0x1f, 0x4d, 0x32, 0xb7, 0xbc, 0x5d, 0x2f, 0x8f, 0x8f,
	0x77, 0xdc, 0x86, 0x47, 0xa5, 0x27, 0x6f, 0x9f, 0x5f, 0x1f, 0x83, 0x79, 0x3c, 0xc5, 0x65, 0x0b,
	0x9e, 0x3f, 0x5f, 0xe6, 0xfc, 0x87, 0x77, 0x1d, 0x5d, 0x14, 0x2f, 0x64, 0x52, 0x6e, 0x94, 0xac,
	0xa1, 0x0f, 0x17, 0x7f, 0x03, 0x4f, 0x10, 0x3e, 0x8b, 0x27, 0x3d, 0xb8, 0x40, 0x98, 0x23, 0xdf,
	0x8f, 0xf1, 0x5b, 0xaf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x05, 0x2e, 0x2c, 0x1c, 0x02,
	0x00, 0x00,
}
