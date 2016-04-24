// Code generated by protoc-gen-go.
// source: freessl_certificate.proto
// DO NOT EDIT!

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

type FreeSSLCertificateRegisterRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *FreeSSLCertificateRegisterRequest) Reset()         { *m = FreeSSLCertificateRegisterRequest{} }
func (m *FreeSSLCertificateRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRegisterRequest) ProtoMessage()    {}
func (*FreeSSLCertificateRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0}
}

type FreeSSLCertificateRegisterResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string         `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *FreeSSLCertificateRegisterResponse) Reset()         { *m = FreeSSLCertificateRegisterResponse{} }
func (m *FreeSSLCertificateRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRegisterResponse) ProtoMessage()    {}
func (*FreeSSLCertificateRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1}
}

func (m *FreeSSLCertificateRegisterResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type FreeSSLCertificate struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	IssuedBy   string `protobuf:"bytes,4,opt,name=issued_by,json=issuedBy" json:"issued_by,omitempty"`
	ValidFrom  string `protobuf:"bytes,5,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	ExpireDate string `protobuf:"bytes,6,opt,name=expire_date,json=expireDate" json:"expire_date,omitempty"`
	// those feilds will not included into list response.
	// only describe response will include the underlying
	// feilds.
	Sans         []string `protobuf:"bytes,7,rep,name=sans" json:"sans,omitempty"`
	Cert         string   `protobuf:"bytes,8,opt,name=cert" json:"cert,omitempty"`
	Key          string   `protobuf:"bytes,9,opt,name=key" json:"key,omitempty"`
	Version      int32    `protobuf:"varint,10,opt,name=version" json:"version,omitempty"`
	SerialNumber string   `protobuf:"bytes,11,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
}

func (m *FreeSSLCertificate) Reset()                    { *m = FreeSSLCertificate{} }
func (m *FreeSSLCertificate) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificate) ProtoMessage()               {}
func (*FreeSSLCertificate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type FreeSSLCertificateCreateRequest struct {
	AccountPhid string   `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName  string   `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	San         []string `protobuf:"bytes,4,rep,name=san" json:"san,omitempty"`
	Bundle      bool     `protobuf:"varint,5,opt,name=bundle" json:"bundle,omitempty"`
	KeyData     string   `protobuf:"bytes,6,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
}

func (m *FreeSSLCertificateCreateRequest) Reset()                    { *m = FreeSSLCertificateCreateRequest{} }
func (m *FreeSSLCertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateCreateRequest) ProtoMessage()               {}
func (*FreeSSLCertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type FreeSSLCertificateCreateResponse struct {
	Status      *dtypes.Status      `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *FreeSSLCertificate `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *FreeSSLCertificateCreateResponse) Reset()         { *m = FreeSSLCertificateCreateResponse{} }
func (m *FreeSSLCertificateCreateResponse) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateCreateResponse) ProtoMessage()    {}
func (*FreeSSLCertificateCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{4}
}

func (m *FreeSSLCertificateCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FreeSSLCertificateCreateResponse) GetCertificate() *FreeSSLCertificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type FreeSSLCertificateRevokeRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	CertPhid    string `protobuf:"bytes,2,opt,name=cert_phid,json=certPhid" json:"cert_phid,omitempty"`
}

func (m *FreeSSLCertificateRevokeRequest) Reset()                    { *m = FreeSSLCertificateRevokeRequest{} }
func (m *FreeSSLCertificateRevokeRequest) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRevokeRequest) ProtoMessage()               {}
func (*FreeSSLCertificateRevokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type FreeSSLCertificateRenewRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CertPhid    string `protobuf:"bytes,3,opt,name=cert_phid,json=certPhid" json:"cert_phid,omitempty"`
}

