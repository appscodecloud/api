// Code generated by protoc-gen-go.
// source: events.proto
// DO NOT EDIT!

package kubernetes

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

type EventRequest struct {
	ClusterName    string                   `protobuf:"bytes,1,opt,name=cluster_name" json:"cluster_name,omitempty"`
	KubeNamespace  string                   `protobuf:"bytes,2,opt,name=kube_namespace" json:"kube_namespace,omitempty"`
	KubeObjectType string                   `protobuf:"bytes,3,opt,name=kube_object_type" json:"kube_object_type,omitempty"`
	KubeObjectName string                   `protobuf:"bytes,4,opt,name=kube_object_name" json:"kube_object_name,omitempty"`
	EventType      string                   `protobuf:"bytes,5,opt,name=event_type" json:"event_type,omitempty"`
	Metadata       *EventRequest_ObjectMeta `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *EventRequest) Reset()                    { *m = EventRequest{} }
func (m *EventRequest) String() string            { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()               {}
func (*EventRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EventRequest) GetMetadata() *EventRequest_ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type EventRequest_ObjectMeta struct {
	Kind                  string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Service               []string          `protobuf:"bytes,2,rep,name=service" json:"service,omitempty"`
	ReplicationController string            `protobuf:"bytes,3,opt,name=replication_controller" json:"replication_controller,omitempty"`
	Namespace             string            `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	PodIp                 string            `protobuf:"bytes,5,opt,name=pod_ip" json:"pod_ip,omitempty"`
	InstanceId            string            `protobuf:"bytes,6,opt,name=instance_id" json:"instance_id,omitempty"`
	Labels                map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EventRequest_ObjectMeta) Reset()                    { *m = EventRequest_ObjectMeta{} }
func (m *EventRequest_ObjectMeta) String() string            { return proto.CompactTextString(m) }
func (*EventRequest_ObjectMeta) ProtoMessage()               {}
func (*EventRequest_ObjectMeta) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *EventRequest_ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type EventResponse struct {
	Status *dtypes.Status           `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Tasks  []*EventResponse_Handler `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *EventResponse) Reset()                    { *m = EventResponse{} }
