// Code generated by protoc-gen-go.
// source: server.proto
// DO NOT EDIT!

package backup

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

type ServerCreateRequest struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Credential  string `protobuf:"bytes,2,opt,name=credential" json:"credential,omitempty"`
	BucketName  string `protobuf:"bytes,3,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	Disk        string `protobuf:"bytes,4,opt,name=disk" json:"disk,omitempty"`
}

func (m *ServerCreateRequest) Reset()                    { *m = ServerCreateRequest{} }
func (m *ServerCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerCreateRequest) ProtoMessage()               {}
func (*ServerCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ServerDeleteRequest struct {
	ClusterName string `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
}

func (m *ServerDeleteRequest) Reset()                    { *m = ServerDeleteRequest{} }
func (m *ServerDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerDeleteRequest) ProtoMessage()               {}
func (*ServerDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*ServerCreateRequest)(nil), "backup.ServerCreateRequest")
	proto.RegisterType((*ServerDeleteRequest)(nil), "backup.ServerDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Server service

type ServerClient interface {
	Create(ctx context.Context, in *ServerCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ServerDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Create(ctx context.Context, in *ServerCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.Server/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) Delete(ctx context.Context, in *ServerDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.Server/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerServer interface {
	Create(context.Context, *ServerCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *ServerDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Server/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Create(ctx, req.(*ServerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Server/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Delete(ctx, req.(*ServerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Server_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Server_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("server.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x4d, 0x4b, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0xbc, 0xf3, 0x16, 0xbc, 0x33, 0xab, 0x28, 0x52, 0xaa, 0xe8, 0x58, 0x10, 0x86,
	0x11, 0x52, 0x3f, 0x36, 0xe2, 0x56, 0xd7, 0x2e, 0x2a, 0xb8, 0x95, 0xb4, 0xbd, 0x8c, 0xa1, 0x35,
	0x89, 0x49, 0x3a, 0xe0, 0xd6, 0xbd, 0x2b, 0x7f, 0x9a, 0x2b, 0xf7, 0xfe, 0x10, 0x6b, 0xae, 0x82,
	0x5f, 0x83, 0x9b, 0x52, 0x4e, 0x6e, 0xce, 0x39, 0xf7, 0x09, 0x8c, 0x1d, 0xda, 0x05, 0x5a, 0x6e,
	0xac, 0xf6, 0x9a, 0xc5, 0xa5, 0xa8, 0x9a, 0xce, 0xa4, 0x9b, 0x73, 0xad, 0xe7, 0x2d, 0xe6, 0xc2,
	0xc8, 0x5c, 0x28, 0xa5, 0xbd, 0xf0, 0x52, 0x2b, 0x47, 0x53, 0xe9, 0xfa, 0x9b, 0x5c, 0xfb, 0x3b,
	0x83, 0x2e, 0x0f, 0x5f, 0xd2, 0xb3, 0x87, 0x08, 0x56, 0x2f, 0x82, 0xdd, 0xa9, 0x45, 0xe1, 0xb1,
	0xc0, 0xdb, 0x0e, 0x9d, 0x67, 0x3b, 0x30, 0xae, 0xda, 0xce, 0x79, 0xb4, 0x57, 0x4a, 0xdc, 0x60,
	0x12, 0x4d, 0xa2, 0xe9, 0x4a, 0x31, 0x7a, 0xd7, 0xce, 0x7b, 0x89, 0x6d, 0x01, 0x54, 0x16, 0x6b,
	0x54, 0x5e, 0x8a, 0x36, 0x19, 0x84, 0x81, 0x4f, 0x0a, 0xdb, 0x86, 0x51, 0xd9, 0x55, 0x0d, 0x7a,
	0x72, 0xf8, 0x47, 0x03, 0x24, 0x05, 0x03, 0x06, 0xc3, 0x5a, 0xba, 0x26, 0x19, 0x86, 0x93, 0xf0,
	0x9f, 0x1d, 0x7f, 0xd4, 0x39, 0xc3, 0x16, 0x97, 0xd7, 0xf9, 0xff, 0xa3, 0xce, 0xe1, 0x73, 0x04,
	0x31, 0x5d, 0x65, 0xd7, 0x10, 0xd3, 0x36, 0x6c, 0x83, 0x13, 0x1d, 0xfe, 0xcb, 0x8e, 0xe9, 0x1a,
	0x27, 0x20, 0xfc, 0x52, 0xcb, 0xba, 0x40, 0x67, 0x7a, 0x5c, 0x98, 0xed, 0xdd, 0x3f, 0xbd, 0x3c,
	0x0e, 0x76, 0xd3, 0x49, 0x4f, 0xd2, 0xb8, 0x4a, 0xd7, 0x84, 0x94, 0x7c, 0xf2, 0xc5, 0x3e, 0x3f,
	0xc8, 0x89, 0xff, 0x49, 0x34, 0x63, 0x08, 0x31, 0x15, 0xfd, 0x9e, 0xf4, 0xa5, 0xfe, 0x92, 0xa4,
	0x69, 0x48, 0xca, 0x66, 0x7f, 0x26, 0x95, 0x71, 0x78, 0xac, 0xa3, 0xd7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0x75, 0x0b, 0xa0, 0xfa, 0x01, 0x00, 0x00,
}
