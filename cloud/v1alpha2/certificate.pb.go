// Code generated by protoc-gen-go. DO NOT EDIT.
// source: certificate.proto

/*
Package v1alpha2 is a generated protocol buffer package.

It is generated from these files:
	certificate.proto
	cloud.proto
	metadata.proto

It has these top-level messages:
	CertificateListRequest
	CertificateListResponse
	CertificateDescribeRequest
	CertificateDescribeResponse
	Certificate
	CertificateLoadRequest
	CertificateDeleteRequest
	CertificateDeployRequest
	CredentialListRequest
	CredentialListResponse
	CredentialDescribeRequest
	CredentialDescribeResponse
	CredentialCreateRequest
	CredentialCreateResponse
	CredentialUpdateRequest
	CredentialUpdateResponse
	CredentialDeleteRequest
	CredentialDeleteResponse
	ClusterListRequest
	ClusterListResponse
	ClusterDescribeRequest
	ClusterDescribeResponse
	ClusterCreateRequest
	ClusterCreateResponse
	ClusterUpdateRequest
	ClusterUpdateResponse
	ClusterDeleteRequest
	ClusterDeleteResponse
	ClusterApplyRequest
	ClusterMetadataRequest
	ClusterClientConfigRequest
	ClusterClientConfigResponse
	ClusterMetadataResponse
	NodeGroupListRequest
	NodeGroupListResponse
	NodeGroupDescribeRequest
	NodeGroupDescribeResponse
	NodeGroupCreateRequest
	NodeGroupCreateResponse
	NodeGroupUpdateRequest
	NodeGroupUpdateResponse
	NodeGroupDeleteRequest
	NodeGroupDeleteResponse
	SSHConfigGetRequest
	SSHConfigGetResponse
	RegionListRequest
	RegionListResponse
	ZoneListRequest
	ZoneListResponse
	BucketListRequest
	BucketListResponse
*/
package v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "appscode.com/api/dtypes"

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

type CertificateListRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
}

func (m *CertificateListRequest) Reset()                    { *m = CertificateListRequest{} }
func (m *CertificateListRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateListRequest) ProtoMessage()               {}
func (*CertificateListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CertificateListRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type CertificateListResponse struct {
	Certificates []*Certificate `protobuf:"bytes,1,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *CertificateListResponse) Reset()                    { *m = CertificateListResponse{} }
func (m *CertificateListResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateListResponse) ProtoMessage()               {}
func (*CertificateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateListResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateDescribeRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDescribeRequest) Reset()                    { *m = CertificateDescribeRequest{} }
func (m *CertificateDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeRequest) ProtoMessage()               {}
func (*CertificateDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CertificateDescribeRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CertificateDescribeRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type CertificateDescribeResponse struct {
	Certificate *Certificate `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *CertificateDescribeResponse) Reset()                    { *m = CertificateDescribeResponse{} }
func (m *CertificateDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeResponse) ProtoMessage()               {}
func (*CertificateDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CertificateDescribeResponse) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type Certificate struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	IssuedBy   string `protobuf:"bytes,4,opt,name=issued_by,json=issuedBy" json:"issued_by,omitempty"`
	ValidFrom  int64  `protobuf:"varint,5,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	ExpireDate int64  `protobuf:"varint,6,opt,name=expire_date,json=expireDate" json:"expire_date,omitempty"`
	// those feilds will not included into list response.
	// only describe response will include the underlying
	// feilds.
	Sans         []string `protobuf:"bytes,7,rep,name=sans" json:"sans,omitempty"`
	Cert         string   `protobuf:"bytes,8,opt,name=cert" json:"cert,omitempty"`
	Key          string   `protobuf:"bytes,9,opt,name=key" json:"key,omitempty"`
	Version      int32    `protobuf:"varint,10,opt,name=version" json:"version,omitempty"`
	SerialNumber string   `protobuf:"bytes,11,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Certificate) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *Certificate) GetIssuedBy() string {
	if m != nil {
		return m.IssuedBy
	}
	return ""
}

func (m *Certificate) GetValidFrom() int64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *Certificate) GetExpireDate() int64 {
	if m != nil {
		return m.ExpireDate
	}
	return 0
}

func (m *Certificate) GetSans() []string {
	if m != nil {
		return m.Sans
	}
	return nil
}

func (m *Certificate) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *Certificate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Certificate) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Certificate) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

type CertificateLoadRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CertData string `protobuf:"bytes,2,opt,name=cert_data,json=certData" json:"cert_data,omitempty"`
	KeyData  string `protobuf:"bytes,3,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	Owner    string `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
}

func (m *CertificateLoadRequest) Reset()                    { *m = CertificateLoadRequest{} }
func (m *CertificateLoadRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateLoadRequest) ProtoMessage()               {}
func (*CertificateLoadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CertificateLoadRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateLoadRequest) GetCertData() string {
	if m != nil {
		return m.CertData
	}
	return ""
}