func (m *EventResponse) String() string            { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()               {}
func (*EventResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *EventResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *EventResponse) GetTasks() []*EventResponse_Handler {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type EventResponse_Handler struct {
	Action   string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Receiver string `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
	Base     string `protobuf:"bytes,3,opt,name=base" json:"base,omitempty"`
	Suffix   string `protobuf:"bytes,4,opt,name=suffix" json:"suffix,omitempty"`
	Verb     string `protobuf:"bytes,5,opt,name=verb" json:"verb,omitempty"`
	Data     []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EventResponse_Handler) Reset()                    { *m = EventResponse_Handler{} }
func (m *EventResponse_Handler) String() string            { return proto.CompactTextString(m) }
func (*EventResponse_Handler) ProtoMessage()               {}
func (*EventResponse_Handler) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

func init() {
	proto.RegisterType((*EventRequest)(nil), "kubernetes.EventRequest")
	proto.RegisterType((*EventRequest_ObjectMeta)(nil), "kubernetes.EventRequest.ObjectMeta")
	proto.RegisterType((*EventResponse)(nil), "kubernetes.EventResponse")
	proto.RegisterType((*EventResponse_Handler)(nil), "kubernetes.EventResponse.Handler")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Events service

type EventsClient interface {
	Constructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	Destructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Constructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Events/Constructive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Destructive(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Events/Destructive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Events service

type EventsServer interface {
	Constructive(context.Context, *EventRequest) (*EventResponse, error)
	Destructive(context.Context, *EventRequest) (*EventResponse, error)
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Constructive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Constructive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Events/Constructive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Constructive(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Destructive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Destructive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Events/Destructive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Destructive(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetes.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Constructive",
			Handler:    _Events_Constructive_Handler,
		},
		{
			MethodName: "Destructive",
			Handler:    _Events_Destructive_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor2 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x8a, 0xd4, 0x4c,
	0x10, 0xc7, 0xc9, 0xcc, 0x4e, 0x66, 0xb7, 0x92, 0x9d, 0x6f, 0xbf, 0x28, 0x43, 0x0c, 0x22, 0x3a,
	0x5e, 0x44, 0x30, 0x19, 0x47, 0x17, 0x64, 0x3c, 0x78, 0xd0, 0x05, 0x0f, 0x8a, 0xa0, 0x0f, 0x10,
	0x3a, 0x49, 0xed, 0xd2, 0x4e, 0xb6, 0x3b, 0xa6, 0x3b, 0xc1, 0x41, 0x14, 0xf1, 0xea, 0xd1, 0x27,
	0xf1, 0x2d, 0xbc, 0xfb, 0x0a, 0xbe, 0x81, 0x2f, 0x60, 0xa7, 0xd2, 0xeb, 0xcc, 0x61, 0x17, 0x16,
	0xbc, 0x84, 0xf4, 0xbf, 0xab, 0xab, 0x7e, 0xf5, 0xaf, 0x6e, 0xf0, 0xb1, 0x45, 0xa1, 0x55, 0x5c,
	0xd5, 0x52, 0xcb, 0x00, 0x56, 0x4d, 0x86, 0xb5, 0x40, 0x8d, 0x2a, 0xba, 0x7e, 0x22, 0xe5, 0x49,
	0x89, 0x09, 0xab, 0x78, 0xc2, 0x84, 0x90, 0x9a, 0x69, 0x2e, 0x85, 0x8d, 0x8c, 0xa6, 0x9d, 0x5c,
	0xe8, 0x75, 0x85, 0x2a, 0xa1, 0x6f, 0xaf, 0xcf, 0xbe, 0x0f, 0xc1, 0x3f, 0xea, 0x52, 0xbe, 0xc6,
	0x77, 0x0d, 0x2a, 0x1d, 0x5c, 0x05, 0x3f, 0x2f, 0x1b, 0xa5, 0xb1, 0x4e, 0x05, 0x3b, 0xc5, 0xd0,
	0xb9, 0xe9, 0xdc, 0xd9, 0x0b, 0xa6, 0x30, 0xe9, 0x4a, 0x91, 0xa4, 0x2a, 0x96, 0x63, 0x38, 0x20,
	0x3d, 0x84, 0x03, 0xd2, 0x65, 0xf6, 0x16, 0x73, 0x9d, 0x76, 0x99, 0xc3, 0xe1, 0x79, 0x3b, 0x94,
	0x6b, 0x87, 0x76, 0x02, 0x00, 0x6a, 0xa2, 0x8f, 0x1e, 0x91, 0x76, 0x08, 0xbb, 0xa7, 0xa8, 0x59,
	0xc1, 0x34, 0x0b, 0xc7, 0x46, 0xf1, 0x16, 0xb7, 0xe3, 0x4d, 0x6f, 0xf1, 0x36, 0x61, 0xfc, 0x8a,
	0x32, 0xbe, 0x34, 0xe1, 0xd1, 0x6f, 0x07, 0x60, 0xb3, 0x0c, 0x7c, 0xd8, 0x59, 0x71, 0x51, 0x58,
	0xe6, 0xff, 0x60, 0xac, 0xb0, 0x6e, 0x39, 0xc1, 0x0e, 0x8d, 0x70, 0x03, 0xa6, 0x35, 0x56, 0x25,
	0xcf, 0xc9, 0x99, 0x34, 0x97, 0x42, 0xd7, 0xb2, 0x2c, 0xb1, 0xb6, 0xc8, 0xff, 0xc3, 0xde, 0xa6,
	0xbf, 0x9e, 0x75, 0x02, 0x6e, 0x25, 0x8b, 0x94, 0x57, 0x96, 0xf3, 0x0a, 0x78, 0x5c, 0x28, 0xcd,
	0x44, 0x8e, 0x29, 0x2f, 0x42, 0x97, 0xc4, 0x27, 0xe0, 0x96, 0x2c, 0xc3, 0x52, 0x19, 0xf4, 0xa1,
	0x41, 0x4f, 0x2e, 0x81, 0x1e, 0xbf, 0xa0, 0x13, 0x47, 0x06, 0x61, 0x1d, 0xdd, 0x03, 0x6f, 0x6b,
	0x19, 0x78, 0x30, 0x5c, 0xe1, 0xda, 0x76, 0xb1, 0x0f, 0xa3, 0x96, 0x95, 0x8d, 0x35, 0x7c, 0x39,
	0x78, 0xe4, 0xcc, 0x7e, 0x38, 0xb0, 0x6f, 0xd3, 0xaa, 0xca, 0x8c, 0x18, 0x4d, 0x67, 0xae, 0x81,
	0xd2, 0x8d, 0xa2, 0x43, 0xde, 0x62, 0x12, 0xf7, 0xa3, 0x8e, 0xdf, 0x90, 0x1a, 0xcc, 0x61, 0xa4,
	0x99, 0x5a, 0x29, 0x32, 0xc2, 0x5b, 0xdc, 0x3a, 0x07, 0xb0, 0xcf, 0x14, 0x3f, 0x67, 0xa2, 0x30,
	0x8e, 0x44, 0x08, 0x63, 0xfb, 0xdb, 0x79, 0xc0, 0xf2, 0xce, 0x31, 0x4b, 0x74, 0x00, 0xbb, 0x35,
	0xe6, 0xc8, 0x5b, 0x63, 0x5c, 0x7f, 0x0b, 0x8c, 0xef, 0x19, 0x53, 0x67, 0x93, 0x37, 0xf1, 0xaa,
	0x39, 0x3e, 0xe6, 0xef, 0xad, 0x87, 0x66, 0xd7, 0x84, 0x66, 0xd6, 0x41, 0xb3, 0xa2, 0x29, 0x77,
	0xd6, 0xf9, 0x8b, 0xaf, 0x03, 0x70, 0x09, 0x40, 0x05, 0x9f, 0x1d, 0xf0, 0x9f, 0x1a, 0x06, 0x5d,
	0x37, 0xa6, 0x5a, 0x8b, 0x41, 0x78, 0x91, 0x8d, 0xd1, 0xb5, 0x0b, 0xf9, 0x67, 0x8f, 0xbf, 0xfc,
	0xfc, 0xf5, 0x6d, 0x70, 0x18, 0xcd, 0x93, 0x4d, 0x48, 0xd2, 0xce, 0xe3, 0xfb, 0x49, 0xff, 0x70,
	0x92, 0x0f, 0xdb, 0xb7, 0xfb, 0x63, 0x92, 0x9f, 0x55, 0x5c, 0x3a, 0x77, 0x83, 0x4f, 0xe0, 0x3d,
	0xc3, 0x7f, 0x04, 0x58, 0x12, 0xc0, 0xc3, 0x28, 0xb9, 0x24, 0x40, 0x81, 0x7f, 0xeb, 0x67, 0x2e,
	0xbd, 0xc9, 0x07, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xff, 0x23, 0x78, 0xe5, 0x03, 0x00,
	0x00,
}
