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
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
}

func (m *BucketListRequest) Reset()                    { *m = BucketListRequest{} }
func (m *BucketListRequest) String() string            { return proto.CompactTextString(m) }
func (*BucketListRequest) ProtoMessage()               {}
func (*BucketListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BucketListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name   []string       `protobuf:"bytes,2,rep,name=name" json:"name,omitempty"`
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

	0xfa, 0x60, 0x12, 0x22, 0xae, 0x64, 0xc7, 0x25, 0xe8, 0x04, 0xd6, 0xef, 0x93, 0x59, 0x5c, 0x12,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc9, 0x25, 0x90, 0x9c, 0x93, 0x5f, 0x9a, 0x12,
	0x9f, 0x5c, 0x94, 0x9a, 0x92, 0x9a, 0x57, 0x92, 0x99, 0x98, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0xc4, 0x0f, 0x16, 0x77, 0x86, 0x0b, 0x2b, 0x05, 0x70, 0x09, 0x21, 0xeb, 0x2f, 0x2e, 0x00,
	0x5a, 0x99, 0x2a, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x0a, 0xd4, 0xc4, 0x0c, 0xd4, 0x04,
	0x66, 0x0b, 0xa9, 0x71, 0xb1, 0x15, 0x03, 0xdd, 0x54, 0x5a, 0x2c, 0xc1, 0x04, 0x34, 0x8a, 0xdb,
	0x88, 0x4f, 0x0f, 0xe2, 0x1c, 0xbd, 0x60, 0xb0, 0x68, 0x10, 0x54, 0xd6, 0x28, 0x99, 0x8b, 0x1d,
	0x62, 0x62, 0xb1, 0x50, 0x04, 0x17, 0x0b, 0xc8, 0x58, 0x21, 0x49, 0x3d, 0xa8, 0x8f, 0x31, 0x9c,
	0x2a, 0x25, 0x85, 0x4d, 0x0a, 0xe2, 0x0a, 0x25, 0xc9, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x09, 0x0b,
	0x09, 0x82, 0xc3, 0x04, 0xa2, 0xae, 0x58, 0xbf, 0xcc, 0x40, 0xcf, 0x30, 0x89, 0x0d, 0xec, 0x7b,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0xad, 0xc3, 0x28, 0x4b, 0x01, 0x00, 0x00,
}
