// Code generated by protoc-gen-go.
// source: namespace.proto
// DO NOT EDIT!

/*
Package namespace is a generated protocol buffer package.

It is generated from these files:
	namespace.proto

It has these top-level messages:
	CheckRequest
	CheckResponse
	CreateRequest
	CreateResponse
	LogRequest
	LogResponse
	Log
	StatusRequest
	StatusResponse
*/
package namespace

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

type CheckRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CheckResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type CreateRequest struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName string   `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Email       string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	UserName    string   `protobuf:"bytes,4,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password    string   `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	InviteEmail []string `protobuf:"bytes,6,rep,name=invite_email,json=inviteEmail" json:"invite_email,omitempty"`
	UserIp      string   `protobuf:"bytes,7,opt,name=user_ip,json=userIp" json:"user_ip,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CreateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type LogRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type LogResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Logs   []*Log         `protobuf:"bytes,2,rep,name=logs" json:"logs,omitempty"`
}

func (m *LogResponse) Reset()                    { *m = LogResponse{} }
func (m *LogResponse) String() string            { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()               {}
func (*LogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LogResponse) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Log struct {
	Info    string `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StatusRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type StatusResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StatusResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "namespace.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "namespace.CheckResponse")
	proto.RegisterType((*CreateRequest)(nil), "namespace.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "namespace.CreateResponse")
	proto.RegisterType((*LogRequest)(nil), "namespace.LogRequest")
	proto.RegisterType((*LogResponse)(nil), "namespace.LogResponse")
	proto.RegisterType((*Log)(nil), "namespace.Log")
	proto.RegisterType((*StatusRequest)(nil), "namespace.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "namespace.StatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Namespace service

type NamespaceClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type namespaceClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceClient(cc *grpc.ClientConn) NamespaceClient {
	return &namespaceClient{cc}
}

func (c *namespaceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/namespace.Namespace/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/namespace.Namespace/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/namespace.Namespace/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := grpc.Invoke(ctx, "/namespace.Namespace/Log", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Namespace service

type NamespaceServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Log(context.Context, *LogRequest) (*LogResponse, error)
}

func RegisterNamespaceServer(s *grpc.Server, srv NamespaceServer) {
	s.RegisterService(&_Namespace_serviceDesc, srv)
}

func _Namespace_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NamespaceServer).Check(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Namespace_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NamespaceServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Namespace_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NamespaceServer).Status(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Namespace_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NamespaceServer).Log(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Namespace_serviceDesc = grpc.ServiceDesc{
	ServiceName: "namespace.Namespace",
	HandlerType: (*NamespaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Namespace_Check_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Namespace_Create_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Namespace_Status_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _Namespace_Log_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x55, 0xba, 0xc9, 0xa6, 0x99, 0xb4, 0x45, 0x1a, 0x41, 0xbb, 0x2c, 0x05, 0x52, 0x17, 0x01,
	0xe2, 0x90, 0x45, 0xed, 0x01, 0xc4, 0x15, 0x71, 0x40, 0x42, 0x3d, 0x84, 0x13, 0x08, 0xa9, 0x72,
	0x1b, 0x37, 0x58, 0xa4, 0x6b, 0x13, 0x3b, 0x45, 0x15, 0xe2, 0xc2, 0x1f, 0x20, 0xbe, 0x0c, 0xf1,
	0x0b, 0x7c, 0x08, 0xf6, 0xd8, 0x5b, 0xb2, 0x09, 0x8d, 0x94, 0x4b, 0xb4, 0x33, 0xef, 0xf9, 0x3d,
	0xcf, 0xf3, 0x28, 0x70, 0xa3, 0xe4, 0xe7, 0xc2, 0x68, 0x7e, 0x2a, 0xfa, 0x7a, 0xa2, 0xac, 0xc2,
	0xce, 0x55, 0x23, 0xdf, 0x1d, 0x29, 0x35, 0x1a, 0x8b, 0x82, 0x6b, 0x59, 0xf0, 0xb2, 0x54, 0x96,
	0x5b, 0xa9, 0x4a, 0x13, 0x88, 0xf9, 0xb6, 0x6f, 0x0f, 0xed, 0xa5, 0x16, 0xa6, 0xa0, 0xdf, 0xd0,
	0x67, 0x0c, 0x36, 0x5e, 0x7e, 0x14, 0xa7, 0x9f, 0x06, 0xe2, 0xf3, 0x54, 0x18, 0x8b, 0x08, 0x4d,
	0x2f, 0x99, 0x35, 0x7a, 0x8d, 0xc7, 0x9d, 0x01, 0x7d, 0xb3, 0x67, 0xb0, 0x19, 0x39, 0x46, 0x3b,
	0x45, 0x81, 0x0f, 0x21, 0x35, 0x4e, 0x7e, 0x6a, 0x88, 0xd6, 0x3d, 0xd8, 0xea, 0x07, 0xe5, 0xfe,
	0x5b, 0xea, 0x0e, 0x22, 0xca, 0x7e, 0x35, 0xdc, 0xc9, 0x89, 0xe0, 0x56, 0x2c, 0x91, 0xc7, 0x3d,
	0xd8, 0x18, 0x4a, 0xa3, 0xc7, 0xfc, 0xf2, 0x98, 0xb0, 0x35, 0xc2, 0xba, 0xb1, 0x77, 0xe4, 0x29,
	0x37, 0xa1, 0x25, 0xce, 0xb9, 0x1c, 0x67, 0x09, 0x61, 0xa1, 0xc0, 0x3b, 0xd0, 0x99, 0x1a, 0x31,
	0x09, 0xa7, 0x9a, 0x84, 0xac, 0xfb, 0x06, 0x1d, 0xc9, 0x61, 0x5d, 0x73, 0x63, 0xbe, 0xa8, 0xc9,
	0x30, 0x6b, 0x05, 0xac, 0xaa, 0xbd, 0xa3, 0x2c, 0x2f, 0xa4, 0x15, 0xc7, 0x41, 0x35, 0xed, 0x25,
	0xde, 0x31, 0xf4, 0x5e, 0x91, 0xf6, 0x0e, 0xb4, 0x49, 0x5b, 0xea, 0xac, 0x4d, 0xa7, 0x53, 0x5f,
	0xbe, 0xd6, 0xec, 0x39, 0x6c, 0x55, 0x23, 0xad, 0x98, 0x46, 0x0f, 0xe0, 0x8d, 0x1a, 0x2d, 0x0b,
	0xfa, 0x1d, 0x74, 0x89, 0xb1, 0x9a, 0x30, 0x32, 0x68, 0x8e, 0xd5, 0xc8, 0xb8, 0xe0, 0x12, 0x62,
	0xfd, 0x5b, 0x12, 0xaf, 0x46, 0x18, 0x3b, 0x84, 0xc4, 0x15, 0xde, 0x55, 0x96, 0x67, 0xaa, 0x72,
	0xf5, 0xdf, 0x98, 0x41, 0x7b, 0x28, 0xac, 0x1b, 0xda, 0xc4, 0xe8, 0xab, 0x92, 0xed, 0xc3, 0x66,
	0xb4, 0x5a, 0x72, 0x69, 0x17, 0x48, 0x45, 0x5a, 0xed, 0xde, 0x07, 0x3f, 0x12, 0xe8, 0x1c, 0x55,
	0x77, 0xc5, 0x0f, 0xd0, 0xa2, 0x2d, 0xc3, 0x9d, 0x99, 0x01, 0x66, 0x77, 0x33, 0xcf, 0x16, 0x81,
	0xe0, 0xc8, 0x7a, 0xdf, 0x7f, 0xff, 0xf9, 0xb9, 0x96, 0x63, 0x46, 0xdb, 0x7f, 0xc5, 0x2a, 0x2e,
	0x9e, 0x16, 0x5f, 0x7d, 0xf5, 0x0d, 0x4f, 0x20, 0x0d, 0xcf, 0x86, 0x35, 0x95, 0xd9, 0xe5, 0xcc,
	0x6f, 0xff, 0x07, 0x89, 0x06, 0xfb, 0x64, 0x70, 0x37, 0xbf, 0xd6, 0xe0, 0x45, 0xe3, 0x09, 0x9e,
	0x41, 0x1a, 0x26, 0xac, 0x79, 0xd4, 0x12, 0xac, 0x79, 0xd4, 0x63, 0x63, 0x8f, 0xc8, 0x63, 0x0f,
	0xef, 0x2f, 0x7a, 0x84, 0xc0, 0xaa, 0x59, 0xde, 0x87, 0xb7, 0xbc, 0x35, 0xf7, 0xd0, 0xd1, 0x61,
	0x7b, 0xbe, 0x1d, 0xe5, 0x1f, 0x90, 0xfc, 0x3d, 0xdc, 0x5d, 0x94, 0x77, 0x1b, 0x12, 0xb5, 0x4f,
	0x52, 0xfa, 0x5b, 0x38, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x00, 0x19, 0xd3, 0x4f, 0x6a, 0x04,
	0x00, 0x00,
}
