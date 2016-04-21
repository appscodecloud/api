// Code generated by protoc-gen-go.
// source: tls_certificate.proto
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

type TLSCertificateListResponse struct {
	Status       *dtypes.Status    `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificates []*TLSCertificate `protobuf:"bytes,2,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *TLSCertificateListResponse) Reset()                    { *m = TLSCertificateListResponse{} }
func (m *TLSCertificateListResponse) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificateListResponse) ProtoMessage()               {}
func (*TLSCertificateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TLSCertificateListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TLSCertificateListResponse) GetCertificates() []*TLSCertificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type TLSCertificateDescribeResponse struct {
	Status      *dtypes.Status  `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *TLSCertificate `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *TLSCertificateDescribeResponse) Reset()                    { *m = TLSCertificateDescribeResponse{} }
func (m *TLSCertificateDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificateDescribeResponse) ProtoMessage()               {}
func (*TLSCertificateDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TLSCertificateDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TLSCertificateDescribeResponse) GetCertificate() *TLSCertificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type TLSCertificate struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name" json:"common_name,omitempty"`
	IssuedBy   string `protobuf:"bytes,4,opt,name=issued_by" json:"issued_by,omitempty"`
	ValidFrom  string `protobuf:"bytes,5,opt,name=valid_from" json:"valid_from,omitempty"`
	ExpireDate string `protobuf:"bytes,6,opt,name=expire_date" json:"expire_date,omitempty"`
	// those feilds will not included into list response.
	// only describe response will include the underlying
	// feilds.
	Sans    []string `protobuf:"bytes,7,rep,name=sans" json:"sans,omitempty"`
	Cert    string   `protobuf:"bytes,8,opt,name=cert" json:"cert,omitempty"`
	Key     string   `protobuf:"bytes,9,opt,name=key" json:"key,omitempty"`
	Version int32    `protobuf:"varint,10,opt,name=version" json:"version,omitempty"`
}

func (m *TLSCertificate) Reset()                    { *m = TLSCertificate{} }
func (m *TLSCertificate) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificate) ProtoMessage()               {}
func (*TLSCertificate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type TLSCertificateCreateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CertData string `protobuf:"bytes,2,opt,name=cert_data" json:"cert_data,omitempty"`
	KeyData  string `protobuf:"bytes,3,opt,name=key_data" json:"key_data,omitempty"`
}

