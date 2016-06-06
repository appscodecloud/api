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
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
}

func (m *VolumeCreateRequest) Reset()                    { *m = VolumeCreateRequest{} }
func (m *VolumeCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeCreateRequest) ProtoMessage()               {}
func (*VolumeCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VolumeDeleteRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
}

func (m *VolumeDeleteRequest) Reset()                    { *m = VolumeDeleteRequest{} }
func (m *VolumeDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeDeleteRequest) ProtoMessage()               {}
func (*VolumeDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type VolumeListRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster" json:"glusterfs_cluster,omitempty"`
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

var fileDescriptor1 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x86, 0x65, 0x8c, 0x4c, 0x99, 0x22, 0x5a, 0xb6, 0xa8, 0xa2, 0x16, 0x45, 0x95, 0x4f, 0x15,
	0x07, 0x6f, 0x0b, 0x97, 0xaa, 0x87, 0x5e, 0x8a, 0x54, 0x21, 0x55, 0x4a, 0x44, 0x24, 0xc4, 0x2d,
	0x32, 0xf6, 0x86, 0x58, 0x38, 0xde, 0x8d, 0x77, 0x8d, 0x14, 0x21, 0x72, 0xc8, 0x2b, 0x24, 0x87,
	0xbc, 0x4b, 0x1e, 0x23, 0xaf, 0x90, 0x07, 0x89, 0xd9, 0x21, 0x0e, 0x88, 0x70, 0x41, 0x5c, 0x10,
	0x9e, 0x19, 0x7f, 0xf3, 0xcf, 0xaf, 0xdf, 0x50, 0x99, 0xf1, 0x28, 0xbd, 0x60, 0xae, 0x48, 0xb8,
	0xe2, 0xa4, 0x3c, 0x89, 0x52, 0xa9, 0x58, 0x72, 0x26, 0xed, 0xe6, 0x84, 0xf3, 0x49, 0xc4, 0xa8,
	0x27, 0x42, 0xea, 0xc5, 0x31, 0x57, 0x9e, 0x0a, 0x79, 0x2c, 0x71, 0xd0, 0xfe, 0xbc, 0x2c, 0x07,
	0xea, 0x4a, 0x30, 0x49, 0xf5, 0x2f, 0xd6, 0x9d, 0x21, 0x7c, 0x1a, 0x6a, 0xe0, 0xdf, 0x84, 0x79,
	0x8a, 0x0d, 0xd8, 0x65, 0xca, 0xa4, 0x22, 0x75, 0xa8, 0x4c, 0xd3, 0x31, 0x3b, 0xf5, 0x11, 0xdf,
	0x30, 0xbe, 0x19, 0xdf, 0xcb, 0xe4, 0x0b, 0xd4, 0xf2, 0x7d, 0x79, 0xab, 0xa0, 0x5b, 0x55, 0xb0,
	0x50, 0x58, 0xc3, 0x5c, 0x3e, 0xbf, 0x72, 0x7b, 0x2c, 0x62, 0x07, 0xe4, 0xf6, 0xa0, 0x86, 0xdc,
	0xff, 0xa1, 0x54, 0xfb, 0x52, 0x9d, 0x11, 0x90, 0x75, 0x8a, 0x14, 0x99, 0x51, 0x8c, 0xb4, 0xc0,
	0x92, 0x99, 0x6b, 0xa9, 0xd4, 0x80, 0xf7, 0x9d, 0xaa, 0x8b, 0x86, 0xb9, 0x27, 0xba, 0x4a, 0x1c,
	0x28, 0xa1, 0x16, 0x99, 0x61, 0xcc, 0x6c, 0xa0, 0xe6, 0xe6, 0x0b, 0x5c, 0xe4, 0x39, 0x7d, 0xb0,
	0xf0, 0x1f, 0xa9, 0x40, 0x51, 0x78, 0xea, 0x7c, 0x25, 0x26, 0xbb, 0x63, 0x9c, 0x84, 0xfe, 0x54,
	0xae, 0xee, 0xfa, 0x08, 0xef, 0x58, 0x1c, 0x08, 0x1e, 0xc6, 0x0a, 0x2f, 0x23, 0x1f, 0xa0, 0x94,
	0x30, 0x11, 0x85, 0xbe, 0xd7, 0x28, 0x66, 0x05, 0xb3, 0xf3, 0x60, 0x42, 0x09, 0x59, 0x92, 0xdc,
	0x19, 0x50, 0x5c, 0x6a, 0x25, 0xcd, 0xad, 0x95, 0x6b, 0x46, 0xd8, 0x5f, 0x77, 0x74, 0xf1, 0x40,
	0xe7, 0xe8, 0xe6, 0xf1, 0xe9, 0xb6, 0xd0, 0x27, 0xff, 0xb2, 0x90, 0x08, 0xe9, 0xf3, 0x00, 0xd3,
	0x92, 0xbf, 0x43, 0x67, 0x3f, 0xdc, 0x9f, 0x74, 0x75, 0x23, 0x9d, 0xaf, 0x7b, 0xba, 0xa0, 0xf3,
	0x2d, 0x33, 0x17, 0xe4, 0x1a, 0x2c, 0xcc, 0x0d, 0x69, 0x6d, 0x6d, 0xde, 0x08, 0x94, 0x5d, 0x7f,
	0xf1, 0x72, 0xc8, 0xc3, 0x20, 0x17, 0xf4, 0x47, 0x0b, 0xfa, 0x65, 0x77, 0xf7, 0x10, 0xf4, 0xdb,
	0x68, 0x93, 0x7b, 0x03, 0x2c, 0x0c, 0xd8, 0x1b, 0x02, 0x36, 0x92, 0xb7, 0x43, 0xc0, 0x48, 0x0b,
	0x18, 0xb4, 0x8f, 0x0f, 0xe4, 0x08, 0x9d, 0xe3, 0xec, 0x62, 0x6c, 0xe9, 0xef, 0xab, 0xfb, 0x1c,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0xe1, 0xa2, 0x67, 0xb0, 0x03, 0x00, 0x00,
}
