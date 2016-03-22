// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

/*
Package glusterfs is a generated protocol buffer package.

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
	VolumeCreateRequest
	VolumeDeleteRequest
	VolumeListRequest
	VolumeListResponse
	Volume
*/
package glusterfs

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

type Cluster struct {
	Phid          string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	KubeCluster   string `protobuf:"bytes,3,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Mood          string `protobuf:"bytes,4,opt,name=mood" json:"mood,omitempty"`
	KubeNamespace string `protobuf:"bytes,5,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	Replica       int64  `protobuf:"varint,6,opt,name=replica" json:"replica,omitempty"`
	Created       string `protobuf:"bytes,7,opt,name=created" json:"created,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClusterListRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
}

func (m *ClusterListRequest) Reset()                    { *m = ClusterListRequest{} }
func (m *ClusterListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterListRequest) ProtoMessage()               {}
func (*ClusterListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClusterListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Glusterfs []*Cluster     `protobuf:"bytes,2,rep,name=glusterfs" json:"glusterfs,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClusterListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetGlusterfs() []*Cluster {
	if m != nil {
		return m.Glusterfs
	}
	return nil
}

type ClusterDescribeRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDescribeRequest) Reset()                    { *m = ClusterDescribeRequest{} }
func (m *ClusterDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeRequest) ProtoMessage()               {}
func (*ClusterDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ClusterDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ClusterDescribeResponse) Reset()                    { *m = ClusterDescribeResponse{} }
func (m *ClusterDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse) ProtoMessage()               {}
func (*ClusterDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClusterDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterCreateRequest struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Node        int64    `protobuf:"varint,2,opt,name=node" json:"node,omitempty"`
	Mood        string   `protobuf:"bytes,3,opt,name=mood" json:"mood,omitempty"`
	Disks       []string `protobuf:"bytes,4,rep,name=disks" json:"disks,omitempty"`
	KubeCluster string   `protobuf:"bytes,5,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ClusterDeleteRequest struct {
	KubeCluster string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDeleteRequest) Reset()                    { *m = ClusterDeleteRequest{} }
func (m *ClusterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDeleteRequest) ProtoMessage()               {}
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Cluster)(nil), "glusterfs.Cluster")
	proto.RegisterType((*ClusterListRequest)(nil), "glusterfs.ClusterListRequest")
	proto.RegisterType((*ClusterListResponse)(nil), "glusterfs.ClusterListResponse")
	proto.RegisterType((*ClusterDescribeRequest)(nil), "glusterfs.ClusterDescribeRequest")
	proto.RegisterType((*ClusterDescribeResponse)(nil), "glusterfs.ClusterDescribeResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "glusterfs.ClusterCreateRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "glusterfs.ClusterDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Clusters service

type ClustersClient interface {
	// Glusterfs cluster list. Needs to work with two modes.
	// First is to list all the glusterfs cluster through the
	// space with out considering the kubernetes cluster. if the
	// cluster_name is provided then list all the glusterfs cluster
	// with respect to the provided kube cluster space.
	List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error)
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type clustersClient struct {
	cc *grpc.ClientConn
}

func NewClustersClient(cc *grpc.ClientConn) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) List(ctx context.Context, in *ClusterListRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error) {
	out := new(ClusterDescribeResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Clusters/Delete", in, out, c.cc, opts...)
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
	Create(context.Context, *ClusterCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *ClusterDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterClustersServer(s *grpc.Server, srv ClustersServer) {
	s.RegisterService(&_Clusters_serviceDesc, srv)
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Clusters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glusterfs.Clusters",
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
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x95, 0xeb, 0xc4, 0x69, 0x6f, 0x69, 0x17, 0x97, 0xa8, 0x58, 0x16, 0xaf, 0x5a, 0x02, 0x41,
	0x91, 0xec, 0x10, 0x90, 0x40, 0xec, 0x50, 0xbb, 0xe4, 0x21, 0x19, 0x89, 0x2d, 0x72, 0xec, 0x4b,
	0xb0, 0x9a, 0x7a, 0x8c, 0x67, 0x8c, 0xa8, 0xaa, 0x6e, 0x10, 0x3b, 0x96, 0x7c, 0x12, 0x9f, 0xc0,
	0x2f, 0xf0, 0x21, 0x78, 0x1e, 0x76, 0x1d, 0xdc, 0xa2, 0x54, 0xdd, 0x44, 0x33, 0xe7, 0xbe, 0xce,
	0x9c, 0x73, 0x63, 0xd8, 0x4a, 0x16, 0x15, 0x17, 0x54, 0x06, 0x45, 0xc9, 0x04, 0xc3, 0x8d, 0xb9,
	0xbe, 0x7e, 0xe4, 0xde, 0xcd, 0x39, 0x63, 0xf3, 0x05, 0x85, 0x71, 0x91, 0x85, 0x71, 0x9e, 0x33,
	0x11, 0x8b, 0x8c, 0xe5, 0x5c, 0x27, 0x7a, 0x3b, 0x12, 0x4e, 0xc5, 0x71, 0x41, 0x3c, 0x54, 0xbf,
	0x1a, 0xf7, 0x7f, 0x59, 0x30, 0xda, 0xd7, 0x3d, 0x10, 0x61, 0x50, 0x7c, 0xca, 0x52, 0xd7, 0xba,
	0x6b, 0x3d, 0xd8, 0x88, 0xd4, 0x59, 0x62, 0x79, 0x7c, 0x44, 0xee, 0x9a, 0xc6, 0xe4, 0x19, 0x77,
	0xe1, 0xda, 0x61, 0x35, 0xa3, 0x0f, 0x86, 0x8a, 0x6b, 0xab, 0xd8, 0xa6, 0xc4, 0x3a, 0xad, 0x8e,
	0x18, 0x4b, 0xdd, 0x81, 0x2e, 0x93, 0x67, 0xbc, 0x07, 0xdb, 0xaa, 0x4c, 0xf6, 0xe0, 0x45, 0x9c,
	0x90, 0x3b, 0x54, 0xd1, 0x2d, 0x89, 0xbe, 0x69, 0x40, 0x74, 0x61, 0x54, 0x52, 0xb1, 0xc8, 0x92,
	0xd8, 0x75, 0xea, 0xb8, 0x1d, 0x35, 0x57, 0x19, 0x49, 0x4a, 0x8a, 0x05, 0xa5, 0xee, 0x48, 0x55,
	0x36, 0x57, 0xff, 0x19, 0xa0, 0x99, 0xfc, 0x2a, 0xe3, 0x22, 0xa2, 0xcf, 0x15, 0x71, 0xd1, 0xe3,
	0x69, 0xf5, 0x78, 0xfa, 0x0c, 0xae, 0x2f, 0x15, 0xf2, 0xa2, 0x96, 0x8c, 0xf0, 0x3e, 0x38, 0xbc,
	0xd6, 0xaf, 0xe2, 0xaa, 0x66, 0x73, 0xba, 0x1d, 0x68, 0xe9, 0x82, 0x77, 0x0a, 0x8d, 0x4c, 0x14,
	0x27, 0x70, 0x66, 0x40, 0x2d, 0x91, 0x5d, 0xa7, 0x62, 0xd0, 0x22, 0x81, 0x69, 0x1d, 0x9d, 0x25,
	0xf9, 0x6f, 0x61, 0xc7, 0xa0, 0x07, 0xc4, 0x93, 0x32, 0x9b, 0xd1, 0xea, 0x6c, 0xcf, 0x33, 0xc3,
	0x7f, 0x09, 0x37, 0x7a, 0x0d, 0x2f, 0xf7, 0x0a, 0xff, 0x87, 0x05, 0x63, 0xd3, 0x63, 0x5f, 0x09,
	0xda, 0x50, 0x6a, 0xe6, 0x59, 0x1d, 0xf3, 0x25, 0xc6, 0x52, 0xcd, 0xc1, 0x8e, 0xd4, 0xb9, 0x75,
	0xdb, 0xee, 0xb8, 0x3d, 0x86, 0x61, 0x9a, 0xf1, 0x43, 0x5e, 0xaf, 0x80, 0x5d, 0x83, 0xfa, 0xd2,
	0x7b, 0xe4, 0xb0, 0x6f, 0xc9, 0xeb, 0x96, 0xcc, 0x01, 0x2d, 0x48, 0x5c, 0x51, 0x9f, 0xe9, 0xf7,
	0x01, 0xac, 0x9b, 0x38, 0xc7, 0xaf, 0x30, 0x90, 0x3e, 0xe3, 0xad, 0xbe, 0x49, 0x9d, 0xc5, 0xf1,
	0x6e, 0x5f, 0x14, 0xd6, 0xc2, 0xfa, 0x8f, 0xbf, 0xfd, 0xfe, 0xf3, 0x73, 0xed, 0x11, 0x3e, 0x54,
	0x7f, 0xb6, 0x36, 0x37, 0xfc, 0x32, 0x09, 0x0d, 0x3b, 0x1e, 0x9e, 0x74, 0xc9, 0x9e, 0x62, 0xad,
	0xf1, 0x7a, 0x63, 0x10, 0xee, 0xf6, 0xfb, 0xff, 0xb3, 0x0d, 0x9e, 0xff, 0xbf, 0x14, 0x43, 0xe3,
	0xb9, 0xa2, 0x31, 0xc5, 0xc9, 0xca, 0x34, 0xc2, 0x13, 0xa9, 0xc9, 0x29, 0x56, 0xe0, 0x68, 0xa7,
	0xf1, 0x4e, 0x7f, 0xce, 0xd2, 0x0e, 0x78, 0xe3, 0x66, 0x69, 0xde, 0xb3, 0x2c, 0x6d, 0x47, 0x3f,
	0x55, 0xa3, 0x03, 0x6f, 0x75, 0x05, 0x5e, 0x58, 0x7b, 0x78, 0x0c, 0x8e, 0xf6, 0xf4, 0xbc, 0xb1,
	0x4b, 0x6e, 0x5f, 0x30, 0xd6, 0xbc, 0x78, 0xef, 0xd2, 0x2f, 0x9e, 0x39, 0xea, 0x73, 0xf7, 0xe4,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xbc, 0xb5, 0x6f, 0x40, 0x05, 0x00, 0x00,
}
