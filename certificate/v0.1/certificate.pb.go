// Code generated by protoc-gen-go.
// source: certificate.proto
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

type CertificateListResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificates []*Certificate `protobuf:"bytes,2,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *CertificateListResponse) Reset()                    { *m = CertificateListResponse{} }
func (m *CertificateListResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateListResponse) ProtoMessage()               {}
func (*CertificateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CertificateListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateListResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateDescribeRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDescribeRequest) Reset()                    { *m = CertificateDescribeRequest{} }
func (m *CertificateDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeRequest) ProtoMessage()               {}
func (*CertificateDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type CertificateDescribeResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *Certificate   `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *CertificateDescribeResponse) Reset()                    { *m = CertificateDescribeResponse{} }
func (m *CertificateDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeResponse) ProtoMessage()               {}
func (*CertificateDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *CertificateDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateDescribeResponse) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type CertificateCreateRequest struct {
	AccountPhid      string       `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name             string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName       string       `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	San              []string     `protobuf:"bytes,4,rep,name=san" json:"san,omitempty"`
	Bundle           bool         `protobuf:"varint,5,opt,name=bundle" json:"bundle,omitempty"`
	KeyData          string       `protobuf:"bytes,6,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	SubjectInfo      *SubjectInfo `protobuf:"bytes,7,opt,name=subject_info,json=subjectInfo" json:"subject_info,omitempty"`
	IssuePrivateCert bool         `protobuf:"varint,8,opt,name=issue_private_cert,json=issuePrivateCert" json:"issue_private_cert,omitempty"`
}

func (m *CertificateCreateRequest) Reset()                    { *m = CertificateCreateRequest{} }
func (m *CertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateRequest) ProtoMessage()               {}
func (*CertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *CertificateCreateRequest) GetSubjectInfo() *SubjectInfo {
	if m != nil {
		return m.SubjectInfo
	}
	return nil
}

// A Name contains the SubjectInfo fields.
type SubjectInfo struct {
	C  string `protobuf:"bytes,1,opt,name=C,json=c" json:"C,omitempty"`
	ST string `protobuf:"bytes,2,opt,name=ST,json=sT" json:"ST,omitempty"`
	L  string `protobuf:"bytes,3,opt,name=L,json=l" json:"L,omitempty"`
	O  string `protobuf:"bytes,4,opt,name=O,json=o" json:"O,omitempty"`
	OU string `protobuf:"bytes,5,opt,name=OU,json=oU" json:"OU,omitempty"`
}

func (m *SubjectInfo) Reset()                    { *m = SubjectInfo{} }
func (m *SubjectInfo) String() string            { return proto.CompactTextString(m) }
func (*SubjectInfo) ProtoMessage()               {}
func (*SubjectInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type CertificateCreateResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *Certificate   `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
	Phid        string         `protobuf:"bytes,3,opt,name=phid" json:"phid,omitempty"`
	Cert        []byte         `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	Key         []byte         `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Root        []byte         `protobuf:"bytes,6,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *CertificateCreateResponse) Reset()                    { *m = CertificateCreateResponse{} }
func (m *CertificateCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateResponse) ProtoMessage()               {}
func (*CertificateCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *CertificateCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateCreateResponse) GetCertificate() *Certificate {
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

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type CertificateImportRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CertData string `protobuf:"bytes,2,opt,name=cert_data,json=certData" json:"cert_data,omitempty"`
	KeyData  string `protobuf:"bytes,3,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
}

func (m *CertificateImportRequest) Reset()                    { *m = CertificateImportRequest{} }
func (m *CertificateImportRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateImportRequest) ProtoMessage()               {}
func (*CertificateImportRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type CertificateDeleteRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDeleteRequest) Reset()                    { *m = CertificateDeleteRequest{} }
func (m *CertificateDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeleteRequest) ProtoMessage()               {}
func (*CertificateDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type CertificateRenewRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Uid         string `protobuf:"bytes,3,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateRenewRequest) Reset()                    { *m = CertificateRenewRequest{} }
func (m *CertificateRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRenewRequest) ProtoMessage()               {}
func (*CertificateRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

type CertificateRenewResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Cert   string         `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
}

