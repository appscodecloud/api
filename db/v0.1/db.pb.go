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
	KubeInfo
	ListRequest
	ListResponse
	Database
	SnapshotListRequest
	SnapshotListResponse
	Snapshot
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
	Cluster        string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type           string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Sku            string `protobuf:"bytes,4,opt,name=sku" json:"sku,omitempty"`
	Version        string `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	Pv             string `protobuf:"bytes,6,opt,name=pv" json:"pv,omitempty"`
	PvSize         int64  `protobuf:"varint,7,opt,name=pv_size,json=pvSize" json:"pv_size,omitempty"`
	AuthSecretName string `protobuf:"bytes,8,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	Node           int64  `protobuf:"varint,9,opt,name=node" json:"node,omitempty"`
	Credential     string `protobuf:"bytes,10,opt,name=credential" json:"credential,omitempty"`
	BucketName     string `protobuf:"bytes,11,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	Region         string `protobuf:"bytes,12,opt,name=region" json:"region,omitempty"`
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
	Cluster        string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type           string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	BucketName     string `protobuf:"bytes,4,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	SnapshotName   string `protobuf:"bytes,5,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	Credential     string `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Region         string `protobuf:"bytes,7,opt,name=region" json:"region,omitempty"`
	AuthSecretName string `protobuf:"bytes,8,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	Wal            bool   `protobuf:"varint,9,opt,name=wal" json:"wal,omitempty"`
	Force          bool   `protobuf:"varint,10,opt,name=force" json:"force,omitempty"`
}

