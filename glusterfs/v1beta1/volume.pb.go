// Code generated by protoc-gen-go.
// source: volume.proto
// DO NOT EDIT!

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

type VolumeListRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace    string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,3,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
}

func (m *VolumeListRequest) Reset()                    { *m = VolumeListRequest{} }
func (m *VolumeListRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeListRequest) ProtoMessage()               {}
func (*VolumeListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VolumeListResponse struct {
	Status  *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Volumes []*Volume               `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *VolumeListResponse) Reset()                    { *m = VolumeListResponse{} }
func (m *VolumeListResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeListResponse) ProtoMessage()               {}
func (*VolumeListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *VolumeListResponse) GetStatus() *appscode_dtypes.Status {
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
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func init() {
	proto.RegisterType((*VolumeListRequest)(nil), "appscode.glusterfs.v1beta1.VolumeListRequest")
	proto.RegisterType((*VolumeListResponse)(nil), "appscode.glusterfs.v1beta1.VolumeListResponse")
	proto.RegisterType((*Volume)(nil), "appscode.glusterfs.v1beta1.Volume")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Volumes service

type VolumesClient interface {
	List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error)
}

type volumesClient struct {
	cc *grpc.ClientConn
}

func NewVolumesClient(cc *grpc.ClientConn) VolumesClient {
	return &volumesClient{cc}
}

func (c *volumesClient) List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error) {
	out := new(VolumeListResponse)
	err := grpc.Invoke(ctx, "/appscode.glusterfs.v1beta1.Volumes/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Volumes service

type VolumesServer interface {
	List(context.Context, *VolumeListRequest) (*VolumeListResponse, error)
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
		FullMethod: "/appscode.glusterfs.v1beta1.Volumes/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).List(ctx, req.(*VolumeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Volumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.glusterfs.v1beta1.Volumes",
	HandlerType: (*VolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Volumes_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x6b, 0x1a, 0x41,
	0x18, 0x65, 0x56, 0xd1, 0x76, 0xd6, 0x96, 0x3a, 0x97, 0x2e, 0x8b, 0xb4, 0x76, 0xa1, 0x20, 0x94,
	0xce, 0xa0, 0xbd, 0x7a, 0xb2, 0xd7, 0x52, 0x64, 0x03, 0x39, 0xe4, 0x12, 0xc6, 0x75, 0x62, 0x96,
	0xe8, 0xce, 0xc4, 0x6f, 0x56, 0x48, 0x8c, 0x17, 0x73, 0x09, 0xb9, 0xe6, 0xb7, 0xe4, 0x97, 0xe4,
	0x2f, 0xe4, 0x07, 0xe4, 0x07, 0xe4, 0x10, 0x76, 0x66, 0x77, 0xa3, 0x88, 0x24, 0x17, 0x91, 0xf7,
	0xde, 0xf7, 0xed, 0x7b, 0xef, 0x1b, 0xdc, 0x58, 0xc8, 0x69, 0x3a, 0x13, 0x54, 0xcd, 0xa5, 0x96,
	0xc4, 0xe7, 0x4a, 0x41, 0x24, 0xc7, 0x82, 0x4e, 0xa6, 0x29, 0x68, 0x31, 0x3f, 0x01, 0xba, 0xe8,
	0x8e, 0x84, 0xe6, 0x5d, 0xbf, 0x35, 0x91, 0x72, 0x32, 0x15, 0x8c, 0xab, 0x98, 0xf1, 0x24, 0x91,
	0x9a, 0xeb, 0x58, 0x26, 0x60, 0x27, 0xfd, 0x6f, 0xc5, 0xe4, 0x1e, 0xfe, 0xfb, 0x16, 0x3f, 0xd6,
	0x17, 0x4a, 0x00, 0x33, 0xbf, 0x56, 0x10, 0xdc, 0x22, 0xdc, 0x3c, 0x34, 0x5e, 0xfe, 0xc5, 0xa0,
	0x43, 0x71, 0x9e, 0x0a, 0xd0, 0xe4, 0x07, 0x6e, 0x9c, 0xa5, 0x23, 0x71, 0x1c, 0x59, 0x3b, 0x1e,
	0x6a, 0xa3, 0xce, 0xc7, 0xd0, 0xcd, 0xb0, 0xbf, 0x16, 0x22, 0x3f, 0xf1, 0x67, 0x23, 0x49, 0xf8,
	0x4c, 0x80, 0xe2, 0x91, 0xf0, 0x1c, 0x23, 0xfa, 0x94, 0xa1, 0xff, 0x0b, 0x90, 0xfc, 0xc2, 0xcd,
	0x32, 0x53, 0xb9, 0xae, 0x62, 0x94, 0x5f, 0x4a, 0x22, 0xdf, 0x19, 0x5c, 0x23, 0x4c, 0x36, 0xcd,
	0x80, 0x92, 0x09, 0x08, 0xc2, 0x70, 0x0d, 0x34, 0xd7, 0x29, 0x18, 0x1f, 0x6e, 0xef, 0x2b, 0x2d,
	0xfb, 0xb2, 0x89, 0xe8, 0x81, 0xa1, 0xc3, 0x5c, 0x46, 0xfa, 0xb8, 0x6e, 0xfb, 0x05, 0xcf, 0x69,
	0x57, 0x3a, 0x6e, 0x2f, 0xa0, 0xfb, 0x1b, 0xa6, 0xf6, 0x8b, 0x61, 0x31, 0x12, 0xb4, 0x70, 0xcd,
	0x42, 0x84, 0xe0, 0xaa, 0xe2, 0xfa, 0x34, 0x8f, 0x6f, 0xfe, 0xf7, 0x9e, 0x11, 0xae, 0x5b, 0x1a,
	0xc8, 0x13, 0xc2, 0xd5, 0xcc, 0x29, 0xf9, 0xfd, 0xf6, 0xfe, 0x8d, 0x7a, 0x7d, 0xfa, 0x5e, 0xb9,
	0x2d, 0x20, 0xb8, 0x41, 0xeb, 0x7b, 0xcf, 0xf9, 0x80, 0xd6, 0x0f, 0x8f, 0x77, 0xce, 0x15, 0xb9,
	0x64, 0x5b, 0x57, 0xcd, 0x0a, 0x9f, 0x27, 0x42, 0x0b, 0x60, 0xf9, 0x0e, 0x96, 0x17, 0x0d, 0x6c,
	0xb9, 0x79, 0xc6, 0x15, 0x2b, 0x8f, 0x55, 0x30, 0x25, 0xb0, 0x62, 0xa5, 0x13, 0xb6, 0xdc, 0xb9,
	0xd8, 0x8a, 0xe5, 0xe5, 0x0c, 0xfa, 0x38, 0x88, 0xe4, 0xec, 0xd5, 0x3f, 0x57, 0xf1, 0x6e, 0x86,
	0x81, 0x6b, 0x43, 0x0c, 0xb3, 0x27, 0x36, 0x44, 0x47, 0xf5, 0x1c, 0x1f, 0xd5, 0xcc, 0xa3, 0xfb,
	0xf3, 0x12, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xf9, 0x3d, 0x08, 0xff, 0x02, 0x00, 0x00,
}
