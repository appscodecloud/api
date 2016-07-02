// Code generated by protoc-gen-go.
// source: database.proto
// DO NOT EDIT!

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	database.proto

It has these top-level messages:
	Database
	Snapshot
	SnapshotInfo
	KubeInfo
	ListRequest
	ListResponse
	CreateRequest
	AddStandbyRequest
	DescribeRequest
	DescribeResponse
	DeleteRequest
	BackupScheduleRequest
	BackupUnscheduleRequest
	RestoreRequest
	SnapshotListRequest
	SnapshotListResponse
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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	LastBackup   string `protobuf:"bytes,12,opt,name=last_backup,json=lastBackup" json:"last_backup,omitempty"`
	Status       string `protobuf:"bytes,13,opt,name=status" json:"status,omitempty"`
	Created      string `protobuf:"bytes,14,opt,name=created" json:"created,omitempty"`
	Deleted      string `protobuf:"bytes,15,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Snapshot struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Phid              string `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	DbName            string `protobuf:"bytes,3,opt,name=db_name,json=dbName" json:"db_name,omitempty"`
	DbType            string `protobuf:"bytes,4,opt,name=db_type,json=dbType" json:"db_type,omitempty"`
	DbCluster         string `protobuf:"bytes,5,opt,name=db_cluster,json=dbCluster" json:"db_cluster,omitempty"`
	Provider          string `protobuf:"bytes,6,opt,name=provider" json:"provider,omitempty"`
	CloudCred         string `protobuf:"bytes,7,opt,name=cloud_cred,json=cloudCred" json:"cloud_cred,omitempty"`
	Bucket            string `protobuf:"bytes,8,opt,name=bucket" json:"bucket,omitempty"`
	Region            string `protobuf:"bytes,9,opt,name=region" json:"region,omitempty"`
	Zone              string `protobuf:"bytes,10,opt,name=zone" json:"zone,omitempty"`
	Status            string `protobuf:"bytes,11,opt,name=status" json:"status,omitempty"`
	Process           string `protobuf:"bytes,12,opt,name=process" json:"process,omitempty"`
	Type              string `protobuf:"bytes,13,opt,name=type" json:"type,omitempty"`
	IsScheduledBackup int32  `protobuf:"varint,14,opt,name=is_scheduled_backup,json=isScheduledBackup" json:"is_scheduled_backup,omitempty"`
	RestoreDb         string `protobuf:"bytes,15,opt,name=restore_db,json=restoreDb" json:"restore_db,omitempty"`
	RestoreSnap       string `protobuf:"bytes,16,opt,name=restore_snap,json=restoreSnap" json:"restore_snap,omitempty"`
	Created           string `protobuf:"bytes,17,opt,name=created" json:"created,omitempty"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SnapshotInfo struct {
	ScheduleCronExpr string `protobuf:"bytes,1,opt,name=schedule_cron_expr,json=scheduleCronExpr" json:"schedule_cron_expr,omitempty"`
	LastBackup       string `protobuf:"bytes,2,opt,name=last_backup,json=lastBackup" json:"last_backup,omitempty"`
	BackupAttempt    int64  `protobuf:"varint,3,opt,name=backup_attempt,json=backupAttempt" json:"backup_attempt,omitempty"`
	BackupSuccess    int64  `protobuf:"varint,4,opt,name=backup_success,json=backupSuccess" json:"backup_success,omitempty"`
	RestoreAttempt   int64  `protobuf:"varint,5,opt,name=restore_attempt,json=restoreAttempt" json:"restore_attempt,omitempty"`
	RestoreSuccess   int64  `protobuf:"varint,6,opt,name=restore_success,json=restoreSuccess" json:"restore_success,omitempty"`
}

func (m *SnapshotInfo) Reset()                    { *m = SnapshotInfo{} }
func (m *SnapshotInfo) String() string            { return proto.CompactTextString(m) }
func (*SnapshotInfo) ProtoMessage()               {}
func (*SnapshotInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type KubeInfo struct {
	Svc  int64 `protobuf:"varint,1,opt,name=svc" json:"svc,omitempty"`
	Pods int64 `protobuf:"varint,2,opt,name=pods" json:"pods,omitempty"`
	Rc   int64 `protobuf:"varint,3,opt,name=rc" json:"rc,omitempty"`
}

func (m *KubeInfo) Reset()                    { *m = KubeInfo{} }
func (m *KubeInfo) String() string            { return proto.CompactTextString(m) }
func (*KubeInfo) ProtoMessage()               {}
func (*KubeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListResponse struct {
	Status           *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Databases        []*Database    `protobuf:"bytes,2,rep,name=databases" json:"databases,omitempty"`
	DeletedDatabases []*Database    `protobuf:"bytes,3,rep,name=deleted_databases,json=deletedDatabases" json:"deleted_databases,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

