// Code generated by protoc-gen-go.
// source: dashboard.proto
// DO NOT EDIT!

/*
Package monitoring is a generated protocol buffer package.

It is generated from these files:
	dashboard.proto

It has these top-level messages:
	DashboardCreateRequest
	DashboardCreateResponse
*/
package monitoring

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

type DashboardCreateRequest struct {
	Type    string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Cluster string `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Pod     string `protobuf:"bytes,3,opt,name=pod" json:"pod,omitempty"`
}

func (m *DashboardCreateRequest) Reset()                    { *m = DashboardCreateRequest{} }
func (m *DashboardCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*DashboardCreateRequest) ProtoMessage()               {}
func (*DashboardCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DashboardCreateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Url    string         `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *DashboardCreateResponse) Reset()                    { *m = DashboardCreateResponse{} }
func (m *DashboardCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*DashboardCreateResponse) ProtoMessage()               {}
func (*DashboardCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DashboardCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*DashboardCreateRequest)(nil), "monitoring.DashboardCreateRequest")
	proto.RegisterType((*DashboardCreateResponse)(nil), "monitoring.DashboardCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Dashboard service

type DashboardClient interface {
	Create(ctx context.Context, in *DashboardCreateRequest, opts ...grpc.CallOption) (*DashboardCreateResponse, error)
}

type dashboardClient struct {
	cc *grpc.ClientConn
}

func NewDashboardClient(cc *grpc.ClientConn) DashboardClient {
	return &dashboardClient{cc}
}

func (c *dashboardClient) Create(ctx context.Context, in *DashboardCreateRequest, opts ...grpc.CallOption) (*DashboardCreateResponse, error) {
	out := new(DashboardCreateResponse)
	err := grpc.Invoke(ctx, "/monitoring.Dashboard/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dashboard service

type DashboardServer interface {
	Create(context.Context, *DashboardCreateRequest) (*DashboardCreateResponse, error)
}

func RegisterDashboardServer(s *grpc.Server, srv DashboardServer) {
	s.RegisterService(&_Dashboard_serviceDesc, srv)
}

func _Dashboard_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DashboardCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DashboardServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Dashboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring.Dashboard",
	HandlerType: (*DashboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Dashboard_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0x2c, 0xce,
	0x48, 0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0xcd, 0xcf,
	0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86,
	0xa8, 0x94, 0x12, 0x03, 0x09, 0xa7, 0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88, 0xb8,
	0x92, 0x0b, 0x97, 0x98, 0x0b, 0xcc, 0x50, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x1e, 0x2e, 0x16, 0x90, 0x42, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e,
	0x21, 0x7e, 0x2e, 0xf6, 0xe4, 0x9c, 0xd2, 0xe2, 0x92, 0xd4, 0x22, 0x09, 0x26, 0xb0, 0x00, 0x37,
	0x17, 0x73, 0x41, 0x7e, 0x8a, 0x04, 0x33, 0x88, 0xa3, 0xe4, 0xc6, 0x25, 0x8e, 0x61, 0x4a, 0x71,
	0x01, 0xd0, 0xf6, 0x54, 0x21, 0x39, 0x2e, 0xb6, 0x62, 0xa0, 0x53, 0x4a, 0x8b, 0xc1, 0x06, 0x71,
	0x1b, 0xf1, 0xe9, 0x41, 0x5c, 0xa1, 0x17, 0x0c, 0x16, 0x05, 0x99, 0x53, 0x5a, 0x94, 0x03, 0x31,
	0xd4, 0x68, 0x12, 0x23, 0x17, 0x27, 0xdc, 0x20, 0xa1, 0x16, 0x46, 0x2e, 0x36, 0x88, 0x69, 0x42,
	0x4a, 0x7a, 0x08, 0x9f, 0xea, 0x61, 0x77, 0xb0, 0x94, 0x32, 0x5e, 0x35, 0x10, 0xe7, 0x28, 0x19,
	0x36, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x5b, 0x4a, 0x0d, 0x1c, 0x4e, 0x08, 0x0d, 0xfa, 0x65, 0x06,
	0xfa, 0xf0, 0xd0, 0xd5, 0xaf, 0x86, 0xfa, 0xb5, 0xd6, 0x8a, 0x51, 0x2b, 0x89, 0x0d, 0x1c, 0x52,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x06, 0x9e, 0x34, 0x7e, 0x01, 0x00, 0x00,
}