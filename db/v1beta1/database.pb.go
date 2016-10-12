// Code generated by protoc-gen-go.
// source: database.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	database.proto
	snapshot.proto

It has these top-level messages:
	Database
	DatabaseListRequest
	DatabaseListResponse
	DatabaseCreateRequest
	DatabaseScaleRequest
	DatabaseUpdateRequest
	DatabaseDescribeRequest
	SnapshotSummary
	DatabaseDescribeResponse
	DatabaseDeleteRequest
	DatabaseRecoverRequest
	DatabaseCheckRequest
	DatabaseInfo
	Snapshot
	SnapshotListRequest
	SnapshotListResponse
	BackupScheduleRequest
	BackupUnscheduleRequest
	SnapshotRestoreRequest
	SnapshotCheckRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
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

// Next Id: 19
type Database struct {
	Phid             string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Cluster          string `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type             string `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	Sku              string `protobuf:"bytes,6,opt,name=sku" json:"sku,omitempty"`
	Version          string `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`
	AuthSecret       string `protobuf:"bytes,9,opt,name=auth_secret,json=authSecret" json:"auth_secret,omitempty"`
	ScheduleCronExpr string `protobuf:"bytes,10,opt,name=schedule_cron_expr,json=scheduleCronExpr" json:"schedule_cron_expr,omitempty"`
	PvSizeGb         int32  `protobuf:"varint,11,opt,name=pv_size_gb,json=pvSizeGb" json:"pv_size_gb,omitempty"`
	LastBackupAt     int64  `protobuf:"varint,12,opt,name=last_backup_at,json=lastBackupAt" json:"last_backup_at,omitempty"`
	Status           string `protobuf:"bytes,13,opt,name=status" json:"status,omitempty"`
	CreatedAt        int64  `protobuf:"varint,14,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	DeletedAt        int64  `protobuf:"varint,15,opt,name=deleted_at,json=deletedAt" json:"deleted_at,omitempty"`
	DoNotDelete      bool   `protobuf:"varint,17,opt,name=do_not_delete,json=doNotDelete" json:"do_not_delete,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Next Id: 4
type DatabaseListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   READY
	//   DELETING
	//   DELETED
	//   DESTROYED
	Status []string `protobuf:"bytes,3,rep,name=status" json:"status,omitempty"`
}

func (m *DatabaseListRequest) Reset()                    { *m = DatabaseListRequest{} }
func (m *DatabaseListRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseListRequest) ProtoMessage()               {}
func (*DatabaseListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Next Id: 3
type DatabaseListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Databases []*Database    `protobuf:"bytes,2,rep,name=databases" json:"databases,omitempty"`
}

func (m *DatabaseListResponse) Reset()                    { *m = DatabaseListResponse{} }
func (m *DatabaseListResponse) String() string            { return proto.CompactTextString(m) }
func (*DatabaseListResponse) ProtoMessage()               {}
func (*DatabaseListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DatabaseListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DatabaseListResponse) GetDatabases() []*Database {
	if m != nil {
		return m.Databases
	}
	return nil
}

// Next Id: 18
type DatabaseCreateRequest struct {
	Cluster          string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type             string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Sku              string `protobuf:"bytes,4,opt,name=sku" json:"sku,omitempty"`
	Version          string `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	PvSizeGb         int32  `protobuf:"varint,7,opt,name=pv_size_gb,json=pvSizeGb" json:"pv_size_gb,omitempty"`
	AuthSecretName   string `protobuf:"bytes,8,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	Size             int32  `protobuf:"varint,9,opt,name=size" json:"size,omitempty"`
	SnapshotPhid     string `protobuf:"bytes,14,opt,name=snapshot_phid,json=snapshotPhid" json:"snapshot_phid,omitempty"`
	Hostname         string `protobuf:"bytes,15,opt,name=hostname" json:"hostname,omitempty"`
	StorageClass     string `protobuf:"bytes,16,opt,name=storage_class,json=storageClass" json:"storage_class,omitempty"`
	IgnoreValidation bool   `protobuf:"varint,17,opt,name=ignore_validation,json=ignoreValidation" json:"ignore_validation,omitempty"`
}

