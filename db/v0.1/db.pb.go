// Code generated by protoc-gen-go.
// source: db.proto
// DO NOT EDIT!

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	db.proto
	slaves.proto

It has these top-level messages:
	CreateRequest
	DeleteRequest
	BackupRequest
	RestoreRequest
	DescribeRequest
	DescribeResponse
	DbPodSpecs
	ListResponse
	Database
	SlaveAddRequest
*/
package db

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

type CreateRequest struct {
	Type       string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Sku        string `protobuf:"bytes,2,opt,name=sku" json:"sku,omitempty"`
	Config     string `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
	Cluster    string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
	Pv         string `protobuf:"bytes,5,opt,name=pv" json:"pv,omitempty"`
	PvSize     int64  `protobuf:"varint,6,opt,name=pv_size,json=pvSize" json:"pv_size,omitempty"`
	Node       int32  `protobuf:"varint,7,opt,name=node" json:"node,omitempty"`
	Name       string `protobuf:"bytes,8,opt,name=name" json:"name,omitempty"`
	BucketName string `protobuf:"bytes,9,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	Credential string `protobuf:"bytes,10,opt,name=credential" json:"credential,omitempty"`
	Region     string `protobuf:"bytes,11,opt,name=region" json:"region,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DeleteRequest struct {
	Type    string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Cluster string `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BackupRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	BucketName string `protobuf:"bytes,4,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	SecretName string `protobuf:"bytes,5,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	Credential string `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Region     string `protobuf:"bytes,7,opt,name=region" json:"region,omitempty"`
}

func (m *BackupRequest) Reset()                    { *m = BackupRequest{} }
func (m *BackupRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupRequest) ProtoMessage()               {}
func (*BackupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RestoreRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Pv         string `protobuf:"bytes,4,opt,name=pv" json:"pv,omitempty"`
	PvSize     int64  `protobuf:"varint,5,opt,name=pv_size,json=pvSize" json:"pv_size,omitempty"`
	Config     string `protobuf:"bytes,6,opt,name=config" json:"config,omitempty"`
	Node       int32  `protobuf:"varint,7,opt,name=node" json:"node,omitempty"`
	SnapTime   string `protobuf:"bytes,8,opt,name=snap_time,json=snapTime" json:"snap_time,omitempty"`
	BucketName string `protobuf:"bytes,9,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	Credential string `protobuf:"bytes,10,opt,name=credential" json:"credential,omitempty"`
	Region     string `protobuf:"bytes,11,opt,name=region" json:"region,omitempty"`
	SecretName string `protobuf:"bytes,12,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
}

func (m *RestoreRequest) Reset()                    { *m = RestoreRequest{} }
func (m *RestoreRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreRequest) ProtoMessage()               {}
func (*RestoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DescribeResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pods        []*DbPodSpecs  `protobuf:"bytes,2,rep,name=pods" json:"pods,omitempty"`
	ServiceName string         `protobuf:"bytes,3,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServiceIp   string         `protobuf:"bytes,4,opt,name=service_ip,json=serviceIp" json:"service_ip,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetPods() []*DbPodSpecs {
	if m != nil {
		return m.Pods
	}
	return nil
}

type DbPodSpecs struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Phase     string `protobuf:"bytes,2,opt,name=phase" json:"phase,omitempty"`
	Condition string `protobuf:"bytes,3,opt,name=condition" json:"condition,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Type      string `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
}

