// Code generated by protoc-gen-go.
// source: mailinglist.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	mailinglist.proto

It has these top-level messages:
	SendEmailRequest
	SubscribeRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type SendEmailRequest struct {
	SenderName    string `protobuf:"bytes,1,opt,name=sender_name,json=senderName" json:"sender_name,omitempty"`
	SenderEmail   string `protobuf:"bytes,2,opt,name=sender_email,json=senderEmail" json:"sender_email,omitempty"`
	Subject       string `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	Body          string `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	ReceiverEmail string `protobuf:"bytes,5,opt,name=receiver_email,json=receiverEmail" json:"receiver_email,omitempty"`
}

func (m *SendEmailRequest) Reset()                    { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string            { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()               {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendEmailRequest) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *SendEmailRequest) GetSenderEmail() string {
	if m != nil {
		return m.SenderEmail
	}
	return ""
}

func (m *SendEmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendEmailRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendEmailRequest) GetReceiverEmail() string {
	if m != nil {
		return m.ReceiverEmail
	}
	return ""
}

type SubscribeRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubscribeRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*SendEmailRequest)(nil), "appscode.mailinglist.v1beta1.SendEmailRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "appscode.mailinglist.v1beta1.SubscribeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MailingList service

type MailingListClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type mailingListClient struct {
	cc *grpc.ClientConn
}

func NewMailingListClient(cc *grpc.ClientConn) MailingListClient {
	return &mailingListClient{cc}
}

func (c *mailingListClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.mailinglist.v1beta1.MailingList/SendEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailingListClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.mailinglist.v1beta1.MailingList/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MailingList service

type MailingListServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*appscode_dtypes.VoidResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterMailingListServer(s *grpc.Server, srv MailingListServer) {
	s.RegisterService(&_MailingList_serviceDesc, srv)
}

func _MailingList_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.mailinglist.v1beta1.MailingList/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/appscode.mailinglist.v1beta1.MailingList/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MailingList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.mailinglist.v1beta1.MailingList",
	HandlerType: (*MailingListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _MailingList_SendEmail_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _MailingList_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailinglist.proto",
}

func init() { proto.RegisterFile("mailinglist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x65, 0x54, 0xbb, 0xae, 0xc7, 0x6d, 0x71, 0x87, 0x2e, 0x84, 0x70, 0xeb, 0x56, 0xb4, 0x60,
	0x6c, 0x90, 0xe2, 0x64, 0x67, 0xb2, 0x72, 0xc8, 0x2e, 0x09, 0xc6, 0x86, 0x2c, 0xb2, 0x09, 0x23,
	0xe9, 0x62, 0x26, 0x58, 0x33, 0x8a, 0x66, 0x6c, 0xf0, 0xd6, 0x9b, 0x7c, 0x40, 0x36, 0xf9, 0x87,
	0xec, 0x02, 0xf9, 0x92, 0xfc, 0x42, 0x3e, 0x24, 0x68, 0x46, 0xf2, 0x2b, 0x0f, 0x6f, 0x84, 0xe6,
	0x9e, 0x7b, 0xce, 0x3d, 0xf7, 0x70, 0xf1, 0x8f, 0x98, 0xb2, 0x09, 0xe3, 0xe3, 0x09, 0x93, 0xca,
	0x4b, 0x52, 0xa1, 0x04, 0x69, 0xd0, 0x24, 0x91, 0xa1, 0x88, 0xc0, 0x5b, 0xc7, 0x66, 0xdd, 0x00,
	0x14, 0xed, 0x3a, 0x8d, 0xb1, 0x10, 0xe3, 0x09, 0xf8, 0x34, 0x61, 0x3e, 0xe5, 0x5c, 0x28, 0xaa,
	0x98, 0xe0, 0xd2, 0x70, 0x9d, 0xdf, 0x05, 0xf7, 0x1d, 0xbc, 0xb9, 0x81, 0x47, 0x6a, 0x9e, 0x80,
	0xf4, 0xf5, 0xd7, 0x34, 0xb8, 0xf7, 0x08, 0xd7, 0x47, 0xc0, 0xa3, 0xe3, 0x6c, 0xf6, 0x10, 0xae,
	0xa7, 0x20, 0x15, 0x69, 0xe2, 0x9a, 0x04, 0x1e, 0x41, 0x7a, 0xc9, 0x69, 0x0c, 0x36, 0xfa, 0x83,
	0x5a, 0xd5, 0x21, 0x36, 0xa5, 0x33, 0x1a, 0x03, 0xf9, 0x8b, 0xbf, 0xe6, 0x0d, 0x90, 0xf1, 0x6c,
	0x4b, 0x77, 0xe4, 0x24, 0x2d, 0x45, 0x6c, 0x5c, 0x91, 0xd3, 0xe0, 0x0a, 0x42, 0x65, 0x7f, 0xd2,
	0x68, 0xf1, 0x24, 0x04, 0x97, 0x02, 0x11, 0xcd, 0xed, 0x92, 0x2e, 0xeb, 0x7f, 0xf2, 0x1f, 0x7f,
	0x4f, 0x21, 0x04, 0x36, 0x5b, 0x4a, 0x96, 0x35, 0xfa, 0xad, 0xa8, 0x6a, 0x51, 0xb7, 0x85, 0xeb,
	0xa3, 0x69, 0x20, 0xc3, 0x94, 0x05, 0x50, 0x98, 0xfd, 0x89, 0xcb, 0x86, 0x61, 0x6c, 0x9a, 0xc7,
	0xfe, 0x83, 0x85, 0x6b, 0xa7, 0x26, 0xce, 0x13, 0x26, 0x15, 0xb9, 0x41, 0xb8, 0xba, 0xdc, 0x93,
	0x78, 0xde, 0x47, 0x99, 0x7b, 0xdb, 0x81, 0x38, 0xbf, 0x56, 0xfd, 0x26, 0x43, 0xef, 0x5c, 0xb0,
	0x68, 0x08, 0x32, 0x11, 0x5c, 0x82, 0xdb, 0x59, 0x3c, 0xda, 0xd6, 0x17, 0xb4, 0x78, 0x7a, 0xbe,
	0xb5, 0x9a, 0xae, 0xe3, 0x6f, 0x84, 0x9e, 0xe9, 0xf8, 0xb9, 0x72, 0x0f, 0xb5, 0xc9, 0x5d, 0xe6,
	0xa4, 0x58, 0x62, 0xa7, 0x93, 0xad, 0x6d, 0x77, 0x39, 0x39, 0x5c, 0x73, 0xb2, 0xe7, 0x74, 0x5e,
	0x3b, 0xc9, 0x07, 0x14, 0x86, 0x7c, 0x59, 0x0c, 0xe8, 0xa1, 0x76, 0xff, 0x08, 0xff, 0x0b, 0x45,
	0xbc, 0x9a, 0x40, 0x13, 0xf6, 0x96, 0xab, 0x7e, 0x7d, 0x2d, 0xd9, 0x41, 0x76, 0x46, 0x03, 0x74,
	0x51, 0xc9, 0xc1, 0xe0, 0xb3, 0x3e, 0xac, 0x83, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x2e,
	0x03, 0x41, 0xea, 0x02, 0x00, 0x00,
}
