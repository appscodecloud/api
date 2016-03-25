// Code generated by protoc-gen-go.
// source: health.proto
// DO NOT EDIT!

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	Response
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

type Response struct {
	Status   *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Latest   string         `protobuf:"bytes,2,opt,name=latest" json:"latest,omitempty"`
	Versions []*Version     `protobuf:"bytes,3,rep,name=versions" json:"versions,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Response) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Response) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

type Version struct {
	Version  string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Response)(nil), "health.Response")
	proto.RegisterType((*Version)(nil), "health.Version")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*Response, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *dtypes.VoidRequest) (*Response, error)
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
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x69, 0x2b, 0xdb, 0x74, 0xaa, 0x6d, 0xd9, 0x82, 0x84, 0x20, 0x52, 0x73, 0xf2, 0x94,
	0x40, 0x7d, 0x05, 0x0f, 0x5e, 0xf4, 0x50, 0xa1, 0x37, 0x0f, 0x2b, 0x19, 0xd2, 0x40, 0xdc, 0x5d,
	0x3b, 0x53, 0xc1, 0xab, 0xaf, 0xe0, 0xa3, 0xf9, 0x0a, 0x3e, 0x48, 0xf7, 0x5f, 0x2e, 0xcb, 0xee,
	0x6f, 0xe6, 0xfb, 0xe6, 0xdb, 0x81, 0xcb, 0x03, 0xaa, 0x9e, 0x0f, 0x95, 0x3d, 0x1a, 0x36, 0x52,
	0xc4, 0x57, 0x71, 0xd3, 0x1a, 0xd3, 0xf6, 0x58, 0x2b, 0xdb, 0xd5, 0x4a, 0x6b, 0xc3, 0x8a, 0x3b,
	0xa3, 0x29, 0x76, 0x15, 0xd7, 0x1e, 0x37, 0xfc, 0x6d, 0x91, 0xea, 0x70, 0x46, 0x5e, 0xbe, 0x41,
	0xb6, 0x43, 0xb2, 0xae, 0x11, 0xe5, 0x2d, 0x08, 0x72, 0xaa, 0x13, 0xe5, 0xa3, 0xcd, 0xe8, 0x7e,
	0xbe, 0x5d, 0x54, 0x51, 0x50, 0xbd, 0x06, 0x2a, 0x17, 0x20, 0x7a, 0xc5, 0x48, 0x9c, 0x8f, 0x5d,
	0x7d, 0x26, 0xef, 0x20, 0xfb, 0xc2, 0x23, 0xf9, 0x29, 0xf9, 0x64, 0x33, 0x71, 0x8a, 0x65, 0x95,
	0xa2, 0xed, 0x23, 0x2f, 0x9f, 0x61, 0x9a, 0xae, 0x72, 0x09, 0xd3, 0xd4, 0x1d, 0xec, 0x67, 0xde,
	0x2e, 0x8d, 0x8b, 0x76, 0x2b, 0xc8, 0x50, 0x37, 0xd6, 0x74, 0x9a, 0x9d, 0x9d, 0x27, 0x4e, 0xf2,
	0x81, 0x44, 0xaa, 0xc5, 0xfc, 0xc2, 0x83, 0xed, 0x0b, 0x88, 0xa7, 0x30, 0x40, 0x3e, 0x82, 0x48,
	0xa9, 0xd6, 0x43, 0xca, 0xbd, 0xe9, 0x9a, 0x1d, 0x7e, 0x9e, 0x5c, 0xc0, 0x62, 0x35, 0x04, 0x19,
	0x3e, 0x57, 0xae, 0x7f, 0xfe, 0xfe, 0x7f, 0xc7, 0x57, 0x72, 0x1e, 0x16, 0x14, 0xab, 0xef, 0x22,
	0x2c, 0xe1, 0xe1, 0x1c, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xa8, 0x1d, 0xd9, 0x52, 0x01, 0x00, 0x00,
}