func (m *ListResponse) GetDeletedDatabases() []*Database {
	if m != nil {
		return m.DeletedDatabases
	}
	return nil
}

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
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type AddStandbyRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Config  string `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *AddStandbyRequest) Reset()                    { *m = AddStandbyRequest{} }
func (m *AddStandbyRequest) String() string            { return proto.CompactTextString(m) }
func (*AddStandbyRequest) ProtoMessage()               {}
func (*AddStandbyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DescribeResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name         string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Phid         string         `protobuf:"bytes,3,opt,name=phid" json:"phid,omitempty"`
	Cluster      string         `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
	Provider     string         `protobuf:"bytes,5,opt,name=provider" json:"provider,omitempty"`
	Type         string         `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	Sku          string         `protobuf:"bytes,7,opt,name=sku" json:"sku,omitempty"`
	Version      string         `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`
	SnapshotInfo *SnapshotInfo  `protobuf:"bytes,9,opt,name=snapshot_info,json=snapshotInfo" json:"snapshot_info,omitempty"`
	Kube         *KubeInfo      `protobuf:"bytes,10,opt,name=kube" json:"kube,omitempty"`
	DbStatus     string         `protobuf:"bytes,11,opt,name=db_status,json=dbStatus" json:"db_status,omitempty"`
	AuthSecret   string         `protobuf:"bytes,12,opt,name=auth_secret,json=authSecret" json:"auth_secret,omitempty"`
	Created      string         `protobuf:"bytes,13,opt,name=created" json:"created,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetSnapshotInfo() *SnapshotInfo {
	if m != nil {
		return m.SnapshotInfo
	}
	return nil
}

func (m *DescribeResponse) GetKube() *KubeInfo {
	if m != nil {
		return m.Kube
	}
	return nil
}

type DeleteRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type BackupScheduleRequest struct {
	Cluster          string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type             string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	BucketName       string `protobuf:"bytes,4,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	SnapshotName     string `protobuf:"bytes,5,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	Credential       string `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Region           string `protobuf:"bytes,7,opt,name=region" json:"region,omitempty"`
	AuthSecretName   string `protobuf:"bytes,8,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	ScheduleCronExpr string `protobuf:"bytes,9,opt,name=schedule_cron_expr,json=scheduleCronExpr" json:"schedule_cron_expr,omitempty"`
	Wal              bool   `protobuf:"varint,10,opt,name=wal" json:"wal,omitempty"`
	Dump             bool   `protobuf:"varint,11,opt,name=dump" json:"dump,omitempty"`
	Force            bool   `protobuf:"varint,12,opt,name=force" json:"force,omitempty"`
}

func (m *BackupScheduleRequest) Reset()                    { *m = BackupScheduleRequest{} }
func (m *BackupScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupScheduleRequest) ProtoMessage()               {}
func (*BackupScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type BackupUnscheduleRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *BackupUnscheduleRequest) Reset()                    { *m = BackupUnscheduleRequest{} }
func (m *BackupUnscheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupUnscheduleRequest) ProtoMessage()               {}
func (*BackupUnscheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

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
	Dump           bool   `protobuf:"varint,15,opt,name=dump" json:"dump,omitempty"`
	Force          bool   `protobuf:"varint,16,opt,name=force" json:"force,omitempty"`
}

func (m *RestoreRequest) Reset()                    { *m = RestoreRequest{} }
func (m *RestoreRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreRequest) ProtoMessage()               {}
func (*RestoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type SnapshotListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *SnapshotListRequest) Reset()                    { *m = SnapshotListRequest{} }
func (m *SnapshotListRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListRequest) ProtoMessage()               {}
func (*SnapshotListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type SnapshotListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	List   []*Snapshot    `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
}