func (m *DbPodSpecs) Reset()                    { *m = DbPodSpecs{} }
func (m *DbPodSpecs) String() string            { return proto.CompactTextString(m) }
func (*DbPodSpecs) ProtoMessage()               {}
func (*DbPodSpecs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Databases []*Database    `protobuf:"bytes,2,rep,name=databases" json:"databases,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListResponse) GetDatabases() []*Database {
	if m != nil {
		return m.Databases
	}
	return nil
}

type Database struct {
	Phid                string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesPhid      string `protobuf:"bytes,2,opt,name=kubernetes_phid,json=kubernetesPhid" json:"kubernetes_phid,omitempty"`
	KubernetesService   string `protobuf:"bytes,3,opt,name=kubernetes_service,json=kubernetesService" json:"kubernetes_service,omitempty"`
	Type                string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Sku                 string `protobuf:"bytes,5,opt,name=sku" json:"sku,omitempty"`
	ArchiveBucket       string `protobuf:"bytes,6,opt,name=archive_bucket,json=archiveBucket" json:"archive_bucket,omitempty"`
	ArchiveBucketRegion string `protobuf:"bytes,7,opt,name=archive_bucket_region,json=archiveBucketRegion" json:"archive_bucket_region,omitempty"`
	AutoSnapshot        int32  `protobuf:"varint,8,opt,name=auto_snapshot,json=autoSnapshot" json:"auto_snapshot,omitempty"`
	PvSize              int64  `protobuf:"varint,9,opt,name=pv_size,json=pvSize" json:"pv_size,omitempty"`
	Created             string `protobuf:"bytes,10,opt,name=created" json:"created,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "db.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "db.DeleteRequest")
	proto.RegisterType((*BackupRequest)(nil), "db.BackupRequest")
	proto.RegisterType((*RestoreRequest)(nil), "db.RestoreRequest")
	proto.RegisterType((*DescribeRequest)(nil), "db.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "db.DescribeResponse")
	proto.RegisterType((*DbPodSpecs)(nil), "db.DbPodSpecs")
	proto.RegisterType((*ListResponse)(nil), "db.ListResponse")
	proto.RegisterType((*Database)(nil), "db.Database")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Databases service

type DatabasesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type databasesClient struct {
	cc *grpc.ClientConn
}

func NewDatabasesClient(cc *grpc.ClientConn) DatabasesClient {
	return &databasesClient{cc}
}

func (c *databasesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/Backup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/Restore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/db.Databases/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/db.Databases/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Databases service

type DatabasesServer interface {
	Create(context.Context, *CreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*dtypes.VoidResponse, error)
	Backup(context.Context, *BackupRequest) (*dtypes.VoidResponse, error)
	Restore(context.Context, *RestoreRequest) (*dtypes.VoidResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	List(context.Context, *dtypes.VoidRequest) (*ListResponse, error)
}

func RegisterDatabasesServer(s *grpc.Server, srv DatabasesServer) {
	s.RegisterService(&_Databases_serviceDesc, srv)
}

func _Databases_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DatabasesServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Databases_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DatabasesServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Databases_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DatabasesServer).Backup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Databases_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DatabasesServer).Restore(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Databases_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DatabasesServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Databases_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DatabasesServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Databases_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.Databases",
	HandlerType: (*DatabasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Databases_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Databases_Delete_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Databases_Backup_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _Databases_Restore_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Databases_Describe_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Databases_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x96, 0x1d, 0xc7, 0x89, 0x27, 0x3f, 0x6d, 0xb7, 0x3d, 0x3d, 0x56, 0xda, 0x73, 0x28, 0x2e,
	0xb4, 0x55, 0xa4, 0x26, 0xa8, 0xdc, 0x71, 0x47, 0xe9, 0x0d, 0x12, 0x42, 0xc5, 0x41, 0x48, 0x5c,
	0x20, 0xe3, 0x9f, 0x25, 0xb1, 0x12, 0x6c, 0xe3, 0xb5, 0x23, 0x41, 0x55, 0x09, 0xf1, 0x06, 0x88,
	0x17, 0xe0, 0x19, 0x78, 0x00, 0x5e, 0x80, 0x0b, 0x2e, 0x78, 0x05, 0x1e, 0x04, 0xef, 0x8f, 0x6b,
	0xbb, 0x4d, 0xab, 0x22, 0xc1, 0x4d, 0xb4, 0xfb, 0xcd, 0xf8, 0x9b, 0x9d, 0x6f, 0x67, 0x66, 0x03,
	0x4d, 0xcf, 0x19, 0x44, 0x71, 0x98, 0x84, 0x48, 0xf6, 0x9c, 0xde, 0xe6, 0x38, 0x0c, 0xc7, 0x33,
	0x3c, 0xb4, 0x23, 0x7f, 0x68, 0x07, 0x41, 0x98, 0xd8, 0x89, 0x1f, 0x06, 0x84, 0x7b, 0xf4, 0xd6,
	0x29, 0xec, 0x25, 0x6f, 0x23, 0x4c, 0x86, 0xec, 0x97, 0xe3, 0xc6, 0x47, 0x19, 0x3a, 0x0f, 0x62,
	0x6c, 0x27, 0xd8, 0xc4, 0x6f, 0x52, 0x4c, 0x12, 0x84, 0x40, 0xa1, 0x0e, 0xba, 0xb4, 0x25, 0xed,
	0x69, 0x26, 0x5b, 0xa3, 0x65, 0xa8, 0x91, 0x69, 0xaa, 0xcb, 0x0c, 0xa2, 0x4b, 0xb4, 0x0e, 0xaa,
	0x1b, 0x06, 0xaf, 0xfc, 0xb1, 0x5e, 0x63, 0xa0, 0xd8, 0x21, 0x1d, 0x1a, 0xee, 0x2c, 0x25, 0x09,
	0x8e, 0x75, 0x85, 0x19, 0xf2, 0x2d, 0xea, 0x82, 0x1c, 0xcd, 0xf5, 0x3a, 0x03, 0xb3, 0x15, 0xfa,
	0x17, 0x1a, 0xd1, 0xdc, 0x22, 0xfe, 0x3b, 0xac, 0xab, 0x19, 0x58, 0x33, 0xd5, 0x68, 0x3e, 0xca,
	0x76, 0xf4, 0x00, 0x41, 0xe8, 0x61, 0xbd, 0x91, 0xa1, 0x75, 0x93, 0xad, 0x19, 0x66, 0xbf, 0xc6,
	0x7a, 0x93, 0x1f, 0x8a, 0xae, 0xd1, 0x0d, 0x68, 0x39, 0xa9, 0x3b, 0xc5, 0x89, 0xc5, 0x4c, 0x1a,
	0x33, 0x01, 0x87, 0x1e, 0x53, 0x87, 0xff, 0x01, 0xdc, 0x18, 0x7b, 0x38, 0x48, 0x7c, 0x7b, 0xa6,
	0x03, 0xb7, 0x17, 0x08, 0xcd, 0x21, 0xc6, 0xe3, 0x4c, 0x24, 0xbd, 0xc5, 0x73, 0xe0, 0x3b, 0xe3,
	0x09, 0x74, 0x8e, 0xf0, 0x0c, 0x5f, 0x2d, 0x49, 0x7e, 0x22, 0xb9, 0x74, 0xa2, 0x52, 0xf2, 0xb5,
	0x4a, 0xf2, 0xc6, 0x37, 0x09, 0x3a, 0x87, 0xb6, 0x3b, 0x4d, 0xa3, 0x9c, 0xb3, 0xe4, 0x2b, 0x55,
	0x85, 0x5a, 0xc4, 0x9c, 0x9f, 0xa0, 0x56, 0x3a, 0xc1, 0xb9, 0xfc, 0x95, 0x0b, 0xf9, 0x67, 0x0e,
	0x04, 0x67, 0xf9, 0x0a, 0x07, 0x2e, 0x3d, 0x70, 0x68, 0x81, 0x40, 0xea, 0x15, 0x02, 0x35, 0x2a,
	0x02, 0x7d, 0x95, 0xa1, 0x6b, 0x66, 0x49, 0x84, 0x31, 0xfe, 0x73, 0xe9, 0xf0, 0xfa, 0x50, 0x16,
	0xd5, 0x47, 0xbd, 0x52, 0x1f, 0x45, 0xe9, 0xa9, 0x95, 0xd2, 0x5b, 0x54, 0x37, 0x1b, 0xa0, 0x91,
	0xc0, 0x8e, 0xac, 0xc4, 0x3f, 0x2b, 0x9e, 0x26, 0x05, 0x9e, 0xfa, 0x7f, 0xb1, 0x80, 0xce, 0x0b,
	0xdf, 0x3e, 0x2f, 0xbc, 0x31, 0x82, 0xa5, 0x23, 0x4c, 0xdc, 0xd8, 0x77, 0xae, 0x27, 0x20, 0x13,
	0x4b, 0x5e, 0x50, 0x7d, 0xb5, 0x42, 0x54, 0xe3, 0xb3, 0x04, 0xcb, 0x05, 0x2b, 0x89, 0xb2, 0xe6,
	0xc7, 0x68, 0x07, 0x54, 0x92, 0x4d, 0x82, 0x94, 0x30, 0xd6, 0xd6, 0x41, 0x77, 0xc0, 0x87, 0xc0,
	0x60, 0xc4, 0x50, 0x53, 0x58, 0x91, 0x01, 0x4a, 0x14, 0x7a, 0x24, 0x0b, 0x52, 0xe3, 0x5e, 0xce,
	0xe0, 0xc8, 0x39, 0x0e, 0xbd, 0x51, 0x84, 0x5d, 0x62, 0x32, 0x1b, 0xba, 0x09, 0x6d, 0x82, 0xe3,
	0xb9, 0xef, 0x62, 0xab, 0x14, 0xbc, 0x25, 0x30, 0xa6, 0xd8, 0x7f, 0x00, 0xb9, 0x8b, 0x1f, 0x89,
	0xcb, 0xd4, 0x04, 0xf2, 0x30, 0x32, 0xde, 0x4b, 0x00, 0x05, 0xed, 0x59, 0x16, 0x52, 0xa9, 0x34,
	0xd6, 0xa0, 0x1e, 0x4d, 0x6c, 0x92, 0xa7, 0xcb, 0x37, 0x68, 0x13, 0xb4, 0xec, 0x96, 0x3d, 0x9f,
	0x8e, 0x34, 0x11, 0xb7, 0x00, 0xe8, 0x3d, 0x88, 0x24, 0x79, 0xc4, 0x3c, 0xa9, 0x5c, 0xb9, 0x7a,
	0xa1, 0x9c, 0xe1, 0x40, 0xfb, 0x91, 0x4f, 0x92, 0xdf, 0x16, 0xa8, 0x0f, 0x9a, 0x67, 0x27, 0xb6,
	0x93, 0x9d, 0x26, 0x57, 0xa9, 0xcd, 0x54, 0x12, 0xa0, 0x59, 0x98, 0x8d, 0xef, 0x32, 0x34, 0x73,
	0x9c, 0x1e, 0x22, 0x9a, 0xf8, 0x5e, 0x9e, 0x24, 0x5d, 0xa3, 0x5d, 0x58, 0x9a, 0xa6, 0x0e, 0x8e,
	0x83, 0x6c, 0xc8, 0x10, 0x8b, 0x99, 0x79, 0xba, 0xdd, 0x02, 0x3e, 0xa6, 0x8e, 0xfb, 0x80, 0x4a,
	0x8e, 0x42, 0x48, 0x21, 0xc0, 0x4a, 0x61, 0x19, 0x71, 0xc3, 0x59, 0xc2, 0xca, 0xc5, 0xd9, 0x5d,
	0x2f, 0x66, 0xf7, 0x6d, 0xe8, 0xda, 0xb1, 0x3b, 0xf1, 0xe7, 0xd8, 0xe2, 0xc5, 0x2e, 0x1a, 0xa9,
	0x23, 0xd0, 0x43, 0x06, 0xa2, 0x03, 0xf8, 0xa7, 0xea, 0x66, 0x55, 0x86, 0xc1, 0x6a, 0xc5, 0xdb,
	0xe4, 0x95, 0xbf, 0x0d, 0x1d, 0x3b, 0x4d, 0x42, 0x8b, 0xf6, 0x18, 0x99, 0x84, 0x09, 0xeb, 0xb9,
	0xba, 0xd9, 0xa6, 0xe0, 0x48, 0x60, 0xe5, 0xce, 0xd6, 0x2a, 0x9d, 0x4d, 0x7b, 0x80, 0xbd, 0x45,
	0x9e, 0x68, 0xb6, 0x7c, 0x7b, 0xf0, 0x45, 0x01, 0x2d, 0x57, 0x94, 0xa0, 0xe7, 0xa0, 0xf2, 0x37,
	0x0b, 0xad, 0xd0, 0x2b, 0xa8, 0xbc, 0x5f, 0xbd, 0xb5, 0xfc, 0x02, 0x9f, 0x85, 0xbe, 0x97, 0x5f,
	0xb3, 0xb1, 0xf3, 0xe1, 0xc7, 0xcf, 0x4f, 0xf2, 0x56, 0x6f, 0x83, 0xbd, 0x8f, 0x9e, 0x33, 0x9c,
	0xdf, 0x19, 0x9e, 0x88, 0x0e, 0x3b, 0x1d, 0x9e, 0xd0, 0x2f, 0x4e, 0xef, 0x49, 0x7d, 0xf4, 0x02,
	0x54, 0x3e, 0xfb, 0x39, 0x75, 0xe5, 0x1d, 0xb8, 0x84, 0xba, 0xcf, 0xa8, 0x6f, 0xf5, 0x8d, 0x2b,
	0xa8, 0x87, 0x27, 0xb4, 0xb8, 0x4f, 0xd1, 0x4b, 0x50, 0xf9, 0x33, 0xc0, 0xe9, 0x2b, 0x4f, 0xc2,
	0x25, 0xf4, 0xfb, 0x8c, 0x7e, 0xd7, 0x28, 0xd3, 0x3b, 0xec, 0xbb, 0x85, 0x09, 0xb8, 0xd0, 0x10,
	0xa3, 0x19, 0x21, 0x1a, 0xa2, 0x3a, 0xa7, 0x2f, 0x89, 0x31, 0x60, 0x31, 0xf6, 0x7a, 0xdb, 0xa5,
	0x18, 0x31, 0xff, 0x70, 0x61, 0x10, 0x3b, 0xab, 0x6f, 0x31, 0x69, 0xd0, 0x2a, 0xd7, 0xa9, 0x32,
	0xcd, 0x68, 0x98, 0x32, 0x58, 0x55, 0x0a, 0x5d, 0x47, 0xa9, 0xfb, 0xa0, 0xd0, 0x3e, 0xa5, 0xf4,
	0xe5, 0x03, 0x73, 0xfa, 0x65, 0x4a, 0x5f, 0x6e, 0x63, 0x03, 0x31, 0xea, 0x36, 0x82, 0x82, 0xda,
	0x51, 0xd9, 0x5f, 0x9c, 0xbb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x19, 0x3d, 0xa1, 0x18, 0x28,
	0x09, 0x00, 0x00,
}
