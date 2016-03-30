// Code generated by protoc-gen-go.
// source: pv.proto
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

type PVRegisterRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Identifier string `protobuf:"bytes,3,opt,name=identifier" json:"identifier,omitempty"`
	Provider   string `protobuf:"bytes,4,opt,name=provider" json:"provider,omitempty"`
	Size       int64  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Path       string `protobuf:"bytes,6,opt,name=path" json:"path,omitempty"`
}

func (m *PVRegisterRequest) Reset()                    { *m = PVRegisterRequest{} }
func (m *PVRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVRegisterRequest) ProtoMessage()               {}
func (*PVRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PVUnregisterRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVUnregisterRequest) Reset()                    { *m = PVUnregisterRequest{} }
func (m *PVUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVUnregisterRequest) ProtoMessage()               {}
func (*PVUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }


type PVListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *PVListRequest) Reset()                    { *m = PVListRequest{} }
func (m *PVListRequest) String() string            { return proto.CompactTextString(m) }
func (*PVListRequest) ProtoMessage()               {}
func (*PVListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type PVListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Volumes []*PV          `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *PVListResponse) Reset()                    { *m = PVListResponse{} }
func (m *PVListResponse) String() string            { return proto.CompactTextString(m) }
func (*PVListResponse) ProtoMessage()               {}
func (*PVListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *PVListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVListResponse) GetVolumes() []*PV {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type PV struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size   int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Status string   `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Volume string   `protobuf:"bytes,4,opt,name=volume" json:"volume,omitempty"`
	Claim  string   `protobuf:"bytes,5,opt,name=claim" json:"claim,omitempty"`
	Users  []string `protobuf:"bytes,6,rep,name=users" json:"users,omitempty"`
}