func (m *DatabaseCreateRequest) Reset()                    { *m = DatabaseCreateRequest{} }
func (m *DatabaseCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseCreateRequest) ProtoMessage()               {}
func (*DatabaseCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Next Id: 4
type DatabaseScaleRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	Size    int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *DatabaseScaleRequest) Reset()                    { *m = DatabaseScaleRequest{} }
func (m *DatabaseScaleRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseScaleRequest) ProtoMessage()               {}
func (*DatabaseScaleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Next Id: 4
type DatabaseUpdateRequest struct {
	Cluster     string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid         string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	DoNotDelete bool   `protobuf:"varint,3,opt,name=do_not_delete,json=doNotDelete" json:"do_not_delete,omitempty"`
}

func (m *DatabaseUpdateRequest) Reset()                    { *m = DatabaseUpdateRequest{} }
func (m *DatabaseUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseUpdateRequest) ProtoMessage()               {}
func (*DatabaseUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Next Id: 3
type DatabaseDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *DatabaseDescribeRequest) Reset()                    { *m = DatabaseDescribeRequest{} }
func (m *DatabaseDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseDescribeRequest) ProtoMessage()               {}
func (*DatabaseDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Next Id: 7
type SnapshotSummary struct {
	BackupAttempt  int32 `protobuf:"varint,3,opt,name=backup_attempt,json=backupAttempt" json:"backup_attempt,omitempty"`
	BackupSuccess  int32 `protobuf:"varint,4,opt,name=backup_success,json=backupSuccess" json:"backup_success,omitempty"`
	RestoreAttempt int32 `protobuf:"varint,5,opt,name=restore_attempt,json=restoreAttempt" json:"restore_attempt,omitempty"`
	RestoreSuccess int32 `protobuf:"varint,6,opt,name=restore_success,json=restoreSuccess" json:"restore_success,omitempty"`
}

func (m *SnapshotSummary) Reset()                    { *m = SnapshotSummary{} }
func (m *SnapshotSummary) String() string            { return proto.CompactTextString(m) }
func (*SnapshotSummary) ProtoMessage()               {}
func (*SnapshotSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Next Id: 17
type DatabaseDescribeResponse struct {
	Status          *dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	SnapshotSummary *SnapshotSummary `protobuf:"bytes,9,opt,name=snapshot_summary,json=snapshotSummary" json:"snapshot_summary,omitempty"`
	Database        *Database        `protobuf:"bytes,16,opt,name=database" json:"database,omitempty"`
}

func (m *DatabaseDescribeResponse) Reset()                    { *m = DatabaseDescribeResponse{} }
func (m *DatabaseDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DatabaseDescribeResponse) ProtoMessage()               {}
func (*DatabaseDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DatabaseDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DatabaseDescribeResponse) GetSnapshotSummary() *SnapshotSummary {
	if m != nil {
		return m.SnapshotSummary
	}
	return nil
}

func (m *DatabaseDescribeResponse) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

// Next Id: 4
type DatabaseDeleteRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	Destroy bool   `protobuf:"varint,3,opt,name=destroy" json:"destroy,omitempty"`
}

func (m *DatabaseDeleteRequest) Reset()                    { *m = DatabaseDeleteRequest{} }
func (m *DatabaseDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseDeleteRequest) ProtoMessage()               {}
func (*DatabaseDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Next Id: 3
type DatabaseRecoverRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *DatabaseRecoverRequest) Reset()                    { *m = DatabaseRecoverRequest{} }
func (m *DatabaseRecoverRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseRecoverRequest) ProtoMessage()               {}
func (*DatabaseRecoverRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Next Id: 4
type DatabaseCheckRequest struct {
	Cluster      string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid          string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	PurchasePhid string `protobuf:"bytes,3,opt,name=purchase_phid,json=purchasePhid" json:"purchase_phid,omitempty"`
}

func (m *DatabaseCheckRequest) Reset()                    { *m = DatabaseCheckRequest{} }
func (m *DatabaseCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseCheckRequest) ProtoMessage()               {}
func (*DatabaseCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*Database)(nil), "db.v1beta1.Database")
	proto.RegisterType((*DatabaseListRequest)(nil), "db.v1beta1.DatabaseListRequest")
	proto.RegisterType((*DatabaseListResponse)(nil), "db.v1beta1.DatabaseListResponse")
	proto.RegisterType((*DatabaseCreateRequest)(nil), "db.v1beta1.DatabaseCreateRequest")
	proto.RegisterType((*DatabaseScaleRequest)(nil), "db.v1beta1.DatabaseScaleRequest")
	proto.RegisterType((*DatabaseUpdateRequest)(nil), "db.v1beta1.DatabaseUpdateRequest")
	proto.RegisterType((*DatabaseDescribeRequest)(nil), "db.v1beta1.DatabaseDescribeRequest")
	proto.RegisterType((*SnapshotSummary)(nil), "db.v1beta1.SnapshotSummary")
	proto.RegisterType((*DatabaseDescribeResponse)(nil), "db.v1beta1.DatabaseDescribeResponse")
	proto.RegisterType((*DatabaseDeleteRequest)(nil), "db.v1beta1.DatabaseDeleteRequest")
	proto.RegisterType((*DatabaseRecoverRequest)(nil), "db.v1beta1.DatabaseRecoverRequest")
	proto.RegisterType((*DatabaseCheckRequest)(nil), "db.v1beta1.DatabaseCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Databases service

type DatabasesClient interface {
	List(ctx context.Context, in *DatabaseListRequest, opts ...grpc.CallOption) (*DatabaseListResponse, error)
	Create(ctx context.Context, in *DatabaseCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Scale(ctx context.Context, in *DatabaseScaleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *DatabaseUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Describe(ctx context.Context, in *DatabaseDescribeRequest, opts ...grpc.CallOption) (*DatabaseDescribeResponse, error)
	Delete(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Recover(ctx context.Context, in *DatabaseRecoverRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type databasesClient struct {
	cc *grpc.ClientConn
}

func NewDatabasesClient(cc *grpc.ClientConn) DatabasesClient {
	return &databasesClient{cc}
}

func (c *databasesClient) List(ctx context.Context, in *DatabaseListRequest, opts ...grpc.CallOption) (*DatabaseListResponse, error) {
	out := new(DatabaseListResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Create(ctx context.Context, in *DatabaseCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Scale(ctx context.Context, in *DatabaseScaleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Scale", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Update(ctx context.Context, in *DatabaseUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Describe(ctx context.Context, in *DatabaseDescribeRequest, opts ...grpc.CallOption) (*DatabaseDescribeResponse, error) {
	out := new(DatabaseDescribeResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Delete(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Recover(ctx context.Context, in *DatabaseRecoverRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Recover", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Databases service

type DatabasesServer interface {
	List(context.Context, *DatabaseListRequest) (*DatabaseListResponse, error)
	Create(context.Context, *DatabaseCreateRequest) (*dtypes.VoidResponse, error)
	Scale(context.Context, *DatabaseScaleRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *DatabaseUpdateRequest) (*dtypes.VoidResponse, error)
	Describe(context.Context, *DatabaseDescribeRequest) (*DatabaseDescribeResponse, error)
	Delete(context.Context, *DatabaseDeleteRequest) (*dtypes.VoidResponse, error)
	Recover(context.Context, *DatabaseRecoverRequest) (*dtypes.VoidResponse, error)
}

func RegisterDatabasesServer(s *grpc.Server, srv DatabasesServer) {
	s.RegisterService(&_Databases_serviceDesc, srv)
}

func _Databases_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).List(ctx, req.(*DatabaseListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Create(ctx, req.(*DatabaseCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Scale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Scale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Scale(ctx, req.(*DatabaseScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Update(ctx, req.(*DatabaseUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Describe(ctx, req.(*DatabaseDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Delete(ctx, req.(*DatabaseDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Recover(ctx, req.(*DatabaseRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Databases_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.v1beta1.Databases",
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
			MethodName: "Scale",
			Handler:    _Databases_Scale_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Databases_Update_Handler,
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
			MethodName: "Recover",
			Handler:    _Databases_Recover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1071 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0xe3, 0x26, 0x4d, 0x4e, 0xda, 0x34, 0x3b, 0x94, 0x62, 0x85, 0x85, 0x0d, 0xee, 0x02,
	0x51, 0x41, 0x31, 0x5b, 0xee, 0x56, 0x42, 0xa8, 0x4d, 0x0b, 0x48, 0xa0, 0x8a, 0x4d, 0xa0, 0x42,
	0x20, 0xb0, 0xc6, 0xf6, 0x28, 0xb1, 0x9a, 0x7a, 0x8c, 0x67, 0x1c, 0x6d, 0x77, 0xb5, 0x37, 0xfb,
	0x06, 0x08, 0x04, 0xe2, 0x7a, 0x6f, 0xb8, 0xe5, 0x82, 0x3b, 0xde, 0x82, 0x17, 0xe0, 0x82, 0x47,
	0xe0, 0x01, 0xd0, 0xfc, 0xd8, 0xb1, 0x9b, 0x04, 0x75, 0xbb, 0xbd, 0x89, 0x66, 0xbe, 0x39, 0x39,
	0xe7, 0x9b, 0x33, 0xdf, 0xf9, 0x64, 0x68, 0x05, 0x98, 0x63, 0x0f, 0x33, 0xd2, 0x8f, 0x13, 0xca,
	0x29, 0x82, 0xc0, 0xeb, 0xcf, 0xee, 0x79, 0x84, 0xe3, 0x7b, 0x9d, 0xdb, 0x63, 0x4a, 0xc7, 0x53,
	0xe2, 0xe0, 0x38, 0x74, 0x70, 0x14, 0x51, 0x8e, 0x79, 0x48, 0x23, 0xa6, 0x22, 0x3b, 0xaf, 0xe3,
	0x38, 0x66, 0x3e, 0x0d, 0x56, 0x9d, 0xef, 0x08, 0x38, 0xe0, 0x17, 0x31, 0x61, 0x8e, 0xfc, 0x55,
	0xb8, 0xfd, 0xab, 0x09, 0xf5, 0x23, 0x5d, 0x14, 0x21, 0x58, 0x8b, 0x27, 0x61, 0x60, 0x19, 0x5d,
	0xa3, 0xd7, 0x18, 0xca, 0x35, 0xb2, 0x60, 0xdd, 0x9f, 0xa6, 0x8c, 0x93, 0xc4, 0xaa, 0x48, 0x38,
	0xdb, 0x8a, 0xe8, 0x08, 0x9f, 0x13, 0xcb, 0x54, 0xd1, 0x62, 0x2d, 0x30, 0x91, 0xdd, 0xaa, 0x2a,
	0x4c, 0xac, 0x51, 0x1b, 0x4c, 0x76, 0x96, 0x5a, 0x35, 0x09, 0x89, 0xa5, 0xc8, 0x39, 0x23, 0x09,
	0x0b, 0x69, 0x64, 0xd5, 0x55, 0x4e, 0xbd, 0x45, 0x77, 0xa0, 0x89, 0x53, 0x3e, 0x71, 0x19, 0xf1,
	0x13, 0xc2, 0xad, 0x86, 0x3c, 0x05, 0x01, 0x8d, 0x24, 0x82, 0xde, 0x05, 0xc4, 0xfc, 0x09, 0x09,
	0xd2, 0x29, 0x71, 0xfd, 0x84, 0x46, 0x2e, 0x79, 0x18, 0x27, 0x16, 0xc8, 0xb8, 0x76, 0x76, 0x32,
	0x48, 0x68, 0x74, 0xfc, 0x30, 0x4e, 0xd0, 0x6d, 0x80, 0x78, 0xe6, 0xb2, 0xf0, 0x11, 0x71, 0xc7,
	0x9e, 0xd5, 0xec, 0x1a, 0xbd, 0xea, 0xb0, 0x1e, 0xcf, 0x46, 0xe1, 0x23, 0xf2, 0xb1, 0x87, 0xee,
	0x42, 0x6b, 0x8a, 0x19, 0x77, 0x3d, 0xec, 0x9f, 0xa5, 0xb1, 0x8b, 0xb9, 0xb5, 0xd1, 0x35, 0x7a,
	0xe6, 0x70, 0x43, 0xa0, 0x87, 0x12, 0x3c, 0xe0, 0x68, 0x07, 0x6a, 0x8c, 0x63, 0x9e, 0x32, 0x6b,
	0x53, 0x56, 0xd1, 0x3b, 0xf4, 0x1a, 0x80, 0x9f, 0x10, 0xcc, 0x49, 0x20, 0xfe, 0xd9, 0x92, 0xff,
	0x6c, 0x68, 0xe4, 0x80, 0x8b, 0xe3, 0x80, 0x4c, 0x89, 0x3e, 0xde, 0x52, 0xc7, 0x1a, 0x39, 0xe0,
	0xc8, 0x86, 0xcd, 0x80, 0xba, 0x11, 0xe5, 0xae, 0xc2, 0xac, 0x5b, 0x5d, 0xa3, 0x57, 0x1f, 0x36,
	0x03, 0x7a, 0x42, 0xf9, 0x91, 0x84, 0xec, 0x6f, 0xe0, 0xa5, 0xec, 0x69, 0x3e, 0x0b, 0x19, 0x1f,
	0x92, 0xef, 0x53, 0xc2, 0x78, 0xf1, 0x45, 0x8c, 0x85, 0x17, 0x91, 0xdd, 0xaf, 0x14, 0xba, 0x3f,
	0xa7, 0x6f, 0x76, 0xcd, 0x39, 0x7d, 0x3b, 0x81, 0xed, 0x72, 0x72, 0x16, 0xd3, 0x88, 0x11, 0xf4,
	0x56, 0x1e, 0x2f, 0x92, 0x37, 0xf7, 0x5b, 0x7d, 0xa5, 0x9a, 0xfe, 0x48, 0xa2, 0xf9, 0xf5, 0xf7,
	0xa1, 0x91, 0x89, 0x95, 0x59, 0x95, 0xae, 0xd9, 0x6b, 0xee, 0x6f, 0xf7, 0xe7, 0x72, 0xed, 0x67,
	0xc9, 0x87, 0xf3, 0x30, 0xfb, 0xdf, 0x0a, 0xbc, 0x9c, 0xe1, 0x03, 0xd9, 0xa9, 0x2b, 0xdd, 0x49,
	0xaa, 0xac, 0xb2, 0x44, 0x65, 0xe6, 0xa2, 0xca, 0xd6, 0x96, 0xaa, 0xac, 0x5a, 0x56, 0x59, 0x59,
	0x16, 0xeb, 0x97, 0x64, 0xd1, 0x83, 0x76, 0x41, 0x83, 0xae, 0xac, 0xae, 0x64, 0xda, 0x9a, 0x0b,
	0xf1, 0x44, 0xf3, 0x10, 0x49, 0xa4, 0x4c, 0xab, 0x43, 0xb9, 0x46, 0xbb, 0xb0, 0xc9, 0x22, 0x1c,
	0xb3, 0x09, 0xe5, 0xae, 0x1c, 0xa6, 0x96, 0xfc, 0xeb, 0x46, 0x06, 0x7e, 0x2e, 0x86, 0xaa, 0x03,
	0xf5, 0x09, 0x65, 0x5c, 0xa6, 0xde, 0x92, 0xe7, 0xf9, 0x5e, 0x26, 0xe0, 0x34, 0xc1, 0x63, 0xe2,
	0xfa, 0x53, 0xcc, 0x98, 0xd5, 0xd6, 0x09, 0x14, 0x38, 0x10, 0x18, 0x7a, 0x07, 0x6e, 0x85, 0xe3,
	0x88, 0x26, 0xc4, 0x9d, 0xe1, 0x69, 0x18, 0xc8, 0x51, 0xd7, 0x12, 0x6a, 0xab, 0x83, 0xd3, 0x1c,
	0xb7, 0x4f, 0xe7, 0x4f, 0x3d, 0xf2, 0xf1, 0xf4, 0x0a, 0x4d, 0x6f, 0x83, 0x99, 0x86, 0x81, 0xee,
	0xb9, 0x58, 0xe6, 0x57, 0x35, 0xe7, 0x57, 0xb5, 0xc7, 0xf3, 0xd7, 0xfc, 0x32, 0x0e, 0xae, 0xf4,
	0x9a, 0x8b, 0x89, 0x17, 0x06, 0xc1, 0x5c, 0x1c, 0x84, 0x63, 0x78, 0x25, 0x2b, 0x74, 0x44, 0x98,
	0x9f, 0x84, 0xde, 0x75, 0x4a, 0xd9, 0xbf, 0x1b, 0xb0, 0x35, 0xd2, 0xcf, 0x30, 0x4a, 0xcf, 0xcf,
	0x71, 0x72, 0x81, 0xde, 0x84, 0x56, 0x3e, 0xfe, 0x9c, 0x9c, 0xc7, 0x5c, 0xdf, 0x70, 0xd3, 0xd3,
	0xf3, 0x2f, 0xc1, 0x42, 0x18, 0x4b, 0x7d, 0x9f, 0x30, 0x26, 0x85, 0x96, 0x87, 0x8d, 0x14, 0x88,
	0xde, 0x86, 0xad, 0x84, 0x88, 0x87, 0x22, 0x79, 0xba, 0xaa, 0x8c, 0x6b, 0x69, 0x38, 0xcb, 0x57,
	0x08, 0xcc, 0x12, 0xd6, 0x4a, 0x81, 0x3a, 0xa3, 0xfd, 0xa7, 0x01, 0xd6, 0xe2, 0xdd, 0x9f, 0x73,
	0x56, 0x3f, 0x82, 0x76, 0xae, 0x49, 0xa6, 0x2e, 0x2e, 0x35, 0xdb, 0xdc, 0x7f, 0xb5, 0x38, 0xb2,
	0x97, 0x7a, 0x33, 0xdc, 0x62, 0x97, 0x9a, 0xf5, 0x1e, 0xd4, 0xb3, 0x61, 0x96, 0xaa, 0x5c, 0x35,
	0xf2, 0x79, 0x94, 0xfd, 0xed, 0x5c, 0x22, 0xea, 0x2d, 0xaf, 0x23, 0x11, 0x0b, 0xd6, 0x03, 0xc2,
	0x78, 0x42, 0x2f, 0xb4, 0x38, 0xb2, 0xad, 0x7d, 0x04, 0x3b, 0x79, 0x51, 0xe2, 0xd3, 0x19, 0x49,
	0xae, 0xa3, 0x8b, 0xf1, 0x7c, 0x3e, 0x06, 0x13, 0xe2, 0x9f, 0x5d, 0x87, 0xe3, 0x2e, 0x6c, 0xc6,
	0x69, 0xe2, 0x4f, 0x30, 0x23, 0x6a, 0xec, 0x95, 0x37, 0x6d, 0x64, 0xa0, 0x18, 0xfb, 0xfd, 0xbf,
	0xeb, 0xd0, 0xc8, 0x2a, 0x31, 0xf4, 0xb3, 0x01, 0x6b, 0xc2, 0x7a, 0xd1, 0x9d, 0x65, 0x4d, 0x2c,
	0x38, 0x7e, 0xa7, 0xbb, 0x3a, 0x40, 0x29, 0xc1, 0xfe, 0xf4, 0xe9, 0x1f, 0x56, 0xa5, 0x6e, 0x3c,
	0xfd, 0xeb, 0x9f, 0x1f, 0x2b, 0x1f, 0xa2, 0x0f, 0x9c, 0xd2, 0xd7, 0xc0, 0x59, 0xea, 0x91, 0x24,
	0x22, 0x9c, 0x30, 0x47, 0x27, 0x71, 0xf4, 0x45, 0x98, 0xf3, 0x58, 0xaf, 0x9e, 0x38, 0xb9, 0x4d,
	0xa3, 0x1f, 0x0c, 0xa8, 0x29, 0x7b, 0x46, 0x6f, 0x2c, 0xab, 0x5c, 0xb2, 0xee, 0xce, 0x76, 0x26,
	0xba, 0x53, 0x1a, 0x06, 0x39, 0xa1, 0x93, 0x02, 0xa1, 0x43, 0xfb, 0xc5, 0x08, 0xdd, 0x37, 0xf6,
	0xd0, 0x33, 0x03, 0xaa, 0xd2, 0xbc, 0xd0, 0xd2, 0x66, 0x14, 0x7d, 0x6d, 0x05, 0x23, 0xaf, 0xc0,
	0xe8, 0xb4, 0xf3, 0xe0, 0x85, 0x18, 0x39, 0x8f, 0xd3, 0x30, 0x78, 0xe2, 0x60, 0x5f, 0x7e, 0x62,
	0x39, 0x4c, 0x54, 0x17, 0x2c, 0x7f, 0x31, 0xa0, 0xa6, 0xac, 0x70, 0x79, 0xe7, 0x4a, 0x36, 0xb9,
	0x82, 0xe7, 0x17, 0x05, 0x9e, 0x9f, 0x74, 0x06, 0x37, 0xc0, 0x53, 0x30, 0xfb, 0xcd, 0x80, 0x7a,
	0xe6, 0x1f, 0x68, 0x77, 0x19, 0xb7, 0x4b, 0xce, 0xda, 0xb9, 0xfb, 0xff, 0x41, 0x9a, 0xed, 0x83,
	0x02, 0xdb, 0x63, 0x74, 0x13, 0x6c, 0xd1, 0x4f, 0x06, 0xd4, 0x94, 0x59, 0x2c, 0x6f, 0x62, 0xc9,
	0x48, 0x56, 0x34, 0xb1, 0x44, 0x6b, 0xef, 0x46, 0x68, 0x3d, 0x33, 0x60, 0x5d, 0x9b, 0x0c, 0xb2,
	0x97, 0xda, 0x5e, 0xc9, 0x81, 0x56, 0x10, 0xfb, 0x4e, 0x52, 0xfa, 0xaa, 0x33, 0xba, 0x49, 0xfd,
	0x25, 0xaa, 0xf2, 0x7d, 0x63, 0xef, 0xb0, 0xf1, 0xf5, 0xba, 0x4e, 0xe0, 0xd5, 0xe4, 0x17, 0xfe,
	0xfb, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x53, 0x34, 0xa7, 0x6e, 0x55, 0x0c, 0x00, 0x00,
}
