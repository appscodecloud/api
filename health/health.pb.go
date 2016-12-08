// Code generated by protoc-gen-go.
// source: health.proto
// DO NOT EDIT!

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	StatusResponse
	Metadata
*/
package health

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"
import appscode_version "github.com/appscode/api/version"

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

type StatusResponse struct {
	Status   *appscode_dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Version  *appscode_version.Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Metadata *Metadata                 `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatusResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StatusResponse) GetVersion() *appscode_version.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *StatusResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Metadata struct {
	Env        string `protobuf:"bytes,1,opt,name=env" json:"env,omitempty"`
	Protocol   string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
	RootDomain string `protobuf:"bytes,3,opt,name=root_domain,json=rootDomain" json:"root_domain,omitempty"`
	TeamDomain string `protobuf:"bytes,4,opt,name=team_domain,json=teamDomain" json:"team_domain,omitempty"`
	TeamId     string `protobuf:"bytes,5,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Metadata) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *Metadata) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Metadata) GetRootDomain() string {
	if m != nil {
		return m.RootDomain
	}
	return ""
}

func (m *Metadata) GetTeamDomain() string {
	if m != nil {
		return m.TeamDomain
	}
	return ""
}

func (m *Metadata) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusResponse)(nil), "appscode.health.StatusResponse")
	proto.RegisterType((*Metadata)(nil), "appscode.health.Metadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/appscode.health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *appscode_dtypes.VoidRequest) (*StatusResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.health.Health/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Status(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Health_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}

func init() { proto.RegisterFile("health.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x49, 0xfb, 0x7d, 0x69, 0x3a, 0xfd, 0xf8, 0x94, 0x05, 0x49, 0x0c, 0xc5, 0x4a, 0xbc,
	0x78, 0x4a, 0xa4, 0xc5, 0x17, 0x28, 0x1e, 0xf4, 0x20, 0x94, 0x08, 0x3d, 0x78, 0x29, 0x6b, 0xb3,
	0xb4, 0x81, 0x64, 0x67, 0xed, 0x6e, 0x0b, 0x5e, 0xfb, 0x0a, 0x7d, 0x0e, 0x8f, 0x3e, 0x89, 0xaf,
	0xe0, 0x83, 0x48, 0x76, 0x37, 0x2d, 0xb1, 0x78, 0xc9, 0x66, 0xff, 0xf3, 0x9b, 0x99, 0xff, 0xcc,
	0xc2, 0xbf, 0x25, 0xa3, 0x85, 0x5a, 0xc6, 0x62, 0x85, 0x0a, 0xc9, 0x09, 0x15, 0x42, 0xce, 0x31,
	0x63, 0xb1, 0x91, 0xc3, 0xfe, 0x02, 0x71, 0x51, 0xb0, 0x84, 0x8a, 0x3c, 0xa1, 0x9c, 0xa3, 0xa2,
	0x2a, 0x47, 0x2e, 0x0d, 0x1e, 0x5e, 0xd4, 0xf8, 0x2f, 0xf1, 0x41, 0x23, 0x9e, 0xa9, 0x37, 0xc1,
	0x64, 0xa2, 0xbf, 0x16, 0x88, 0x1a, 0xc0, 0x86, 0xad, 0x64, 0x8e, 0xbc, 0x3e, 0x0d, 0x13, 0xbd,
	0x3b, 0xf0, 0xff, 0x49, 0x51, 0xb5, 0x96, 0x29, 0x93, 0x02, 0xb9, 0x64, 0x24, 0x01, 0x57, 0x6a,
	0x25, 0x70, 0x2e, 0x9d, 0xeb, 0xde, 0xd0, 0x8f, 0xf7, 0xbe, 0x4d, 0x93, 0xd8, 0x26, 0x58, 0x8c,
	0x8c, 0xa0, 0x63, 0x8b, 0x06, 0x2d, 0x9d, 0x71, 0x7e, 0xc8, 0xa8, 0xbb, 0x4d, 0xcd, 0x99, 0xd6,
	0x24, 0xb9, 0x05, 0xaf, 0x64, 0x8a, 0x66, 0x54, 0xd1, 0xa0, 0xfd, 0x33, 0xcb, 0xae, 0xed, 0xd1,
	0x02, 0xe9, 0x1e, 0x8d, 0x76, 0x0e, 0x78, 0xb5, 0x4c, 0x4e, 0xa1, 0xcd, 0xf8, 0x46, 0xdb, 0xec,
	0xa6, 0xd5, 0x2f, 0x09, 0xc1, 0xd3, 0x73, 0xcd, 0xb1, 0xd0, 0x5e, 0xba, 0xe9, 0xfe, 0x4e, 0x06,
	0xd0, 0x5b, 0x21, 0xaa, 0x59, 0x86, 0x25, 0xcd, 0xb9, 0x6e, 0xda, 0x4d, 0xa1, 0x92, 0xee, 0xb4,
	0x52, 0x01, 0x8a, 0xd1, 0xb2, 0x06, 0xfe, 0x18, 0xa0, 0x92, 0x2c, 0xe0, 0x43, 0x47, 0x03, 0x79,
	0x16, 0xfc, 0xd5, 0x41, 0xb7, 0xba, 0x3e, 0x64, 0xc3, 0x0d, 0xb8, 0xf7, 0xda, 0x32, 0x29, 0xc0,
	0x35, 0xdb, 0x21, 0xfd, 0xa3, 0xb5, 0x4d, 0x31, 0xcf, 0x52, 0xf6, 0xba, 0x66, 0x52, 0x85, 0x83,
	0xa3, 0x61, 0x9b, 0xaf, 0x10, 0x5d, 0x6d, 0x3f, 0x82, 0x96, 0xe7, 0x6c, 0x3f, 0xbf, 0x76, 0x2d,
	0x9f, 0x9c, 0x25, 0xb3, 0xc6, 0x63, 0x9a, 0x9c, 0xf1, 0x0d, 0xf8, 0x73, 0x2c, 0x0f, 0xa5, 0xa8,
	0xc8, 0x6d, 0xb9, 0x71, 0xcf, 0x18, 0x9a, 0x54, 0xd3, 0x4f, 0x9c, 0x67, 0xd7, 0xc8, 0x2f, 0xae,
	0x5e, 0xc7, 0xe8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x8a, 0x22, 0x80, 0x9a, 0x02, 0x00, 0x00,
}
