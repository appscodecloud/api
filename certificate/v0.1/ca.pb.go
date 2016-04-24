// Code generated by protoc-gen-go.
// source: ca.proto
// DO NOT EDIT!

/*
Package certificate is a generated protocol buffer package.

It is generated from these files:
	ca.proto
	freessl_certificate.proto
	tls_certificate.proto

It has these top-level messages:
	CertificateCreateRequest
	CertificateCreateResponse
	FreeSSLCertificateRegisterRequest
	FreeSSLCertificateRegisterResponse
	FreeSSLCertificate
	FreeSSLCertificateCreateRequest
	FreeSSLCertificateCreateResponse
	FreeSSLCertificateRevokeRequest
	FreeSSLCertificateRenewRequest
	FreeSSLCertificateRenewResponse
	FreeSSLCertificateDescribeRequest
	DNSCheckRequest
	TLSCertificateListResponse
	TLSCertificateDescribeResponse
	TLSCertificate
	TLSCertificateCreateRequest
	TLSCertificateDeleteRequest
	TLSCertificateDescribeRequest
	TLSCertificateDeployRequest
*/
package certificate

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
	proto.RegisterType((*CertificateCreateRequest)(nil), "certificate.CertificateCreateRequest")
	proto.RegisterType((*CertificateCreateResponse)(nil), "certificate.CertificateCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

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
	err := grpc.Invoke(ctx, "/certificate.AuthorityCertificate/Create", in, out, c.cc, opts...)
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

func _AuthorityCertificate_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityCertificateServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.AuthorityCertificate/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityCertificateServer).Create(ctx, req.(*CertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorityCertificate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.AuthorityCertificate",
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
	err := grpc.Invoke(ctx, "/certificate.ClientCertificate/Create", in, out, c.cc, opts...)
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

func _ClientCertificate_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientCertificateServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.ClientCertificate/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientCertificateServer).Create(ctx, req.(*CertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientCertificate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.ClientCertificate",
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
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0xc9, 0x36, 0x8b, 0x66, 0x22, 0x1a, 0x44, 0x62, 0x19, 0xa8, 0x45, 0xa7, 0xec, 0x90,
	0x6a, 0xbd, 0x79, 0x52, 0xfa, 0x06, 0xf5, 0x09, 0x62, 0x17, 0xb7, 0xe0, 0x6c, 0x6a, 0x92, 0x0a,
	0xbd, 0x89, 0x3e, 0x81, 0x08, 0x9e, 0x7c, 0x2b, 0x5f, 0xc1, 0x07, 0xb1, 0xf9, 0x67, 0x62, 0x0f,
	0x13, 0x4f, 0xbb, 0x84, 0x8f, 0xaf, 0x5f, 0xbf, 0xff, 0x2f, 0xff, 0x16, 0xaf, 0xe6, 0x9c, 0x95,
	0x5a, 0x59, 0x45, 0xfa, 0xb9, 0xd0, 0x56, 0xde, 0xca, 0x9c, 0x5b, 0x11, 0x0e, 0x26, 0x4a, 0x4d,
	0x66, 0x22, 0xe6, 0xa5, 0x8c, 0x79, 0x51, 0x28, 0xcb, 0xad, 0x54, 0x85, 0xf1, 0xd1, 0x70, 0xc7,
	0xd9, 0x63, 0x5b, 0x97, 0xc2, 0xc4, 0x70, 0x7a, 0x3f, 0xba, 0xc4, 0x34, 0xfd, 0x2d, 0x49, 0xb5,
	0x68, 0xce, 0x4c, 0x3c, 0x54, 0xc2, 0x58, 0x42, 0x70, 0xaf, 0xe0, 0xf7, 0x82, 0xa2, 0x7d, 0x74,
	0xb2, 0x96, 0x81, 0x26, 0x9b, 0xb8, 0x9b, 0x1b, 0x4d, 0x3b, 0x60, 0x39, 0x19, 0xbd, 0x22, 0xbc,
	0xbb, 0xa0, 0xc2, 0x94, 0xcd, 0x70, 0x41, 0x86, 0x38, 0x30, 0x0d, 0x49, 0x65, 0xa0, 0xa5, 0x9f,
	0x6c, 0x30, 0x0f, 0xc1, 0xae, 0xc1, 0xcd, 0xe6, 0x4f, 0xdd, 0xac, 0x72, 0x2a, 0xc7, 0xf3, 0x62,
	0xd0, 0xce, 0x73, 0x17, 0xa4, 0xdd, 0xc6, 0x5b, 0xcf, 0x40, 0xbb, 0xf9, 0x77, 0xa2, 0xa6, 0x3d,
	0xb0, 0x9c, 0x74, 0x29, 0xad, 0x94, 0xa5, 0x2b, 0x3e, 0xe5, 0x74, 0xf2, 0x81, 0xf0, 0xf6, 0x55,
	0x65, 0xa7, 0x4a, 0x4b, 0x5b, 0xb7, 0xe0, 0xc8, 0x0b, 0xc2, 0x81, 0x27, 0x24, 0x47, 0xac, 0xb5,
	0x3d, 0xf6, 0xd7, 0x12, 0xc2, 0xe1, 0x7f, 0x31, 0x7f, 0xd1, 0x68, 0xf4, 0xfc, 0xf9, 0xf5, 0xd6,
	0x39, 0x0c, 0xf7, 0xe0, 0x03, 0xb4, 0xde, 0x89, 0x1f, 0x4f, 0xd9, 0x59, 0xcc, 0x7f, 0x68, 0x2e,
	0xd0, 0x28, 0x79, 0x47, 0x78, 0x2b, 0x9d, 0x49, 0x51, 0xd8, 0x36, 0xdb, 0xd3, 0xd2, 0xd8, 0x8e,
	0x81, 0xed, 0x20, 0x1c, 0x2c, 0x66, 0xcb, 0x01, 0xa5, 0x01, 0xbb, 0x09, 0xe0, 0xa7, 0x38, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x96, 0xa6, 0x05, 0x13, 0x63, 0x02, 0x00, 0x00,
}
