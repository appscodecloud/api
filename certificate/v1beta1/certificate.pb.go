// Code generated by protoc-gen-go.
// source: certificate.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CertificateListRequest struct {
}

func (m *CertificateListRequest) Reset()                    { *m = CertificateListRequest{} }
func (m *CertificateListRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateListRequest) ProtoMessage()               {}
func (*CertificateListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type CertificateListResponse struct {
	Status       *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificates []*Certificate          `protobuf:"bytes,2,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *CertificateListResponse) Reset()                    { *m = CertificateListResponse{} }
func (m *CertificateListResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateListResponse) ProtoMessage()               {}
func (*CertificateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CertificateListResponse) GetStatus() *appscode_dtypes.Status {
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
func (*CertificateDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type CertificateDescribeResponse struct {
	Status      *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *Certificate            `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *CertificateDescribeResponse) Reset()                    { *m = CertificateDescribeResponse{} }
func (m *CertificateDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeResponse) ProtoMessage()               {}
func (*CertificateDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *CertificateDescribeResponse) GetStatus() *appscode_dtypes.Status {
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
	Sans             []string     `protobuf:"bytes,4,rep,name=sans" json:"sans,omitempty"`
	Bundle           bool         `protobuf:"varint,5,opt,name=bundle" json:"bundle,omitempty"`
	KeyData          string       `protobuf:"bytes,6,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	SubjectInfo      *SubjectInfo `protobuf:"bytes,7,opt,name=subject_info,json=subjectInfo" json:"subject_info,omitempty"`
	IssuePrivateCert bool         `protobuf:"varint,8,opt,name=issue_private_cert,json=issuePrivateCert" json:"issue_private_cert,omitempty"`
}

func (m *CertificateCreateRequest) Reset()                    { *m = CertificateCreateRequest{} }
func (m *CertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateRequest) ProtoMessage()               {}
func (*CertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CertificateCreateRequest) GetSubjectInfo() *SubjectInfo {
	if m != nil {
		return m.SubjectInfo
	}
	return nil
}

// A Name contains the SubjectInfo fields.
type SubjectInfo struct {
	C  string `protobuf:"bytes,1,opt,name=C" json:"C,omitempty"`
	ST string `protobuf:"bytes,2,opt,name=ST" json:"ST,omitempty"`
	L  string `protobuf:"bytes,3,opt,name=L" json:"L,omitempty"`
	O  string `protobuf:"bytes,4,opt,name=O" json:"O,omitempty"`
	OU string `protobuf:"bytes,5,opt,name=OU" json:"OU,omitempty"`
}

func (m *SubjectInfo) Reset()                    { *m = SubjectInfo{} }
func (m *SubjectInfo) String() string            { return proto.CompactTextString(m) }
func (*SubjectInfo) ProtoMessage()               {}
func (*SubjectInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type CertificateCreateResponse struct {
	Status      *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *Certificate            `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
	Phid        string                  `protobuf:"bytes,3,opt,name=phid" json:"phid,omitempty"`
	Cert        []byte                  `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	Key         []byte                  `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Root        []byte                  `protobuf:"bytes,6,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *CertificateCreateResponse) Reset()                    { *m = CertificateCreateResponse{} }
func (m *CertificateCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateResponse) ProtoMessage()               {}
func (*CertificateCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *CertificateCreateResponse) GetStatus() *appscode_dtypes.Status {
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
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type CertificateObtainRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CertData string `protobuf:"bytes,2,opt,name=cert_data,json=certData" json:"cert_data,omitempty"`
	KeyData  string `protobuf:"bytes,3,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
}

func (m *CertificateObtainRequest) Reset()                    { *m = CertificateObtainRequest{} }
func (m *CertificateObtainRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateObtainRequest) ProtoMessage()               {}
func (*CertificateObtainRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type CertificateDeleteRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDeleteRequest) Reset()                    { *m = CertificateDeleteRequest{} }
func (m *CertificateDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeleteRequest) ProtoMessage()               {}
func (*CertificateDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

type CertificateRenewRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Uid         string `protobuf:"bytes,3,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateRenewRequest) Reset()                    { *m = CertificateRenewRequest{} }
func (m *CertificateRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRenewRequest) ProtoMessage()               {}
func (*CertificateRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

type CertificateRenewResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Cert   string                  `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
}

func (m *CertificateRenewResponse) Reset()                    { *m = CertificateRenewResponse{} }
func (m *CertificateRenewResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateRenewResponse) ProtoMessage()               {}
func (*CertificateRenewResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *CertificateRenewResponse) GetStatus() *appscode_dtypes.Status {
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
func (*CertificateRevokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

type CertificateDeployRequest struct {
	Uid         string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *CertificateDeployRequest) Reset()                    { *m = CertificateDeployRequest{} }
func (m *CertificateDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeployRequest) ProtoMessage()               {}
func (*CertificateDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

type DNSCheckRequest struct {
	Domain  string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	KeyAuth string `protobuf:"bytes,2,opt,name=key_auth,json=keyAuth" json:"key_auth,omitempty"`
}

func (m *DNSCheckRequest) Reset()                    { *m = DNSCheckRequest{} }
func (m *DNSCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*DNSCheckRequest) ProtoMessage()               {}
func (*DNSCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func init() {
	proto.RegisterType((*CertificateListRequest)(nil), "appscode.certificate.v1beta1.CertificateListRequest")
	proto.RegisterType((*CertificateListResponse)(nil), "appscode.certificate.v1beta1.CertificateListResponse")
	proto.RegisterType((*CertificateDescribeRequest)(nil), "appscode.certificate.v1beta1.CertificateDescribeRequest")
	proto.RegisterType((*CertificateDescribeResponse)(nil), "appscode.certificate.v1beta1.CertificateDescribeResponse")
	proto.RegisterType((*CertificateCreateRequest)(nil), "appscode.certificate.v1beta1.CertificateCreateRequest")
	proto.RegisterType((*SubjectInfo)(nil), "appscode.certificate.v1beta1.SubjectInfo")
	proto.RegisterType((*CertificateCreateResponse)(nil), "appscode.certificate.v1beta1.CertificateCreateResponse")
	proto.RegisterType((*Certificate)(nil), "appscode.certificate.v1beta1.Certificate")
	proto.RegisterType((*CertificateObtainRequest)(nil), "appscode.certificate.v1beta1.CertificateObtainRequest")
	proto.RegisterType((*CertificateDeleteRequest)(nil), "appscode.certificate.v1beta1.CertificateDeleteRequest")
	proto.RegisterType((*CertificateRenewRequest)(nil), "appscode.certificate.v1beta1.CertificateRenewRequest")
	proto.RegisterType((*CertificateRenewResponse)(nil), "appscode.certificate.v1beta1.CertificateRenewResponse")
	proto.RegisterType((*CertificateRevokeRequest)(nil), "appscode.certificate.v1beta1.CertificateRevokeRequest")
	proto.RegisterType((*CertificateDeployRequest)(nil), "appscode.certificate.v1beta1.CertificateDeployRequest")
	proto.RegisterType((*DNSCheckRequest)(nil), "appscode.certificate.v1beta1.DNSCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Certificates service

type CertificatesClient interface {
	List(ctx context.Context, in *CertificateListRequest, opts ...grpc.CallOption) (*CertificateListResponse, error)
	Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error)
	Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error)
	Obtain(ctx context.Context, in *CertificateObtainRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Renew(ctx context.Context, in *CertificateRenewRequest, opts ...grpc.CallOption) (*CertificateRenewResponse, error)
	Revoke(ctx context.Context, in *CertificateRevokeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
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
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error) {
	out := new(CertificateDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error) {
	out := new(CertificateCreateResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Obtain(ctx context.Context, in *CertificateObtainRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Obtain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Renew(ctx context.Context, in *CertificateRenewRequest, opts ...grpc.CallOption) (*CertificateRenewResponse, error) {
	out := new(CertificateRenewResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Revoke(ctx context.Context, in *CertificateRevokeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Revoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.Certificates/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Certificates service

type CertificatesServer interface {
	List(context.Context, *CertificateListRequest) (*CertificateListResponse, error)
	Describe(context.Context, *CertificateDescribeRequest) (*CertificateDescribeResponse, error)
	Create(context.Context, *CertificateCreateRequest) (*CertificateCreateResponse, error)
	Obtain(context.Context, *CertificateObtainRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *CertificateDeleteRequest) (*appscode_dtypes.VoidResponse, error)
	Renew(context.Context, *CertificateRenewRequest) (*CertificateRenewResponse, error)
	Revoke(context.Context, *CertificateRevokeRequest) (*appscode_dtypes.VoidResponse, error)
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
		FullMethod: "/appscode.certificate.v1beta1.Certificates/List",
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
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Describe",
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
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Create(ctx, req.(*CertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Obtain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateObtainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Obtain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Obtain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Obtain(ctx, req.(*CertificateObtainRequest))
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
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Delete",
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
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Renew",
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
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Revoke",
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
		FullMethod: "/appscode.certificate.v1beta1.Certificates/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Deploy(ctx, req.(*CertificateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Certificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.certificate.v1beta1.Certificates",
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
			MethodName: "Obtain",
			Handler:    _Certificates_Obtain_Handler,
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
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x97, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xc0, 0x35, 0xeb, 0xd4, 0xb1, 0xdf, 0x1a, 0x28, 0x73, 0x68, 0xb7, 0x4e, 0x4a, 0xd3, 0x85,
	0x43, 0x88, 0x2a, 0x9b, 0xa6, 0x25, 0x40, 0x85, 0x84, 0x88, 0x0d, 0xa2, 0x22, 0xc4, 0x61, 0x93,
	0x80, 0xc4, 0x65, 0x35, 0xde, 0x9d, 0x34, 0x4b, 0xec, 0x9d, 0x65, 0x67, 0x36, 0x60, 0xa1, 0x5e,
	0xfa, 0x01, 0xb8, 0x70, 0x42, 0x1c, 0xb9, 0x72, 0x01, 0x29, 0x1c, 0xb8, 0x72, 0xe1, 0xce, 0x95,
	0x23, 0x12, 0x5f, 0x03, 0xcd, 0x1f, 0xc7, 0xb3, 0x69, 0xda, 0x64, 0x93, 0x0b, 0x17, 0x6b, 0xe6,
	0xbd, 0x79, 0x33, 0xbf, 0xf7, 0xe6, 0xbd, 0x37, 0x5e, 0x78, 0x39, 0xa2, 0xb9, 0x48, 0xf6, 0x92,
	0x88, 0x08, 0xda, 0xc9, 0x72, 0x26, 0x18, 0x5e, 0x24, 0x59, 0xc6, 0x23, 0x16, 0xd3, 0x8e, 0xad,
	0x3b, 0xbc, 0x3b, 0xa4, 0x82, 0xdc, 0x6d, 0x2f, 0x3e, 0x62, 0xec, 0xd1, 0x88, 0x76, 0x49, 0x96,
	0x74, 0x49, 0x9a, 0x32, 0x41, 0x44, 0xc2, 0x52, 0xae, 0x6d, 0xdb, 0xaf, 0x4c, 0x6d, 0x9f, 0xa1,
	0xbf, 0x55, 0xd2, 0xc7, 0x62, 0x92, 0x51, 0xde, 0x55, 0xbf, 0x7a, 0x81, 0xef, 0xc1, 0xb5, 0xde,
	0xec, 0xd4, 0x8d, 0x84, 0x8b, 0x80, 0x7e, 0x55, 0x50, 0x2e, 0xfc, 0x1f, 0x10, 0x5c, 0x7f, 0x4a,
	0xc5, 0x33, 0x96, 0x72, 0x8a, 0xbb, 0x50, 0xe7, 0x82, 0x88, 0x82, 0x7b, 0x68, 0x09, 0x2d, 0xbb,
	0xab, 0xd7, 0x3b, 0xc7, 0x3e, 0xe8, 0x33, 0x3a, 0xdb, 0x4a, 0x1d, 0x98, 0x65, 0xf8, 0x13, 0x68,
	0x59, 0xce, 0x71, 0xcf, 0x59, 0xaa, 0x2d, 0xbb, 0xab, 0xaf, 0x77, 0x9e, 0xe7, 0x7a, 0xc7, 0x3a,
	0x3d, 0x28, 0x99, 0xfb, 0x1d, 0x68, 0x5b, 0xca, 0x3e, 0xe5, 0x51, 0x9e, 0x0c, 0xa9, 0x21, 0xc7,
	0x57, 0xa1, 0x56, 0x24, 0xb1, 0x42, 0x6b, 0x06, 0x72, 0xe8, 0xff, 0x88, 0x60, 0xe1, 0x54, 0x83,
	0x8b, 0xfa, 0xf3, 0x31, 0xb8, 0x16, 0x90, 0xe7, 0x28, 0xab, 0x0a, 0xee, 0xd8, 0xd6, 0xfe, 0x91,
	0x03, 0x9e, 0xa5, 0xec, 0xe5, 0x54, 0x2e, 0x31, 0xce, 0xdc, 0x86, 0x16, 0x89, 0x22, 0x56, 0xa4,
	0x22, 0xcc, 0xf6, 0x8f, 0xbd, 0x72, 0x8d, 0x6c, 0x6b, 0x3f, 0x89, 0x31, 0x86, 0xb9, 0x94, 0x8c,
	0x35, 0x45, 0x33, 0x50, 0x63, 0x7c, 0x0b, 0xdc, 0x88, 0x8d, 0xc7, 0x2c, 0x0d, 0x95, 0xaa, 0xa6,
	0x54, 0xa0, 0x45, 0x9b, 0x72, 0x01, 0x86, 0x39, 0x4e, 0x52, 0xee, 0xcd, 0x2d, 0xd5, 0xa4, 0x91,
	0x1c, 0xe3, 0x6b, 0x50, 0x1f, 0x16, 0x69, 0x3c, 0xa2, 0xde, 0x95, 0x25, 0xb4, 0xdc, 0x08, 0xcc,
	0x0c, 0xdf, 0x80, 0xc6, 0x01, 0x9d, 0x84, 0x31, 0x11, 0xc4, 0xab, 0xab, 0x9d, 0xe6, 0x0f, 0xe8,
	0xa4, 0x4f, 0x04, 0xc1, 0x1b, 0xd0, 0xe2, 0xc5, 0xf0, 0x4b, 0x1a, 0x89, 0x30, 0x49, 0xf7, 0x98,
	0x37, 0x7f, 0x9e, 0x48, 0x6c, 0x6b, 0x8b, 0x87, 0xe9, 0x1e, 0x0b, 0x5c, 0x3e, 0x9b, 0xe0, 0x3b,
	0x80, 0x13, 0xce, 0x0b, 0x1a, 0x66, 0x79, 0x72, 0x48, 0x04, 0x0d, 0xa5, 0xb5, 0xd7, 0x50, 0x30,
	0x57, 0x95, 0x66, 0x4b, 0x2b, 0x64, 0xb8, 0xfc, 0xcf, 0xc1, 0xb5, 0x76, 0xc2, 0x2d, 0x40, 0x3d,
	0x13, 0x1e, 0xd4, 0xc3, 0x2f, 0x82, 0xb3, 0xbd, 0x63, 0x42, 0xe2, 0x6c, 0xef, 0x48, 0xed, 0x86,
	0x09, 0x03, 0xda, 0x90, 0xb3, 0x81, 0x37, 0xa7, 0x67, 0x03, 0xb9, 0x76, 0xb0, 0xab, 0x7c, 0x6e,
	0x06, 0xce, 0x60, 0xd7, 0xff, 0x17, 0xc1, 0x8d, 0x53, 0x2e, 0xe4, 0xff, 0x90, 0x2c, 0xf2, 0xde,
	0x54, 0x1e, 0x68, 0x57, 0xd4, 0x58, 0xca, 0x54, 0xa0, 0xa4, 0x43, 0xad, 0x40, 0x8d, 0x65, 0x11,
	0x1c, 0xd0, 0x89, 0x72, 0xaa, 0x15, 0xc8, 0xa1, 0x5c, 0x95, 0x33, 0x26, 0xd4, 0x0d, 0xb6, 0x02,
	0x35, 0xf6, 0x7f, 0x76, 0xc0, 0xed, 0x9d, 0xb2, 0x3b, 0x2a, 0xef, 0x5e, 0x3d, 0xbd, 0x16, 0xa0,
	0xa9, 0xee, 0x2b, 0x0e, 0x87, 0x13, 0x13, 0xe8, 0x86, 0x16, 0xac, 0x4f, 0xf0, 0x4d, 0x80, 0x43,
	0x32, 0x4a, 0xe2, 0x70, 0x2f, 0x67, 0x63, 0x85, 0x58, 0x0b, 0x9a, 0x4a, 0xf2, 0x61, 0xce, 0xc6,
	0x72, 0x73, 0xfa, 0x4d, 0x96, 0xe4, 0x54, 0x66, 0x1c, 0x55, 0xbc, 0xb5, 0x00, 0xb4, 0xa8, 0x6f,
	0x28, 0x55, 0xee, 0xce, 0x5b, 0xb9, 0x3b, 0x8d, 0x41, 0x43, 0x53, 0xda, 0x31, 0x68, 0xea, 0x46,
	0x20, 0x63, 0xe0, 0xc1, 0xfc, 0x21, 0xcd, 0x79, 0xc2, 0x52, 0x0f, 0x96, 0xd0, 0xf2, 0x95, 0x60,
	0x3a, 0xc5, 0xaf, 0xc2, 0x0b, 0x9c, 0xe6, 0x09, 0x19, 0x85, 0x69, 0x31, 0x1e, 0xd2, 0xdc, 0x73,
	0x95, 0x55, 0x4b, 0x0b, 0x37, 0x95, 0xcc, 0xdf, 0x2b, 0x15, 0xea, 0x60, 0x28, 0x48, 0x92, 0x4e,
	0x0b, 0x75, 0x1a, 0x26, 0x64, 0x85, 0x69, 0x01, 0x9a, 0x12, 0x44, 0x57, 0x8e, 0x8e, 0x5f, 0x43,
	0x0a, 0x54, 0xe9, 0xd8, 0x55, 0x55, 0x2b, 0x55, 0x95, 0x7f, 0xa7, 0x74, 0x4e, 0x9f, 0x8e, 0xa8,
	0x78, 0x4e, 0x77, 0x1b, 0x96, 0x1a, 0x75, 0x40, 0x53, 0xfa, 0xf5, 0x25, 0xbb, 0x87, 0x39, 0xa3,
	0x36, 0x3b, 0x23, 0x2c, 0x11, 0x99, 0x33, 0x2e, 0x5a, 0x10, 0xd3, 0xbb, 0x72, 0x66, 0x77, 0xe5,
	0x0f, 0x4e, 0x1c, 0x70, 0xc8, 0x0e, 0xaa, 0xf4, 0x40, 0x43, 0xec, 0xcc, 0x88, 0xbf, 0x43, 0x27,
	0x82, 0x98, 0x8d, 0xd8, 0xe4, 0x99, 0x41, 0x94, 0x49, 0xc7, 0x69, 0x94, 0x53, 0x11, 0x5a, 0xd1,
	0x00, 0x2d, 0x52, 0x19, 0x7d, 0x1b, 0x5a, 0xd1, 0xa8, 0xe0, 0x82, 0xe6, 0x76, 0xce, 0xbb, 0x46,
	0xa6, 0x96, 0x2c, 0x42, 0x53, 0xaa, 0x78, 0x46, 0x22, 0x6a, 0x92, 0x7e, 0x26, 0xf0, 0xfb, 0xf0,
	0x52, 0x7f, 0x73, 0xbb, 0xb7, 0x4f, 0xa3, 0x83, 0x29, 0xc6, 0x35, 0xa8, 0xc7, 0x6c, 0x4c, 0x92,
	0xd4, 0x90, 0x98, 0xd9, 0x34, 0x35, 0x48, 0x21, 0xf6, 0x0d, 0x89, 0x4c, 0x8d, 0xf7, 0x0b, 0xb1,
	0xbf, 0xfa, 0xb7, 0x0b, 0x2d, 0xcb, 0x2d, 0x8e, 0x7f, 0x41, 0x30, 0x27, 0x1f, 0x67, 0x7c, 0xff,
	0xdc, 0x1d, 0xc5, 0x7a, 0xe6, 0xdb, 0x6f, 0x56, 0xb4, 0xd2, 0x77, 0xee, 0xbf, 0xfb, 0xe4, 0xc8,
	0x73, 0x1a, 0xe8, 0xc9, 0x5f, 0xff, 0x7c, 0xef, 0xbc, 0x81, 0x3b, 0xdd, 0xd2, 0xff, 0x0c, 0x6b,
	0xa3, 0xae, 0xd9, 0xc8, 0x96, 0x71, 0xfc, 0x07, 0x82, 0xc6, 0xf4, 0x11, 0xc6, 0x6f, 0x9f, 0x9b,
	0xe0, 0xc4, 0x43, 0xdf, 0x7e, 0xe7, 0x02, 0x96, 0x86, 0x7f, 0xdd, 0xe2, 0x5f, 0xc3, 0xf7, 0xab,
	0xf1, 0x77, 0xbf, 0x2d, 0x92, 0xf8, 0x31, 0xfe, 0x1d, 0x41, 0x5d, 0xbf, 0x0d, 0x78, 0xed, 0xdc,
	0x24, 0xa5, 0xd7, 0xbd, 0xfd, 0x56, 0x65, 0x3b, 0xc3, 0xff, 0x9e, 0xc5, 0x7f, 0xcf, 0xaf, 0x18,
	0xff, 0x07, 0x68, 0x05, 0xff, 0x86, 0xa0, 0xae, 0x1b, 0x58, 0x05, 0xf8, 0x52, 0xc7, 0x6b, 0xdf,
	0x7c, 0xaa, 0xce, 0x3f, 0x63, 0x49, 0x7c, 0x8c, 0xb8, 0x6b, 0x21, 0x3e, 0x6c, 0xf7, 0xab, 0x86,
	0x58, 0x16, 0xce, 0xe3, 0x2e, 0x89, 0xd4, 0xdf, 0xd9, 0x6e, 0x32, 0xce, 0x58, 0x2e, 0x24, 0xf8,
	0x4f, 0x08, 0xea, 0xba, 0x23, 0x56, 0x00, 0x2f, 0xb5, 0xd0, 0xb3, 0xc0, 0x4b, 0xb9, 0xb1, 0x72,
	0xb1, 0xdc, 0xf8, 0x13, 0xc1, 0x15, 0xd5, 0x25, 0xf1, 0xf9, 0x0b, 0xcc, 0xee, 0xdc, 0xed, 0xb5,
	0xaa, 0x66, 0x06, 0x7e, 0xcb, 0x82, 0xef, 0xfb, 0xeb, 0x17, 0x81, 0x3f, 0x0e, 0x7a, 0xae, 0x1c,
	0xf8, 0x15, 0x41, 0x5d, 0xf7, 0x63, 0x5c, 0x05, 0xca, 0x6a, 0xe0, 0x67, 0x05, 0xfc, 0x53, 0x8b,
	0xf9, 0x83, 0x76, 0xef, 0x92, 0xcc, 0x8a, 0xf4, 0x48, 0x65, 0x89, 0x6c, 0xf9, 0x95, 0xb2, 0xc4,
	0x7a, 0x23, 0xce, 0x82, 0xde, 0xb1, 0xa0, 0x3f, 0xba, 0x24, 0x74, 0xac, 0x0e, 0x7e, 0x80, 0x56,
	0xd6, 0x7b, 0xf0, 0x5a, 0xc4, 0xc6, 0xb3, 0x93, 0x49, 0x96, 0x9c, 0x46, 0xbd, 0x7e, 0xd5, 0xc2,
	0xde, 0x92, 0x5f, 0x72, 0x5b, 0xe8, 0x8b, 0x79, 0xa3, 0x1c, 0xd6, 0xd5, 0xb7, 0xdd, 0xbd, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0x56, 0x19, 0xea, 0x6d, 0x0e, 0x00, 0x00,
}
