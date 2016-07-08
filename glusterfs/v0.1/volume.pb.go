// Code generated by protoc-gen-go.
// source: volume.proto
// DO NOT EDIT!

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

type VolumeCreateRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
	EnableBacula     int32  `protobuf:"varint,4,opt,name=enable_bacula,json=enableBacula" json:"enable_bacula,omitempty"`
}

func (m *VolumeCreateRequest) Reset()                    { *m = VolumeCreateRequest{} }
func (m *VolumeCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeCreateRequest) ProtoMessage()               {}
func (*VolumeCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VolumeDeleteRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
}

func (m *VolumeDeleteRequest) Reset()                    { *m = VolumeDeleteRequest{} }
func (m *VolumeDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeDeleteRequest) ProtoMessage()               {}
func (*VolumeDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type VolumeListRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
}

func (m *VolumeListRequest) Reset()                    { *m = VolumeListRequest{} }
func (m *VolumeListRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeListRequest) ProtoMessage()               {}
func (*VolumeListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type VolumeListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Volumes []*Volume      `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *VolumeListResponse) Reset()                    { *m = VolumeListResponse{} }
func (m *VolumeListResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeListResponse) ProtoMessage()               {}
func (*VolumeListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *VolumeListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VolumeListResponse) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type Volume struct {
	Path     string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Bricks   string `protobuf:"bytes,2,opt,name=bricks" json:"bricks,omitempty"`
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
	Replica  int64  `protobuf:"varint,4,opt,name=replica" json:"replica,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*VolumeCreateRequest)(nil), "glusterfs.VolumeCreateRequest")
	proto.RegisterType((*VolumeDeleteRequest)(nil), "glusterfs.VolumeDeleteRequest")
	proto.RegisterType((*VolumeListRequest)(nil), "glusterfs.VolumeListRequest")
	proto.RegisterType((*VolumeListResponse)(nil), "glusterfs.VolumeListResponse")
	proto.RegisterType((*Volume)(nil), "glusterfs.Volume")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Volumes service

type VolumesClient interface {
	List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error)
	Create(ctx context.Context, in *VolumeCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type volumesClient struct {
	cc *grpc.ClientConn
}

func NewVolumesClient(cc *grpc.ClientConn) VolumesClient {
	return &volumesClient{cc}
}

func (c *volumesClient) List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error) {
	out := new(VolumeListResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) Create(ctx context.Context, in *VolumeCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Volumes service

type VolumesServer interface {
	List(context.Context, *VolumeListRequest) (*VolumeListResponse, error)
	Create(context.Context, *VolumeCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *VolumeDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterVolumesServer(s *grpc.Server, srv VolumesServer) {
	s.RegisterService(&_Volumes_serviceDesc, srv)
}

func _Volumes_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).List(ctx, req.(*VolumeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).Create(ctx, req.(*VolumeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).Delete(ctx, req.(*VolumeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Volumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glusterfs.Volumes",
	HandlerType: (*VolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Volumes_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Volumes_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Volumes_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x25, 0x3b, 0x63, 0xc6, 0xad, 0x1d, 0xc5, 0x29, 0x65, 0x19, 0xc2, 0x2a, 0x1a, 0x41, 0x64,
	0x85, 0x44, 0x77, 0x2f, 0xe2, 0xc1, 0x83, 0x2b, 0x88, 0x20, 0x28, 0x11, 0x16, 0x6f, 0x4b, 0x27,
	0x53, 0x8e, 0xed, 0xc6, 0x74, 0x9b, 0xee, 0x2c, 0xc8, 0x30, 0x1e, 0xfc, 0x05, 0x3d, 0x78, 0xf6,
	0x37, 0xfc, 0x0c, 0x7f, 0xc1, 0x0f, 0x31, 0xe9, 0x4a, 0xc2, 0x2c, 0xe3, 0x5e, 0x44, 0xbd, 0x84,
	0xd4, 0xab, 0x97, 0x7a, 0xaf, 0xab, 0x2a, 0x0d, 0xe3, 0x13, 0x95, 0x57, 0xef, 0x28, 0xd2, 0xa5,
	0xb2, 0x0a, 0x37, 0xe7, 0x79, 0x65, 0x2c, 0x95, 0xaf, 0x4d, 0xb0, 0x33, 0x57, 0x6a, 0x9e, 0x53,
	0x2c, 0xb4, 0x8c, 0x45, 0x51, 0x28, 0x2b, 0xac, 0x54, 0x85, 0x61, 0x62, 0xb0, 0xdd, 0xc0, 0x33,
	0xfb, 0x41, 0x93, 0x89, 0xdd, 0x93, 0xf1, 0xf0, 0x9b, 0x07, 0x97, 0x0f, 0x5d, 0xc5, 0x83, 0x92,
	0x84, 0xa5, 0x84, 0xde, 0x57, 0x64, 0x2c, 0xde, 0x80, 0xf1, 0x71, 0x95, 0xd2, 0x51, 0xc6, 0xf5,
	0xa7, 0xde, 0x75, 0xef, 0xf6, 0x66, 0xb2, 0xd5, 0x60, 0x07, 0x0c, 0xe1, 0x1d, 0x98, 0xf4, 0xea,
	0x3d, 0x6f, 0xc3, 0xf1, 0x2e, 0xf5, 0x89, 0x8e, 0xbc, 0x0d, 0x3e, 0x1b, 0x9f, 0x0e, 0x1c, 0xa3,
	0x8d, 0xf0, 0x26, 0x5c, 0xa0, 0x42, 0xa4, 0x39, 0x1d, 0xa5, 0x22, 0xab, 0x72, 0x31, 0x1d, 0xd6,
	0xe9, 0x73, 0xc9, 0x98, 0xc1, 0x47, 0x0e, 0x0b, 0x97, 0x9d, 0xc7, 0xc7, 0x94, 0xd3, 0x7f, 0xf7,
	0x18, 0x66, 0x30, 0x61, 0xf9, 0x67, 0xd2, 0xd8, 0x7f, 0x24, 0x1e, 0x4a, 0xc0, 0x55, 0x11, 0xa3,
	0xeb, 0xd9, 0x11, 0xde, 0x02, 0xdf, 0xd4, 0x83, 0xac, 0x8c, 0xab, 0xbf, 0xb5, 0x77, 0x31, 0xe2,
	0x19, 0x46, 0x2f, 0x1d, 0x9a, 0xb4, 0xd9, 0x5a, 0x6a, 0xc4, 0x66, 0x4d, 0x2d, 0x30, 0xa8, 0x89,
	0x93, 0xa8, 0x57, 0x88, 0xb8, 0x6e, 0xd2, 0x31, 0xc2, 0xb7, 0xe0, 0x33, 0x84, 0x08, 0x43, 0x2d,
	0xec, 0x9b, 0xd6, 0xbc, 0x7b, 0x6f, 0xba, 0x90, 0x96, 0x32, 0x3b, 0x36, 0xad, 0xd5, 0x36, 0xc2,
	0x00, 0xce, 0x53, 0x31, 0xd3, 0x4a, 0x16, 0xb6, 0xed, 0x4f, 0x1f, 0xe3, 0x14, 0x46, 0x25, 0xe9,
	0x5c, 0x66, 0x3c, 0xbf, 0x41, 0xd2, 0x85, 0x7b, 0xdf, 0x07, 0x30, 0x62, 0x31, 0x83, 0x5f, 0x3c,
	0x18, 0x36, 0xa7, 0xc3, 0x9d, 0x35, 0x73, 0x2b, 0x9d, 0x0d, 0xae, 0x9e, 0x91, 0xe5, 0x96, 0x84,
	0xcf, 0x3f, 0xfd, 0xf8, 0xf9, 0x79, 0xe3, 0x29, 0x3e, 0xa9, 0x37, 0x5d, 0x9b, 0x4c, 0xcd, 0x78,
	0xe5, 0xfb, 0x6f, 0xe2, 0x93, 0xbb, 0xd1, 0xbd, 0xb8, 0x3d, 0x6b, 0xbc, 0x58, 0x1d, 0xd2, 0x32,
	0x5e, 0xac, 0x0d, 0x64, 0x89, 0x1f, 0xc1, 0xe7, 0xdd, 0xc7, 0x6b, 0x6b, 0xca, 0xa7, 0x7e, 0x8a,
	0xe0, 0x4a, 0xd7, 0xfd, 0x43, 0x25, 0x67, 0xbd, 0xa1, 0x87, 0xce, 0xd0, 0xfd, 0x60, 0xff, 0x0f,
	0x0c, 0x3d, 0xf0, 0x76, 0xf1, 0xab, 0x07, 0x3e, 0x2f, 0xf6, 0x6f, 0x0c, 0x9c, 0xda, 0xf8, 0x33,
	0x0c, 0xbc, 0x72, 0x06, 0x92, 0xdd, 0x17, 0x7f, 0xa9, 0x23, 0xf1, 0x82, 0xb9, 0xcb, 0xd4, 0x77,
	0x97, 0xc4, 0xfe, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x07, 0x2f, 0xef, 0x75, 0x04, 0x00,
	0x00,
}
