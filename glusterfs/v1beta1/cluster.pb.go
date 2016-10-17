// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	cluster.proto
	volume.proto

It has these top-level messages:
	Cluster
	ClusterListRequest
	ClusterListResponse
	ClusterDescribeRequest
	ClusterDescribeResponse
	ClusterCreateRequest
	ClusterDeleteRequest
	VolumeListRequest
	VolumeListResponse
	Volume
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

type Cluster struct {
	Phid          string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	KubeCluster   string `protobuf:"bytes,3,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string `protobuf:"bytes,4,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Replica       int32  `protobuf:"varint,5,opt,name=replica" json:"replica,omitempty"`
	Endpoint      string `protobuf:"bytes,6,opt,name=endpoint" json:"endpoint,omitempty"`
	CreatedAt     int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Status        string `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClusterListRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   READY
	//   DELETED
	Status []string `protobuf:"bytes,2,rep,name=status" json:"status,omitempty"`
}

func (m *ClusterListRequest) Reset()                    { *m = ClusterListRequest{} }
func (m *ClusterListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterListRequest) ProtoMessage()               {}
func (*ClusterListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClusterListResponse struct {
	Status   *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Clusters []*Cluster              `protobuf:"bytes,2,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type ClusterDescribeRequest struct {
	KubeCluster   string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDescribeRequest) Reset()                    { *m = ClusterDescribeRequest{} }
func (m *ClusterDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeRequest) ProtoMessage()               {}
func (*ClusterDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ClusterDescribeResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Glusterfs *Cluster                `protobuf:"bytes,2,opt,name=glusterfs" json:"glusterfs,omitempty"`
}

func (m *ClusterDescribeResponse) Reset()                    { *m = ClusterDescribeResponse{} }
func (m *ClusterDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse) ProtoMessage()               {}
func (*ClusterDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClusterDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterDescribeResponse) GetGlusterfs() *Cluster {
	if m != nil {
		return m.Glusterfs
	}
	return nil
}

type ClusterCreateRequest struct {
	Name          string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Node          int32    `protobuf:"varint,2,opt,name=node" json:"node,omitempty"`
	Disks         []string `protobuf:"bytes,3,rep,name=disks" json:"disks,omitempty"`
	KubeCluster   string   `protobuf:"bytes,4,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string   `protobuf:"bytes,5,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ClusterDeleteRequest struct {
	KubeCluster   string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDeleteRequest) Reset()                    { *m = ClusterDeleteRequest{} }
func (m *ClusterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDeleteRequest) ProtoMessage()               {}
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Cluster)(nil), "appscode.glusterfs.v1beta1.Cluster")
	proto.RegisterType((*ClusterListRequest)(nil), "appscode.glusterfs.v1beta1.ClusterListRequest")
	proto.RegisterType((*ClusterListResponse)(nil), "appscode.glusterfs.v1beta1.ClusterListResponse")
	proto.RegisterType((*ClusterDescribeRequest)(nil), "appscode.glusterfs.v1beta1.ClusterDescribeRequest")
	proto.RegisterType((*ClusterDescribeResponse)(nil), "appscode.glusterfs.v1beta1.ClusterDescribeResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "appscode.glusterfs.v1beta1.ClusterCreateRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "appscode.glusterfs.v1beta1.ClusterDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Clusters service

type ClustersClient interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error)
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type clustersClient struct {
	cc *grpc.ClientConn
}