func (m *SnapshotListResponse) Reset()                    { *m = SnapshotListResponse{} }
func (m *SnapshotListResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListResponse) ProtoMessage()               {}
func (*SnapshotListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

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

func init() {
	proto.RegisterType((*Database)(nil), "db.Database")
	proto.RegisterType((*Snapshot)(nil), "db.Snapshot")
	proto.RegisterType((*SnapshotInfo)(nil), "db.SnapshotInfo")
	proto.RegisterType((*KubeInfo)(nil), "db.KubeInfo")
	proto.RegisterType((*ListRequest)(nil), "db.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "db.ListResponse")
	proto.RegisterType((*CreateRequest)(nil), "db.CreateRequest")
	proto.RegisterType((*AddStandbyRequest)(nil), "db.AddStandbyRequest")
	proto.RegisterType((*DescribeRequest)(nil), "db.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "db.DescribeResponse")
	proto.RegisterType((*DeleteRequest)(nil), "db.DeleteRequest")
	proto.RegisterType((*BackupScheduleRequest)(nil), "db.BackupScheduleRequest")
	proto.RegisterType((*BackupUnscheduleRequest)(nil), "db.BackupUnscheduleRequest")
	proto.RegisterType((*RestoreRequest)(nil), "db.RestoreRequest")
	proto.RegisterType((*SnapshotListRequest)(nil), "db.SnapshotListRequest")
	proto.RegisterType((*SnapshotListResponse)(nil), "db.SnapshotListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Databases service

type DatabasesClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	AddStandby(ctx context.Context, in *AddStandbyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	BackupSchedule(ctx context.Context, in *BackupScheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	BackupUnschedule(ctx context.Context, in *BackupUnscheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	SnapshotList(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error)
}

type databasesClient struct {
	cc *grpc.ClientConn
}

func NewDatabasesClient(cc *grpc.ClientConn) DatabasesClient {
	return &databasesClient{cc}
}

func (c *databasesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/db.Databases/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) AddStandby(ctx context.Context, in *AddStandbyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/AddStandby", in, out, c.cc, opts...)
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

func (c *databasesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) BackupSchedule(ctx context.Context, in *BackupScheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/BackupSchedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) BackupUnschedule(ctx context.Context, in *BackupUnscheduleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Databases/BackupUnschedule", in, out, c.cc, opts...)
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

func (c *databasesClient) SnapshotList(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error) {
	out := new(SnapshotListResponse)
	err := grpc.Invoke(ctx, "/db.Databases/SnapshotList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Databases service

type DatabasesServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *CreateRequest) (*dtypes.VoidResponse, error)
	AddStandby(context.Context, *AddStandbyRequest) (*dtypes.VoidResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	Delete(context.Context, *DeleteRequest) (*dtypes.VoidResponse, error)
	BackupSchedule(context.Context, *BackupScheduleRequest) (*dtypes.VoidResponse, error)
	BackupUnschedule(context.Context, *BackupUnscheduleRequest) (*dtypes.VoidResponse, error)
	Restore(context.Context, *RestoreRequest) (*dtypes.VoidResponse, error)
	SnapshotList(context.Context, *SnapshotListRequest) (*SnapshotListResponse, error)
}

func RegisterDatabasesServer(s *grpc.Server, srv DatabasesServer) {
	s.RegisterService(&_Databases_serviceDesc, srv)
}

func _Databases_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
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
		return srv.(DatabasesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _Databases_AddStandby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStandbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).AddStandby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/AddStandby",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).AddStandby(ctx, req.(*AddStandbyRequest))
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

func _Databases_BackupSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).BackupSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/BackupSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).BackupSchedule(ctx, req.(*BackupScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_BackupUnschedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupUnscheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).BackupUnschedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Databases/BackupUnschedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).BackupUnschedule(ctx, req.(*BackupUnscheduleRequest))
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

var _Databases_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.Databases",
	HandlerType: (*DatabasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Databases_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Databases_Create_Handler,
		},
		{
			MethodName: "AddStandby",
			Handler:    _Databases_AddStandby_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Databases_Describe_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Databases_Delete_Handler,
		},
		{
			MethodName: "BackupSchedule",
			Handler:    _Databases_BackupSchedule_Handler,
		},
		{
			MethodName: "BackupUnschedule",
			Handler:    _Databases_BackupUnschedule_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _Databases_Restore_Handler,
		},
		{
			MethodName: "SnapshotList",
			Handler:    _Databases_SnapshotList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x58, 0xcd, 0x6f, 0x1c, 0xc5,
	0x12, 0x97, 0x77, 0xf6, 0xb3, 0xf7, 0xc3, 0xeb, 0xce, 0x87, 0xe7, 0x6d, 0x92, 0x97, 0xbc, 0xc9,
	0x7b, 0x79, 0x91, 0x05, 0xbb, 0xc1, 0x51, 0x10, 0x41, 0x42, 0x10, 0x6c, 0x04, 0x08, 0x84, 0xd0,
	0x18, 0x90, 0x38, 0xa0, 0x61, 0x3e, 0x3a, 0xf6, 0xc8, 0xeb, 0x99, 0x61, 0x7b, 0x66, 0x21, 0x09,
	0xb9, 0x70, 0xe3, 0x9c, 0x33, 0xe2, 0x2f, 0xe1, 0xaf, 0x80, 0x13, 0x67, 0x6e, 0xfc, 0x0d, 0x48,
	0x74, 0x55, 0x77, 0xef, 0xf4, 0xec, 0xae, 0x13, 0x27, 0xe4, 0x92, 0x74, 0x57, 0x75, 0x57, 0x55,
	0xd7, 0xaf, 0xea, 0x57, 0xb3, 0x26, 0x83, 0xc8, 0xcf, 0xfd, 0xc0, 0xe7, 0x6c, 0x9c, 0xcd, 0xd2,
	0x3c, 0xa5, 0xb5, 0x28, 0x18, 0x5d, 0x3e, 0x4c, 0xd3, 0xc3, 0x29, 0x9b, 0xf8, 0x59, 0x3c, 0xf1,
	0x93, 0x24, 0xcd, 0xfd, 0x3c, 0x4e, 0x13, 0x2e, 0x4f, 0x8c, 0x2e, 0x82, 0x38, 0xca, 0x1f, 0x64,
	0x8c, 0x4f, 0xf0, 0x5f, 0x29, 0x77, 0x9e, 0x58, 0xa4, 0xbd, 0xaf, 0x8c, 0x51, 0x4a, 0xea, 0xd9,
	0x51, 0x1c, 0xd9, 0x1b, 0xd7, 0x36, 0x6e, 0x76, 0x5c, 0x5c, 0xd3, 0xff, 0x90, 0x5e, 0x38, 0x2d,
	0x78, 0xce, 0x66, 0x5e, 0xe2, 0x9f, 0x30, 0xbb, 0x86, 0xba, 0xae, 0x92, 0x7d, 0x22, 0x44, 0x70,
	0x0d, 0x55, 0x96, 0xbc, 0x06, 0x6b, 0x7a, 0x99, 0x74, 0xe0, 0x7f, 0x9e, 0xf9, 0x21, 0xb3, 0xeb,
	0xa8, 0x28, 0x05, 0x70, 0x03, 0x82, 0xb0, 0x1b, 0xf2, 0x06, 0xac, 0xe9, 0x90, 0x58, 0xfc, 0xb8,
	0xb0, 0x9b, 0x28, 0x82, 0x25, 0x9c, 0x3a, 0x49, 0x23, 0x66, 0xb7, 0xe4, 0x29, 0x58, 0x53, 0x9b,
	0xb4, 0xe6, 0x6c, 0xc6, 0xc5, 0xcb, 0xec, 0x36, 0x8a, 0xf5, 0x96, 0x5e, 0x25, 0x5d, 0xbf, 0xc8,
	0x8f, 0x3c, 0xce, 0xc2, 0x19, 0xcb, 0xed, 0x0e, 0x6a, 0x09, 0x88, 0x0e, 0x50, 0x42, 0xaf, 0x93,
	0xbe, 0xd8, 0xa5, 0x1e, 0x4f, 0xfc, 0x8c, 0x1f, 0xa5, 0xb9, 0x4d, 0xc4, 0x91, 0x86, 0xdb, 0x03,
	0xe1, 0x81, 0x92, 0xd1, 0x6d, 0xd2, 0xca, 0xe6, 0x1e, 0x8f, 0x1f, 0x32, 0xbb, 0x2b, 0xd4, 0x96,
	0xdb, 0xcc, 0xe6, 0x07, 0x62, 0x07, 0xe6, 0xa7, 0x3e, 0xcf, 0xbd, 0xc0, 0x0f, 0x8f, 0x8b, 0xcc,
	0xee, 0x49, 0xf3, 0x20, 0x7a, 0x17, 0x25, 0xf4, 0x22, 0x69, 0x72, 0x91, 0xf3, 0x82, 0xdb, 0x7d,
	0xd4, 0xa9, 0x1d, 0x44, 0x2c, 0xdc, 0xfb, 0x39, 0x8b, 0xec, 0x81, 0x8c, 0x58, 0x6d, 0x41, 0x13,
	0xb1, 0x29, 0x03, 0xcd, 0xa6, 0xd4, 0xa8, 0xad, 0xf3, 0x9b, 0x40, 0x65, 0x11, 0x92, 0x4e, 0xef,
	0x86, 0x91, 0x5e, 0x8d, 0x54, 0xcd, 0x40, 0x4a, 0x84, 0x1e, 0x05, 0x9e, 0x81, 0x44, 0x33, 0x0a,
	0x10, 0x1f, 0xa9, 0xc0, 0x84, 0xd7, 0xb5, 0xe2, 0x33, 0x48, 0xf9, 0x15, 0x42, 0x84, 0x42, 0x41,
	0xa9, 0xc0, 0xe8, 0x44, 0xc1, 0x9e, 0x14, 0xd0, 0x11, 0x69, 0x8b, 0x22, 0x99, 0xc7, 0x91, 0x50,
	0x4a, 0x58, 0x16, 0x7b, 0xb8, 0x1a, 0x4e, 0xd3, 0x22, 0xf2, 0xc4, 0x63, 0x22, 0x85, 0x50, 0x07,
	0x25, 0x7b, 0x42, 0x00, 0xc9, 0x08, 0x8a, 0xf0, 0x58, 0xe0, 0x20, 0x51, 0x52, 0x3b, 0x90, 0xcf,
	0xd8, 0x21, 0xa0, 0x27, 0xf1, 0x51, 0x3b, 0x78, 0xcf, 0xc3, 0x34, 0x61, 0x08, 0x89, 0x78, 0x0f,
	0xac, 0x8d, 0x84, 0x76, 0x97, 0x13, 0x2a, 0xc2, 0x08, 0x19, 0xe7, 0x0a, 0x05, 0xbd, 0x5d, 0x94,
	0x55, 0xdf, 0x28, 0xab, 0x31, 0x39, 0x17, 0x73, 0x8f, 0x87, 0x47, 0x2c, 0x2a, 0xa6, 0x2c, 0xd2,
	0xf8, 0x0d, 0x10, 0xfb, 0xad, 0x98, 0x1f, 0x68, 0x8d, 0x82, 0x51, 0x3c, 0x6c, 0xc6, 0x78, 0x9e,
	0xce, 0x98, 0x17, 0x05, 0x0a, 0x97, 0x8e, 0x92, 0xec, 0x07, 0xd0, 0x0e, 0x5a, 0x0d, 0x75, 0x64,
	0x0f, 0x65, 0x3b, 0x28, 0x19, 0x60, 0x66, 0x02, 0xbe, 0x55, 0x01, 0xdc, 0xf9, 0x6b, 0x83, 0xf4,
	0x34, 0xac, 0x1f, 0x26, 0xf7, 0x53, 0xfa, 0x0a, 0xa1, 0x3a, 0x32, 0x91, 0xc8, 0x34, 0xf1, 0xd8,
	0x77, 0xd9, 0x4c, 0x01, 0x3d, 0xd4, 0x9a, 0x3d, 0xa1, 0x78, 0x4f, 0xc8, 0x97, 0x4b, 0xb0, 0xb6,
	0x52, 0x82, 0xff, 0x23, 0x03, 0xa9, 0xf3, 0xfc, 0x3c, 0x67, 0x27, 0x59, 0x8e, 0x85, 0x60, 0xb9,
	0x7d, 0x29, 0xbd, 0x27, 0x85, 0xc6, 0x31, 0x5e, 0x84, 0x98, 0xc7, 0xba, 0x79, 0xec, 0x40, 0x0a,
	0xe9, 0xff, 0xc9, 0xa6, 0x7e, 0xaa, 0x36, 0xd7, 0xc0, 0x73, 0x03, 0x25, 0xd6, 0xf6, 0x8c, 0x83,
	0xda, 0x60, 0xb3, 0x72, 0x50, 0x59, 0x74, 0xde, 0x21, 0xed, 0x8f, 0x8a, 0x80, 0xe1, 0xd3, 0xa1,
	0xdd, 0xe7, 0x21, 0xbe, 0xd5, 0x72, 0x61, 0x89, 0x35, 0x9d, 0x46, 0x1c, 0xdf, 0x65, 0xb9, 0xb8,
	0xa6, 0x03, 0x52, 0x9b, 0x85, 0xea, 0x15, 0x62, 0xe5, 0xf4, 0x49, 0xf7, 0xe3, 0x98, 0xe7, 0x2e,
	0xfb, 0xa6, 0x10, 0xa6, 0x9d, 0x9f, 0x44, 0x42, 0xe5, 0x9e, 0x67, 0x82, 0xeb, 0x18, 0xbd, 0xb1,
	0xa8, 0x19, 0x30, 0xdc, 0xdd, 0x1d, 0x8c, 0x25, 0xe7, 0x8d, 0x0f, 0x50, 0xba, 0xa8, 0xa1, 0x1d,
	0xd2, 0xd1, 0x14, 0x0a, 0x0e, 0x2d, 0x71, 0xb4, 0x37, 0x8e, 0x82, 0xb1, 0xa6, 0x42, 0xb7, 0x54,
	0xd3, 0xbb, 0x64, 0x4b, 0xf5, 0xa5, 0x57, 0xde, 0xb1, 0xd6, 0xdc, 0x19, 0xaa, 0x63, 0x5a, 0xc0,
	0x9d, 0x5f, 0x6a, 0xa4, 0xbf, 0x87, 0xe0, 0xab, 0x88, 0xb1, 0x38, 0x54, 0xbf, 0x6d, 0xa8, 0xe2,
	0x50, 0xdd, 0xa6, 0xdb, 0xbc, 0x56, 0x6d, 0x73, 0x2c, 0x68, 0x6b, 0x95, 0x27, 0xeb, 0x25, 0x4f,
	0x1a, 0x9c, 0xd8, 0xa8, 0x72, 0xa2, 0x48, 0x5f, 0x36, 0x57, 0xbd, 0x2b, 0x56, 0x26, 0xbb, 0xb5,
	0x2a, 0xec, 0x76, 0x93, 0x0c, 0x0d, 0xf2, 0x94, 0x24, 0x22, 0x3b, 0x77, 0x50, 0x32, 0xe8, 0x82,
	0xec, 0x81, 0x94, 0x3b, 0x12, 0x25, 0x58, 0xd3, 0x7f, 0x0b, 0x32, 0x10, 0x5d, 0xcf, 0x92, 0x3c,
	0xf6, 0xa7, 0xaa, 0x87, 0x0d, 0x09, 0x14, 0xae, 0xec, 0x7f, 0x69, 0x58, 0xb6, 0x33, 0x91, 0x22,
	0x34, 0x5a, 0xd2, 0x42, 0xcf, 0xa4, 0x05, 0xe7, 0x4b, 0xb2, 0x75, 0x2f, 0x8a, 0x04, 0x76, 0x49,
	0x14, 0x3c, 0x78, 0xb1, 0x14, 0x0a, 0xd3, 0x61, 0x9a, 0xdc, 0x8f, 0x0f, 0x35, 0x29, 0xca, 0x9d,
	0xf3, 0x36, 0xd9, 0xdc, 0x67, 0x3c, 0x9c, 0xc5, 0xc1, 0x8b, 0x61, 0xe3, 0xfc, 0x68, 0x91, 0x61,
	0x69, 0xe1, 0x39, 0xeb, 0xef, 0x14, 0xb0, 0x91, 0xd3, 0x2d, 0x83, 0xd3, 0x8d, 0x90, 0xea, 0xd5,
	0x90, 0x4c, 0x72, 0x6e, 0x2c, 0x91, 0xb3, 0x2e, 0x9b, 0xe6, 0x6a, 0xd9, 0xb4, 0xd6, 0x96, 0xcd,
	0xd2, 0x28, 0xbd, 0x43, 0xfa, 0x7a, 0x48, 0x7a, 0xb1, 0x68, 0x56, 0x04, 0xbb, 0xbb, 0x3b, 0x84,
	0x6a, 0x37, 0xf9, 0xcb, 0xed, 0x71, 0x93, 0xcd, 0xae, 0x91, 0xfa, 0xb1, 0x68, 0x6f, 0x2c, 0x00,
	0xd5, 0x1b, 0xba, 0xdd, 0x5d, 0xd4, 0xd0, 0x4b, 0xa2, 0xed, 0x02, 0xaf, 0xc2, 0xea, 0xed, 0x28,
	0x90, 0xb9, 0x59, 0x1e, 0xe0, 0xbd, 0x95, 0x01, 0x6e, 0x10, 0x6b, 0xbf, 0x4a, 0xac, 0x6f, 0x91,
	0xfe, 0x3e, 0xf6, 0xde, 0x8b, 0x41, 0xf9, 0x67, 0x8d, 0x5c, 0x90, 0x14, 0xaa, 0xa7, 0xc1, 0xcb,
	0x6b, 0xd7, 0xa5, 0xda, 0xaf, 0xaf, 0xd4, 0xfe, 0x75, 0x23, 0xd9, 0x78, 0x44, 0xa2, 0xb9, 0x48,
	0x2d, 0x1e, 0xaa, 0x76, 0x58, 0x73, 0xa5, 0xc3, 0xca, 0x06, 0x6a, 0x55, 0xe6, 0xea, 0xd9, 0xfb,
	0x7a, 0xfd, 0x28, 0xea, 0x9c, 0x32, 0x8a, 0x44, 0x35, 0x7d, 0xab, 0x5a, 0xbd, 0xed, 0xc2, 0x12,
	0xde, 0x1e, 0x15, 0x27, 0x19, 0xa2, 0xda, 0x76, 0x71, 0x4d, 0xcf, 0x93, 0xc6, 0xfd, 0x74, 0x26,
	0x3e, 0x00, 0x7b, 0x28, 0x94, 0x1b, 0xe7, 0x7d, 0xb2, 0x2d, 0x93, 0xfd, 0x79, 0xc2, 0xff, 0x49,
	0xba, 0x9d, 0x9f, 0x2d, 0x32, 0x70, 0xe5, 0x84, 0x79, 0x79, 0x78, 0x49, 0xca, 0xac, 0xaf, 0xa3,
	0xcc, 0x46, 0x85, 0x32, 0xd7, 0x7e, 0xaf, 0x26, 0xfa, 0x7b, 0x55, 0x53, 0xa3, 0x89, 0x2e, 0x76,
	0x77, 0xbb, 0x8a, 0xee, 0xa7, 0xd0, 0xe5, 0x4b, 0x35, 0xd2, 0x59, 0xa9, 0x91, 0x67, 0x11, 0x6c,
	0x09, 0x7f, 0xb7, 0x02, 0xbf, 0xd1, 0xe2, 0xbd, 0x6a, 0x8b, 0xaf, 0x2b, 0x8c, 0xfe, 0xda, 0xc2,
	0x50, 0x50, 0x0f, 0x56, 0xa1, 0xde, 0x5c, 0x07, 0xf5, 0xd0, 0x84, 0xfa, 0x1e, 0x39, 0xa7, 0xf9,
	0xc2, 0x18, 0xdb, 0x4f, 0x41, 0x49, 0x38, 0x2b, 0x16, 0x9f, 0xb5, 0xb0, 0x74, 0xbe, 0x26, 0xe7,
	0xab, 0x26, 0x9e, 0x93, 0x69, 0x05, 0x29, 0x4d, 0xc5, 0x3d, 0x73, 0xc8, 0x6b, 0x7b, 0x2e, 0x6a,
	0x76, 0x7f, 0x6f, 0x91, 0xce, 0x62, 0x64, 0xd3, 0x0f, 0x48, 0x1d, 0xfc, 0xd0, 0x4d, 0x38, 0x69,
	0x04, 0x3d, 0x1a, 0x96, 0x02, 0x19, 0x82, 0x73, 0xe5, 0x87, 0x5f, 0xff, 0x78, 0x52, 0xdb, 0xa6,
	0x17, 0xc4, 0x6f, 0xae, 0x8c, 0x87, 0x02, 0x6e, 0xfc, 0xf1, 0x15, 0x05, 0x93, 0xf9, 0xad, 0xf1,
	0x6b, 0xf4, 0x2b, 0xd2, 0x94, 0xb3, 0x9f, 0x6e, 0xc1, 0xd5, 0xca, 0x77, 0xc0, 0xe8, 0xbc, 0x0e,
	0xf7, 0x8b, 0x34, 0x8e, 0x16, 0x16, 0x77, 0xd0, 0xe2, 0x7f, 0x47, 0x57, 0xd7, 0x5a, 0x9c, 0x3c,
	0x52, 0x69, 0x7a, 0xfc, 0xe6, 0xc6, 0x0e, 0xcd, 0x09, 0x29, 0x67, 0x23, 0xbd, 0x00, 0x2e, 0x56,
	0x66, 0xe5, 0x29, 0x6e, 0xee, 0xa2, 0x9b, 0xdb, 0xce, 0xf8, 0x19, 0x6e, 0x26, 0x8f, 0xa0, 0x2e,
	0x1e, 0x4f, 0xb8, 0x34, 0x0a, 0x5e, 0x99, 0xf8, 0xb9, 0xa8, 0x86, 0x1e, 0x3d, 0x87, 0x5f, 0x3f,
	0xd5, 0x21, 0x0a, 0x1e, 0x4d, 0xa1, 0xf2, 0x38, 0x46, 0x8f, 0x37, 0xe9, 0x8d, 0xb3, 0x79, 0xa4,
	0x3e, 0x69, 0x4a, 0x42, 0x97, 0xb9, 0xab, 0x90, 0xfb, 0x29, 0x8f, 0x52, 0x2e, 0x76, 0xce, 0xea,
	0xe2, 0x11, 0x19, 0x54, 0x39, 0x9f, 0xfe, 0x0b, 0x5c, 0xad, 0x9d, 0x03, 0xa7, 0xb8, 0x7c, 0x03,
	0x5d, 0xee, 0x3a, 0xaf, 0x9e, 0x31, 0x8f, 0xf2, 0xfb, 0x1a, 0xd2, 0xf8, 0x3d, 0x19, 0x2e, 0x73,
	0x20, 0xbd, 0x54, 0xba, 0x5f, 0x61, 0xc6, 0x53, 0x02, 0xb8, 0x83, 0x01, 0x4c, 0x46, 0xcf, 0x17,
	0x00, 0x4d, 0x48, 0x4b, 0xf1, 0x26, 0xa5, 0xe0, 0xb4, 0x4a, 0xa2, 0x4f, 0x2f, 0x9a, 0xd1, 0x59,
	0x8b, 0x46, 0x7d, 0xfb, 0xc3, 0x6b, 0x1f, 0x94, 0x3f, 0x7b, 0xb0, 0xb7, 0xb6, 0xcd, 0x2e, 0x34,
	0x7b, 0xcc, 0x5e, 0x55, 0x28, 0xef, 0xaf, 0xa3, 0xf7, 0x5b, 0xf4, 0xd9, 0xde, 0x05, 0x69, 0x88,
	0x8a, 0x55, 0x36, 0x82, 0x26, 0xfe, 0x99, 0xe3, 0xf6, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e,
	0x70, 0x92, 0x0c, 0x32, 0x11, 0x00, 0x00,
}