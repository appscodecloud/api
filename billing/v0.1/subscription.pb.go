// Code generated by protoc-gen-go.
// source: subscription.proto
// DO NOT EDIT!

package billing

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

type SubscriptionSubscribeRequest struct {
	ProductId string `protobuf:"bytes,1,opt,name=product_id" json:"product_id,omitempty"`
}

func (m *SubscriptionSubscribeRequest) Reset()                    { *m = SubscriptionSubscribeRequest{} }
func (m *SubscriptionSubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscriptionSubscribeRequest) ProtoMessage()               {}
func (*SubscriptionSubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterType((*SubscriptionSubscribeRequest)(nil), "billing.SubscriptionSubscribeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Subscription service

type SubscriptionClient interface {
	Subscribe(ctx context.Context, in *SubscriptionSubscribeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type subscriptionClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionClient(cc *grpc.ClientConn) SubscriptionClient {
	return &subscriptionClient{cc}
}

func (c *subscriptionClient) Subscribe(ctx context.Context, in *SubscriptionSubscribeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/billing.Subscription/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Subscription service

type SubscriptionServer interface {
	Subscribe(context.Context, *SubscriptionSubscribeRequest) (*dtypes.VoidResponse, error)
}

func RegisterSubscriptionServer(s *grpc.Server, srv SubscriptionServer) {
	s.RegisterService(&_Subscription_serviceDesc, srv)
}

func _Subscription_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Subscription/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServer).Subscribe(ctx, req.(*SubscriptionSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Subscription_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.Subscription",
	HandlerType: (*SubscriptionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Subscription_Subscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor4 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2e, 0x4d, 0x2a,
	0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4f, 0xca, 0xcc, 0xc9, 0xc9, 0xcc, 0x4b, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x04, 0xa9, 0x2a, 0x86, 0x28, 0x93,
	0x12, 0x03, 0x09, 0xa7, 0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88, 0xb8, 0x92, 0x11,
	0x97, 0x4c, 0x30, 0x92, 0xa1, 0x50, 0x76, 0x52, 0x6a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x90, 0x10, 0x17, 0x17, 0x50, 0x61, 0x4a, 0x69, 0x72, 0x49, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xa7, 0x51, 0x0f, 0x23, 0x17, 0x0f, 0xb2, 0x26, 0xa1, 0x1a, 0x2e, 0x4e, 0xb8, 0x46,
	0x21, 0x55, 0x3d, 0xa8, 0x8b, 0xf4, 0xf0, 0x19, 0x2c, 0x25, 0xa2, 0x07, 0x71, 0x8d, 0x5e, 0x58,
	0x7e, 0x66, 0x4a, 0x50, 0x6a, 0x71, 0x01, 0xd0, 0xad, 0xa9, 0x4a, 0x46, 0x4d, 0x97, 0x9f, 0x4c,
	0x66, 0xd2, 0x51, 0x52, 0x07, 0x7a, 0xa3, 0xa0, 0x38, 0x39, 0x3f, 0x05, 0xe2, 0x1f, 0xa8, 0x89,
	0xfa, 0x65, 0x06, 0x7a, 0x86, 0xfa, 0xc8, 0x81, 0x60, 0xc5, 0xa8, 0x95, 0xc4, 0x06, 0xf6, 0x89,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x62, 0xd3, 0xd9, 0xd5, 0x1e, 0x01, 0x00, 0x00,
}
