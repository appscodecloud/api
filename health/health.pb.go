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
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Version *Version       `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
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

func (m *Response) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Version struct {
	Name       string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Version    string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Status     string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Endpoint   string `protobuf:"bytes,4,opt,name=endpoint" json:"endpoint,omitempty"`
	CommitHash string `protobuf:"bytes,5,opt,name=commit_hash" json:"commit_hash,omitempty"`
	// Only development mode informations.
	GitBranch       string `protobuf:"bytes,6,opt,name=git_branch" json:"git_branch,omitempty"`
	GitTag          string `protobuf:"bytes,7,opt,name=git_tag" json:"git_tag,omitempty"`
	CommitTimestamp string `protobuf:"bytes,8,opt,name=commit_timestamp" json:"commit_timestamp,omitempty"`
	BuildTimestamp  string `protobuf:"bytes,9,opt,name=build_timestamp" json:"build_timestamp,omitempty"`
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

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
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x45, 0xd5, 0x52, 0xdc, 0x76, 0x5a, 0xda, 0xca, 0x95, 0xc0, 0xaa, 0x10, 0xaa, 0xb2, 0x62,
	0x95, 0x48, 0xe5, 0x0a, 0x2c, 0x58, 0xa0, 0x2e, 0x8a, 0xd4, 0x2d, 0x72, 0x1a, 0x2b, 0xb1, 0x94,
	0xd8, 0x26, 0x76, 0x90, 0xd8, 0x72, 0x05, 0x0e, 0xc3, 0x41, 0xb8, 0x02, 0x07, 0xc1, 0x19, 0x27,
	0x08, 0x36, 0x51, 0xe6, 0xfd, 0xf9, 0x33, 0xfe, 0x03, 0xf3, 0x42, 0xf0, 0xd2, 0x15, 0xb1, 0xa9,
	0xb5, 0xd3, 0x94, 0x84, 0x6a, 0x73, 0x9d, 0x6b, 0x9d, 0x97, 0x22, 0xe1, 0x46, 0x26, 0x5c, 0x29,
	0xed, 0xb8, 0x93, 0x5a, 0xd9, 0xd0, 0xb5, 0xb9, 0x6c, 0x71, 0xe6, 0xde, 0x8c, 0xb0, 0x09, 0x7e,
	0x03, 0x8f, 0x1e, 0x61, 0x72, 0x10, 0xd6, 0xf8, 0x46, 0x41, 0x6f, 0x80, 0x58, 0xef, 0x6a, 0x2c,
	0x1b, 0x6c, 0x07, 0xb7, 0xb3, 0xdd, 0x22, 0x0e, 0x86, 0xf8, 0x09, 0x29, 0xdd, 0xc2, 0xf8, 0x55,
	0xd4, 0xd6, 0x4f, 0x65, 0x43, 0x6c, 0x58, 0xc6, 0xdd, 0x4b, 0x8e, 0x01, 0x47, 0x9f, 0x03, 0x18,
	0x77, 0xff, 0x74, 0x0e, 0xa3, 0x3d, 0xaf, 0x04, 0xce, 0x9a, 0xd2, 0xe5, 0x7f, 0xef, 0x94, 0x2e,
	0x7e, 0x97, 0x9d, 0x61, 0xbd, 0x82, 0x89, 0x50, 0x99, 0xd1, 0x52, 0x39, 0x36, 0x42, 0xb2, 0x86,
	0xd9, 0x49, 0x57, 0x95, 0x74, 0xcf, 0x05, 0xb7, 0x05, 0x3b, 0x47, 0x48, 0x01, 0x72, 0x4f, 0xd2,
	0x9a, 0xab, 0x53, 0xc1, 0x48, 0x3f, 0xbb, 0x65, 0x8e, 0xe7, 0x6c, 0x8c, 0x80, 0xc1, 0xaa, 0x73,
	0x3a, 0x59, 0x09, 0xbf, 0xa6, 0x32, 0x6c, 0x82, 0xca, 0x15, 0x2c, 0xd3, 0x46, 0x96, 0xd9, 0x1f,
	0x61, 0xda, 0x0a, 0xbb, 0x3d, 0x90, 0x07, 0xcc, 0x42, 0xef, 0x81, 0x74, 0x79, 0xd7, 0x7d, 0xfe,
	0xa3, 0x96, 0xd9, 0x41, 0xbc, 0x34, 0xde, 0xb1, 0x59, 0xf5, 0x99, 0xfb, 0xb3, 0x45, 0xeb, 0xf7,
	0xaf, 0xef, 0x8f, 0xe1, 0x05, 0x9d, 0xe1, 0xe9, 0x83, 0x9a, 0x12, 0x3c, 0xef, 0xdd, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0x64, 0x54, 0x2d, 0xac, 0x01, 0x00, 0x00,
}