func (m *BackupRequest) Reset()                    { *m = BackupRequest{} }
func (m *BackupRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupRequest) ProtoMessage()               {}
func (*BackupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RestoreRequest struct {
	Cluster        string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type           string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Pv             string `protobuf:"bytes,4,opt,name=pv" json:"pv,omitempty"`
	PvSize         int64  `protobuf:"varint,5,opt,name=pv_size,json=pvSize" json:"pv_size,omitempty"`
	Sku            string `protobuf:"bytes,6,opt,name=sku" json:"sku,omitempty"`
	Node           int64  `protobuf:"varint,7,opt,name=node" json:"node,omitempty"`
	SnapshotPhid   string `protobuf:"bytes,8,opt,name=snapshot_phid,json=snapshotPhid" json:"snapshot_phid,omitempty"`
	BucketName     string `protobuf:"bytes,9,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	Credential     string `protobuf:"bytes,10,opt,name=credential" json:"credential,omitempty"`
	Region         string `protobuf:"bytes,11,opt,name=region" json:"region,omitempty"`
	Version        string `protobuf:"bytes,12,opt,name=version" json:"version,omitempty"`
	AuthSecretName string `protobuf:"bytes,13,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	Wal            bool   `protobuf:"varint,14,opt,name=wal" json:"wal,omitempty"`
	Force          bool   `protobuf:"varint,15,opt,name=force" json:"force,omitempty"`
}

func (m *RestoreRequest) Reset()                    { *m = RestoreRequest{} }
func (m *RestoreRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreRequest) ProtoMessage()               {}
func (*RestoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DescribeResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name         string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Phid         string         `protobuf:"bytes,3,opt,name=phid" json:"phid,omitempty"`
	Cluster      string         `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
	Type         string         `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	Sku          string         `protobuf:"bytes,6,opt,name=sku" json:"sku,omitempty"`
	Version      string         `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
	AutoSnapshot int32          `protobuf:"varint,8,opt,name=auto_snapshot,json=autoSnapshot" json:"auto_snapshot,omitempty"`
	LastBackup   string         `protobuf:"bytes,9,opt,name=last_backup,json=lastBackup" json:"last_backup,omitempty"`
	Kube         *KubeInfo      `protobuf:"bytes,10,opt,name=kube" json:"kube,omitempty"`
	DbStatus     string         `protobuf:"bytes,11,opt,name=db_status,json=dbStatus" json:"db_status,omitempty"`
	Created      string         `protobuf:"bytes,12,opt,name=created" json:"created,omitempty"`
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

func (m *DescribeResponse) GetKube() *KubeInfo {
	if m != nil {
		return m.Kube
	}
	return nil
}

type KubeInfo struct {
	Svc  int64 `protobuf:"varint,1,opt,name=svc" json:"svc,omitempty"`
	Pods int64 `protobuf:"varint,2,opt,name=pods" json:"pods,omitempty"`
	Rc   int64 `protobuf:"varint,3,opt,name=rc" json:"rc,omitempty"`
}

func (m *KubeInfo) Reset()                    { *m = KubeInfo{} }
func (m *KubeInfo) String() string            { return proto.CompactTextString(m) }
func (*KubeInfo) ProtoMessage()               {}
func (*KubeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Databases []*Database    `protobuf:"bytes,2,rep,name=databases" json:"databases,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
	Phid         string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	ClusterName  string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Namespace    string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	Type         string `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	Sku          string `protobuf:"bytes,6,opt,name=sku" json:"sku,omitempty"`
	Mode         string `protobuf:"bytes,7,opt,name=mode" json:"mode,omitempty"`
	Version      string `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`
	AuthSecret   string `protobuf:"bytes,9,opt,name=auth_secret,json=authSecret" json:"auth_secret,omitempty"`
	AutoSnapshot int32  `protobuf:"varint,10,opt,name=auto_snapshot,json=autoSnapshot" json:"auto_snapshot,omitempty"`
	PvSize       int64  `protobuf:"varint,11,opt,name=pv_size,json=pvSize" json:"pv_size,omitempty"`
	Created      string `protobuf:"bytes,12,opt,name=created" json:"created,omitempty"`
	LastBackup   string `protobuf:"bytes,13,opt,name=last_backup,json=lastBackup" json:"last_backup,omitempty"`
	Status       string `protobuf:"bytes,14,opt,name=status" json:"status,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SnapshotListRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SnapshotListRequest) Reset()                    { *m = SnapshotListRequest{} }
func (m *SnapshotListRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListRequest) ProtoMessage()               {}
func (*SnapshotListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type SnapshotListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	List   []*Snapshot    `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
}

func (m *SnapshotListResponse) Reset()                    { *m = SnapshotListResponse{} }
func (m *SnapshotListResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListResponse) ProtoMessage()               {}
func (*SnapshotListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SnapshotListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SnapshotListResponse) GetList() []*Snapshot {
	if m != nil {
		return m.List
	}
	return nil
}

type Snapshot struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Phid        string `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	DbName      string `protobuf:"bytes,3,opt,name=db_name,json=dbName" json:"db_name,omitempty"`
	DbCluster   string `protobuf:"bytes,4,opt,name=db_cluster,json=dbCluster" json:"db_cluster,omitempty"`
	Provider    string `protobuf:"bytes,5,opt,name=provider" json:"provider,omitempty"`
	CloudCred   string `protobuf:"bytes,6,opt,name=cloud_cred,json=cloudCred" json:"cloud_cred,omitempty"`
	Bucket      string `protobuf:"bytes,7,opt,name=bucket" json:"bucket,omitempty"`
	Region      string `protobuf:"bytes,8,opt,name=region" json:"region,omitempty"`
	Zone        string `protobuf:"bytes,9,opt,name=zone" json:"zone,omitempty"`
	Status      string `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
	Process     string `protobuf:"bytes,11,opt,name=process" json:"process,omitempty"`
	RestoreDb   string `protobuf:"bytes,12,opt,name=restore_db,json=restoreDb" json:"restore_db,omitempty"`
	RestoreSnap string `protobuf:"bytes,13,opt,name=restore_snap,json=restoreSnap" json:"restore_snap,omitempty"`
	Created     string `protobuf:"bytes,14,opt,name=created" json:"created,omitempty"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "db.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "db.DeleteRequest")
	proto.RegisterType((*BackupRequest)(nil), "db.BackupRequest")
	proto.RegisterType((*RestoreRequest)(nil), "db.RestoreRequest")
	proto.RegisterType((*DescribeRequest)(nil), "db.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "db.DescribeResponse")
	proto.RegisterType((*KubeInfo)(nil), "db.kubeInfo")
	proto.RegisterType((*ListRequest)(nil), "db.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "db.ListResponse")
	proto.RegisterType((*Database)(nil), "db.Database")
	proto.RegisterType((*SnapshotListRequest)(nil), "db.SnapshotListRequest")
	proto.RegisterType((*SnapshotListResponse)(nil), "db.SnapshotListResponse")
	proto.RegisterType((*Snapshot)(nil), "db.Snapshot")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Databases service

type DatabasesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	SnapshotList(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error)
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

func (c *databasesClient) SnapshotList(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error) {
	out := new(SnapshotListResponse)
	err := grpc.Invoke(ctx, "/db.Databases/SnapshotList", in, out, c.cc, opts...)
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
	SnapshotList(context.Context, *SnapshotListRequest) (*SnapshotListResponse, error)
	Restore(context.Context, *RestoreRequest) (*dtypes.VoidResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	List(context.Context, *dtypes.VoidRequest) (*ListResponse, error)
}

func RegisterDatabasesServer(s *grpc.Server, srv DatabasesServer) {
	s.RegisterService(&_Databases_serviceDesc, srv)
}

func _Databases_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Backup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_SnapshotList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).SnapshotList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/SnapshotList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).SnapshotList(ctx, req.(*SnapshotListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Restore(ctx, req.(*RestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).List(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "SnapshotList",
			Handler:    _Databases_SnapshotList_Handler,
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
	// 1153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x57, 0xcd, 0x6e, 0x23, 0x45,
	0x17, 0x95, 0xff, 0xda, 0xf6, 0xf5, 0xcf, 0xe4, 0xab, 0xe4, 0x4b, 0x5a, 0x66, 0x06, 0x86, 0x46,
	0x1a, 0x85, 0x20, 0x6c, 0x26, 0x23, 0x24, 0x60, 0x03, 0x62, 0xb2, 0x41, 0x1a, 0x21, 0xe8, 0x48,
	0x6c, 0x4d, 0xff, 0xd4, 0x24, 0x96, 0x3d, 0xdd, 0x3d, 0x5d, 0x6d, 0x23, 0x26, 0xca, 0x86, 0x2d,
	0x4b, 0x9e, 0x80, 0x07, 0xe1, 0x29, 0x58, 0xb2, 0x61, 0xc1, 0x3b, 0x20, 0x76, 0xdc, 0xba, 0x55,
	0x65, 0x57, 0xc7, 0x6d, 0x48, 0xd0, 0x6c, 0x92, 0xaa, 0x53, 0xdd, 0x75, 0xeb, 0x9e, 0x73, 0xee,
	0xed, 0x32, 0x74, 0xe2, 0x70, 0x9c, 0xe5, 0x69, 0x91, 0xb2, 0x7a, 0x1c, 0x8e, 0xee, 0x5f, 0xa4,
	0xe9, 0xc5, 0x82, 0x4f, 0x82, 0x6c, 0x36, 0x09, 0x92, 0x24, 0x2d, 0x82, 0x62, 0x96, 0x26, 0x42,
	0x3d, 0x31, 0x3a, 0x94, 0x70, 0x5c, 0x7c, 0x9f, 0x71, 0x31, 0xa1, 0xbf, 0x0a, 0xf7, 0x7e, 0xa9,
	0xc3, 0xe0, 0x69, 0xce, 0x83, 0x82, 0xfb, 0xfc, 0xe5, 0x92, 0x8b, 0x82, 0xb9, 0xd0, 0x8e, 0x16,
	0x4b, 0x51, 0xf0, 0xdc, 0xad, 0x3d, 0xac, 0x1d, 0x77, 0x7d, 0x33, 0x65, 0x0c, 0x9a, 0x49, 0xf0,
	0x82, 0xbb, 0x75, 0x82, 0x69, 0x2c, 0x31, 0xb9, 0x9d, 0xdb, 0x50, 0x98, 0x1c, 0xb3, 0x3d, 0x68,
	0x88, 0xf9, 0xd2, 0x6d, 0x12, 0x24, 0x87, 0x72, 0xcf, 0x15, 0xcf, 0x05, 0x9e, 0xc7, 0x6d, 0xa9,
	0x3d, 0xf5, 0x94, 0x0d, 0xa1, 0x9e, 0xad, 0x5c, 0x87, 0x40, 0x1c, 0xb1, 0x23, 0x68, 0x67, 0xab,
	0xa9, 0x98, 0xbd, 0xe2, 0x6e, 0x1b, 0xc1, 0x86, 0xef, 0x64, 0xab, 0x73, 0x9c, 0xb1, 0x63, 0xd8,
	0x0b, 0x96, 0xc5, 0xe5, 0x54, 0xf0, 0x28, 0xe7, 0xc5, 0x94, 0x0e, 0xd2, 0xa1, 0xd7, 0x86, 0x12,
	0x3f, 0x27, 0xf8, 0x4b, 0x7d, 0xa4, 0x24, 0x8d, 0xb9, 0xdb, 0xa5, 0xf7, 0x69, 0xcc, 0xde, 0x04,
	0xc0, 0xf5, 0x98, 0x27, 0xc5, 0x2c, 0x58, 0xb8, 0x40, 0xef, 0x59, 0x08, 0x7b, 0x0b, 0x7a, 0xe1,
	0x32, 0x9a, 0x9b, 0x8d, 0x7b, 0xea, 0x01, 0x05, 0xd1, 0xa6, 0x87, 0xe0, 0xe4, 0xfc, 0x42, 0x26,
	0xd0, 0xa7, 0x35, 0x3d, 0xf3, 0xbe, 0x86, 0xc1, 0x19, 0x5f, 0xf0, 0x0d, 0x7d, 0x86, 0x90, 0x9a,
	0x45, 0x48, 0x15, 0x71, 0x16, 0xcd, 0x8d, 0x12, 0xcd, 0xde, 0xcf, 0x28, 0xc9, 0xe7, 0x41, 0x34,
	0x5f, 0x66, 0xaf, 0x4f, 0x92, 0x1b, 0xf9, 0x35, 0xb7, 0xf2, 0x7b, 0x07, 0x06, 0x22, 0x09, 0x32,
	0x71, 0x99, 0xea, 0x47, 0x94, 0x4e, 0x7d, 0x03, 0xd2, 0x43, 0x65, 0x16, 0x9d, 0x2d, 0x16, 0x37,
	0x24, 0xb5, 0x6d, 0x92, 0xee, 0xa0, 0x1d, 0x5a, 0xe7, 0x3b, 0xdc, 0x5a, 0x4a, 0xd7, 0xf1, 0xe5,
	0x90, 0x1d, 0x40, 0xeb, 0x79, 0x9a, 0x47, 0x9c, 0x44, 0xeb, 0xf8, 0x6a, 0xe2, 0xfd, 0xd8, 0x80,
	0xa1, 0x8f, 0xd4, 0xa4, 0xf9, 0x6b, 0xf4, 0xad, 0xf2, 0x62, 0xb3, 0xca, 0x8b, 0xad, 0x92, 0x17,
	0xb5, 0xc1, 0x9d, 0x8d, 0xc1, 0x8d, 0xe7, 0xda, 0x96, 0xe7, 0x6c, 0x4a, 0xb3, 0xcb, 0x59, 0xac,
	0x53, 0x5e, 0x53, 0xfa, 0x15, 0x62, 0x37, 0x85, 0xe9, 0x6e, 0x09, 0xf3, 0x6f, 0xce, 0xdd, 0x70,
	0xde, 0x2b, 0x71, 0x6e, 0x95, 0x5c, 0xbf, 0x5c, 0x72, 0x55, 0x6a, 0x0c, 0xfe, 0x49, 0x8d, 0x61,
	0x85, 0x1a, 0xf7, 0x6c, 0x35, 0x3e, 0x85, 0x7b, 0x67, 0x5c, 0x44, 0xf9, 0x2c, 0xfc, 0x6f, 0x6a,
	0x78, 0xbf, 0xd5, 0x61, 0x6f, 0xb3, 0x83, 0xc8, 0xb0, 0x6f, 0x71, 0xf6, 0x08, 0x1c, 0x81, 0x4d,
	0x6c, 0x29, 0x68, 0x87, 0xde, 0xe9, 0x70, 0xac, 0xfa, 0xd7, 0xf8, 0x9c, 0x50, 0x5f, 0xaf, 0xee,
	0x92, 0x97, 0x28, 0xd7, 0xf2, 0xca, 0xb1, 0x7d, 0xa4, 0xe6, 0xd6, 0x91, 0xc8, 0x0c, 0xad, 0xed,
	0x26, 0xe6, 0x54, 0x36, 0xb1, 0x76, 0x99, 0x51, 0x54, 0x1a, 0x99, 0x4b, 0xa7, 0x46, 0x59, 0x52,
	0xba, 0xe5, 0xf7, 0x25, 0x78, 0xae, 0x31, 0xa9, 0xf4, 0x22, 0x10, 0xc5, 0x34, 0xa4, 0xd2, 0x36,
	0x4a, 0x4b, 0x48, 0x15, 0x3b, 0x7b, 0x08, 0xcd, 0xf9, 0x32, 0x54, 0x46, 0xef, 0x9d, 0xf6, 0xc7,
	0xd8, 0xdd, 0xe5, 0xfc, 0x8b, 0xe4, 0x79, 0xea, 0xd3, 0x0a, 0x7b, 0x03, 0xba, 0x71, 0x38, 0xd5,
	0xa4, 0x28, 0xb9, 0xf1, 0x1b, 0xa0, 0xe8, 0xa0, 0xf4, 0xa8, 0x91, 0xc7, 0x46, 0x70, 0x3d, 0xf5,
	0x3e, 0x83, 0x8e, 0xd9, 0x88, 0xd2, 0x5a, 0x45, 0xc4, 0x68, 0xc3, 0x97, 0x43, 0xa2, 0x2a, 0x8d,
	0x05, 0xd1, 0x87, 0xd6, 0x95, 0x63, 0x59, 0x09, 0x79, 0x44, 0xe4, 0x35, 0x7c, 0x1c, 0x79, 0x03,
	0xe8, 0x3d, 0x9b, 0x89, 0x42, 0x8b, 0xeb, 0x85, 0xd0, 0x57, 0xd3, 0x3b, 0x2a, 0x75, 0x82, 0xe7,
	0x0f, 0x8a, 0x20, 0x0c, 0x04, 0x97, 0xf1, 0x1a, 0x26, 0xcd, 0x33, 0x0d, 0xfa, 0x9b, 0x65, 0xef,
	0xaf, 0x3a, 0x74, 0x0c, 0xbe, 0x96, 0xb3, 0x66, 0xc9, 0xf9, 0x36, 0xf4, 0xb5, 0x7e, 0x53, 0x4b,
	0xfe, 0x9e, 0xc6, 0xd6, 0x5f, 0x02, 0xb9, 0xd4, 0xb0, 0x9c, 0x71, 0x1f, 0xba, 0xf2, 0xbf, 0xc8,
	0x82, 0xc8, 0xf4, 0xc1, 0x0d, 0x70, 0x4b, 0x27, 0xe0, 0x53, 0x2f, 0x4c, 0xb5, 0xe3, 0x53, 0x72,
	0x6c, 0xbb, 0xa3, 0x53, 0x76, 0x07, 0x0a, 0x6f, 0xd5, 0x9b, 0x11, 0x7e, 0x53, 0x6a, 0xdb, 0xf6,
	0x81, 0x0a, 0xfb, 0x58, 0xcd, 0xa8, 0x57, 0x6a, 0x46, 0x3b, 0x75, 0xbf, 0xe9, 0xb8, 0xc1, 0x96,
	0xe3, 0x0e, 0xd7, 0xba, 0x0d, 0x55, 0xef, 0x50, 0x33, 0xef, 0x5d, 0xd8, 0x37, 0x71, 0x2d, 0xd9,
	0xd7, 0x74, 0xd6, 0xac, 0xca, 0xfd, 0x16, 0x0e, 0xca, 0x8f, 0xde, 0xd1, 0x12, 0x68, 0xfa, 0x05,
	0xbe, 0x67, 0xbb, 0xc1, 0xec, 0xe7, 0xd3, 0x8a, 0xf7, 0x27, 0x1a, 0x61, 0xcd, 0x42, 0xc5, 0x11,
	0xd6, 0xe6, 0xa8, 0x5b, 0xe6, 0x40, 0xb6, 0xb0, 0x52, 0x2c, 0xf1, 0x9d, 0x38, 0x24, 0x4b, 0x3c,
	0x00, 0xc0, 0x85, 0x72, 0x1f, 0xc0, 0xa2, 0x7a, 0xaa, 0x3b, 0xc1, 0x08, 0x3a, 0x78, 0x2f, 0x5a,
	0xcd, 0x62, 0x5c, 0x54, 0x1e, 0x58, 0xcf, 0xe5, 0xab, 0xd1, 0x22, 0x5d, 0xc6, 0x53, 0xd9, 0x7d,
	0xb5, 0x1d, 0xba, 0x84, 0xe0, 0x05, 0x2a, 0x96, 0x64, 0xaa, 0xb6, 0x6d, 0x3e, 0x7e, 0x6a, 0x66,
	0x35, 0xe8, 0x4e, 0xa9, 0x41, 0xe3, 0xb1, 0x5f, 0xa5, 0x89, 0x69, 0xf9, 0x34, 0xb6, 0x04, 0x01,
	0x5b, 0x10, 0xa9, 0x31, 0x1e, 0x23, 0xe2, 0xc2, 0x94, 0xbd, 0x99, 0xca, 0x43, 0xe5, 0xea, 0x3b,
	0x38, 0x8d, 0x43, 0x6d, 0x80, 0xae, 0x46, 0xce, 0x42, 0x59, 0x24, 0x66, 0x59, 0xba, 0x4b, 0x7b,
	0xa0, 0xa7, 0x31, 0x49, 0xab, 0xed, 0x9f, 0x61, 0xc9, 0x3f, 0xa7, 0xbf, 0xb7, 0xa0, 0x6b, 0x4a,
	0x50, 0xb0, 0x18, 0x1c, 0x75, 0x51, 0x64, 0xff, 0x93, 0x2a, 0x95, 0x2e, 0x8d, 0xa3, 0x03, 0x23,
	0xef, 0x37, 0xe9, 0x2c, 0x36, 0x26, 0xf0, 0x1e, 0xff, 0xf0, 0xeb, 0x1f, 0x3f, 0xd5, 0xdf, 0x1b,
	0x3d, 0xc2, 0x4b, 0x69, 0x26, 0x22, 0x2c, 0x13, 0xba, 0x9d, 0xc6, 0xe1, 0x64, 0xf5, 0xc1, 0xf8,
	0xf1, 0xe4, 0x4a, 0xcb, 0x70, 0x3d, 0xb9, 0x92, 0x2f, 0x5f, 0x7f, 0x52, 0x3b, 0x61, 0x97, 0xe0,
	0xa8, 0xfb, 0x94, 0x8a, 0x52, 0xba, 0x5b, 0xed, 0x88, 0xf2, 0x21, 0x45, 0x99, 0x9c, 0xbc, 0x7f,
	0xbb, 0x28, 0x93, 0x2b, 0x69, 0x8b, 0x6b, 0x36, 0x07, 0x47, 0x97, 0x01, 0x45, 0x2a, 0xdd, 0xb8,
	0x76, 0x44, 0xfa, 0x88, 0x22, 0x9d, 0x7a, 0x3b, 0x22, 0xa9, 0x3a, 0xab, 0x4c, 0xeb, 0x25, 0xf4,
	0xed, 0x32, 0x61, 0x47, 0xb6, 0xd1, 0xad, 0x1a, 0x1b, 0xb9, 0xdb, 0x0b, 0x3a, 0xf8, 0x98, 0x82,
	0x1f, 0xb3, 0x1d, 0x64, 0x9a, 0xe6, 0x21, 0x4c, 0x7e, 0x09, 0xb4, 0xf5, 0x0d, 0x89, 0x31, 0xb9,
	0x69, 0xf9, 0xba, 0xb4, 0x23, 0xc3, 0x8f, 0x29, 0xc8, 0x93, 0xd1, 0xb8, 0x3a, 0x88, 0xf6, 0x4d,
	0x65, 0x8a, 0x1c, 0xfb, 0xb5, 0xfe, 0x84, 0xb3, 0x7d, 0xa5, 0x5d, 0xe9, 0x4a, 0x20, 0x23, 0xda,
	0xe0, 0xed, 0xd2, 0xb2, 0x22, 0xa9, 0xb4, 0x9e, 0x41, 0x93, 0x18, 0xdc, 0x2f, 0x9f, 0x5f, 0x85,
	0xd8, 0x93, 0x21, 0x4a, 0xac, 0x3d, 0xa0, 0xed, 0x8f, 0xd8, 0xff, 0x2b, 0xb7, 0x0f, 0x1d, 0xfa,
	0x15, 0xf4, 0xe4, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xe7, 0x03, 0x47, 0x4b, 0x0d, 0x00,
	0x00,
}
