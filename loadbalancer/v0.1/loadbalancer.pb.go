// Code generated by protoc-gen-go.
// source: loadbalancer.proto
// DO NOT EDIT!

/*
Package loadbalancer is a generated protocol buffer package.

It is generated from these files:
	loadbalancer.proto

It has these top-level messages:
	ListRequest
	ListResponse
	DescribeRequest
	DescribeResponse
	CreateRequest
	UpdateRequest
	DeleteRequest
	LoadBalancer
	Spec
	Status
	LoadBalancerStatus
	LoadBalancerBackend
	LoadBalancerRule
	HTTPLoadBalancerRule
	TCPLoadBalancerRule
*/
package loadbalancer

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

type ListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListResponse struct {
	Status       *dtypes.Status  `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancer []*LoadBalancer `protobuf:"bytes,2,rep,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListResponse) GetLoadBalancer() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type DescribeRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DescribeResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancer *LoadBalancer  `protobuf:"bytes,2,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type CreateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace    string        `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Cluster      string        `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,4,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type UpdateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cluster      string        `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,3,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type DeleteRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LoadBalancer struct {
	// 'kind' defines is it the regular kubernetes instance or the
	// appscode superset called Extended Ingress. This field will
	// strictly contains only those two values
	// 'ingress' - default kubernetes ingress object.
	// 'extendedIngress' - appscode superset of ingress.
	// when creating a Loadbalancer from UI this field will always
	// be only 'extendedIngress.' List, Describe, Update and Delete
	// will support both two modes.
	// Create will support only extendedIngress.
	// For Creating or Updating an regular ingress one must use the
	// kubectl or direct API calls directly to kubernetes.
	Kind              string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name              string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace         string            `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	CreationTimestamp string            `protobuf:"bytes,4,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	Options           map[string]string `protobuf:"bytes,5,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec              *Spec             `protobuf:"bytes,6,opt,name=spec" json:"spec,omitempty"`
	Status            *Status           `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
}

func (m *LoadBalancer) Reset()                    { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()               {}
func (*LoadBalancer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoadBalancer) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *LoadBalancer) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *LoadBalancer) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type Spec struct {
	Backend *HTTPLoadBalancerRule `protobuf:"bytes,1,opt,name=backend" json:"backend,omitempty"`
	Rules   []*LoadBalancerRule   `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Spec) GetBackend() *HTTPLoadBalancerRule {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Spec) GetRules() []*LoadBalancerRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Status struct {
	Status []*LoadBalancerStatus `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Status) GetStatus() []*LoadBalancerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type LoadBalancerStatus struct {
	IP   string `protobuf:"bytes,1,opt,name=IP,json=iP" json:"IP,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *LoadBalancerStatus) Reset()                    { *m = LoadBalancerStatus{} }
func (m *LoadBalancerStatus) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerStatus) ProtoMessage()               {}
func (*LoadBalancerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type LoadBalancerBackend struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServicePort string `protobuf:"bytes,2,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
}

func (m *LoadBalancerBackend) Reset()                    { *m = LoadBalancerBackend{} }
func (m *LoadBalancerBackend) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerBackend) ProtoMessage()               {}
func (*LoadBalancerBackend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type LoadBalancerRule struct {
	Host string                  `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Http []*HTTPLoadBalancerRule `protobuf:"bytes,2,rep,name=http" json:"http,omitempty"`
	Tcp  []*TCPLoadBalancerRule  `protobuf:"bytes,3,rep,name=tcp" json:"tcp,omitempty"`
}

func (m *LoadBalancerRule) Reset()                    { *m = LoadBalancerRule{} }
func (m *LoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerRule) ProtoMessage()               {}
func (*LoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *LoadBalancerRule) GetHttp() []*HTTPLoadBalancerRule {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *LoadBalancerRule) GetTcp() []*TCPLoadBalancerRule {
	if m != nil {
		return m.Tcp
	}
	return nil
}

