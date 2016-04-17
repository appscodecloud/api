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
	ClusterName    string                   `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	KubeNamespace  string                   `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	KubeObjectType string                   `protobuf:"bytes,3,opt,name=kube_object_type,json=kubeObjectType" json:"kube_object_type,omitempty"`
	KubeObjectName string                   `protobuf:"bytes,4,opt,name=kube_object_name,json=kubeObjectName" json:"kube_object_name,omitempty"`
	EventType      string                   `protobuf:"bytes,5,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
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
	ReplicationController string            `protobuf:"bytes,3,opt,name=replication_controller,json=replicationController" json:"replication_controller,omitempty"`
	Namespace             string            `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	PodIp                 string            `protobuf:"bytes,5,opt,name=pod_ip,json=podIp" json:"pod_ip,omitempty"`
	InstanceId            string            `protobuf:"bytes,6,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
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
const _ = grpc.SupportPackageIsVersion1

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

func _Events_Constructive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(EventsServer).Constructive(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Events_Destructive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(EventsServer).Destructive(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x6b, 0xd4, 0x4e,
	0x14, 0x25, 0xbb, 0xdd, 0xb4, 0xbd, 0xd9, 0x96, 0x32, 0xfc, 0x5a, 0xf2, 0x0b, 0x15, 0x6d, 0x45,
	0x29, 0x3e, 0x24, 0xba, 0x22, 0x6a, 0x51, 0xc4, 0x3f, 0xa5, 0x2d, 0x68, 0x85, 0xe8, 0xfb, 0x32,
	0x9b, 0xdc, 0x96, 0xb8, 0x69, 0x12, 0x33, 0x93, 0xe0, 0x22, 0xbe, 0x88, 0x6f, 0xe2, 0x83, 0xf8,
	0x1d, 0xfc, 0x32, 0x3e, 0xfa, 0xe6, 0xb3, 0x1f, 0xc4, 0x99, 0x9b, 0xc9, 0xee, 0x52, 0x2c, 0x88,
	0xbe, 0x2c, 0x73, 0xcf, 0x3d, 0x73, 0xef, 0xb9, 0xe7, 0x4e, 0x16, 0xfa, 0x58, 0x63, 0x26, 0x85,
	0x5f, 0x94, 0xb9, 0xcc, 0x19, 0x8c, 0xab, 0x11, 0x96, 0x19, 0x4a, 0x14, 0xde, 0xe6, 0x49, 0x9e,
	0x9f, 0xa4, 0x18, 0xf0, 0x22, 0x09, 0x78, 0x96, 0xe5, 0x92, 0xcb, 0x24, 0xcf, 0x0c, 0xd3, 0xdb,
	0xd0, 0x70, 0x2c, 0x27, 0x05, 0x8a, 0x80, 0x7e, 0x1b, 0x7c, 0xfb, 0xc7, 0x02, 0xf4, 0xf7, 0x74,
	0xc9, 0x10, 0x5f, 0x57, 0x28, 0x24, 0xdb, 0x82, 0x7e, 0x94, 0x56, 0x42, 0x62, 0x39, 0xcc, 0xf8,
	0x29, 0xba, 0xd6, 0x25, 0x6b, 0x67, 0x39, 0x74, 0x0c, 0x76, 0xa4, 0x20, 0x76, 0x05, 0x56, 0x75,
	0x5f, 0xca, 0x8b, 0x82, 0x47, 0xe8, 0x76, 0x88, 0xb4, 0xa2, 0xd1, 0xa3, 0x16, 0x64, 0x3b, 0xb0,
	0x46, 0xb4, 0x7c, 0xf4, 0x0a, 0x23, 0x39, 0xd4, 0x5d, 0xdd, 0x2e, 0x11, 0xe9, 0xfa, 0x73, 0x82,
	0x5f, 0x2a, 0xf4, 0x2c, 0x93, 0xfa, 0x2e, 0x9c, 0x65, 0x52, 0xeb, 0x0b, 0x00, 0x64, 0x40, 0x53,
	0xad, 0x47, 0x9c, 0x65, 0x42, 0xa8, 0xd0, 0x03, 0x58, 0x3a, 0x45, 0xc9, 0x63, 0x2e, 0xb9, 0xbb,
	0xa8, 0x92, 0xce, 0xe0, 0xb2, 0x3f, 0xb3, 0xc8, 0x9f, 0x1f, 0xd4, 0x6f, 0xaa, 0x3e, 0x53, 0xf4,
	0x70, 0x7a, 0xc9, 0xfb, 0xd6, 0x01, 0x98, 0x25, 0x18, 0x83, 0x85, 0x71, 0x92, 0xc5, 0xc6, 0x04,
	0x3a, 0x33, 0x17, 0x16, 0x05, 0x96, 0x75, 0x42, 0x63, 0x77, 0x15, 0xdc, 0x86, 0xec, 0x16, 0x6c,
	0x94, 0x58, 0xa4, 0x49, 0x44, 0xce, 0x0f, 0xa3, 0x3c, 0x93, 0x65, 0x9e, 0xa6, 0x58, 0x9a, 0xb1,
	0xd7, 0xe7, 0xb2, 0x8f, 0xa7, 0x49, 0xb6, 0x09, 0xcb, 0x33, 0x27, 0x9b, 0xb1, 0x67, 0x00, 0x5b,
	0x07, 0xbb, 0xc8, 0xe3, 0x61, 0x52, 0x98, 0x69, 0x7b, 0x2a, 0x3a, 0x2c, 0xd8, 0x45, 0x70, 0x92,
	0x4c, 0x48, 0x9e, 0x45, 0x38, 0x4c, 0x62, 0xd7, 0xa6, 0x1c, 0xb4, 0xd0, 0x61, 0xcc, 0xf6, 0xc1,
	0x4e, 0xf9, 0x08, 0x53, 0xa1, 0x8c, 0xe8, 0x2a, 0x23, 0x82, 0x3f, 0x30, 0xc2, 0x7f, 0x4a, 0x37,
	0xf6, 0x94, 0xb4, 0x49, 0x68, 0xae, 0x7b, 0x77, 0xc1, 0x99, 0x83, 0xd9, 0x1a, 0x74, 0xc7, 0x38,
	0x31, 0x8e, 0xe8, 0x23, 0xfb, 0x0f, 0x7a, 0x35, 0x4f, 0xab, 0xf6, 0x15, 0x34, 0xc1, 0x6e, 0xe7,
	0x8e, 0xb5, 0xfd, 0xa1, 0x03, 0x2b, 0xa6, 0x95, 0x28, 0xd4, 0x5b, 0x44, 0x76, 0x15, 0x6c, 0xa5,
	0x50, 0x56, 0x82, 0x0a, 0x38, 0x83, 0x55, 0xbf, 0x79, 0x93, 0xfe, 0x0b, 0x42, 0x43, 0x93, 0x65,
	0xb7, 0xa1, 0x27, 0xb9, 0x18, 0x0b, 0xb2, 0xd8, 0x19, 0x6c, 0xfd, 0x46, 0x7c, 0x53, 0xd1, 0x3f,
	0xe0, 0x59, 0xac, 0x5c, 0x0c, 0x1b, 0xbe, 0xf7, 0xd9, 0x82, 0x45, 0x03, 0xb1, 0x0d, 0xb0, 0x79,
	0xa4, 0xcd, 0x36, 0x6a, 0x4d, 0xc4, 0x3c, 0x58, 0x2a, 0x31, 0xc2, 0xa4, 0x56, 0x9b, 0x69, 0x34,
	0x4f, 0x63, 0xbd, 0xf1, 0x11, 0x17, 0xed, 0x43, 0xa5, 0xb3, 0xae, 0x23, 0xaa, 0xe3, 0xe3, 0xe4,
	0x8d, 0xd9, 0x8e, 0x89, 0x34, 0x57, 0x5d, 0x19, 0x99, 0xc5, 0xd0, 0x59, 0x63, 0xf4, 0xfa, 0xf4,
	0x42, 0xfa, 0x21, 0x9d, 0x07, 0x5f, 0x3b, 0x60, 0x93, 0x68, 0xc1, 0x3e, 0x59, 0xd0, 0x57, 0xab,
	0x17, 0xb2, 0xac, 0x94, 0x96, 0x1a, 0x99, 0x7b, 0xde, 0x5a, 0xbc, 0xff, 0xcf, 0x9d, 0x79, 0xfb,
	0xe0, 0xfd, 0xf7, 0x9f, 0x5f, 0x3a, 0x8f, 0xbc, 0xfb, 0xf4, 0xb1, 0xcf, 0x68, 0x41, 0x7d, 0xdd,
	0xbf, 0x11, 0x98, 0x4f, 0x55, 0x04, 0x6f, 0xe7, 0x3f, 0xe4, 0x77, 0x41, 0xf3, 0xcf, 0x11, 0x44,
	0x6d, 0xfb, 0x5d, 0xeb, 0x1a, 0xfb, 0x68, 0x81, 0xf3, 0x04, 0xff, 0x51, 0xce, 0x3e, 0xc9, 0x79,
	0xe8, 0xdd, 0xfb, 0x1b, 0x39, 0x31, 0x4e, 0xd5, 0x8c, 0x6c, 0xfa, 0x4f, 0xba, 0xf9, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x42, 0x17, 0x87, 0xe5, 0x04, 0x00, 0x00,
}