func (m *CertificateLoadRequest) GetKeyData() string {
	if m != nil {
		return m.KeyData
	}
	return ""
}

func (m *CertificateLoadRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type CertificateDeleteRequest struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDeleteRequest) Reset()                    { *m = CertificateDeleteRequest{} }
func (m *CertificateDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeleteRequest) ProtoMessage()               {}
func (*CertificateDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CertificateDeleteRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CertificateDeleteRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type CertificateDeployRequest struct {
	Uid         string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	Owner       string `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
}

func (m *CertificateDeployRequest) Reset()                    { *m = CertificateDeployRequest{} }
func (m *CertificateDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeployRequest) ProtoMessage()               {}
func (*CertificateDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CertificateDeployRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CertificateDeployRequest) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *CertificateDeployRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *CertificateDeployRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CertificateDeployRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*CertificateListRequest)(nil), "appscode.cloud.v1alpha2.CertificateListRequest")
	proto.RegisterType((*CertificateListResponse)(nil), "appscode.cloud.v1alpha2.CertificateListResponse")
	proto.RegisterType((*CertificateDescribeRequest)(nil), "appscode.cloud.v1alpha2.CertificateDescribeRequest")
	proto.RegisterType((*CertificateDescribeResponse)(nil), "appscode.cloud.v1alpha2.CertificateDescribeResponse")
	proto.RegisterType((*Certificate)(nil), "appscode.cloud.v1alpha2.Certificate")
	proto.RegisterType((*CertificateLoadRequest)(nil), "appscode.cloud.v1alpha2.CertificateLoadRequest")
	proto.RegisterType((*CertificateDeleteRequest)(nil), "appscode.cloud.v1alpha2.CertificateDeleteRequest")
	proto.RegisterType((*CertificateDeployRequest)(nil), "appscode.cloud.v1alpha2.CertificateDeployRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Certificates service

type CertificatesClient interface {
	List(ctx context.Context, in *CertificateListRequest, opts ...grpc.CallOption) (*CertificateListResponse, error)
	Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error)
	Load(ctx context.Context, in *CertificateLoadRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type certificatesClient struct {
	cc *grpc.ClientConn
}

func NewCertificatesClient(cc *grpc.ClientConn) CertificatesClient {
	return &certificatesClient{cc}
}

func (c *certificatesClient) List(ctx context.Context, in *CertificateListRequest, opts ...grpc.CallOption) (*CertificateListResponse, error) {
	out := new(CertificateListResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha2.Certificates/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error) {
	out := new(CertificateDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha2.Certificates/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Load(ctx context.Context, in *CertificateLoadRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha2.Certificates/Load", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha2.Certificates/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.cloud.v1alpha2.Certificates/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Certificates service

type CertificatesServer interface {
	List(context.Context, *CertificateListRequest) (*CertificateListResponse, error)
	Describe(context.Context, *CertificateDescribeRequest) (*CertificateDescribeResponse, error)
	Load(context.Context, *CertificateLoadRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *CertificateDeleteRequest) (*appscode_dtypes.VoidResponse, error)
	Deploy(context.Context, *CertificateDeployRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterCertificatesServer(s *grpc.Server, srv CertificatesServer) {
	s.RegisterService(&_Certificates_serviceDesc, srv)
}

func _Certificates_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha2.Certificates/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).List(ctx, req.(*CertificateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha2.Certificates/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Describe(ctx, req.(*CertificateDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha2.Certificates/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Load(ctx, req.(*CertificateLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha2.Certificates/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Delete(ctx, req.(*CertificateDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.cloud.v1alpha2.Certificates/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Deploy(ctx, req.(*CertificateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Certificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.cloud.v1alpha2.Certificates",
	HandlerType: (*CertificatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Certificates_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Certificates_Describe_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _Certificates_Load_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Certificates_Delete_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Certificates_Deploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate.proto",
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xe5, 0x7c, 0x34, 0xc9, 0x49, 0xae, 0xd4, 0x3b, 0xba, 0xba, 0x35, 0x6e, 0x4b, 0x43,
	0xca, 0x22, 0xea, 0x22, 0xa6, 0x2d, 0x12, 0x02, 0x09, 0x55, 0x0a, 0xa1, 0x42, 0xa2, 0x54, 0x25,
	0x0b, 0x16, 0x48, 0x10, 0x4d, 0xec, 0x69, 0x3b, 0xc4, 0xf6, 0x18, 0x8f, 0x53, 0x30, 0x55, 0x37,
	0x7d, 0x05, 0x5e, 0x01, 0xc4, 0x86, 0x05, 0x1b, 0x90, 0xd8, 0xf2, 0x0a, 0xbc, 0x02, 0x0f, 0x82,
	0x66, 0xc6, 0x6e, 0x26, 0x85, 0x52, 0xb7, 0x6c, 0xa2, 0xc9, 0xf9, 0xf2, 0xef, 0xfc, 0x7d, 0xce,
	0x18, 0xfe, 0x75, 0x48, 0x14, 0xd3, 0x5d, 0xea, 0xe0, 0x98, 0x74, 0xc2, 0x88, 0xc5, 0x0c, 0xcd,
	0xe1, 0x30, 0xe4, 0x0e, 0x73, 0x49, 0xc7, 0xf1, 0xd8, 0xd8, 0xed, 0x1c, 0xac, 0x62, 0x2f, 0xdc,
	0xc7, 0x6b, 0xd6, 0xc2, 0x1e, 0x63, 0x7b, 0x1e, 0xb1, 0x71, 0x48, 0x6d, 0x1c, 0x04, 0x2c, 0xc6,
	0x31, 0x65, 0x01, 0x57, 0x69, 0xd6, 0xd5, 0x2c, 0xed, 0x0c, 0xff, 0xf2, 0xa4, 0x2c, 0xf3, 0x65,
	0x8c, 0x1b, 0x27, 0x21, 0xe1, 0xb6, 0xfc, 0x55, 0x41, 0xad, 0x0e, 0xfc, 0x7f, 0x6f, 0x02, 0xb4,
	0x45, 0x79, 0xdc, 0x27, 0x2f, 0xc7, 0x84, 0xc7, 0xe8, 0x3f, 0x28, 0xb3, 0x57, 0x01, 0x89, 0x4c,
	0xa3, 0x69, 0xb4, 0x6b, 0x7d, 0xf5, 0xa7, 0xe5, 0xc0, 0xdc, 0x2f, 0xf1, 0x3c, 0x64, 0x01, 0x27,
	0xe8, 0x01, 0x34, 0xb4, 0xde, 0xb8, 0x69, 0x34, 0x8b, 0xed, 0xfa, 0xda, 0xf5, 0xce, 0x19, 0xdd,
	0x75, 0xb4, 0x3a, 0xfd, 0xa9, 0xcc, 0x56, 0x0f, 0x2c, 0xcd, 0xd9, 0x23, 0xdc, 0x89, 0xe8, 0x90,
	0xfc, 0x11, 0x0c, 0xcd, 0x42, 0x71, 0x4c, 0x5d, 0xb3, 0x20, 0x6d, 0xe2, 0xd8, 0x22, 0x30, 0xff,
	0xdb, 0x2a, 0x29, 0xee, 0x26, 0xd4, 0xb5, 0x87, 0xca, 0x62, 0x79, 0x69, 0xf5, 0xc4, 0xd6, 0xc7,
	0x02, 0xd4, 0x35, 0x27, 0x42, 0x50, 0x0a, 0xf7, 0xa9, 0x9b, 0xd2, 0xc9, 0xb3, 0xb0, 0x05, 0xd8,
	0x27, 0x29, 0x9d, 0x3c, 0xa3, 0x25, 0xa8, 0x3b, 0xcc, 0xf7, 0x59, 0x30, 0x90, 0xae, 0xa2, 0x74,
	0x81, 0x32, 0x6d, 0x8b, 0x80, 0x79, 0xa8, 0x51, 0xce, 0xc7, 0xc4, 0x1d, 0x0c, 0x13, 0xb3, 0x24,
	0xdd, 0x55, 0x65, 0xe8, 0x26, 0x68, 0x11, 0xe0, 0x00, 0x7b, 0xd4, 0x1d, 0xec, 0x46, 0xcc, 0x37,
	0xcb, 0x4d, 0xa3, 0x5d, 0xec, 0xd7, 0xa4, 0x65, 0x33, 0x62, 0xbe, 0x28, 0x4e, 0x5e, 0x87, 0x34,
	0x22, 0x03, 0x57, 0x34, 0x37, 0x23, 0xfd, 0xa0, 0x4c, 0xbd, 0x94, 0x92, 0xe3, 0x80, 0x9b, 0x95,
	0x66, 0x51, 0x10, 0x89, 0xb3, 0xb0, 0x89, 0xc6, 0xcc, 0xaa, 0xa2, 0x14, 0x67, 0x21, 0xeb, 0x88,
	0x24, 0x66, 0x4d, 0xc9, 0x3a, 0x22, 0x09, 0x32, 0xa1, 0x72, 0x40, 0x22, 0x4e, 0x59, 0x60, 0x42,
	0xd3, 0x68, 0x97, 0xfb, 0xd9, 0x5f, 0xb4, 0x0c, 0xff, 0x70, 0x12, 0x51, 0xec, 0x0d, 0x82, 0xb1,
	0x3f, 0x24, 0x91, 0x59, 0x97, 0x59, 0x0d, 0x65, 0xdc, 0x96, 0xb6, 0xd6, 0x9b, 0xe9, 0x81, 0x63,
	0xd8, 0xcd, 0xde, 0x6b, 0x26, 0x92, 0xa1, 0x89, 0x34, 0x0f, 0x35, 0x81, 0x21, 0xba, 0xc0, 0xa9,
	0x7a, 0x55, 0x61, 0xe8, 0xe1, 0x18, 0xa3, 0x2b, 0x50, 0x1d, 0x91, 0x44, 0xf9, 0x94, 0x7c, 0x95,
	0x11, 0x49, 0xa4, 0xeb, 0x64, 0x46, 0x4a, 0xfa, 0xf0, 0x76, 0xc1, 0x9c, 0x9a, 0x08, 0x8f, 0xc4,
	0x17, 0x9e, 0xaa, 0x77, 0xc6, 0xa9, 0x22, 0xa1, 0xc7, 0x92, 0xac, 0x48, 0x1a, 0x6e, 0x9c, 0x84,
	0x8b, 0x17, 0xc1, 0x89, 0x13, 0x91, 0x78, 0xa0, 0x0d, 0x00, 0x28, 0x93, 0x7c, 0xcb, 0xd7, 0xa0,
	0xe1, 0x78, 0x63, 0x1e, 0x93, 0x48, 0x9f, 0x83, 0x7a, 0x6a, 0x93, 0x21, 0x0b, 0x50, 0x13, 0x2e,
	0x1e, 0x62, 0x87, 0xa4, 0x0d, 0x4d, 0x0c, 0x13, 0xf0, 0xb2, 0x06, 0xbe, 0xf6, 0xbe, 0x02, 0x0d,
	0x0d, 0x93, 0xa3, 0x4f, 0x06, 0x94, 0xc4, 0xba, 0x22, 0x3b, 0xcf, 0x88, 0x6b, 0x17, 0x81, 0x75,
	0x23, 0x7f, 0x82, 0x5a, 0xad, 0xd6, 0xfd, 0xe3, 0xcf, 0x66, 0xa1, 0x6a, 0x1c, 0x7f, 0xff, 0xf1,
	0xb6, 0x70, 0x1b, 0xdd, 0xb2, 0x07, 0x53, 0x37, 0x95, 0x2c, 0x62, 0x67, 0x45, 0xec, 0x43, 0x09,
	0x7c, 0x64, 0xeb, 0x37, 0x80, 0xfd, 0x82, 0xb3, 0x00, 0x7d, 0x33, 0xa0, 0x9a, 0xad, 0x2d, 0x5a,
	0xcf, 0x43, 0x71, 0xea, 0xaa, 0xb0, 0x6e, 0x5e, 0x2c, 0x29, 0xc5, 0x7f, 0xa8, 0xe1, 0x6f, 0xa0,
	0xbb, 0x97, 0xc0, 0x3f, 0x1c, 0x53, 0xf7, 0x48, 0x35, 0xf1, 0x45, 0xe8, 0xce, 0xb0, 0x9b, 0x53,
	0xf7, 0xc9, 0x3e, 0x58, 0x8b, 0x93, 0x04, 0x75, 0x71, 0x77, 0x9e, 0x30, 0xea, 0x9e, 0x50, 0x3e,
	0xd7, 0x28, 0xfb, 0xd6, 0xa3, 0xcb, 0x50, 0x8a, 0x21, 0x3a, 0xb2, 0xb1, 0x23, 0xbf, 0x1a, 0x36,
	0x1b, 0xc6, 0x98, 0x06, 0x92, 0xfa, 0x8e, 0xb1, 0x82, 0x3e, 0x18, 0x30, 0xa3, 0x56, 0x04, 0xad,
	0xe6, 0x93, 0x51, 0x5b, 0xa7, 0xf3, 0xe0, 0xa7, 0x24, 0x5e, 0xf9, 0x4b, 0x89, 0xbf, 0x4a, 0x52,
	0xb1, 0x87, 0x79, 0x49, 0xb5, 0x9d, 0x3d, 0x8f, 0xf4, 0x99, 0x46, 0xfa, 0xd8, 0xda, 0xba, 0x34,
	0x69, 0xa6, 0xb2, 0x2b, 0x1f, 0x9e, 0xa9, 0xdc, 0xdd, 0x80, 0x25, 0x87, 0xf9, 0x13, 0x04, 0x1c,
	0xd2, 0x53, 0xe4, 0xdd, 0x59, 0x0d, 0x7d, 0x47, 0x7c, 0xb4, 0x77, 0x8c, 0xa7, 0xd5, 0xcc, 0x3b,
	0x9c, 0x91, 0xdf, 0xf1, 0xf5, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0xe3, 0xbd, 0xb4, 0x58,
	0x08, 0x00, 0x00,
}
