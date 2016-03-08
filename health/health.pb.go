// Code generated by protoc-gen-go.
// source: health.proto
// DO NOT EDIT!

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	HealthResponse
	Version
*/
package health

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

type HealthResponse struct {
	Status   *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Versions []*Version     `protobuf:"bytes,2,rep,name=versions" json:"versions,omitempty"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HealthResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *HealthResponse) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

type Version struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HealthResponse)(nil), "health.HealthResponse")
	proto.RegisterType((*Version)(nil), "health.Version")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := grpc.Invoke(ctx, "/health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *dtypes.VoidRequest) (*HealthResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HealthServer).Status(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Health_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xaa, 0xa4, 0xc4, 0x40, 0xc2, 0x29, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0xfa, 0x60, 0x12, 0x22, 0xae, 0x14, 0xcc, 0xc5, 0xe7, 0x01, 0xd6, 0x1f, 0x94, 0x5a, 0x5c, 0x00,
	0x54, 0x9e, 0x2a, 0x24, 0xc7, 0xc5, 0x56, 0x0c, 0xd4, 0x5b, 0x5a, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x6d, 0xc4, 0xa7, 0x07, 0xd1, 0xa6, 0x17, 0x0c, 0x16, 0x15, 0x52, 0xe4, 0xe2, 0x28, 0x4b,
	0x2d, 0x2a, 0x06, 0x99, 0x2d, 0xc1, 0xa4, 0xc0, 0x0c, 0x54, 0xc1, 0xaf, 0x07, 0x75, 0x50, 0x18,
	0x44, 0x5c, 0xc9, 0x9a, 0x8b, 0x1d, 0xca, 0x14, 0xe2, 0xe7, 0x62, 0x87, 0xaa, 0x06, 0x1b, 0xc7,
	0x29, 0xc4, 0x07, 0x37, 0x9e, 0x09, 0xcc, 0x07, 0x2a, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x95, 0x60, 0x06, 0x09, 0x18, 0x85, 0x70, 0xb1, 0x41, 0x5c, 0x24, 0xe4, 0xc5, 0xc5, 0x06, 0xb5,
	0x53, 0x18, 0xe6, 0x86, 0xb0, 0xfc, 0xcc, 0x94, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29,
	0x31, 0x98, 0xb5, 0xa8, 0x1e, 0x50, 0x12, 0x6e, 0xba, 0xfc, 0x64, 0x32, 0x13, 0xaf, 0x10, 0x37,
	0x38, 0x28, 0x20, 0x6a, 0x92, 0xd8, 0xc0, 0xde, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x58,
	0x3a, 0xf2, 0xa3, 0x3c, 0x01, 0x00, 0x00,
}