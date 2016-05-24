// Code generated by protoc-gen-go.
// source: charge.proto
// DO NOT EDIT!

/*
Package billing is a generated protocol buffer package.

It is generated from these files:
	charge.proto
	paymentmethod.proto
	purchase.proto
	subscription.proto

It has these top-level messages:
	ChargeRequest
	ChargeResponse
	ListPaymentResponse
	Purchase
	PurchaseOpenRequest
	PurchaseCloseRequest
	SubscriptionCreateRequest
	SubscriptionDescribeRequest
	SubscriptionDescribeResponse
	Subscription
	Quota
	PhabricatorQuota
	CIQuota
	ArtifactoryQuota
	ClusterQuota
	DBQuota
	SubscriptionQoutaRequest
	SubscriptionQutaResponse
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ChargeRequest struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ChargeRequest) Reset()                    { *m = ChargeRequest{} }
func (m *ChargeRequest) String() string            { return proto.CompactTextString(m) }
func (*ChargeRequest) ProtoMessage()               {}
func (*ChargeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ChargeResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	CostUsd string         `protobuf:"bytes,2,opt,name=cost_usd,json=costUsd" json:"cost_usd,omitempty"`
}

func (m *ChargeResponse) Reset()                    { *m = ChargeResponse{} }
func (m *ChargeResponse) String() string            { return proto.CompactTextString(m) }
func (*ChargeResponse) ProtoMessage()               {}
func (*ChargeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChargeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ChargeRequest)(nil), "billing.ChargeRequest")
	proto.RegisterType((*ChargeResponse)(nil), "billing.ChargeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Charge service

type ChargeClient interface {
	Calculate(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error)
}

type chargeClient struct {
	cc *grpc.ClientConn
}

func NewChargeClient(cc *grpc.ClientConn) ChargeClient {
	return &chargeClient{cc}
}

func (c *chargeClient) Calculate(ctx context.Context, in *ChargeRequest, opts ...grpc.CallOption) (*ChargeResponse, error) {
	out := new(ChargeResponse)
	err := grpc.Invoke(ctx, "/billing.Charge/Calculate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Charge service

type ChargeServer interface {
	Calculate(context.Context, *ChargeRequest) (*ChargeResponse, error)
}

func RegisterChargeServer(s *grpc.Server, srv ChargeServer) {
	s.RegisterService(&_Charge_serviceDesc, srv)
}

func _Charge_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChargeServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Charge/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChargeServer).Calculate(ctx, req.(*ChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Charge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.Charge",
	HandlerType: (*ChargeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Charge_Calculate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x48, 0x2c,
	0x4a, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xca, 0xcc, 0xc9, 0xc9, 0xcc,
	0x4b, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc,
	0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x93, 0x12, 0x03, 0x09, 0xa7,
	0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88, 0xb8, 0x92, 0x32, 0x17, 0xaf, 0x33, 0xd8,
	0xb8, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x16, 0x90, 0xbc, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x14, 0xcc, 0xc5, 0x07, 0x53, 0x54, 0x5c, 0x00, 0x34,
	0x33, 0x55, 0x48, 0x8d, 0x8b, 0xad, 0x18, 0x68, 0x41, 0x69, 0x31, 0x58, 0x1d, 0xb7, 0x11, 0x9f,
	0x1e, 0xc4, 0x6c, 0xbd, 0x60, 0xb0, 0x68, 0x10, 0x54, 0x56, 0x48, 0x92, 0x8b, 0x23, 0x39, 0xbf,
	0xb8, 0x24, 0xbe, 0xb4, 0x38, 0x45, 0x82, 0x09, 0x6c, 0x22, 0x3b, 0x88, 0x1f, 0x5a, 0x9c, 0x62,
	0x54, 0xc8, 0xc5, 0x06, 0x31, 0x54, 0x28, 0x9d, 0x8b, 0xd3, 0x39, 0x31, 0x27, 0xb9, 0x34, 0x27,
	0xb1, 0x24, 0x55, 0x48, 0x4c, 0x0f, 0xea, 0x21, 0x3d, 0x14, 0x77, 0x49, 0x89, 0x63, 0x88, 0x43,
	0x9c, 0xa2, 0xa4, 0xd9, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x65, 0x21, 0x45, 0xa0, 0xcf, 0x0b, 0x8a,
	0x93, 0xf3, 0x53, 0x20, 0x41, 0x00, 0x55, 0xad, 0x5f, 0x66, 0xa0, 0x67, 0xa8, 0x0f, 0x09, 0xb1,
	0x24, 0x36, 0xb0, 0x9f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x41, 0xf4, 0xf1, 0x42,
	0x01, 0x00, 0x00,
}