func (m *FreeSSLCertificateRenewRequest) Reset()                    { *m = FreeSSLCertificateRenewRequest{} }
func (m *FreeSSLCertificateRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRenewRequest) ProtoMessage()               {}
func (*FreeSSLCertificateRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type FreeSSLCertificateRenewResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Cert   string         `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
}

func (m *FreeSSLCertificateRenewResponse) Reset()                    { *m = FreeSSLCertificateRenewResponse{} }
func (m *FreeSSLCertificateRenewResponse) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRenewResponse) ProtoMessage()               {}
func (*FreeSSLCertificateRenewResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *FreeSSLCertificateRenewResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type FreeSSLCertificateDescribeRequest struct {
	Cert string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
}

func (m *FreeSSLCertificateDescribeRequest) Reset()         { *m = FreeSSLCertificateDescribeRequest{} }
func (m *FreeSSLCertificateDescribeRequest) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateDescribeRequest) ProtoMessage()    {}
func (*FreeSSLCertificateDescribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{8}
}

type DNSCheckRequest struct {
	Domain  string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	KeyAuth string `protobuf:"bytes,2,opt,name=key_auth,json=keyAuth" json:"key_auth,omitempty"`
}

func (m *DNSCheckRequest) Reset()                    { *m = DNSCheckRequest{} }
func (m *DNSCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*DNSCheckRequest) ProtoMessage()               {}
func (*DNSCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func init() {
	proto.RegisterType((*FreeSSLCertificateRegisterRequest)(nil), "certificate.FreeSSLCertificateRegisterRequest")
	proto.RegisterType((*FreeSSLCertificateRegisterResponse)(nil), "certificate.FreeSSLCertificateRegisterResponse")
	proto.RegisterType((*FreeSSLCertificate)(nil), "certificate.FreeSSLCertificate")
	proto.RegisterType((*FreeSSLCertificateCreateRequest)(nil), "certificate.FreeSSLCertificateCreateRequest")
	proto.RegisterType((*FreeSSLCertificateCreateResponse)(nil), "certificate.FreeSSLCertificateCreateResponse")
	proto.RegisterType((*FreeSSLCertificateRevokeRequest)(nil), "certificate.FreeSSLCertificateRevokeRequest")
	proto.RegisterType((*FreeSSLCertificateRenewRequest)(nil), "certificate.FreeSSLCertificateRenewRequest")
	proto.RegisterType((*FreeSSLCertificateRenewResponse)(nil), "certificate.FreeSSLCertificateRenewResponse")
	proto.RegisterType((*FreeSSLCertificateDescribeRequest)(nil), "certificate.FreeSSLCertificateDescribeRequest")
	proto.RegisterType((*DNSCheckRequest)(nil), "certificate.DNSCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for FreeSSLCertificates service

type FreeSSLCertificatesClient interface {
	Register(ctx context.Context, in *FreeSSLCertificateRegisterRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRegisterResponse, error)
	Create(ctx context.Context, in *FreeSSLCertificateCreateRequest, opts ...grpc.CallOption) (*FreeSSLCertificateCreateResponse, error)
	Revoke(ctx context.Context, in *FreeSSLCertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Renew(ctx context.Context, in *FreeSSLCertificateRenewRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRenewResponse, error)
}

type freeSSLCertificatesClient struct {
	cc *grpc.ClientConn
}

func NewFreeSSLCertificatesClient(cc *grpc.ClientConn) FreeSSLCertificatesClient {
	return &freeSSLCertificatesClient{cc}
}

func (c *freeSSLCertificatesClient) Register(ctx context.Context, in *FreeSSLCertificateRegisterRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRegisterResponse, error) {
	out := new(FreeSSLCertificateRegisterResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Create(ctx context.Context, in *FreeSSLCertificateCreateRequest, opts ...grpc.CallOption) (*FreeSSLCertificateCreateResponse, error) {
	out := new(FreeSSLCertificateCreateResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Revoke(ctx context.Context, in *FreeSSLCertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Revoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Renew(ctx context.Context, in *FreeSSLCertificateRenewRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRenewResponse, error) {
	out := new(FreeSSLCertificateRenewResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FreeSSLCertificates service

type FreeSSLCertificatesServer interface {
	Register(context.Context, *FreeSSLCertificateRegisterRequest) (*FreeSSLCertificateRegisterResponse, error)
	Create(context.Context, *FreeSSLCertificateCreateRequest) (*FreeSSLCertificateCreateResponse, error)
	Revoke(context.Context, *FreeSSLCertificateRevokeRequest) (*dtypes.VoidResponse, error)
	Renew(context.Context, *FreeSSLCertificateRenewRequest) (*FreeSSLCertificateRenewResponse, error)
}

func RegisterFreeSSLCertificatesServer(s *grpc.Server, srv FreeSSLCertificatesServer) {
	s.RegisterService(&_FreeSSLCertificates_serviceDesc, srv)
}

func _FreeSSLCertificates_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Register(ctx, req.(*FreeSSLCertificateRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Create(ctx, req.(*FreeSSLCertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Revoke(ctx, req.(*FreeSSLCertificateRevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Renew(ctx, req.(*FreeSSLCertificateRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FreeSSLCertificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.FreeSSLCertificates",
	HandlerType: (*FreeSSLCertificatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _FreeSSLCertificates_Register_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FreeSSLCertificates_Create_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _FreeSSLCertificates_Revoke_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _FreeSSLCertificates_Renew_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0x93, 0x26, 0x4d, 0x6e, 0xfa, 0xfd, 0x68, 0xbe, 0x7e, 0xc5, 0x35, 0x94, 0xa6, 0xae,
	0x84, 0xa0, 0x94, 0x18, 0xca, 0x02, 0xc1, 0xae, 0x34, 0xea, 0x0a, 0x55, 0xc8, 0x91, 0xd8, 0x21,
	0x33, 0xb1, 0x6f, 0x5b, 0xab, 0x8e, 0x1d, 0x3c, 0x76, 0x4a, 0x84, 0xd8, 0xb0, 0x41, 0x42, 0x62,
	0xc5, 0xa3, 0x20, 0xb6, 0xbc, 0x04, 0xaf, 0xc0, 0x83, 0x30, 0x3f, 0x76, 0xe2, 0x50, 0x43, 0x1a,
	0xd8, 0x58, 0x33, 0x67, 0xee, 0x9d, 0x73, 0xef, 0x99, 0x33, 0x63, 0x58, 0x3f, 0x8e, 0x11, 0x19,
	0x0b, 0x1c, 0x17, 0xe3, 0xc4, 0x3f, 0xf6, 0x5d, 0x9a, 0x60, 0x67, 0x18, 0x47, 0x49, 0x44, 0x5a,
	0x05, 0xc8, 0xb8, 0x76, 0x12, 0x45, 0x27, 0x01, 0x5a, 0x74, 0xe8, 0x5b, 0x34, 0x0c, 0xa3, 0x84,
	0x26, 0x7e, 0x14, 0x32, 0x15, 0x6a, 0xac, 0x09, 0xd8, 0x4b, 0xc6, 0x43, 0x64, 0x96, 0xfc, 0x2a,
	0xdc, 0x7c, 0x08, 0x5b, 0x87, 0x7c, 0xff, 0x5e, 0xef, 0xc9, 0xc1, 0x74, 0x2f, 0x1b, 0x4f, 0x7c,
	0x96, 0x60, 0x6c, 0xe3, 0xcb, 0x14, 0x59, 0x42, 0x56, 0xa1, 0x86, 0x03, 0xea, 0x07, 0xba, 0xd6,
	0xd6, 0x6e, 0x36, 0x6d, 0x35, 0x31, 0x5f, 0x80, 0xf9, 0xab, 0x54, 0x36, 0xe4, 0xec, 0x48, 0x6e,
	0x40, 0x9d, 0xf1, 0x52, 0x52, 0x26, 0x93, 0x5b, 0x7b, 0x7f, 0x77, 0x54, 0x15, 0x9d, 0x9e, 0x44,
	0xed, 0x6c, 0x95, 0x10, 0x58, 0x1a, 0x9e, 0xfa, 0x9e, 0x5e, 0x91, 0x14, 0x72, 0x6c, 0x7e, 0xaa,
	0x00, 0xb9, 0x48, 0x31, 0x09, 0xd5, 0xa6, 0xa1, 0x02, 0x0b, 0xe9, 0x00, 0xf3, 0x74, 0x31, 0x26,
	0x9b, 0xd0, 0x72, 0xa3, 0xc1, 0x20, 0x0a, 0x1d, 0xb9, 0x54, 0x95, 0x4b, 0xa0, 0xa0, 0x23, 0x11,
	0x70, 0x15, 0x9a, 0x3e, 0x63, 0x29, 0x7a, 0x4e, 0x7f, 0xac, 0x2f, 0xc9, 0xe5, 0x86, 0x02, 0x1e,
	0x8f, 0xc9, 0x06, 0xc0, 0x88, 0x06, 0xbe, 0xe7, 0x1c, 0xc7, 0xd1, 0x40, 0xaf, 0xc9, 0xd5, 0xa6,
	0x44, 0x0e, 0x39, 0x20, 0x36, 0xc7, 0x57, 0x43, 0x3f, 0x46, 0xc7, 0xe3, 0x35, 0xe9, 0x75, 0xb5,
	0xb9, 0x82, 0xba, 0x59, 0x95, 0x8c, 0x86, 0x4c, 0x5f, 0x6e, 0x57, 0x45, 0x45, 0x62, 0x2c, 0x30,
	0x71, 0x64, 0x7a, 0x43, 0x55, 0x29, 0xc6, 0xe4, 0x5f, 0xa8, 0x9e, 0xe1, 0x58, 0x6f, 0x4a, 0x48,
	0x0c, 0x89, 0x0e, 0xcb, 0x23, 0x8c, 0x19, 0x3f, 0x3d, 0x1d, 0x38, 0x5a, 0xb3, 0xf3, 0x29, 0xd9,
	0x86, 0xbf, 0x18, 0xc6, 0x3e, 0x0d, 0x9c, 0x30, 0x1d, 0xf4, 0x31, 0xd6, 0x5b, 0x32, 0x6b, 0x45,
	0x81, 0x47, 0x12, 0x33, 0xbf, 0x68, 0xb0, 0x79, 0x51, 0xb5, 0x83, 0x18, 0xe5, 0xf1, 0xa8, 0x13,
	0xdd, 0x82, 0x15, 0xea, 0xba, 0x51, 0x1a, 0x26, 0x4e, 0x41, 0xca, 0x56, 0x86, 0x3d, 0xfd, 0x6d,
	0x45, 0x79, 0x33, 0xbc, 0x51, 0xae, 0xa5, 0xe8, 0x59, 0x0c, 0xc9, 0x1a, 0xd4, 0xfb, 0x69, 0xe8,
	0x05, 0x28, 0x25, 0x6c, 0xd8, 0xd9, 0x8c, 0xac, 0x43, 0x83, 0xf7, 0x2a, 0xc4, 0xa3, 0x99, 0x78,
	0xcb, 0x7c, 0xce, 0x95, 0xa3, 0xe6, 0x07, 0x0d, 0xda, 0x3f, 0x6f, 0x60, 0x41, 0x5f, 0xed, 0x43,
	0xf1, 0x96, 0xc8, 0x6e, 0x5a, 0x7b, 0x9b, 0x9d, 0xe2, 0x65, 0x2a, 0x71, 0x71, 0x31, 0xc7, 0xa4,
	0x65, 0x7a, 0xda, 0x38, 0x8a, 0xce, 0x16, 0xd1, 0x93, 0x9b, 0x4d, 0x6c, 0xea, 0x14, 0x5c, 0xde,
	0x10, 0x80, 0x58, 0x34, 0x13, 0xb8, 0x5e, 0x46, 0x11, 0xe2, 0xf9, 0x1f, 0x9e, 0xd8, 0x0c, 0x6b,
	0xf5, 0x07, 0xd6, 0xe7, 0xe5, 0x8d, 0x49, 0xd6, 0xc5, 0xaf, 0xaf, 0x74, 0x76, 0x65, 0xea, 0x6c,
	0xf3, 0x41, 0xd9, 0xdb, 0xd2, 0x45, 0xe6, 0xc6, 0x7e, 0x7f, 0xa2, 0x5c, 0x9e, 0xa8, 0x15, 0x12,
	0xbb, 0xf0, 0x4f, 0xf7, 0xa8, 0x77, 0x70, 0x8a, 0xee, 0x59, 0x1e, 0xc6, 0x6d, 0xe4, 0x45, 0xfc,
	0xd9, 0x09, 0xb3, 0xc0, 0x6c, 0x96, 0xdb, 0x88, 0xa6, 0xc9, 0x69, 0xc6, 0x2d, 0x6c, 0xb4, 0xcf,
	0xa7, 0x7b, 0x9f, 0x97, 0xe0, 0xbf, 0x8b, 0xfc, 0x8c, 0xbc, 0xd7, 0xa0, 0x91, 0x3f, 0x53, 0xa4,
	0x33, 0xcf, 0x09, 0xb3, 0x4f, 0xa1, 0x61, 0x5d, 0x3a, 0x5e, 0x09, 0x68, 0x6e, 0xbc, 0xfd, 0xfa,
	0xed, 0x63, 0xe5, 0x0a, 0xf9, 0x5f, 0x3e, 0xcc, 0xd9, 0x5b, 0x6e, 0x8d, 0xee, 0x76, 0xee, 0x59,
	0x49, 0xc0, 0xc8, 0x3b, 0x0d, 0xea, 0xca, 0xd9, 0x64, 0x77, 0xce, 0xd6, 0x33, 0x37, 0xd8, 0xb8,
	0x73, 0xc9, 0xe8, 0xac, 0x8c, 0xb6, 0x2c, 0xc3, 0x30, 0xca, 0xcb, 0x78, 0xa4, 0xed, 0x90, 0x73,
	0xa8, 0x2b, 0x4f, 0xcf, 0x2d, 0x64, 0xc6, 0xfa, 0xc6, 0x6a, 0xee, 0x88, 0x67, 0x91, 0xef, 0x4d,
	0xf8, 0x6e, 0x49, 0xbe, 0xed, 0x9d, 0xad, 0x52, 0x3e, 0xeb, 0xf5, 0xc4, 0x94, 0x6f, 0x08, 0xbf,
	0xee, 0x35, 0x69, 0x3a, 0x72, 0x7b, 0x2e, 0xf1, 0xf4, 0x42, 0x18, 0xbb, 0x97, 0x0b, 0x5e, 0xb8,
	0x9e, 0x7e, 0x5d, 0xfe, 0x19, 0xef, 0x7f, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x78, 0xaf, 0xdb, 0x9d,
	0x79, 0x07, 0x00, 0x00,
}