func (m *CertificateRenewResponse) Reset()                    { *m = CertificateRenewResponse{} }
func (m *CertificateRenewResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateRenewResponse) ProtoMessage()               {}
func (*CertificateRenewResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *CertificateRenewResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type CertificateRevokeRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Uid         string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateRevokeRequest) Reset()                    { *m = CertificateRevokeRequest{} }
func (m *CertificateRevokeRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRevokeRequest) ProtoMessage()               {}
func (*CertificateRevokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

type CertificateDeployRequest struct {
	Uid         string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *CertificateDeployRequest) Reset()                    { *m = CertificateDeployRequest{} }
func (m *CertificateDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeployRequest) ProtoMessage()               {}
func (*CertificateDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

type DNSCheckRequest struct {
	Domain  string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	KeyAuth string `protobuf:"bytes,2,opt,name=key_auth,json=keyAuth" json:"key_auth,omitempty"`
}

func (m *DNSCheckRequest) Reset()                    { *m = DNSCheckRequest{} }
func (m *DNSCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*DNSCheckRequest) ProtoMessage()               {}
func (*DNSCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func init() {
	proto.RegisterType((*CertificateListResponse)(nil), "certificate.CertificateListResponse")
	proto.RegisterType((*CertificateDescribeRequest)(nil), "certificate.CertificateDescribeRequest")
	proto.RegisterType((*CertificateDescribeResponse)(nil), "certificate.CertificateDescribeResponse")
	proto.RegisterType((*CertificateCreateRequest)(nil), "certificate.CertificateCreateRequest")
	proto.RegisterType((*SubjectInfo)(nil), "certificate.SubjectInfo")
	proto.RegisterType((*CertificateCreateResponse)(nil), "certificate.CertificateCreateResponse")
	proto.RegisterType((*Certificate)(nil), "certificate.Certificate")
	proto.RegisterType((*CertificateImportRequest)(nil), "certificate.CertificateImportRequest")
	proto.RegisterType((*CertificateDeleteRequest)(nil), "certificate.CertificateDeleteRequest")
	proto.RegisterType((*CertificateRenewRequest)(nil), "certificate.CertificateRenewRequest")
	proto.RegisterType((*CertificateRenewResponse)(nil), "certificate.CertificateRenewResponse")
	proto.RegisterType((*CertificateRevokeRequest)(nil), "certificate.CertificateRevokeRequest")
	proto.RegisterType((*CertificateDeployRequest)(nil), "certificate.CertificateDeployRequest")
	proto.RegisterType((*DNSCheckRequest)(nil), "certificate.DNSCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Certificates service

type CertificatesClient interface {
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CertificateListResponse, error)
	Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error)
	Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error)
	Import(ctx context.Context, in *CertificateImportRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Renew(ctx context.Context, in *CertificateRenewRequest, opts ...grpc.CallOption) (*CertificateRenewResponse, error)
	Revoke(ctx context.Context, in *CertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type certificatesClient struct {
	cc *grpc.ClientConn
}

func NewCertificatesClient(cc *grpc.ClientConn) CertificatesClient {
	return &certificatesClient{cc}
}

func (c *certificatesClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CertificateListResponse, error) {
	out := new(CertificateListResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error) {
	out := new(CertificateDescribeResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error) {
	out := new(CertificateCreateResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Import(ctx context.Context, in *CertificateImportRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Import", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Renew(ctx context.Context, in *CertificateRenewRequest, opts ...grpc.CallOption) (*CertificateRenewResponse, error) {
	out := new(CertificateRenewResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Revoke(ctx context.Context, in *CertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Revoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Certificates service

type CertificatesServer interface {
	List(context.Context, *dtypes.VoidRequest) (*CertificateListResponse, error)
	Describe(context.Context, *CertificateDescribeRequest) (*CertificateDescribeResponse, error)
	Create(context.Context, *CertificateCreateRequest) (*CertificateCreateResponse, error)
	Import(context.Context, *CertificateImportRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CertificateDeleteRequest) (*dtypes.VoidResponse, error)
	Renew(context.Context, *CertificateRenewRequest) (*CertificateRenewResponse, error)
	Revoke(context.Context, *CertificateRevokeRequest) (*dtypes.VoidResponse, error)
	Deploy(context.Context, *CertificateDeployRequest) (*dtypes.VoidResponse, error)
}

func RegisterCertificatesServer(s *grpc.Server, srv CertificatesServer) {
	s.RegisterService(&_Certificates_serviceDesc, srv)
}

func _Certificates_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).List(ctx, req.(*dtypes.VoidRequest))
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
		FullMethod: "/certificate.Certificates/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Describe(ctx, req.(*CertificateDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Create(ctx, req.(*CertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Import_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Import(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Import",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Import(ctx, req.(*CertificateImportRequest))
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
		FullMethod: "/certificate.Certificates/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Delete(ctx, req.(*CertificateDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Renew(ctx, req.(*CertificateRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Revoke(ctx, req.(*CertificateRevokeRequest))
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
		FullMethod: "/certificate.Certificates/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Deploy(ctx, req.(*CertificateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Certificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.Certificates",
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
			MethodName: "Create",
			Handler:    _Certificates_Create_Handler,
		},
		{
			MethodName: "Import",
			Handler:    _Certificates_Import_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Certificates_Delete_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _Certificates_Renew_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _Certificates_Revoke_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Certificates_Deploy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor2 = []byte{
	// 1029 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0x97, 0x93, 0x5e, 0x2e, 0x99, 0x84, 0x52, 0x16, 0x54, 0x5c, 0xb7, 0x88, 0x62, 0xfa, 0xe7,
	0x74, 0x2a, 0xf1, 0xb5, 0xfc, 0x93, 0xae, 0xbc, 0x40, 0x22, 0xa4, 0x4a, 0xd5, 0x5d, 0xe5, 0x6b,
	0xcb, 0xa3, 0xe5, 0xd8, 0x7b, 0x3d, 0x73, 0x8e, 0xd7, 0x78, 0xed, 0x94, 0x08, 0x15, 0x44, 0x1f,
	0x78, 0x42, 0xf0, 0xc0, 0x37, 0xe0, 0x0d, 0x89, 0xaf, 0xc2, 0x13, 0x5f, 0x81, 0x0f, 0xc2, 0xee,
	0xec, 0xa6, 0x59, 0x9f, 0xce, 0x77, 0x89, 0x90, 0x78, 0x89, 0x76, 0x7f, 0x33, 0x9e, 0xf9, 0xed,
	0x6f, 0x66, 0x67, 0x03, 0x6f, 0x44, 0xb4, 0x28, 0x93, 0xc3, 0x24, 0x0a, 0x4b, 0x3a, 0xcc, 0x0b,
	0x56, 0x32, 0xd2, 0x37, 0x20, 0xe7, 0xda, 0x33, 0xc6, 0x9e, 0xa5, 0xd4, 0x0b, 0xf3, 0xc4, 0x0b,
	0xb3, 0x8c, 0x95, 0x61, 0x99, 0xb0, 0x8c, 0x2b, 0x57, 0xe7, 0xb2, 0x84, 0xe3, 0x72, 0x9e, 0x53,
	0xee, 0xe1, 0xaf, 0xc2, 0xdd, 0x1f, 0xe0, 0xed, 0xd1, 0x32, 0xc8, 0xc3, 0x84, 0x97, 0x3e, 0xe5,
	0xb9, 0xf8, 0x8e, 0x92, 0x5b, 0xd0, 0xe1, 0x22, 0x48, 0xc5, 0x6d, 0xeb, 0xba, 0xb5, 0xd5, 0xbf,
	0x77, 0x71, 0xa8, 0xbe, 0x1f, 0x1e, 0x20, 0xea, 0x6b, 0x2b, 0xf9, 0x0c, 0x06, 0x06, 0x0f, 0x6e,
	0xb7, 0xae, 0xb7, 0x85, 0xb7, 0x3d, 0x34, 0xf9, 0x1a, 0x39, 0xfc, 0x9a, 0xb7, 0x3b, 0x04, 0xc7,
	0x30, 0x8e, 0x29, 0x8f, 0x8a, 0x64, 0x42, 0x7d, 0xfa, 0x4d, 0x45, 0x79, 0x49, 0x2e, 0x41, 0xbb,
	0x4a, 0x62, 0x24, 0xd0, 0xf3, 0xe5, 0xd2, 0xfd, 0xd1, 0x82, 0xab, 0xa7, 0x7e, 0xb0, 0x26, 0xeb,
	0x5d, 0x30, 0xd5, 0x13, 0xa4, 0xad, 0x33, 0x49, 0x9b, 0xce, 0xee, 0xef, 0x2d, 0xb0, 0x0d, 0xe3,
	0xa8, 0xa0, 0xd2, 0x45, 0x53, 0x7e, 0x0f, 0x06, 0x61, 0x14, 0xb1, 0x2a, 0x2b, 0x83, 0xfc, 0xe8,
	0x15, 0xf7, 0xbe, 0xc6, 0x1e, 0x09, 0x88, 0x10, 0xb8, 0x90, 0x85, 0x53, 0x95, 0xb4, 0xe7, 0xe3,
	0x9a, 0xbc, 0x2b, 0xf8, 0xb0, 0xe9, 0x94, 0x65, 0x01, 0x9a, 0xda, 0x68, 0x02, 0x05, 0xed, 0x49,
	0x07, 0x21, 0x05, 0x0f, 0x33, 0xfb, 0x82, 0x50, 0x57, 0x48, 0x21, 0x96, 0xe4, 0x32, 0x74, 0x26,
	0x55, 0x16, 0xa7, 0xd4, 0xde, 0x10, 0xde, 0x5d, 0x5f, 0xef, 0xc8, 0x15, 0xe8, 0x1e, 0xd3, 0x79,
	0x10, 0x87, 0x65, 0x68, 0x77, 0x30, 0xce, 0xa6, 0xd8, 0x8f, 0xc5, 0x96, 0xdc, 0x87, 0x01, 0xaf,
	0x26, 0x5f, 0xd3, 0xa8, 0x0c, 0x92, 0xec, 0x90, 0xd9, 0x9b, 0xa7, 0x1c, 0xfb, 0x40, 0x39, 0x3c,
	0x10, 0x76, 0xbf, 0xcf, 0x97, 0x1b, 0x72, 0x07, 0x48, 0xc2, 0x79, 0x45, 0x83, 0xbc, 0x48, 0x66,
	0xc2, 0x33, 0x90, 0x5f, 0xd9, 0x5d, 0xcc, 0x7d, 0x09, 0x2d, 0x8f, 0x94, 0x41, 0x6a, 0xe3, 0x7e,
	0x05, 0x7d, 0x23, 0x12, 0x19, 0x80, 0x35, 0xd2, 0x5a, 0x58, 0x11, 0xb9, 0x08, 0xad, 0x83, 0xc7,
	0xfa, 0xfc, 0x2d, 0xfe, 0x58, 0x5a, 0x1f, 0xea, 0x33, 0x5b, 0xa9, 0xdc, 0xed, 0x8b, 0x83, 0xe2,
	0x8e, 0x49, 0xdf, 0xfd, 0x27, 0x78, 0x44, 0xe1, 0xcb, 0x9e, 0xb8, 0x7f, 0x59, 0x70, 0xe5, 0x14,
	0xf5, 0xff, 0xbf, 0xfa, 0xcb, 0xfa, 0x61, 0x69, 0x15, 0x61, 0x5c, 0x4b, 0x0c, 0xe5, 0x90, 0xb4,
	0x07, 0x3e, 0xae, 0x65, 0xc9, 0x84, 0xf0, 0x48, 0x7d, 0xe0, 0xcb, 0xa5, 0xf4, 0x2a, 0x18, 0x2b,
	0xb1, 0x2c, 0xc2, 0x4b, 0xae, 0xdd, 0x3f, 0x5b, 0xd0, 0x1f, 0x9d, 0x12, 0xdd, 0xaa, 0x47, 0x5f,
	0xbf, 0x63, 0xae, 0x42, 0x0f, 0xab, 0x12, 0x07, 0x93, 0xb9, 0x96, 0xb3, 0xab, 0x80, 0x2f, 0xe6,
	0xe4, 0x1d, 0x80, 0x59, 0x98, 0x26, 0x71, 0x70, 0x58, 0xb0, 0xa9, 0x56, 0xb7, 0x87, 0xc8, 0x97,
	0x02, 0x90, 0xc1, 0xe9, 0xb7, 0x79, 0x52, 0x50, 0xd9, 0x46, 0x54, 0xb7, 0x11, 0x28, 0x68, 0xac,
	0x59, 0x8a, 0x1e, 0xe4, 0xa2, 0x83, 0x64, 0x3f, 0xe2, 0xfa, 0x95, 0x06, 0x5d, 0xc5, 0xd2, 0xd4,
	0xa0, 0xa7, 0x6e, 0xb0, 0xd4, 0xc0, 0x86, 0xcd, 0x19, 0x2d, 0xb8, 0x18, 0x4e, 0x36, 0x08, 0x74,
	0xc3, 0x5f, 0x6c, 0xc9, 0xfb, 0xf0, 0x1a, 0xa7, 0x45, 0x12, 0xa6, 0x41, 0x56, 0x4d, 0x27, 0xb4,
	0xb0, 0xfb, 0xf8, 0xd5, 0x40, 0x81, 0x7b, 0x88, 0xb9, 0x87, 0xb5, 0xbb, 0xf7, 0x60, 0x9a, 0xb3,
	0xa2, 0x5c, 0xdc, 0xbd, 0x85, 0x4c, 0x96, 0x21, 0x93, 0x50, 0x41, 0x12, 0x51, 0xd7, 0x41, 0xe9,
	0xd7, 0x95, 0x00, 0xde, 0x07, 0xf3, 0xaa, 0xb4, 0x6b, 0x57, 0xc5, 0xbd, 0x53, 0xcb, 0x33, 0xa6,
	0x29, 0x2d, 0xcf, 0x18, 0x4b, 0x93, 0xda, 0x1c, 0xf5, 0x69, 0x46, 0x9f, 0xff, 0xc7, 0x81, 0xa0,
	0x73, 0xb4, 0x97, 0x39, 0x9e, 0xd6, 0x18, 0xe9, 0x1c, 0x6b, 0xb6, 0xfd, 0xa2, 0x44, 0xad, 0x65,
	0x89, 0xdc, 0xfd, 0x13, 0x71, 0x67, 0xec, 0x78, 0x9d, 0x69, 0xa6, 0x89, 0xb6, 0x96, 0x44, 0x7f,
	0xb1, 0x4e, 0x68, 0x97, 0xa7, 0x6c, 0xde, 0xa8, 0x9d, 0xec, 0x35, 0x4e, 0xa3, 0x82, 0x96, 0x81,
	0x21, 0x02, 0x28, 0x08, 0x1b, 0x59, 0x90, 0x88, 0xd2, 0x8a, 0x97, 0xb4, 0x30, 0x5b, 0xbd, 0xaf,
	0x31, 0x74, 0xb9, 0x06, 0x3d, 0x69, 0xe2, 0x79, 0x18, 0x51, 0xdd, 0xeb, 0x4b, 0xc0, 0x1d, 0xc3,
	0xeb, 0xe3, 0xbd, 0x83, 0xd1, 0x11, 0x8d, 0x8e, 0x17, 0x34, 0xc4, 0xf0, 0x8c, 0xd9, 0x34, 0x4c,
	0x32, 0xcd, 0x44, 0xef, 0x16, 0x1d, 0x11, 0x56, 0xe5, 0x91, 0x66, 0x22, 0x3b, 0xe2, 0x73, 0xb1,
	0xbd, 0xf7, 0x47, 0x17, 0x06, 0xc6, 0xb1, 0x38, 0x99, 0xc2, 0x05, 0xf9, 0x62, 0x92, 0x37, 0x17,
	0x62, 0x3f, 0x65, 0x49, 0xac, 0x13, 0x38, 0x37, 0x9a, 0x66, 0x89, 0xf9, 0xc8, 0xba, 0xdb, 0x2f,
	0xff, 0xfe, 0xe7, 0xb7, 0xd6, 0x0d, 0xe2, 0x8a, 0x77, 0x3b, 0xe7, 0x11, 0x8b, 0xd5, 0x03, 0xce,
	0x79, 0xea, 0x19, 0x9f, 0x7b, 0xb3, 0x9d, 0xe1, 0x5d, 0xf2, 0xab, 0x05, 0xdd, 0xc5, 0x7b, 0x47,
	0x6e, 0x37, 0x85, 0x3f, 0xf1, 0x84, 0x3a, 0x5b, 0xe7, 0x3b, 0x6a, 0x2e, 0x3b, 0xc8, 0x65, 0x9b,
	0x6c, 0x9d, 0xcf, 0xc5, 0xfb, 0x4e, 0x14, 0xee, 0x05, 0xf9, 0xc9, 0x82, 0x8e, 0x9a, 0xbf, 0xe4,
	0x66, 0x53, 0x9a, 0xda, 0xeb, 0xe8, 0xdc, 0x3a, 0xcf, 0x4d, 0x73, 0xf9, 0x00, 0xb9, 0xdc, 0x76,
	0x56, 0xd0, 0x65, 0xd7, 0xda, 0x26, 0x2f, 0x05, 0x11, 0x35, 0x0a, 0x9a, 0x89, 0xd4, 0x46, 0x85,
	0xf3, 0x56, 0xbd, 0x66, 0x3a, 0xed, 0x7d, 0x4c, 0xfb, 0xb1, 0xb3, 0xb3, 0x8a, 0x04, 0xb2, 0xbf,
	0x5e, 0x78, 0x09, 0x86, 0x95, 0x24, 0x9e, 0x43, 0x47, 0x8d, 0x89, 0x66, 0x0e, 0xb5, 0x31, 0xd2,
	0xc0, 0x41, 0x97, 0x61, 0x7b, 0xf5, 0x32, 0xfc, 0x6c, 0xc1, 0x06, 0x8e, 0x03, 0xd2, 0xd8, 0x74,
	0xe6, 0x44, 0x72, 0x6e, 0x9e, 0xe3, 0xa5, 0x89, 0x7c, 0x82, 0x44, 0x76, 0xdc, 0xe1, 0xaa, 0x44,
	0xbc, 0x02, 0x49, 0x7c, 0x0f, 0x1d, 0x35, 0x44, 0xc8, 0x19, 0x89, 0x8c, 0x21, 0xd3, 0xa0, 0xc3,
	0xa7, 0x98, 0xfe, 0xae, 0xe3, 0xad, 0x91, 0x1e, 0xb3, 0x8a, 0xbf, 0x88, 0x1d, 0x35, 0x73, 0xce,
	0x2a, 0x84, 0x31, 0x93, 0x1a, 0x08, 0xec, 0x22, 0x81, 0x8f, 0xdc, 0xd5, 0x09, 0xc4, 0x18, 0x55,
	0xf4, 0xc2, 0xa4, 0x83, 0x7f, 0xaf, 0x3f, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xeb, 0x6d,
	0x11, 0xb6, 0x0b, 0x00, 0x00,
}
