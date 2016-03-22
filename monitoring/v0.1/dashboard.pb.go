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
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0x2c, 0xce,
	0x48, 0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0xcd, 0xcf,
	0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86,
	0xa8, 0x94, 0x12, 0x03, 0x09, 0xa7, 0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88, 0xb8,
	0x52, 0x04, 0x97, 0x98, 0x0b, 0xcc, 0x50, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x16, 0x90, 0x42, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x39, 0xa7, 0xb4, 0xb8, 0x24, 0xb5, 0x48, 0x82, 0x09,
	0x2c, 0x0c, 0xe3, 0x0a, 0x09, 0x70, 0x31, 0x17, 0xe4, 0xa7, 0x48, 0x30, 0x83, 0x45, 0x41, 0x4c,
	0xa5, 0x60, 0x2e, 0x71, 0x0c, 0x93, 0x8b, 0x0b, 0x80, 0x2e, 0x4a, 0x15, 0x52, 0xe3, 0x62, 0x2b,
	0x06, 0x3a, 0xaf, 0xb4, 0x18, 0x6c, 0x38, 0xb7, 0x11, 0x9f, 0x1e, 0xc4, 0x65, 0x7a, 0xc1, 0x60,
	0xd1, 0x20, 0xa8, 0x2c, 0xc8, 0xd0, 0xd2, 0xa2, 0x1c, 0xa8, 0x55, 0x20, 0xa6, 0xd1, 0x24, 0x46,
	0x2e, 0x4e, 0xb8, 0xa9, 0x42, 0x2d, 0x8c, 0x5c, 0x6c, 0x10, 0xa3, 0x85, 0x94, 0xf4, 0x10, 0x41,
	0xa1, 0x87, 0xdd, 0x47, 0x52, 0xca, 0x78, 0xd5, 0x40, 0xdc, 0xa6, 0x64, 0xd8, 0x74, 0xf9, 0xc9,
	0x64, 0x26, 0x6d, 0x29, 0x35, 0x70, 0x40, 0x22, 0x34, 0xe8, 0x97, 0x19, 0xe8, 0xc3, 0x83, 0x5f,
	0xbf, 0x1a, 0xea, 0xef, 0x5a, 0x2b, 0x46, 0xad, 0x24, 0x36, 0x70, 0x50, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x39, 0xd0, 0xe8, 0xce, 0x9f, 0x01, 0x00, 0x00,
}