func NewClustersClient(cc *grpc.ClientConn) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error) {
	out := new(ClusterDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Clusters/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clusters service

type ClustersServer interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(context.Context, *ClusterListRequest) (*ClusterListResponse, error)
	Describe(context.Context, *ClusterDescribeRequest) (*ClusterDescribeResponse, error)
	Create(context.Context, *ClusterCreateRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *ClusterDeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterClustersServer(s *grpc.Server, srv ClustersServer) {
	s.RegisterService(&_Clusters_serviceDesc, srv)
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).List(ctx, req.(*ClusterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Describe(ctx, req.(*ClusterDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Create(ctx, req.(*ClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.glusterfs.v1beta1.Clusters/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Delete(ctx, req.(*ClusterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clusters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.glusterfs.v1beta1.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Clusters_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Clusters_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Clusters_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Clusters_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xd6, 0xe5, 0xc3, 0x71, 0xdf, 0xb6, 0x0c, 0x47, 0xd5, 0x5a, 0x16, 0x85, 0x60, 0x84, 0x54,
	0x75, 0xb0, 0x69, 0xba, 0x21, 0x21, 0xd4, 0x34, 0x23, 0x82, 0xc8, 0x48, 0x0c, 0x0c, 0x54, 0x17,
	0xfb, 0x08, 0x56, 0x53, 0x9f, 0xf1, 0x5d, 0x2a, 0xa1, 0xaa, 0x4b, 0x17, 0x26, 0x26, 0x76, 0xfe,
	0x05, 0x03, 0x7f, 0x81, 0x01, 0x06, 0xfe, 0x02, 0x43, 0x7f, 0x06, 0xf2, 0xdd, 0xd9, 0xb1, 0x71,
	0x2b, 0xdc, 0x4a, 0x74, 0x89, 0xee, 0x9e, 0xf7, 0xf3, 0x79, 0xdf, 0xe7, 0x62, 0x58, 0x0d, 0x66,
	0x73, 0x2e, 0x68, 0xea, 0x26, 0x29, 0x13, 0x0c, 0xdb, 0x24, 0x49, 0x78, 0xc0, 0x42, 0xea, 0x4e,
	0x15, 0xfe, 0x96, 0xbb, 0xc7, 0x3b, 0x13, 0x2a, 0xc8, 0x8e, 0x7d, 0x67, 0xca, 0xd8, 0x74, 0x46,
	0x3d, 0x92, 0x44, 0x1e, 0x89, 0x63, 0x26, 0x88, 0x88, 0x58, 0xcc, 0x55, 0xa4, 0x7d, 0x37, 0x8f,
	0xbc, 0xc4, 0x7e, 0xaf, 0x62, 0x0f, 0xc5, 0x87, 0x84, 0x72, 0x4f, 0xfe, 0x2a, 0x07, 0xe7, 0x1c,
	0x41, 0x6f, 0x5f, 0x15, 0xc5, 0x18, 0x3a, 0xc9, 0xbb, 0x28, 0xb4, 0x50, 0x1f, 0x6d, 0x2d, 0xf9,
	0xf2, 0x9c, 0x61, 0x31, 0x39, 0xa2, 0x56, 0x4b, 0x61, 0xd9, 0x19, 0xdf, 0x87, 0x95, 0xc3, 0xf9,
	0x84, 0x1e, 0x68, 0x12, 0x56, 0x5b, 0xda, 0x96, 0x33, 0x2c, 0x4f, 0xf5, 0x10, 0x6e, 0x49, 0x97,
	0xcc, 0x9f, 0x27, 0x24, 0xa0, 0x56, 0x47, 0x3a, 0xad, 0x66, 0xe8, 0xf3, 0x1c, 0xc4, 0x16, 0xf4,
	0x52, 0x9a, 0xcc, 0xa2, 0x80, 0x58, 0xdd, 0x3e, 0xda, 0xea, 0xfa, 0xf9, 0x15, 0xdb, 0x60, 0xd2,
	0x38, 0x4c, 0x58, 0x14, 0x0b, 0xcb, 0x90, 0xa1, 0xc5, 0x1d, 0x6f, 0x02, 0x04, 0x29, 0x25, 0x82,
	0x86, 0x07, 0x44, 0x58, 0xbd, 0x3e, 0xda, 0x6a, 0xfb, 0x4b, 0x1a, 0xd9, 0x13, 0x78, 0x1d, 0x0c,
	0x2e, 0x88, 0x98, 0x73, 0xcb, 0x94, 0x81, 0xfa, 0xe6, 0xbc, 0x00, 0xac, 0xdb, 0x7b, 0x16, 0x71,
	0xe1, 0xd3, 0xf7, 0x73, 0xca, 0x45, 0x8d, 0x0c, 0xaa, 0x93, 0x59, 0x24, 0x6c, 0xf5, 0xdb, 0xa5,
	0x84, 0x1f, 0x11, 0xdc, 0xae, 0x64, 0xe4, 0x09, 0x8b, 0x39, 0xc5, 0x5e, 0xe1, 0x9f, 0x25, 0x5b,
	0x1e, 0x6c, 0xb8, 0xc5, 0x7e, 0xd5, 0x06, 0xdc, 0x97, 0xd2, 0x9c, 0x27, 0xc2, 0x4f, 0xc1, 0xd4,
	0xe5, 0x55, 0x89, 0xe5, 0xc1, 0x03, 0xf7, 0x72, 0x49, 0xb8, 0xba, 0xa6, 0x5f, 0x04, 0x39, 0xc7,
	0xb0, 0xae, 0xc1, 0x11, 0xe5, 0x41, 0x1a, 0x4d, 0xe8, 0x15, 0xe8, 0xd5, 0x77, 0xd5, 0xba, 0x68,
	0x57, 0xb9, 0x12, 0xda, 0x0b, 0x25, 0x38, 0x9f, 0x10, 0x6c, 0xd4, 0x0a, 0x5f, 0x77, 0x0a, 0x7b,
	0xb0, 0x54, 0x70, 0x95, 0x2d, 0x34, 0x1c, 0xc3, 0x22, 0xca, 0xf9, 0x82, 0x60, 0x4d, 0xc3, 0xfb,
	0x52, 0x0f, 0xf9, 0x18, 0xf2, 0xe6, 0x51, 0x49, 0xc6, 0x19, 0xc6, 0x42, 0xc5, 0xb6, 0xeb, 0xcb,
	0x33, 0x5e, 0x83, 0x6e, 0x18, 0xf1, 0x43, 0x6e, 0xb5, 0xe5, 0xa6, 0xd5, 0xa5, 0x36, 0xc4, 0x4e,
	0x93, 0x21, 0x76, 0x2f, 0x18, 0xa2, 0x23, 0x8a, 0xfe, 0x46, 0x74, 0x46, 0xc5, 0xcd, 0xac, 0x69,
	0xf0, 0xd3, 0x00, 0x53, 0xa7, 0xe1, 0xf8, 0x1b, 0x82, 0x4e, 0x26, 0x57, 0xec, 0x36, 0x18, 0x6e,
	0xe9, 0xa5, 0xd8, 0x5e, 0x63, 0x7f, 0xa5, 0x00, 0x67, 0x7c, 0xf6, 0xd5, 0x6a, 0x99, 0xe8, 0xec,
	0xd7, 0xef, 0xcf, 0xad, 0x11, 0x1e, 0x7a, 0x95, 0xff, 0xa2, 0xac, 0xeb, 0x34, 0xa6, 0x82, 0x72,
	0x4f, 0xe7, 0xf0, 0x72, 0x39, 0x7b, 0x27, 0xe5, 0x59, 0x9c, 0x7a, 0x45, 0x35, 0x7c, 0x8e, 0xc0,
	0xcc, 0x85, 0x86, 0x07, 0x0d, 0xfa, 0xf9, 0xeb, 0x39, 0xd8, 0xbb, 0x57, 0x8a, 0xd1, 0x3c, 0x78,
	0x89, 0xc7, 0x14, 0xd3, 0x6b, 0xf3, 0x28, 0x36, 0x97, 0x5b, 0x0a, 0xa0, 0xc4, 0xd1, 0x3b, 0xc9,
	0xd0, 0x53, 0xfc, 0x1d, 0x81, 0xa1, 0x34, 0x8c, 0x1f, 0x35, 0x68, 0xba, 0x22, 0x77, 0x7b, 0xb3,
	0xf6, 0xd6, 0x5e, 0xb1, 0x28, 0x2c, 0x08, 0xb1, 0x12, 0xa1, 0xc0, 0x79, 0xf3, 0x7f, 0x09, 0x3d,
	0x46, 0xdb, 0xf8, 0x07, 0x02, 0x43, 0x09, 0xbe, 0x11, 0x99, 0xca, 0xdb, 0xf8, 0x17, 0x99, 0xca,
	0x76, 0xb6, 0x6f, 0x66, 0x3b, 0xc3, 0x27, 0xe0, 0x04, 0xec, 0x68, 0xd1, 0x18, 0x49, 0xa2, 0x3a,
	0x9d, 0xe1, 0x8a, 0xe6, 0x33, 0xce, 0x3e, 0xb5, 0x63, 0xf4, 0xba, 0xa7, 0x0d, 0x13, 0x43, 0x7e,
	0x7c, 0x77, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x75, 0xd9, 0x28, 0xee, 0x08, 0x08, 0x00, 0x00,
}