func (m *PV) Reset()                    { *m = PV{} }
func (m *PV) String() string            { return proto.CompactTextString(m) }
func (*PV) ProtoMessage()               {}
func (*PV) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*PVRegisterRequest)(nil), "pv.PVRegisterRequest")
	proto.RegisterType((*PVUnregisterRequest)(nil), "pv.PVUnregisterRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for PersistentVolumes service

type PersistentVolumesClient interface {
	// rpc List(PVListRequest) returns (PVListResponse) {
	// option (google.api.http) = {
	// get: "/api/pv/v0.1/{cluster}"
	// };
	// }
	Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type persistentVolumesClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumesClient(cc *grpc.ClientConn) PersistentVolumesClient {
	return &persistentVolumesClient{cc}
}

func (c *persistentVolumesClient) Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumes/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumes/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumes service

type PersistentVolumesServer interface {
	// rpc List(PVListRequest) returns (PVListResponse) {
	// option (google.api.http) = {
	// get: "/api/pv/v0.1/{cluster}"
	// };
	// }
	Register(context.Context, *PVRegisterRequest) (*dtypes.VoidResponse, error)
	Unregister(context.Context, *PVUnregisterRequest) (*dtypes.VoidResponse, error)
}

func RegisterPersistentVolumesServer(s *grpc.Server, srv PersistentVolumesServer) {
	s.RegisterService(&_PersistentVolumes_serviceDesc, srv)
}

func _PersistentVolumes_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PVRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PersistentVolumesServer).Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PersistentVolumes_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PVUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PersistentVolumesServer).Unregister(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PersistentVolumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.PersistentVolumes",
	HandlerType: (*PersistentVolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PersistentVolumes_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumes_Unregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}


var fileDescriptor1 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x93, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0x3b, 0xe7, 0xbb, 0x1b, 0xc4, 0x49, 0x37, 0x1c, 0xc1, 0xb2, 0x20, 0x44, 0x96,
	0x40, 0x21, 0x85, 0x0d, 0xa1, 0xa3, 0xa5, 0x05, 0xc9, 0x32, 0xc2, 0x05, 0x9d, 0x93, 0x2c, 0x61,
	0x25, 0xc7, 0xbb, 0x78, 0xd7, 0x96, 0x20, 0x4a, 0x43, 0x45, 0xcf, 0x33, 0x50, 0xf0, 0x3c, 0xbc,
	0x02, 0x0f, 0xc2, 0x7e, 0xd8, 0xc6, 0x88, 0xaf, 0x82, 0x26, 0x9a, 0xf9, 0xef, 0xec, 0xfc, 0x27,
	0xbf, 0x1d, 0xc3, 0x19, 0x6f, 0x63, 0x5e, 0x33, 0xc9, 0xd0, 0xe5, 0x6d, 0x78, 0x7b, 0xc7, 0xd8,
	0xae, 0x24, 0x49, 0xc1, 0x69, 0x52, 0x54, 0x15, 0x93, 0x85, 0xa4, 0xac, 0x12, 0xb6, 0x22, 0x9c,
	0x6a, 0x79, 0x2b, 0xdf, 0x71, 0x22, 0x12, 0xf3, 0x6b, 0xf5, 0xe8, 0xb3, 0x03, 0x97, 0x69, 0x9e,
	0x91, 0x1d, 0x15, 0x92, 0xd4, 0x19, 0x79, 0xdb, 0x10, 0x21, 0x31, 0x80, 0xd3, 0x4d, 0xd9, 0x68,
	0x25, 0x70, 0xe6, 0xce, 0xe2, 0x3c, 0xeb, 0x53, 0x44, 0x98, 0x54, 0xc5, 0x9e, 0x04, 0xae, 0x91,
	0x4d, 0x8c, 0x33, 0x00, 0xba, 0x25, 0x95, 0xa4, 0xaf, 0xa9, 0xba, 0xe0, 0x99, 0x93, 0x91, 0x82,
	0xa1, 0x9a, 0xb4, 0x66, 0xad, 0x52, 0xea, 0x60, 0x62, 0x4e, 0x87, 0x5c, 0xf7, 0x13, 0xf4, 0x3d,
	0x09, 0x4e, 0x94, 0xee, 0x65, 0x26, 0xd6, 0x1a, 0x2f, 0xe4, 0x9b, 0xc0, 0xb7, 0x1e, 0x3a, 0x8e,
	0x9e, 0xc2, 0x8d, 0x34, 0x7f, 0x59, 0xd5, 0xff, 0x33, 0x68, 0xf4, 0x00, 0xae, 0xa7, 0xf9, 0x33,
	0xd5, 0xe0, 0x9f, 0xd7, 0xa3, 0x57, 0x70, 0xd1, 0x97, 0x0a, 0xae, 0x30, 0x12, 0xbc, 0x0f, 0xbe,
	0x50, 0x4c, 0x1b, 0x61, 0x4a, 0xaf, 0xad, 0x2e, 0x62, 0x8b, 0x33, 0x7e, 0x61, 0xd4, 0xac, 0x3b,
	0xc5, 0x39, 0x9c, 0xb6, 0xac, 0x6c, 0xf6, 0x44, 0x28, 0x6f, 0x4f, 0x15, 0xfa, 0xb1, 0x7a, 0x27,
	0xc5, 0xb8, 0x97, 0xa3, 0x8f, 0x0e, 0xb8, 0x69, 0x3e, 0x4c, 0xe8, 0x8c, 0x50, 0xf6, 0x38, 0xdc,
	0x11, 0x8e, 0xe9, 0x60, 0x6c, 0xd1, 0xf6, 0x46, 0x4a, 0xb7, 0x1d, 0x3b, 0xa8, 0x5d, 0x86, 0x57,
	0x70, 0xb2, 0x29, 0x0b, 0xba, 0x37, 0x4c, 0xcf, 0x33, 0x9b, 0x68, 0xb5, 0x11, 0xa4, 0x16, 0x8a,
	0xaa, 0xa7, 0x55, 0x93, 0xac, 0xbe, 0xb8, 0xea, 0xf9, 0x55, 0xa0, 0xa1, 0x56, 0x32, 0xb7, 0x03,
	0xe2, 0x73, 0x98, 0xe8, 0xbf, 0x8e, 0x97, 0x76, 0xf2, 0x11, 0xb1, 0x10, 0xc7, 0x92, 0x25, 0x13,
	0xcd, 0x3e, 0x7c, 0xfd, 0xf6, 0xc9, 0x0d, 0x70, 0x6a, 0x76, 0x8f, 0xb7, 0x49, 0xfb, 0x30, 0x7e,
	0x94, 0x1c, 0x3a, 0x94, 0x47, 0x2c, 0xe0, 0xac, 0x5f, 0x30, 0xbc, 0xd9, 0xc1, 0xf8, 0xf9, 0x1d,
	0xc3, 0xab, 0x1e, 0x66, 0xce, 0xe8, 0x76, 0x68, 0xbc, 0x30, 0x8d, 0xa3, 0xf0, 0xce, 0xef, 0x1b,
	0x27, 0x07, 0xcd, 0xec, 0xf8, 0xc4, 0x59, 0xe2, 0x1a, 0xe0, 0xc7, 0x72, 0xe0, 0x2d, 0x6b, 0xf2,
	0xcb, 0xba, 0xfc, 0xc1, 0xe6, 0x9e, 0xb1, 0xb9, 0xbb, 0xfc, 0xbb, 0xcd, 0xda, 0x37, 0x5f, 0xcc,
	0xe3, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x83, 0x76, 0x2f, 0x1d, 0x77, 0x03, 0x00, 0x00,
}
