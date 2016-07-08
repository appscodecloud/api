// Code generated by protoc-gen-go.
// source: credential.proto
// DO NOT EDIT!

/*
Package credential is a generated protocol buffer package.

It is generated from these files:
	credential.proto

It has these top-level messages:
	CloudCredentialCreateRequest
	CloudCredentialUpdateRequest
	CloudCredentialDeleteRequest
	CloudCredentialListResponse
	Credential
*/
package credential

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CloudCredentialCreateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialCreateRequest) Reset()                    { *m = CloudCredentialCreateRequest{} }
func (m *CloudCredentialCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialCreateRequest) ProtoMessage()               {}
func (*CloudCredentialCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CloudCredentialCreateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialUpdateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialUpdateRequest) Reset()                    { *m = CloudCredentialUpdateRequest{} }
func (m *CloudCredentialUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialUpdateRequest) ProtoMessage()               {}
func (*CloudCredentialUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CloudCredentialUpdateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CloudCredentialDeleteRequest) Reset()                    { *m = CloudCredentialDeleteRequest{} }
func (m *CloudCredentialDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialDeleteRequest) ProtoMessage()               {}
func (*CloudCredentialDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CloudCredentialListResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Credentials []*Credential  `protobuf:"bytes,2,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *CloudCredentialListResponse) Reset()                    { *m = CloudCredentialListResponse{} }
func (m *CloudCredentialListResponse) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListResponse) ProtoMessage()               {}
func (*CloudCredentialListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CloudCredentialListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CloudCredentialListResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type Credential struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider    string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Information string `protobuf:"bytes,3,opt,name=information" json:"information,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*CloudCredentialCreateRequest)(nil), "credential.CloudCredentialCreateRequest")
	proto.RegisterType((*CloudCredentialUpdateRequest)(nil), "credential.CloudCredentialUpdateRequest")
	proto.RegisterType((*CloudCredentialDeleteRequest)(nil), "credential.CloudCredentialDeleteRequest")
	proto.RegisterType((*CloudCredentialListResponse)(nil), "credential.CloudCredentialListResponse")
	proto.RegisterType((*Credential)(nil), "credential.Credential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for CloudCredential service

type CloudCredentialClient interface {
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error)
	Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type cloudCredentialClient struct {
	cc *grpc.ClientConn
}

func NewCloudCredentialClient(cc *grpc.ClientConn) CloudCredentialClient {
	return &cloudCredentialClient{cc}
}

func (c *cloudCredentialClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error) {
	out := new(CloudCredentialListResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialClient) Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialClient) Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialClient) Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.CloudCredential/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudCredential service

type CloudCredentialServer interface {
	List(context.Context, *dtypes.VoidRequest) (*CloudCredentialListResponse, error)
	Create(context.Context, *CloudCredentialCreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *CloudCredentialUpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CloudCredentialDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterCloudCredentialServer(s *grpc.Server, srv CloudCredentialServer) {
	s.RegisterService(&_CloudCredential_serviceDesc, srv)
}

func _CloudCredential_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).List(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredential_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).Create(ctx, req.(*CloudCredentialCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredential_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).Update(ctx, req.(*CloudCredentialUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredential_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.CloudCredential/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialServer).Delete(ctx, req.(*CloudCredentialDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudCredential_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credential.CloudCredential",
	HandlerType: (*CloudCredentialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CloudCredential_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CloudCredential_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CloudCredential_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudCredential_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("credential.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0x53, 0x83, 0xf7, 0x04, 0xf4, 0x32, 0x5e, 0x2e, 0x25, 0xde, 0x45, 0x09, 0x72,
	0xad, 0x57, 0x4d, 0x34, 0x77, 0xe1, 0xa5, 0xdb, 0x56, 0x57, 0xae, 0x22, 0xba, 0x14, 0xc6, 0xce,
	0x58, 0x82, 0x69, 0x26, 0x66, 0x26, 0x85, 0x20, 0x45, 0xf0, 0x15, 0x7c, 0x30, 0x17, 0xbe, 0x82,
	0x4f, 0xe1, 0xca, 0xf9, 0x13, 0xd3, 0xa4, 0x6a, 0x88, 0x20, 0x6e, 0xca, 0x99, 0x33, 0x73, 0xe6,
	0x77, 0xbe, 0xd3, 0x6f, 0x02, 0xc7, 0xab, 0x82, 0x12, 0x9a, 0x89, 0x04, 0xa7, 0x41, 0x5e, 0x30,
	0xc1, 0x10, 0xec, 0x33, 0xde, 0xd9, 0x9a, 0xb1, 0x75, 0x4a, 0x43, 0x9c, 0x27, 0x21, 0xce, 0x32,
	0x26, 0xb0, 0x48, 0x58, 0xc6, 0xcd, 0x49, 0xef, 0x54, 0xa5, 0x89, 0xa8, 0x72, 0xca, 0x43, 0xfd,
	0x6b, 0xf2, 0xfe, 0x17, 0x0b, 0xce, 0x16, 0x29, 0x2b, 0xc9, 0xa2, 0xb9, 0x49, 0x46, 0x58, 0xd0,
	0x98, 0xbe, 0x2f, 0x29, 0x17, 0x08, 0xc1, 0x38, 0xc3, 0x1b, 0x3a, 0xb1, 0xa6, 0xd6, 0xec, 0x28,
	0xd6, 0x31, 0xf2, 0xe0, 0xba, 0xac, 0xde, 0x26, 0x84, 0x16, 0x93, 0x91, 0xce, 0x37, 0x6b, 0xf4,
	0x0c, 0xc6, 0x04, 0x0b, 0x3c, 0xb1, 0xa7, 0xf6, 0xcc, 0x8d, 0xa2, 0xa0, 0xd5, 0x73, 0x1f, 0x27,
	0x58, 0xca, 0xa2, 0xa7, 0x99, 0x28, 0xaa, 0x58, 0xd7, 0x7b, 0x4f, 0xe0, 0xa8, 0x49, 0xa1, 0x63,
	0xb0, 0xdf, 0xd1, 0xaa, 0xee, 0x41, 0x85, 0xe8, 0x04, 0xae, 0x6d, 0x71, 0x5a, 0xd2, 0x9a, 0x6f,
	0x16, 0xf3, 0xd1, 0x95, 0xf5, 0x3b, 0x45, 0x2f, 0x73, 0xf2, 0x5f, 0x14, 0x75, 0x38, 0xff, 0x4e,
	0x51, 0xf4, 0x8b, 0xa0, 0x25, 0x4d, 0x69, 0xaf, 0x20, 0xff, 0x23, 0xdc, 0x3e, 0xa8, 0x79, 0x9e,
	0x70, 0x11, 0x53, 0x9e, 0x4b, 0x4f, 0x50, 0x74, 0x0e, 0x0e, 0x97, 0x06, 0x29, 0xb9, 0x2e, 0x72,
	0xa3, 0x1b, 0x81, 0xf1, 0x46, 0xf0, 0x42, 0x67, 0xe3, 0x7a, 0x17, 0x5d, 0x81, 0xbb, 0x97, 0xcb,
	0x65, 0x6b, 0x6a, 0x04, 0xa7, 0x9d, 0x11, 0x34, 0x61, 0xdc, 0x3e, 0xea, 0xbf, 0x06, 0xd8, 0x6f,
	0xfd, 0xf5, 0xcc, 0xa7, 0xe0, 0x26, 0xd9, 0x5b, 0x56, 0x6c, 0xb4, 0x89, 0xe5, 0xe8, 0xd5, 0x76,
	0x3b, 0x15, 0x7d, 0xb7, 0xe1, 0xe6, 0x81, 0x42, 0x54, 0xc0, 0x58, 0xa9, 0x44, 0xb7, 0x7e, 0xaa,
	0x79, 0xc5, 0x12, 0x52, 0x4f, 0xc9, 0xbb, 0xdb, 0xf3, 0xc7, 0xb5, 0x67, 0xe3, 0x3f, 0xf8, 0xf4,
	0xf5, 0xdb, 0xe7, 0xd1, 0x39, 0xba, 0x23, 0x9f, 0x52, 0xce, 0x57, 0x8c, 0x98, 0x37, 0xd5, 0x12,
	0x17, 0x6e, 0x1f, 0x05, 0x8f, 0xc3, 0x95, 0xba, 0x03, 0x55, 0xe0, 0x18, 0x23, 0xa3, 0xd9, 0x50,
	0xaf, 0x7b, 0x27, 0xdd, 0xfe, 0x6a, 0x6e, 0xa8, 0xb9, 0xf7, 0xbc, 0x41, 0xdc, 0xb9, 0x75, 0xa1,
	0xd0, 0xc6, 0x71, 0xbd, 0xe8, 0x8e, 0x29, 0xfb, 0xd1, 0xfe, 0x60, 0xf4, 0x0e, 0x1c, 0xe3, 0xc1,
	0x5e, 0x74, 0xc7, 0xa6, 0x7f, 0x40, 0x5f, 0x6a, 0xf4, 0xc3, 0x8b, 0xfb, 0x43, 0xd0, 0xe1, 0x07,
	0xe5, 0x9c, 0xdd, 0x1b, 0x47, 0x7f, 0xbc, 0x2e, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x28,
	0x26, 0x11, 0x12, 0x05, 0x00, 0x00,
}
