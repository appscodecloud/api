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

func _Volumes_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(VolumeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(VolumesServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Volumes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(VolumeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(VolumesServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Volumes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(VolumeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(VolumesServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x91, 0x65, 0xe4, 0x7a, 0x6c, 0x4a, 0xbd, 0x2d, 0x46, 0x08, 0xb7, 0xb4, 0x3a, 0x14,
	0xe3, 0x52, 0x6d, 0xb1, 0x6f, 0xa5, 0xf4, 0x62, 0x97, 0x5e, 0xda, 0x8b, 0x02, 0xbe, 0x06, 0x59,
	0xde, 0x38, 0x1b, 0x2b, 0x5a, 0x45, 0xbb, 0x32, 0x04, 0xa3, 0x4b, 0x6e, 0x39, 0xe4, 0x94, 0x5b,
	0x5e, 0x2b, 0xaf, 0x90, 0x07, 0x89, 0xb4, 0x2b, 0x29, 0x36, 0x8e, 0x21, 0x10, 0x92, 0x8b, 0xf1,
	0xcc, 0xfc, 0xcc, 0xf7, 0xef, 0xcc, 0x08, 0xda, 0x2b, 0x16, 0x24, 0xa7, 0xc4, 0x89, 0x62, 0x26,
	0x18, 0x6a, 0x2e, 0x82, 0x84, 0x0b, 0x12, 0x1f, 0x71, 0xab, 0xb7, 0x60, 0x6c, 0x11, 0x10, 0xec,
	0x45, 0x14, 0x7b, 0x61, 0xc8, 0x84, 0x27, 0x28, 0x0b, 0xb9, 0x12, 0x5a, 0xdd, 0x3c, 0x3d, 0x17,
	0xe7, 0x11, 0xe1, 0x58, 0xfe, 0xaa, 0xbc, 0x9d, 0xc2, 0xfb, 0xa9, 0x6c, 0x38, 0x8e, 0x89, 0x27,
	0x88, 0x4b, 0xce, 0x12, 0xc2, 0x05, 0xfa, 0x02, 0xed, 0x65, 0x32, 0x23, 0x87, 0xbe, 0x6a, 0x6f,
	0x6a, 0x9f, 0xb5, 0x7e, 0xd3, 0x6d, 0xe5, 0xb9, 0xb1, 0x4a, 0xa1, 0x6f, 0xd0, 0xa9, 0xe0, 0x95,
	0xae, 0x26, 0x75, 0xef, 0xaa, 0x42, 0x29, 0xee, 0x82, 0xa1, 0x7c, 0x9b, 0xba, 0x54, 0x14, 0xd1,
	0x03, 0x7e, 0x42, 0x02, 0xf2, 0xfa, 0x78, 0x1f, 0x3a, 0x0a, 0xff, 0x8f, 0x72, 0xf1, 0x42, 0x70,
	0x9b, 0x02, 0xda, 0x84, 0xf0, 0x28, 0xdb, 0x0a, 0x41, 0x5f, 0xc1, 0xe0, 0xd9, 0x8a, 0x12, 0x2e,
	0xfb, 0xb7, 0x86, 0x6f, 0x1d, 0xb5, 0x1d, 0xe7, 0x40, 0x66, 0xdd, 0xa2, 0x9a, 0xa1, 0x1a, 0xca,
	0x2c, 0xcf, 0x00, 0x7a, 0x26, 0xec, 0x38, 0x15, 0xc1, 0x51, 0x7d, 0xdd, 0x52, 0x61, 0x9f, 0x80,
	0xa1, 0x52, 0x08, 0x41, 0x3d, 0xf2, 0xc4, 0x71, 0x61, 0x5e, 0xfe, 0xcf, 0xa7, 0x30, 0x8b, 0xa9,
	0xbf, 0xe4, 0x85, 0xd5, 0x22, 0x42, 0x16, 0xbc, 0x21, 0xe1, 0x3c, 0x62, 0x34, 0x14, 0xc5, 0x7c,
	0xaa, 0x18, 0x99, 0xd0, 0x88, 0x49, 0x14, 0x50, 0xdf, 0x33, 0xeb, 0x59, 0x49, 0x77, 0xcb, 0x70,
	0x78, 0xa3, 0x43, 0x43, 0xc1, 0x38, 0xba, 0xd4, 0xa0, 0x9e, 0xbf, 0x0e, 0xf5, 0x76, 0xcc, 0x6d,
	0x4c, 0xd6, 0xfa, 0xb8, 0xa7, 0xaa, 0x46, 0x62, 0x4f, 0x2e, 0x6e, 0xef, 0xae, 0x6b, 0xbf, 0xd1,
	0x2f, 0x79, 0xc3, 0x95, 0x14, 0xaf, 0x7e, 0xe0, 0xe2, 0x81, 0x78, 0xbd, 0xb9, 0x99, 0x14, 0xaf,
	0x77, 0xb6, 0x90, 0x22, 0x0e, 0x86, 0xba, 0x65, 0xf4, 0x69, 0x07, 0xb7, 0x75, 0xe4, 0xd6, 0x87,
	0x72, 0xe4, 0x53, 0x46, 0xe7, 0x95, 0x8b, 0x91, 0x74, 0xf1, 0xdd, 0xea, 0x3f, 0xd5, 0xc5, 0x4f,
	0x6d, 0x80, 0xae, 0x34, 0x30, 0xd4, 0x09, 0x3f, 0x42, 0xdd, 0xba, 0xed, 0x3d, 0xd4, 0xff, 0x92,
	0xfa, 0x77, 0xf0, 0xe7, 0x39, 0x6f, 0xc7, 0x6b, 0xa5, 0x4d, 0x67, 0x86, 0xfc, 0xba, 0x47, 0xf7,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x05, 0x18, 0x0a, 0x2e, 0x04, 0x00, 0x00,
}
