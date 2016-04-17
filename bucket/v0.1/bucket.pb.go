// Code generated by protoc-gen-go.
// source: bucket.proto
// DO NOT EDIT!

/*
Package bucket is a generated protocol buffer package.

It is generated from these files:
	bucket.proto

It has these top-level messages:
	BucketListRequest
	BucketListResponse
*/
package bucket

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

type BucketListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloudCredential" json:"cloudCredential,omitempty"`
}

func (m *BucketListRequest) Reset()                    { *m = BucketListRequest{} }
func (m *BucketListRequest) String() string            { return proto.CompactTextString(m) }
func (*BucketListRequest) ProtoMessage()               {}
func (*BucketListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BucketListResponse struct {
	Name   []string       `protobuf:"bytes,1,rep,name=name" json:"name,omitempty"`
	Status *dtypes.Status `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *BucketListResponse) Reset()                    { *m = BucketListResponse{} }
func (m *BucketListResponse) String() string            { return proto.CompactTextString(m) }
func (*BucketListResponse) ProtoMessage()               {}
func (*BucketListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BucketListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*BucketListRequest)(nil), "bucket.BucketListRequest")
	proto.RegisterType((*BucketListResponse)(nil), "bucket.BucketListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Buckets service

type BucketsClient interface {
	List(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error)
}

type bucketsClient struct {
	cc *grpc.ClientConn
}

func NewBucketsClient(cc *grpc.ClientConn) BucketsClient {
	return &bucketsClient{cc}
}

func (c *bucketsClient) List(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error) {
	out := new(BucketListResponse)
	err := grpc.Invoke(ctx, "/bucket.Buckets/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Buckets service

type BucketsServer interface {
	List(context.Context, *BucketListRequest) (*BucketListResponse, error)
}

func RegisterBucketsServer(s *grpc.Server, srv BucketsServer) {
	s.RegisterService(&_Buckets_serviceDesc, srv)
}

func _Buckets_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BucketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BucketsServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Buckets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bucket.Buckets",
	HandlerType: (*BucketsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Buckets_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2a, 0x4d, 0xce,
	0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xaa, 0xa4, 0xc4, 0x40, 0xc2, 0x29, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0xfa, 0x60, 0x12, 0x22, 0xae, 0x64, 0xcb, 0x25, 0xe8, 0x04, 0xd6, 0xef, 0x93, 0x59, 0x5c, 0x12,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc1, 0xc5, 0x9f, 0x9c, 0x93, 0x5f, 0x9a, 0xe2,
	0x5c, 0x94, 0x9a, 0x92, 0x9a, 0x57, 0x92, 0x99, 0x98, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x84, 0x2e, 0xac, 0x14, 0xc0, 0x25, 0x84, 0xac, 0xbd, 0xb8, 0x00, 0x68, 0x63, 0xaa, 0x90, 0x10,
	0x17, 0x4b, 0x5e, 0x62, 0x6e, 0x2a, 0x50, 0x13, 0x33, 0x50, 0x13, 0x98, 0x2d, 0xa4, 0xc6, 0xc5,
	0x56, 0x0c, 0x74, 0x52, 0x69, 0xb1, 0x04, 0x13, 0xd0, 0x28, 0x6e, 0x23, 0x3e, 0x3d, 0x88, 0x6b,
	0xf4, 0x82, 0xc1, 0xa2, 0x41, 0x50, 0x59, 0xa3, 0x52, 0x2e, 0x76, 0x88, 0x89, 0xc5, 0x42, 0x59,
	0x5c, 0x2c, 0x20, 0x63, 0x85, 0x24, 0xf5, 0xa0, 0x1e, 0xc6, 0x70, 0xa9, 0x94, 0x14, 0x36, 0x29,
	0x88, 0x2b, 0x94, 0xb4, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0x2a, 0xa4, 0x0c, 0x0e, 0x12, 0x88,
	0xba, 0x62, 0xfd, 0x32, 0x03, 0x3d, 0x43, 0xfd, 0x6a, 0x34, 0x7f, 0xd4, 0x26, 0xb1, 0x81, 0x83,
	0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x8e, 0x3c, 0xb1, 0x5c, 0x01, 0x00, 0x00,
}