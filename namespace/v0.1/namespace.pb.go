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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

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
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0x95, 0x6e, 0xb2, 0x69, 0x26, 0x6d, 0x91, 0x46, 0xd0, 0x2e, 0x4b, 0x05, 0xa9, 0x2b,
	0x2a, 0xe0, 0x90, 0x85, 0xf6, 0x00, 0xe2, 0x8a, 0x38, 0x20, 0xa1, 0x1e, 0xc2, 0x09, 0x09, 0x51,
	0x99, 0xae, 0x49, 0x2d, 0xd2, 0xb5, 0x89, 0x9d, 0xa2, 0x0a, 0x71, 0xe1, 0x15, 0x38, 0xf1, 0x58,
	0x88, 0x57, 0xe0, 0x41, 0xb0, 0xc7, 0xde, 0x92, 0x2d, 0x51, 0xa4, 0x5c, 0xa2, 0x9d, 0xf9, 0xc7,
	0xff, 0xe7, 0x19, 0x8f, 0x02, 0x37, 0x2a, 0x7e, 0x2e, 0x8c, 0xe6, 0xa7, 0x62, 0xa8, 0xa7, 0xca,
	0x2a, 0xec, 0x5d, 0x25, 0xf2, 0xdd, 0xb1, 0x52, 0xe3, 0x89, 0x28, 0xb8, 0x96, 0x05, 0xaf, 0x2a,
	0x65, 0xb9, 0x95, 0xaa, 0x32, 0xa1, 0x30, 0xdf, 0xf6, 0xe9, 0xd2, 0x5e, 0x6a, 0x61, 0x0a, 0xfa,
	0x0d, 0x79, 0xc6, 0x60, 0xe3, 0xc5, 0x99, 0x38, 0xfd, 0x34, 0x12, 0x9f, 0x67, 0xc2, 0x58, 0x44,
	0x68, 0x7b, 0xcb, 0xac, 0x35, 0x68, 0x3d, 0xe8, 0x8d, 0xe8, 0x9b, 0x3d, 0x85, 0xcd, 0x58, 0x63,
	0xb4, 0x73, 0x14, 0x78, 0x00, 0xa9, 0x71, 0xf6, 0x33, 0x43, 0x65, 0xfd, 0xc3, 0xad, 0x61, 0x70,
	0x1e, 0xbe, 0xa1, 0xec, 0x28, 0xaa, 0xec, 0x57, 0xcb, 0x9d, 0x9c, 0x0a, 0x6e, 0xc5, 0x12, 0x7b,
	0xdc, 0x83, 0x8d, 0x52, 0x1a, 0x3d, 0xe1, 0x97, 0x27, 0xa4, 0xad, 0x91, 0xd6, 0x8f, 0xb9, 0x63,
	0x5f, 0x72, 0x13, 0x3a, 0xe2, 0x9c, 0xcb, 0x49, 0x96, 0x90, 0x16, 0x02, 0xbc, 0x03, 0xbd, 0x99,
	0x11, 0xd3, 0x70, 0xaa, 0x4d, 0xca, 0xba, 0x4f, 0xd0, 0x91, 0x1c, 0xd6, 0x35, 0x37, 0xe6, 0x8b,
	0x9a, 0x96, 0x59, 0x27, 0x68, 0x75, 0xec, 0x89, 0xb2, 0xba, 0x90, 0x56, 0x9c, 0x04, 0xd7, 0x74,
	0x90, 0x78, 0x62, 0xc8, 0xbd, 0x24, 0xef, 0x1d, 0xe8, 0x92, 0xb7, 0xd4, 0x59, 0x97, 0x4e, 0xa7,
	0x3e, 0x7c, 0xa5, 0xd9, 0x33, 0xd8, 0xaa, 0x5b, 0x5a, 0x71, 0x1a, 0x03, 0x80, 0xd7, 0x6a, 0xbc,
	0x6c, 0xd0, 0x6f, 0xa1, 0x4f, 0x15, 0xab, 0x19, 0x23, 0x83, 0xf6, 0x44, 0x8d, 0x8d, 0x1b, 0x5c,
	0x42, 0x55, 0xff, 0x96, 0xc4, 0xbb, 0x91, 0xc6, 0x8e, 0x20, 0x71, 0x81, 0xa7, 0xca, 0xea, 0xa3,
	0xaa, 0xa9, 0xfe, 0x1b, 0x33, 0xe8, 0x96, 0xc2, 0xba, 0xa6, 0x4d, 0x1c, 0x7d, 0x1d, 0xb2, 0x7d,
	0xd8, 0x8c, 0xa8, 0x25, 0x97, 0x76, 0x03, 0xa9, 0x8b, 0x56, 0xbb, 0xf7, 0xe1, 0xcf, 0x04, 0x7a,
	0xc7, 0xf5, 0x5d, 0xf1, 0x3d, 0x74, 0x68, 0xcb, 0x70, 0x67, 0xae, 0x81, 0xf9, 0xdd, 0xcc, 0xb3,
	0xff, 0x85, 0x40, 0x64, 0xec, 0xfb, 0xef, 0x3f, 0x3f, 0xd6, 0x76, 0x31, 0xa7, 0xed, 0xbf, 0xaa,
	0x2a, 0x2e, 0x1e, 0x0f, 0x9f, 0x14, 0x5f, 0x7d, 0xfc, 0x0d, 0x4b, 0x48, 0xc3, 0xc3, 0x61, 0xc3,
	0x67, 0x7e, 0x3d, 0xf3, 0xdb, 0x0b, 0x94, 0x88, 0xb8, 0x4f, 0x88, 0x7b, 0xf9, 0x12, 0xc4, 0xf3,
	0xd6, 0x23, 0x3c, 0x83, 0x34, 0x74, 0xd9, 0xa0, 0x34, 0xa6, 0xd8, 0xa0, 0x34, 0x47, 0xc7, 0x1e,
	0x12, 0x65, 0x1f, 0xf7, 0x16, 0x51, 0xc2, 0xd8, 0xea, 0x7e, 0xde, 0x85, 0x17, 0xbd, 0x75, 0xed,
	0xb9, 0x23, 0x63, 0xfb, 0x7a, 0x3a, 0x02, 0x0e, 0x08, 0x30, 0xc0, 0xbb, 0x8b, 0x00, 0x6e, 0x53,
	0xa2, 0xfb, 0x87, 0x94, 0xfe, 0x1e, 0x8e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x67, 0xad,
	0xfb, 0x72, 0x04, 0x00, 0x00,
}