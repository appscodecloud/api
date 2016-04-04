// Code generated by protoc-gen-go.
// source: pvc.proto
// DO NOT EDIT!

package pv

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

type PVCRegisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Size      int64  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCRegisterRequest) Reset()                    { *m = PVCRegisterRequest{} }
func (m *PVCRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCRegisterRequest) ProtoMessage()               {}
func (*PVCRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PVCUnregisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCUnregisterRequest) Reset()                    { *m = PVCUnregisterRequest{} }
func (m *PVCUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCUnregisterRequest) ProtoMessage()               {}
func (*PVCUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type PVCDescribeRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCDescribeRequest) Reset()                    { *m = PVCDescribeRequest{} }
func (m *PVCDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeRequest) ProtoMessage()               {}
func (*PVCDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type PVCInfo struct {
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size      int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Namespace string   `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Status    string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Volume    string   `protobuf:"bytes,5,opt,name=volume" json:"volume,omitempty"`
	Users     []string `protobuf:"bytes,6,rep,name=users" json:"users,omitempty"`
}

func (m *PVCInfo) Reset()                    { *m = PVCInfo{} }
func (m *PVCInfo) String() string            { return proto.CompactTextString(m) }
func (*PVCInfo) ProtoMessage()               {}
func (*PVCInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type PVCDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pvc    *PVCInfo       `protobuf:"bytes,2,opt,name=pvc" json:"pvc,omitempty"`
}

func (m *PVCDescribeResponse) Reset()                    { *m = PVCDescribeResponse{} }
func (m *PVCDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeResponse) ProtoMessage()               {}
func (*PVCDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PVCDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVCDescribeResponse) GetPvc() *PVCInfo {
	if m != nil {
		return m.Pvc
	}
	return nil
}

func init() {
	proto.RegisterType((*PVCRegisterRequest)(nil), "pv.PVCRegisterRequest")
	proto.RegisterType((*PVCUnregisterRequest)(nil), "pv.PVCUnregisterRequest")
	proto.RegisterType((*PVCDescribeRequest)(nil), "pv.PVCDescribeRequest")
	proto.RegisterType((*PVCInfo)(nil), "pv.PVCInfo")
	proto.RegisterType((*PVCDescribeResponse)(nil), "pv.PVCDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for PersistentVolumeClaims service

type PersistentVolumeClaimsClient interface {
	Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error)
	Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type persistentVolumeClaimsClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumeClaimsClient(cc *grpc.ClientConn) PersistentVolumeClaimsClient {
	return &persistentVolumeClaimsClient{cc}
}

func (c *persistentVolumeClaimsClient) Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error) {
	out := new(PVCDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumeClaims service

type PersistentVolumeClaimsServer interface {
	Describe(context.Context, *PVCDescribeRequest) (*PVCDescribeResponse, error)
	Register(context.Context, *PVCRegisterRequest) (*dtypes.VoidResponse, error)
	Unregister(context.Context, *PVCUnregisterRequest) (*dtypes.VoidResponse, error)
}

func RegisterPersistentVolumeClaimsServer(s *grpc.Server, srv PersistentVolumeClaimsServer) {
	s.RegisterService(&_PersistentVolumeClaims_serviceDesc, srv)
}

func _PersistentVolumeClaims_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PVCDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PersistentVolumeClaimsServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PersistentVolumeClaims_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PVCRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PersistentVolumeClaimsServer).Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PersistentVolumeClaims_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PVCUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PersistentVolumeClaimsServer).Unregister(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PersistentVolumeClaims_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.PersistentVolumeClaims",
	HandlerType: (*PersistentVolumeClaimsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _PersistentVolumeClaims_Describe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PersistentVolumeClaims_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumeClaims_Unregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0xeb, 0xd3, 0x30,
	0x18, 0xc7, 0x69, 0xbb, 0xbf, 0xcf, 0xc0, 0x43, 0x1c, 0xb5, 0x94, 0x09, 0xb3, 0x87, 0x21, 0x43,
	0x5b, 0x9d, 0x78, 0xf1, 0x3a, 0x2f, 0xde, 0x46, 0xc5, 0x9d, 0x3c, 0xd8, 0x75, 0xd9, 0x28, 0x74,
	0x4d, 0x6c, 0xd2, 0x82, 0x8e, 0x5d, 0x7c, 0x0b, 0xe2, 0x2b, 0xf3, 0x2d, 0x78, 0xf1, 0x5d, 0x98,
	0xa4, 0x69, 0xc7, 0x3a, 0x87, 0xc8, 0x8f, 0xdf, 0x65, 0xc9, 0xf3, 0x5d, 0x92, 0xcf, 0xf3, 0xe7,
	0x5b, 0x18, 0xd2, 0x32, 0xf6, 0x69, 0x4e, 0x38, 0x41, 0x26, 0x2d, 0xdd, 0xc9, 0x9e, 0x90, 0x7d,
	0x8a, 0x83, 0x88, 0x26, 0x41, 0x94, 0x65, 0x84, 0x47, 0x3c, 0x21, 0x19, 0xab, 0x4e, 0xb8, 0xb6,
	0x94, 0xb7, 0xfc, 0x0b, 0xc5, 0x2c, 0x50, 0xbf, 0x95, 0xee, 0x71, 0x40, 0xab, 0xf5, 0x32, 0xc4,
	0xfb, 0x84, 0x71, 0x9c, 0x87, 0xf8, 0x73, 0x81, 0x19, 0x47, 0x0e, 0xf4, 0xe3, 0xb4, 0x90, 0x8a,
	0x63, 0x4c, 0x8d, 0xa7, 0xc3, 0xb0, 0x0e, 0x11, 0x82, 0x4e, 0x16, 0x1d, 0xb0, 0x63, 0x2a, 0x59,
	0xed, 0xa5, 0xc6, 0x92, 0xaf, 0xd8, 0xb1, 0x84, 0x66, 0x85, 0x6a, 0x8f, 0x26, 0x30, 0x94, 0xff,
	0x31, 0x1a, 0xc5, 0xd8, 0xe9, 0xa8, 0xc3, 0x67, 0xc1, 0xdb, 0xc0, 0x58, 0x50, 0x3f, 0x64, 0xf9,
	0x9d, 0xb8, 0x17, 0x0c, 0xab, 0xcd, 0xf8, 0xa4, 0x2a, 0x7b, 0x8b, 0x59, 0x9c, 0x27, 0x1b, 0x7c,
	0x1f, 0x84, 0x1f, 0x06, 0xf4, 0x05, 0xe2, 0x5d, 0xb6, 0x23, 0xcd, 0x6d, 0xe3, 0x2f, 0x7d, 0x31,
	0x6f, 0xf5, 0xa5, 0xfd, 0x22, 0xb2, 0xa1, 0xc7, 0xc4, 0xdc, 0x0a, 0xa6, 0x5b, 0xa6, 0x23, 0xa9,
	0x97, 0x24, 0x2d, 0xc4, 0xfb, 0xdd, 0x4a, 0xaf, 0x22, 0x34, 0x86, 0x6e, 0xc1, 0x70, 0xce, 0x9c,
	0xde, 0xd4, 0x12, 0x72, 0x15, 0x78, 0x1f, 0xe1, 0xe1, 0x45, 0xe5, 0x8c, 0x0a, 0x1f, 0x60, 0x34,
	0x6b, 0x1e, 0x97, 0x49, 0x8e, 0x16, 0x0f, 0xfc, 0xca, 0x0f, 0xfe, 0x7b, 0xa5, 0x36, 0xb0, 0xc7,
	0x60, 0x09, 0x67, 0xa9, 0xac, 0x47, 0x8b, 0x91, 0x4f, 0x4b, 0x5f, 0x17, 0x19, 0x4a, 0x7d, 0xf1,
	0xdb, 0x04, 0x7b, 0x25, 0x30, 0x72, 0x70, 0x19, 0x5f, 0xab, 0x44, 0x96, 0x69, 0x94, 0x1c, 0x18,
	0xca, 0x61, 0x50, 0x53, 0x91, 0xad, 0x2f, 0xb6, 0x06, 0xe0, 0x3e, 0xba, 0xd2, 0xab, 0xf4, 0xbc,
	0xd7, 0xdf, 0x7e, 0xfe, 0xfa, 0x6e, 0x06, 0xe8, 0xb9, 0x72, 0x30, 0x2d, 0x83, 0xf2, 0x85, 0xff,
	0x52, 0xac, 0x71, 0x70, 0xd4, 0x23, 0x3a, 0x05, 0x47, 0xd9, 0x2a, 0xbd, 0xa8, 0x8e, 0x9d, 0xd0,
	0x0e, 0x06, 0xb5, 0x7b, 0x1b, 0x66, 0xcb, 0xce, 0xee, 0xb8, 0xae, 0x74, 0x4d, 0x92, 0x6d, 0x03,
	0x7c, 0xa6, 0x80, 0x33, 0xf7, 0xc9, 0x3f, 0x81, 0x6f, 0x8c, 0x39, 0x62, 0x00, 0x67, 0xbf, 0x22,
	0x47, 0x93, 0xae, 0x2c, 0x7c, 0x83, 0xa5, 0x8b, 0x9b, 0xff, 0x5f, 0x71, 0x9b, 0x9e, 0xfa, 0x48,
	0x5f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x05, 0x01, 0x03, 0xdb, 0xeb, 0x03, 0x00, 0x00,
}