type HTTPLoadBalancerRule struct {
	Path          string               `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Backend       *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	HeaderRule    []string             `protobuf:"bytes,3,rep,name=header_rule,json=headerRule" json:"header_rule,omitempty"`
	RewriteRule   []string             `protobuf:"bytes,4,rep,name=rewrite_rule,json=rewriteRule" json:"rewrite_rule,omitempty"`
	SSLSecretName string               `protobuf:"bytes,5,opt,name=SSL_secret_name,json=sSLSecretName" json:"SSL_secret_name,omitempty"`
}

func (m *HTTPLoadBalancerRule) Reset()                    { *m = HTTPLoadBalancerRule{} }
func (m *HTTPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*HTTPLoadBalancerRule) ProtoMessage()               {}
func (*HTTPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *HTTPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

type TCPLoadBalancerRule struct {
	Port       string               `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
	Backend    *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	SecretName string               `protobuf:"bytes,3,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	PEMName    string               `protobuf:"bytes,4,opt,name=PEM_name,json=pEMName" json:"PEM_name,omitempty"`
}

func (m *TCPLoadBalancerRule) Reset()                    { *m = TCPLoadBalancerRule{} }
func (m *TCPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*TCPLoadBalancerRule) ProtoMessage()               {}
func (*TCPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *TCPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "loadbalancer.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "loadbalancer.ListResponse")
	proto.RegisterType((*DescribeRequest)(nil), "loadbalancer.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "loadbalancer.DescribeResponse")
	proto.RegisterType((*CreateRequest)(nil), "loadbalancer.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "loadbalancer.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "loadbalancer.DeleteRequest")
	proto.RegisterType((*LoadBalancer)(nil), "loadbalancer.LoadBalancer")
	proto.RegisterType((*Spec)(nil), "loadbalancer.Spec")
	proto.RegisterType((*Status)(nil), "loadbalancer.Status")
	proto.RegisterType((*LoadBalancerStatus)(nil), "loadbalancer.LoadBalancerStatus")
	proto.RegisterType((*LoadBalancerBackend)(nil), "loadbalancer.LoadBalancerBackend")
	proto.RegisterType((*LoadBalancerRule)(nil), "loadbalancer.LoadBalancerRule")
	proto.RegisterType((*HTTPLoadBalancerRule)(nil), "loadbalancer.HTTPLoadBalancerRule")
	proto.RegisterType((*TCPLoadBalancerRule)(nil), "loadbalancer.TCPLoadBalancerRule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for LoadBalancers service

type LoadBalancersClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type loadBalancersClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancersClient(cc *grpc.ClientConn) LoadBalancersClient {
	return &loadBalancersClient{cc}
}

func (c *loadBalancersClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.LoadBalancers/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.LoadBalancers/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.LoadBalancers/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.LoadBalancers/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/loadbalancer.LoadBalancers/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadBalancers service

type LoadBalancersServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	Create(context.Context, *CreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *UpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterLoadBalancersServer(s *grpc.Server, srv LoadBalancersServer) {
	s.RegisterService(&_LoadBalancers_serviceDesc, srv)
}

func _LoadBalancers_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.LoadBalancers/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.LoadBalancers/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.LoadBalancers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.LoadBalancers/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadbalancer.LoadBalancers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loadbalancer.LoadBalancers",
	HandlerType: (*LoadBalancersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LoadBalancers_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _LoadBalancers_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LoadBalancers_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LoadBalancers_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LoadBalancers_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 884 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0x9d, 0xbf, 0xed, 0x49, 0xb2, 0x5b, 0x66, 0x23, 0x48, 0xbd, 0xdd, 0xd2, 0x58, 0x68,
	0x8b, 0x2a, 0x48, 0xa0, 0x45, 0xa8, 0x2a, 0x48, 0x88, 0xfe, 0x48, 0x20, 0xa5, 0x10, 0x39, 0x81,
	0x0b, 0x90, 0x88, 0x1c, 0x67, 0xd4, 0x5a, 0x75, 0x63, 0xd7, 0x33, 0x69, 0x55, 0xaa, 0x72, 0xc1,
	0x2b, 0xf4, 0x96, 0x2b, 0x5e, 0xa9, 0x37, 0x3c, 0x00, 0x4f, 0xc1, 0x15, 0xf3, 0x9b, 0xd8, 0x49,
	0xdc, 0x52, 0x15, 0xed, 0x4d, 0x3b, 0x3e, 0xe7, 0xcc, 0x77, 0xbe, 0xef, 0xcc, 0x99, 0x39, 0x01,
	0x14, 0x84, 0xee, 0x70, 0xe0, 0x06, 0xee, 0xc8, 0xc3, 0x71, 0x33, 0x8a, 0x43, 0x1a, 0xa2, 0x4a,
	0xd2, 0x66, 0xad, 0x1e, 0x87, 0xe1, 0x71, 0x80, 0x5b, 0x6e, 0xe4, 0xb7, 0xdc, 0xd1, 0x28, 0xa4,
	0x2e, 0xf5, 0xc3, 0x11, 0x91, 0xb1, 0xd6, 0xbb, 0xdc, 0x3c, 0xa4, 0x57, 0x11, 0x26, 0x2d, 0xf1,
	0x57, 0xda, 0xed, 0x0d, 0x28, 0xb7, 0x7d, 0x42, 0x1d, 0x7c, 0x3e, 0xc6, 0x84, 0xa2, 0x3a, 0x94,
	0xbc, 0x60, 0x4c, 0x28, 0x8e, 0xeb, 0xc6, 0xba, 0xf1, 0xe1, 0x92, 0xa3, 0x3f, 0xed, 0x4b, 0xa8,
	0xc8, 0x40, 0x12, 0x31, 0x54, 0x8c, 0xde, 0x40, 0x91, 0xb0, 0x14, 0x63, 0x22, 0x02, 0xcb, 0x5b,
	0xcf, 0x9b, 0x12, 0xbd, 0xd9, 0x15, 0x56, 0x47, 0x79, 0xd1, 0x57, 0x50, 0xe5, 0x34, 0xfb, 0x9a,
	0x67, 0xdd, 0x5c, 0xcf, 0xb1, 0x70, 0xab, 0x99, 0x12, 0xd4, 0x66, 0x1f, 0x7b, 0xea, 0xc3, 0x11,
	0xba, 0xf4, 0x97, 0x7d, 0x0e, 0x2f, 0x0e, 0x30, 0xf1, 0x62, 0x7f, 0x80, 0x35, 0x4b, 0x04, 0xf9,
	0x53, 0x7f, 0x34, 0x54, 0x14, 0xc5, 0x9a, 0xdb, 0x46, 0xee, 0x19, 0x66, 0xf0, 0xc2, 0xc6, 0xd7,
	0x68, 0x15, 0x96, 0xf8, 0x7f, 0x12, 0xb9, 0x1e, 0xae, 0xe7, 0x84, 0x63, 0x6a, 0x48, 0x6a, 0xcd,
	0xa7, 0xb5, 0x5e, 0xc3, 0xf2, 0x34, 0xe5, 0xd3, 0xf5, 0x1a, 0x8f, 0xd2, 0xfb, 0x87, 0x01, 0xd5,
	0xfd, 0x18, 0xbb, 0x34, 0x29, 0x57, 0x48, 0x33, 0xb2, 0xa4, 0x99, 0xf7, 0x48, 0xcb, 0xa5, 0xa4,
	0xcd, 0xd3, 0xcb, 0x3f, 0x92, 0xde, 0x6f, 0x50, 0xfd, 0x21, 0x1a, 0x3e, 0xc0, 0x2e, 0x91, 0xdf,
	0x7c, 0x20, 0x7f, 0xee, 0x91, 0xf9, 0x43, 0xa8, 0x1e, 0xe0, 0x00, 0xd3, 0xb7, 0xd6, 0x0c, 0x7f,
	0x99, 0xac, 0xf3, 0x13, 0x0c, 0xfe, 0xa7, 0x84, 0x1f, 0x03, 0xf2, 0xf8, 0x29, 0xb3, 0x3b, 0xda,
	0xa7, 0x3e, 0xb3, 0x52, 0xf7, 0x2c, 0x52, 0xb9, 0xdf, 0xd1, 0x9e, 0x9e, 0x76, 0xa0, 0xaf, 0xa1,
	0x14, 0x46, 0xe2, 0x42, 0xd7, 0x0b, 0xe2, 0x02, 0x6d, 0x64, 0x57, 0xac, 0xf9, 0xbd, 0x8c, 0x3c,
	0x1c, 0xd1, 0xf8, 0xca, 0xd1, 0xfb, 0x58, 0x07, 0xe7, 0x49, 0x84, 0xbd, 0x7a, 0x51, 0x54, 0x1c,
	0xa5, 0xf7, 0x77, 0x99, 0xc7, 0x11, 0x7e, 0xf4, 0xd1, 0xa4, 0xd3, 0x4b, 0x22, 0xb2, 0x36, 0x13,
	0x99, 0xea, 0x77, 0x6b, 0x17, 0x2a, 0xc9, 0x74, 0x68, 0x19, 0x72, 0xa7, 0xf8, 0x4a, 0x15, 0x87,
	0x2f, 0x51, 0x0d, 0x0a, 0x17, 0x6e, 0x30, 0xd6, 0xc5, 0x91, 0x1f, 0xbb, 0xe6, 0x8e, 0x61, 0xff,
	0x0a, 0x79, 0x9e, 0x17, 0x7d, 0x09, 0xa5, 0x81, 0xeb, 0x9d, 0x62, 0x55, 0xd4, 0xf2, 0x96, 0x9d,
	0x4e, 0xf9, 0x4d, 0xaf, 0xd7, 0x49, 0xb5, 0xc4, 0x38, 0xc0, 0x8e, 0xde, 0x82, 0x3e, 0x83, 0x42,
	0xcc, 0x0c, 0x44, 0xbd, 0x2c, 0x6b, 0xf7, 0xb4, 0x12, 0xdf, 0x27, 0x83, 0xed, 0x3d, 0x28, 0x4a,
	0x25, 0x68, 0x27, 0x71, 0xb3, 0x39, 0xc0, 0x7a, 0x36, 0x40, 0x5a, 0xbb, 0xbd, 0x03, 0x68, 0xde,
	0x8b, 0x9e, 0x83, 0xf9, 0x6d, 0x47, 0x15, 0xc0, 0xf4, 0x3b, 0xbc, 0x37, 0x4e, 0x42, 0x42, 0x75,
	0x6f, 0xf0, 0xb5, 0xfd, 0x33, 0xbc, 0x4c, 0xee, 0xdc, 0x53, 0x52, 0x1a, 0x50, 0x21, 0x38, 0xbe,
	0xf0, 0x3d, 0xdc, 0x4f, 0xdc, 0xa9, 0xb2, 0xb2, 0x7d, 0xc7, 0xbb, 0x2a, 0x11, 0x12, 0x85, 0xb1,
	0x46, 0xd5, 0x21, 0x1d, 0x66, 0xb2, 0x6f, 0x0d, 0x58, 0x9e, 0x95, 0x3d, 0x61, 0x61, 0x4c, 0x59,
	0xa0, 0xcf, 0x99, 0x8d, 0xd2, 0x48, 0x15, 0xee, 0xbf, 0x14, 0x5d, 0xc4, 0xa3, 0x6d, 0xc8, 0x51,
	0x2f, 0x62, 0x3d, 0xcd, 0xb7, 0x35, 0xd2, 0xdb, 0x7a, 0xfb, 0xf3, 0xbb, 0x78, 0xb4, 0x7d, 0x67,
	0x40, 0x6d, 0x11, 0x26, 0x67, 0x16, 0xb9, 0xf4, 0x44, 0x33, 0xe3, 0x6b, 0xf4, 0xc5, 0xb4, 0x23,
	0xe4, 0xfb, 0xd9, 0xc8, 0x3e, 0x14, 0x55, 0xbc, 0x69, 0x43, 0xbc, 0x0f, 0xe5, 0x13, 0xec, 0x0e,
	0x71, 0xdc, 0xe7, 0x47, 0x2d, 0x68, 0x2e, 0x39, 0x20, 0x4d, 0x22, 0x23, 0xab, 0x61, 0x8c, 0x2f,
	0x63, 0x9f, 0x62, 0x19, 0x91, 0x17, 0x11, 0x65, 0x65, 0x13, 0x21, 0x6f, 0xe0, 0x45, 0xb7, 0xdb,
	0xee, 0x13, 0xcc, 0xae, 0x22, 0x95, 0x87, 0x51, 0x10, 0xfc, 0xaa, 0xa4, 0xdb, 0xee, 0x0a, 0x2b,
	0x3f, 0x0e, 0xfb, 0x4f, 0x03, 0x5e, 0x2e, 0x90, 0x2c, 0x44, 0xf1, 0xe3, 0xd1, 0xa2, 0xd8, 0xfa,
	0xc9, 0xa2, 0x92, 0x64, 0xe4, 0x7b, 0x02, 0x64, 0xc2, 0x04, 0xad, 0xc0, 0xb3, 0xce, 0xe1, 0x91,
	0xf4, 0xaa, 0x27, 0x2c, 0x3a, 0x3c, 0xe2, 0xae, 0xad, 0x7f, 0x72, 0x50, 0x4d, 0x82, 0x13, 0xe4,
	0x41, 0x9e, 0x4f, 0x73, 0xb4, 0x32, 0xc3, 0x60, 0xfa, 0x53, 0xc0, 0xb2, 0x16, 0xb9, 0xe4, 0x30,
	0xb4, 0x3f, 0xf8, 0xfd, 0xee, 0xef, 0x5b, 0x73, 0x0d, 0xad, 0xb6, 0x92, 0x31, 0xad, 0x8b, 0x4f,
	0x9a, 0x9f, 0xb6, 0xae, 0xd5, 0xc3, 0x79, 0x83, 0x02, 0x78, 0xa6, 0xc7, 0x28, 0x7a, 0x9d, 0x46,
	0x9b, 0x99, 0xe8, 0xd6, 0x5a, 0x96, 0x5b, 0x25, 0x6c, 0x88, 0x84, 0xaf, 0xd0, 0xca, 0xa2, 0x84,
	0x5c, 0xf2, 0x0d, 0xfa, 0x09, 0x8a, 0x72, 0x6c, 0xa2, 0x57, 0x69, 0xb0, 0xd4, 0x30, 0xb5, 0x6a,
	0x7a, 0x6e, 0xff, 0x18, 0xfa, 0xc3, 0x09, 0xfe, 0x6b, 0x81, 0xff, 0x9e, 0x85, 0xe6, 0xf1, 0x77,
	0x8d, 0x4d, 0x8e, 0x2d, 0x87, 0xde, 0x2c, 0x76, 0x6a, 0x14, 0xde, 0x8f, 0x6d, 0x67, 0x60, 0xff,
	0x02, 0x45, 0x39, 0xd0, 0x66, 0xb1, 0x53, 0x63, 0x2e, 0x03, 0x5b, 0xd5, 0x65, 0x33, 0xbb, 0x2e,
	0x83, 0xa2, 0xf8, 0xa1, 0xb7, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x43, 0xc0, 0xd7,
	0x42, 0x0a, 0x00, 0x00,
}
