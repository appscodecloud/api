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
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x91, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xb4, 0x44, 0xe0, 0x56, 0x48, 0xb5, 0xaa, 0xca, 0x8d, 0xc2, 0x8f, 0xb2, 0x42,
	0x2c, 0x12, 0x14, 0xc4, 0x86, 0x1d, 0xca, 0x0d, 0xe0, 0x04, 0xc6, 0x9d, 0xb6, 0x16, 0x25, 0x0e,
	0xf6, 0x04, 0x91, 0x0d, 0x0b, 0xae, 0xc0, 0xd1, 0xb8, 0x02, 0x07, 0xc1, 0x3f, 0x42, 0xca, 0x82,
	0xee, 0xd8, 0x8c, 0xf4, 0xde, 0x64, 0xde, 0x37, 0x19, 0x93, 0x43, 0xc1, 0x8b, 0x56, 0x2b, 0x54,
	0x34, 0x16, 0x3c, 0xcd, 0x36, 0x4a, 0x6d, 0x76, 0x50, 0xf2, 0x56, 0x96, 0xbc, 0x69, 0x14, 0x72,
	0x94, 0xaa, 0x31, 0xe1, 0x8b, 0x74, 0xe1, 0xec, 0x15, 0xf6, 0x2d, 0x98, 0xd2, 0xd7, 0xe0, 0xe7,
	0x37, 0x84, 0xd5, 0xa0, 0x51, 0xae, 0xa5, 0xe0, 0x08, 0xb5, 0x06, 0x5b, 0xef, 0xe1, 0xa5, 0x03,
	0x83, 0x74, 0x4a, 0xc6, 0x0d, 0x7f, 0x06, 0x16, 0x9d, 0x47, 0x17, 0x47, 0x74, 0x42, 0x46, 0xc2,
	0x68, 0x16, 0x3b, 0x91, 0x6b, 0xb2, 0xfc, 0x63, 0xcc, 0xb4, 0x16, 0x08, 0xf4, 0x94, 0x24, 0xc6,
	0xd2, 0x3b, 0xe3, 0x27, 0x27, 0xd5, 0x71, 0x11, 0xc0, 0xc5, 0x83, 0x77, 0x5d, 0x6e, 0xbb, 0x95,
	0xab, 0x10, 0xe5, 0x94, 0xb0, 0x51, 0x6c, 0x64, 0xd5, 0xd4, 0x51, 0x9e, 0xa0, 0x67, 0x63, 0x2f,
	0x6c, 0x4b, 0x2b, 0x85, 0xec, 0xc0, 0xa9, 0xea, 0x9d, 0xcc, 0xef, 0x3a, 0xdc, 0x2a, 0x2d, 0xb1,
	0x1f, 0xc0, 0xe9, 0x9a, 0x24, 0x61, 0x01, 0x9a, 0x15, 0xf6, 0x22, 0xfb, 0x7e, 0x27, 0x3d, 0xd9,
	0xd3, 0x0d, 0x5b, 0xe7, 0x67, 0x1f, 0x5f, 0xdf, 0x9f, 0xf1, 0x32, 0x9d, 0xfb, 0x0b, 0x0a, 0x5e,
	0xbe, 0x5e, 0x95, 0xfc, 0x97, 0x77, 0x1b, 0x5d, 0x56, 0x6f, 0x64, 0x56, 0xef, 0x24, 0x34, 0x38,
	0x84, 0x8b, 0xff, 0x81, 0x67, 0x1e, 0xbe, 0x48, 0x67, 0x03, 0xb8, 0xf0, 0x30, 0x4b, 0x7e, 0x4c,
	0xfc, 0x5b, 0x5d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0xc6, 0xba, 0x27, 0xf1, 0x01, 0x00,
	0x00,
}