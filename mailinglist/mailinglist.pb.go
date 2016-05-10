// Code generated by protoc-gen-go.
// source: mailinglist.proto
// DO NOT EDIT!

/*
Package mailinglist is a generated protocol buffer package.

It is generated from these files:
	mailinglist.proto

It has these top-level messages:
	SubscribeRequest
*/
package mailinglist

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

type SubscribeRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "mailinglist.SubscribeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for MailingList service

type MailingListClient interface {
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type mailingListClient struct {
	cc *grpc.ClientConn
}

func NewMailingListClient(cc *grpc.ClientConn) MailingListClient {
	return &mailingListClient{cc}
}

func (c *mailingListClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/mailinglist.MailingList/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MailingList service

type MailingListServer interface {
	Subscribe(context.Context, *SubscribeRequest) (*dtypes.VoidResponse, error)
}

func RegisterMailingListServer(s *grpc.Server, srv MailingListServer) {
	s.RegisterService(&_MailingList_serviceDesc, srv)
}

func _MailingList_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailinglist.MailingList/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MailingList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mailinglist.MailingList",
	HandlerType: (*MailingListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _MailingList_Subscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4d, 0xcc, 0xcc,
	0xc9, 0xcc, 0x4b, 0xcf, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x92, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc,
	0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x95, 0x12, 0x03, 0x09, 0xa7,
	0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88, 0xb8, 0x92, 0x06, 0x97, 0x40, 0x70, 0x69,
	0x52, 0x71, 0x72, 0x51, 0x66, 0x52, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08,
	0x17, 0x6b, 0x2a, 0xc8, 0x64, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0xc7, 0xa8, 0x84,
	0x8b, 0xdb, 0x17, 0x62, 0x9d, 0x0f, 0xd0, 0x3a, 0xa1, 0x54, 0x2e, 0x4e, 0xb8, 0x46, 0x21, 0x59,
	0x3d, 0x64, 0xc7, 0xa1, 0x1b, 0x28, 0x25, 0xa2, 0x07, 0xb1, 0x59, 0x2f, 0x2c, 0x3f, 0x33, 0x25,
	0x28, 0xb5, 0xb8, 0x00, 0xe8, 0xae, 0x54, 0x25, 0xe5, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xc9, 0x4a,
	0x49, 0xe8, 0x23, 0x69, 0xd6, 0x2f, 0x33, 0xd0, 0x33, 0xd4, 0xcf, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d,
	0xb2, 0x62, 0xd4, 0x4a, 0x62, 0x03, 0x3b, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xba, 0x04,
	0xdf, 0xfb, 0xfe, 0x00, 0x00, 0x00,
}