func (m *TLSCertificateCreateRequest) Reset()                    { *m = TLSCertificateCreateRequest{} }
func (m *TLSCertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificateCreateRequest) ProtoMessage()               {}
func (*TLSCertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type TLSCertificateDeleteRequest struct {
	Cert string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
}

func (m *TLSCertificateDeleteRequest) Reset()                    { *m = TLSCertificateDeleteRequest{} }
func (m *TLSCertificateDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificateDeleteRequest) ProtoMessage()               {}
func (*TLSCertificateDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type TLSCertificateDescribeRequest struct {
	Cert string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
}

func (m *TLSCertificateDescribeRequest) Reset()                    { *m = TLSCertificateDescribeRequest{} }
func (m *TLSCertificateDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificateDescribeRequest) ProtoMessage()               {}
func (*TLSCertificateDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type TLSCertificateDeployRequest struct {
	Cert        string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name" json:"secret_name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *TLSCertificateDeployRequest) Reset()                    { *m = TLSCertificateDeployRequest{} }
func (m *TLSCertificateDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificateDeployRequest) ProtoMessage()               {}
func (*TLSCertificateDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func init() {
	proto.RegisterType((*TLSCertificateListResponse)(nil), "certificate.TLSCertificateListResponse")
	proto.RegisterType((*TLSCertificateDescribeResponse)(nil), "certificate.TLSCertificateDescribeResponse")
	proto.RegisterType((*TLSCertificate)(nil), "certificate.TLSCertificate")
	proto.RegisterType((*TLSCertificateCreateRequest)(nil), "certificate.TLSCertificateCreateRequest")
	proto.RegisterType((*TLSCertificateDeleteRequest)(nil), "certificate.TLSCertificateDeleteRequest")
	proto.RegisterType((*TLSCertificateDescribeRequest)(nil), "certificate.TLSCertificateDescribeRequest")
	proto.RegisterType((*TLSCertificateDeployRequest)(nil), "certificate.TLSCertificateDeployRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for TLSCertificates service

type TLSCertificatesClient interface {
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*TLSCertificateListResponse, error)
	Describe(ctx context.Context, in *TLSCertificateDescribeRequest, opts ...grpc.CallOption) (*TLSCertificateDescribeResponse, error)
	Create(ctx context.Context, in *TLSCertificateCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *TLSCertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Deploy(ctx context.Context, in *TLSCertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type tLSCertificatesClient struct {
	cc *grpc.ClientConn
}

func NewTLSCertificatesClient(cc *grpc.ClientConn) TLSCertificatesClient {
	return &tLSCertificatesClient{cc}
}

func (c *tLSCertificatesClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*TLSCertificateListResponse, error) {
	out := new(TLSCertificateListResponse)
	err := grpc.Invoke(ctx, "/certificate.TLSCertificates/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tLSCertificatesClient) Describe(ctx context.Context, in *TLSCertificateDescribeRequest, opts ...grpc.CallOption) (*TLSCertificateDescribeResponse, error) {
	out := new(TLSCertificateDescribeResponse)
	err := grpc.Invoke(ctx, "/certificate.TLSCertificates/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tLSCertificatesClient) Create(ctx context.Context, in *TLSCertificateCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.TLSCertificates/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tLSCertificatesClient) Delete(ctx context.Context, in *TLSCertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.TLSCertificates/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tLSCertificatesClient) Deploy(ctx context.Context, in *TLSCertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.TLSCertificates/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TLSCertificates service

type TLSCertificatesServer interface {
	List(context.Context, *dtypes.VoidRequest) (*TLSCertificateListResponse, error)
	Describe(context.Context, *TLSCertificateDescribeRequest) (*TLSCertificateDescribeResponse, error)
	Create(context.Context, *TLSCertificateCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *TLSCertificateDeleteRequest) (*dtypes.VoidResponse, error)
	Deploy(context.Context, *TLSCertificateDeployRequest) (*dtypes.VoidResponse, error)
}

func RegisterTLSCertificatesServer(s *grpc.Server, srv TLSCertificatesServer) {
	s.RegisterService(&_TLSCertificates_serviceDesc, srv)
}

func _TLSCertificates_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TLSCertificatesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.TLSCertificates/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TLSCertificatesServer).List(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TLSCertificates_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSCertificateDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TLSCertificatesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.TLSCertificates/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TLSCertificatesServer).Describe(ctx, req.(*TLSCertificateDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TLSCertificates_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSCertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TLSCertificatesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.TLSCertificates/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TLSCertificatesServer).Create(ctx, req.(*TLSCertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TLSCertificates_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSCertificateDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TLSCertificatesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.TLSCertificates/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TLSCertificatesServer).Delete(ctx, req.(*TLSCertificateDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TLSCertificates_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSCertificateDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TLSCertificatesServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.TLSCertificates/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TLSCertificatesServer).Deploy(ctx, req.(*TLSCertificateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TLSCertificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.TLSCertificates",
	HandlerType: (*TLSCertificatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TLSCertificates_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _TLSCertificates_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TLSCertificates_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TLSCertificates_Delete_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _TLSCertificates_Deploy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x94, 0xdb, 0x34, 0x4d, 0x5e, 0xaa, 0x16, 0xb6, 0x05, 0x19, 0x07, 0xa2, 0x60, 0x21, 0x35,
	0x6a, 0xc1, 0x6e, 0xc3, 0x8d, 0x6b, 0x7b, 0xec, 0x01, 0xb5, 0x88, 0xab, 0xe5, 0xd8, 0xaf, 0x61,
	0x85, 0xe3, 0x35, 0xbb, 0x9b, 0x88, 0x08, 0xb8, 0x70, 0xe7, 0xc4, 0xc7, 0xf0, 0x21, 0xfc, 0x02,
	0x47, 0x3e, 0x82, 0xdd, 0x75, 0x0c, 0x76, 0xa8, 0x93, 0x5c, 0x2c, 0xed, 0xe4, 0xed, 0xcc, 0xbc,
	0xf1, 0x38, 0xf0, 0x40, 0x26, 0x22, 0x88, 0x90, 0x4b, 0x7a, 0x4b, 0xa3, 0x50, 0xa2, 0x97, 0x71,
	0x26, 0x19, 0xe9, 0x94, 0x20, 0xe7, 0xf1, 0x98, 0xb1, 0x71, 0x82, 0x7e, 0x98, 0x51, 0x3f, 0x4c,
	0x53, 0x26, 0x43, 0x49, 0x59, 0x2a, 0xf2, 0x51, 0xe7, 0xa1, 0x86, 0x63, 0x39, 0xcf, 0x50, 0xf8,
	0xe6, 0x99, 0xe3, 0x2e, 0x03, 0xe7, 0xcd, 0xd5, 0xcd, 0xc5, 0x3f, 0x9e, 0x2b, 0x2a, 0xe4, 0x35,
	0x8a, 0x4c, 0x5d, 0x45, 0xd2, 0x83, 0xa6, 0x50, 0x3c, 0x53, 0x61, 0x5b, 0x7d, 0x6b, 0xd0, 0x19,
	0xee, 0x7b, 0x39, 0x85, 0x77, 0x63, 0x50, 0x72, 0x0e, 0x7b, 0x25, 0x0b, 0xc2, 0xde, 0xea, 0x6f,
	0xab, 0xa9, 0xae, 0x57, 0xb6, 0x5a, 0xa5, 0x77, 0x39, 0xf4, 0xaa, 0xc8, 0x25, 0x8a, 0x88, 0xd3,
	0x11, 0x6e, 0x2c, 0x7a, 0x06, 0xe5, 0xbd, 0x95, 0xa6, 0xb5, 0x4e, 0xf3, 0x87, 0x05, 0xfb, 0x55,
	0x88, 0xec, 0x41, 0x23, 0x7b, 0x47, 0x63, 0x23, 0xd1, 0xd6, 0xa7, 0x34, 0x9c, 0xe4, 0x5c, 0x6d,
	0x72, 0xa8, 0x04, 0xd8, 0x64, 0xc2, 0xd2, 0xc0, 0x80, 0xdb, 0x06, 0xbc, 0x0f, 0x6d, 0x2a, 0xc4,
	0x14, 0xe3, 0x60, 0x34, 0xb7, 0x1b, 0x06, 0x22, 0x00, 0xb3, 0x30, 0xa1, 0x71, 0x70, 0xcb, 0xd9,
	0xc4, 0xde, 0x29, 0xee, 0xe2, 0xc7, 0x8c, 0x72, 0x0c, 0x62, 0x6d, 0xae, 0x59, 0xd0, 0x8b, 0x30,
	0x15, 0xf6, 0xae, 0x8a, 0xc7, 0x9c, 0xb4, 0x57, 0xbb, 0x65, 0x7e, 0xeb, 0xc0, 0xf6, 0x7b, 0x9c,
	0xdb, 0x6d, 0x73, 0x38, 0x80, 0xdd, 0x19, 0x72, 0xa1, 0xde, 0x9b, 0x0d, 0x0a, 0xd8, 0x71, 0x5f,
	0x43, 0xb7, 0x6a, 0xfc, 0x82, 0xa3, 0x7a, 0x5e, 0xe3, 0x87, 0x29, 0x0a, 0xf9, 0xd7, 0xb7, 0x55,
	0x58, 0xd4, 0xc4, 0x5a, 0x39, 0x5c, 0xac, 0x72, 0x0f, 0x5a, 0x8a, 0x3d, 0x47, 0xcc, 0x1e, 0xee,
	0xe9, 0x32, 0xe3, 0x25, 0x26, 0x58, 0x61, 0x34, 0xe6, 0x0c, 0xa3, 0xfb, 0x02, 0x9e, 0xd4, 0xbd,
	0xac, 0xbb, 0xc6, 0xe9, 0xff, 0xdc, 0x59, 0xc2, 0xe6, 0x77, 0x0e, 0xeb, 0xa4, 0x04, 0x46, 0x1c,
	0x65, 0x50, 0x8a, 0xfe, 0x48, 0x15, 0x2a, 0x99, 0x0a, 0x89, 0x7c, 0x29, 0x7b, 0x7d, 0x12, 0x59,
	0x18, 0x61, 0x9e, 0xfd, 0xf0, 0x77, 0x03, 0x0e, 0xaa, 0x5a, 0x82, 0x8c, 0xa1, 0xa1, 0xdb, 0x4b,
	0x0e, 0x8b, 0xc2, 0xbc, 0x65, 0x34, 0x5e, 0x88, 0x3b, 0xc7, 0x2b, 0x0a, 0x52, 0xee, 0xbc, 0xfb,
	0xf4, 0xeb, 0xcf, 0x5f, 0xdf, 0xb7, 0xba, 0xe4, 0x91, 0xf9, 0x92, 0x4a, 0x97, 0xfc, 0xd9, 0x99,
	0x77, 0xee, 0xab, 0x2f, 0x91, 0x7c, 0xb3, 0xa0, 0x55, 0x24, 0x41, 0x4e, 0x56, 0x10, 0x2f, 0xc5,
	0xe5, 0x9c, 0x6e, 0x34, 0xbb, 0x30, 0x32, 0x30, 0x46, 0x5c, 0xd2, 0xaf, 0x35, 0xe2, 0x7f, 0xd2,
	0xe8, 0x17, 0x92, 0x42, 0x33, 0xef, 0x05, 0x19, 0xac, 0x10, 0xa8, 0x54, 0xc7, 0x39, 0xaa, 0x86,
	0xb4, 0xd0, 0x7c, 0x66, 0x34, 0x7b, 0x4e, 0xfd, 0xf2, 0xaf, 0xac, 0x13, 0xc2, 0xa1, 0x99, 0xb7,
	0x66, 0xa5, 0x5e, 0xa5, 0x58, 0x35, 0x7a, 0x9b, 0xef, 0xf8, 0x59, 0x6b, 0xea, 0x36, 0xad, 0xd1,
	0x2c, 0x15, 0xae, 0x46, 0x73, 0x68, 0x34, 0x9f, 0x3b, 0xc7, 0xeb, 0x34, 0xfd, 0xd8, 0xb0, 0xa9,
	0x8d, 0x47, 0x4d, 0xf3, 0x6f, 0xf9, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xcc, 0x1e,
	0x58, 0x89, 0x05, 0x00, 0x00,
}
