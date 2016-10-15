// Code generated by protoc-gen-go.
// source: acmeuser.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	acmeuser.proto
	ca.proto
	certificate.proto

It has these top-level messages:
	AcmeUserRegisterRequest
	AcmeUserRegisterResponse
	CACreateRequest
	CACreateResponse
	CertificateListRequest
	CertificateListResponse
	CertificateDescribeRequest
	CertificateDescribeResponse
	CertificateCreateRequest
	SubjectInfo
	CertificateCreateResponse
	Certificate
	CertificateObtainRequest
	CertificateDeleteRequest
	CertificateRenewRequest
	CertificateRenewResponse
	CertificateRevokeRequest
	CertificateDeployRequest
	DNSCheckRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AcmeUserRegisterRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *AcmeUserRegisterRequest) Reset()                    { *m = AcmeUserRegisterRequest{} }
func (m *AcmeUserRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*AcmeUserRegisterRequest) ProtoMessage()               {}
func (*AcmeUserRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AcmeUserRegisterResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string                  `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *AcmeUserRegisterResponse) Reset()                    { *m = AcmeUserRegisterResponse{} }
func (m *AcmeUserRegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*AcmeUserRegisterResponse) ProtoMessage()               {}
func (*AcmeUserRegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AcmeUserRegisterResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AcmeUserRegisterRequest)(nil), "appscode.certificate.v1beta1.AcmeUserRegisterRequest")
	proto.RegisterType((*AcmeUserRegisterResponse)(nil), "appscode.certificate.v1beta1.AcmeUserRegisterResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AcmeUsers service

type AcmeUsersClient interface {
	Register(ctx context.Context, in *AcmeUserRegisterRequest, opts ...grpc.CallOption) (*AcmeUserRegisterResponse, error)
}

type acmeUsersClient struct {
	cc *grpc.ClientConn
}

func NewAcmeUsersClient(cc *grpc.ClientConn) AcmeUsersClient {
	return &acmeUsersClient{cc}
}

func (c *acmeUsersClient) Register(ctx context.Context, in *AcmeUserRegisterRequest, opts ...grpc.CallOption) (*AcmeUserRegisterResponse, error) {
	out := new(AcmeUserRegisterResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.AcmeUsers/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AcmeUsers service

type AcmeUsersServer interface {
	Register(context.Context, *AcmeUserRegisterRequest) (*AcmeUserRegisterResponse, error)
}

func RegisterAcmeUsersServer(s *grpc.Server, srv AcmeUsersServer) {
	s.RegisterService(&_AcmeUsers_serviceDesc, srv)
}

func _AcmeUsers_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcmeUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcmeUsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.AcmeUsers/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcmeUsersServer).Register(ctx, req.(*AcmeUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AcmeUsers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.certificate.v1beta1.AcmeUsers",
	HandlerType: (*AcmeUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AcmeUsers_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("acmeuser.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x0a, 0x0a, 0x35, 0x82, 0x85, 0x85, 0x68, 0x15, 0x75, 0x81, 0x22, 0x16, 0x2c,
	0xc0, 0x56, 0xcb, 0x63, 0x0d, 0x3d, 0x41, 0x15, 0xc4, 0x86, 0x0d, 0x72, 0x9d, 0x21, 0x58, 0x6a,
	0x62, 0x93, 0x99, 0x20, 0xb1, 0xe5, 0x0a, 0x1c, 0x84, 0x13, 0x70, 0x0a, 0xae, 0xc0, 0x41, 0x50,
	0xec, 0x04, 0x90, 0x78, 0x48, 0x6c, 0xfc, 0x98, 0xf1, 0x3f, 0x9f, 0x67, 0x7e, 0xb6, 0xa5, 0x74,
	0x01, 0x35, 0x42, 0x25, 0x5c, 0x65, 0xc9, 0xf2, 0xb1, 0x72, 0x0e, 0xb5, 0xcd, 0x40, 0x68, 0xa8,
	0xc8, 0xdc, 0x18, 0xad, 0x08, 0xc4, 0xfd, 0x64, 0x01, 0xa4, 0x26, 0xf1, 0x38, 0xb7, 0x36, 0x5f,
	0x82, 0x54, 0xce, 0x48, 0x55, 0x96, 0x96, 0x14, 0x19, 0x5b, 0x62, 0xd0, 0xc6, 0x3b, 0x4d, 0x38,
	0xa3, 0x07, 0x07, 0x28, 0xfd, 0x1a, 0xe2, 0x89, 0x64, 0xc3, 0x73, 0x5d, 0xc0, 0x25, 0x42, 0x95,
	0x42, 0x6e, 0x90, 0x9a, 0xfd, 0xae, 0x06, 0x24, 0xbe, 0xcd, 0x56, 0xa1, 0x50, 0x66, 0x39, 0x8a,
	0x76, 0xa3, 0xfd, 0x41, 0x1a, 0x2e, 0xc9, 0x35, 0x1b, 0x7d, 0x17, 0xa0, 0xb3, 0x25, 0x02, 0x97,
	0xac, 0x8f, 0xa4, 0xa8, 0x46, 0x2f, 0xd9, 0x98, 0x0e, 0xc5, 0xc7, 0x8f, 0x03, 0x5a, 0x5c, 0xf8,
	0x74, 0xda, 0x3e, 0xe3, 0x9c, 0xad, 0xb8, 0x5b, 0x93, 0x8d, 0x7a, 0x9e, 0xe0, 0xcf, 0xd3, 0x97,
	0x88, 0x0d, 0x3a, 0x02, 0xf2, 0xe7, 0x88, 0xad, 0x77, 0x1c, 0x7e, 0x22, 0xfe, 0x9a, 0x80, 0xf8,
	0xa5, 0x91, 0xf8, 0xf4, 0xbf, 0xb2, 0xd0, 0x4e, 0x72, 0xfc, 0xf8, 0xfa, 0xf6, 0xd4, 0x13, 0xf1,
	0x81, 0xec, 0xf4, 0x7e, 0xb8, 0x5f, 0x6a, 0xc8, 0xb6, 0x86, 0x6c, 0x9c, 0x3a, 0x6c, 0xac, 0xc2,
	0xd9, 0x19, 0xdb, 0xd3, 0xb6, 0xf8, 0x44, 0x2a, 0x67, 0x7e, 0xc2, 0xce, 0x36, 0x3b, 0xee, 0xbc,
	0x31, 0x62, 0x1e, 0x5d, 0xad, 0xb5, 0x99, 0x45, 0xdf, 0x5b, 0x73, 0xf4, 0x1e, 0x00, 0x00, 0xff,
	0xff, 0x35, 0xf3, 0x4b, 0xc1, 0x00, 0x02, 0x00, 0x00,
}
