// Code generated by protoc-gen-go.
// source: gearman.proto
// DO NOT EDIT!

/*
Package gearman is a generated protocol buffer package.

It is generated from these files:
	gearman.proto

It has these top-level messages:
	Auth
	Metadata
	Operation
	DataBucketDeleteRequest
*/
package gearman

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import certificate "github.com/appscode/api/certificate/v0.1"
import ci "github.com/appscode/api/ci/v0.1"
import kubernetes "github.com/appscode/api/kubernetes/v0.1"
import db "github.com/appscode/api/db/v0.1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Next Id: 11
type OperationType int32

const (
	OperationType_UNKNOWN            OperationType = 0
	OperationType_CLUSTER_CREATE     OperationType = 1
	OperationType_CLUSTER_SCALE      OperationType = 2
	OperationType_CLUSTER_DELETE     OperationType = 3
	OperationType_CLUSTER_UPDATE     OperationType = 4
	OperationType_BUILD_AGENT_CREATE OperationType = 5
	OperationType_BUILD_AGENT_DELETE OperationType = 6
	OperationType_CHECK_DNS          OperationType = 7
	OperationType_NAMESPACE_CREATE   OperationType = 8
	OperationType_DATA_BUCKET_DELETE OperationType = 9
	OperationType_DATABASE_BACKUP    OperationType = 10
)

var OperationType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CLUSTER_CREATE",
	2:  "CLUSTER_SCALE",
	3:  "CLUSTER_DELETE",
	4:  "CLUSTER_UPDATE",
	5:  "BUILD_AGENT_CREATE",
	6:  "BUILD_AGENT_DELETE",
	7:  "CHECK_DNS",
	8:  "NAMESPACE_CREATE",
	9:  "DATA_BUCKET_DELETE",
	10: "DATABASE_BACKUP",
}
var OperationType_value = map[string]int32{
	"UNKNOWN":            0,
	"CLUSTER_CREATE":     1,
	"CLUSTER_SCALE":      2,
	"CLUSTER_DELETE":     3,
	"CLUSTER_UPDATE":     4,
	"BUILD_AGENT_CREATE": 5,
	"BUILD_AGENT_DELETE": 6,
	"CHECK_DNS":          7,
	"NAMESPACE_CREATE":   8,
	"DATA_BUCKET_DELETE": 9,
	"DATABASE_BACKUP":    10,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Auth struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	TokenType string `protobuf:"bytes,4,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Metadata struct {
	PurchasePhid []string `protobuf:"bytes,1,rep,name=purchase_phid,json=purchasePhid" json:"purchase_phid,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Next Id: 10
type Operation struct {
	Phid        string        `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Auth        *Auth         `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	RequestTime int64         `protobuf:"varint,3,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	Metadata    *Metadata     `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
	Type        OperationType `protobuf:"varint,5,opt,name=type,enum=gearman.OperationType" json:"type,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*Operation_ClusterCreateRequest
	//	*Operation_ClusterScaleRequest
	//	*Operation_ClusterDeleteRequest
	//	*Operation_ClusterUpdateRequest
	//	*Operation_CiSlaveCreateRequest
	//	*Operation_CiSlaveDeleteRequest
	//	*Operation_DnsCheckRequest
	//	*Operation_DataBucketDeleteRequest
	//	*Operation_DbBackupScheduleRequest
	Request isOperation_Request `protobuf_oneof:"request"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isOperation_Request interface {
	isOperation_Request()
}

type Operation_ClusterCreateRequest struct {
	ClusterCreateRequest *kubernetes.ClusterCreateRequest `protobuf:"bytes,6,opt,name=cluster_create_request,json=clusterCreateRequest,oneof"`
}
type Operation_ClusterScaleRequest struct {
	ClusterScaleRequest *kubernetes.ClusterScaleRequest `protobuf:"bytes,7,opt,name=cluster_scale_request,json=clusterScaleRequest,oneof"`
}
type Operation_ClusterDeleteRequest struct {
	ClusterDeleteRequest *kubernetes.ClusterDeleteRequest `protobuf:"bytes,8,opt,name=cluster_delete_request,json=clusterDeleteRequest,oneof"`
}
type Operation_ClusterUpdateRequest struct {
	ClusterUpdateRequest *kubernetes.ClusterUpdateRequest `protobuf:"bytes,9,opt,name=cluster_update_request,json=clusterUpdateRequest,oneof"`
}
type Operation_CiSlaveCreateRequest struct {
	CiSlaveCreateRequest *ci.SlaveCreateRequest `protobuf:"bytes,10,opt,name=ci_slave_create_request,json=ciSlaveCreateRequest,oneof"`
}
type Operation_CiSlaveDeleteRequest struct {
	CiSlaveDeleteRequest *ci.SlaveDeleteRequest `protobuf:"bytes,11,opt,name=ci_slave_delete_request,json=ciSlaveDeleteRequest,oneof"`
}
type Operation_DnsCheckRequest struct {
	DnsCheckRequest *certificate.DNSCheckRequest `protobuf:"bytes,12,opt,name=dns_check_request,json=dnsCheckRequest,oneof"`
}
type Operation_DataBucketDeleteRequest struct {
	DataBucketDeleteRequest *DataBucketDeleteRequest `protobuf:"bytes,13,opt,name=data_bucket_delete_request,json=dataBucketDeleteRequest,oneof"`
}
type Operation_DbBackupScheduleRequest struct {
	DbBackupScheduleRequest *db.BackupScheduleRequest `protobuf:"bytes,14,opt,name=db_backup_schedule_request,json=dbBackupScheduleRequest,oneof"`
}

func (*Operation_ClusterCreateRequest) isOperation_Request()    {}
func (*Operation_ClusterScaleRequest) isOperation_Request()     {}
func (*Operation_ClusterDeleteRequest) isOperation_Request()    {}
func (*Operation_ClusterUpdateRequest) isOperation_Request()    {}
func (*Operation_CiSlaveCreateRequest) isOperation_Request()    {}
func (*Operation_CiSlaveDeleteRequest) isOperation_Request()    {}
func (*Operation_DnsCheckRequest) isOperation_Request()         {}
func (*Operation_DataBucketDeleteRequest) isOperation_Request() {}
func (*Operation_DbBackupScheduleRequest) isOperation_Request() {}

func (m *Operation) GetRequest() isOperation_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Operation) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Operation) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetClusterCreateRequest() *kubernetes.ClusterCreateRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterCreateRequest); ok {
		return x.ClusterCreateRequest
	}
	return nil
}

func (m *Operation) GetClusterScaleRequest() *kubernetes.ClusterScaleRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterScaleRequest); ok {
		return x.ClusterScaleRequest
	}
	return nil
}

func (m *Operation) GetClusterDeleteRequest() *kubernetes.ClusterDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterDeleteRequest); ok {
		return x.ClusterDeleteRequest
	}
	return nil
}

func (m *Operation) GetClusterUpdateRequest() *kubernetes.ClusterUpdateRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterUpdateRequest); ok {
		return x.ClusterUpdateRequest
	}
	return nil
}

func (m *Operation) GetCiSlaveCreateRequest() *ci.SlaveCreateRequest {
	if x, ok := m.GetRequest().(*Operation_CiSlaveCreateRequest); ok {
		return x.CiSlaveCreateRequest
	}
	return nil
}

func (m *Operation) GetCiSlaveDeleteRequest() *ci.SlaveDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_CiSlaveDeleteRequest); ok {
		return x.CiSlaveDeleteRequest
	}
	return nil
}

func (m *Operation) GetDnsCheckRequest() *certificate.DNSCheckRequest {
	if x, ok := m.GetRequest().(*Operation_DnsCheckRequest); ok {
		return x.DnsCheckRequest
	}
	return nil
}

func (m *Operation) GetDataBucketDeleteRequest() *DataBucketDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_DataBucketDeleteRequest); ok {
		return x.DataBucketDeleteRequest
	}
	return nil
}

func (m *Operation) GetDbBackupScheduleRequest() *db.BackupScheduleRequest {
	if x, ok := m.GetRequest().(*Operation_DbBackupScheduleRequest); ok {
		return x.DbBackupScheduleRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_ClusterCreateRequest)(nil),
		(*Operation_ClusterScaleRequest)(nil),
		(*Operation_ClusterDeleteRequest)(nil),
		(*Operation_ClusterUpdateRequest)(nil),
		(*Operation_CiSlaveCreateRequest)(nil),
		(*Operation_CiSlaveDeleteRequest)(nil),
		(*Operation_DnsCheckRequest)(nil),
		(*Operation_DataBucketDeleteRequest)(nil),
		(*Operation_DbBackupScheduleRequest)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterCreateRequest); err != nil {
			return err
		}
	case *Operation_ClusterScaleRequest:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterScaleRequest); err != nil {
			return err
		}
	case *Operation_ClusterDeleteRequest:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterDeleteRequest); err != nil {
			return err
		}
	case *Operation_ClusterUpdateRequest:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterUpdateRequest); err != nil {
			return err
		}
	case *Operation_CiSlaveCreateRequest:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiSlaveCreateRequest); err != nil {
			return err
		}
	case *Operation_CiSlaveDeleteRequest:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiSlaveDeleteRequest); err != nil {
			return err
		}
	case *Operation_DnsCheckRequest:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DnsCheckRequest); err != nil {
			return err
		}
	case *Operation_DataBucketDeleteRequest:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DataBucketDeleteRequest); err != nil {
			return err
		}
	case *Operation_DbBackupScheduleRequest:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DbBackupScheduleRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Request has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 6: // request.cluster_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterCreateRequest{msg}
		return true, err
	case 7: // request.cluster_scale_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterScaleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterScaleRequest{msg}
		return true, err
	case 8: // request.cluster_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterDeleteRequest{msg}
		return true, err
	case 9: // request.cluster_update_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterUpdateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterUpdateRequest{msg}
		return true, err
	case 10: // request.ci_slave_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.SlaveCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiSlaveCreateRequest{msg}
		return true, err
	case 11: // request.ci_slave_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.SlaveDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiSlaveDeleteRequest{msg}
		return true, err
	case 12: // request.dns_check_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(certificate.DNSCheckRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DnsCheckRequest{msg}
		return true, err
	case 13: // request.data_bucket_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DataBucketDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DataBucketDeleteRequest{msg}
		return true, err
	case 14: // request.db_backup_schedule_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(db.BackupScheduleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DbBackupScheduleRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		s := proto.Size(x.ClusterCreateRequest)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterScaleRequest:
		s := proto.Size(x.ClusterScaleRequest)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterDeleteRequest:
		s := proto.Size(x.ClusterDeleteRequest)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterUpdateRequest:
		s := proto.Size(x.ClusterUpdateRequest)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiSlaveCreateRequest:
		s := proto.Size(x.CiSlaveCreateRequest)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiSlaveDeleteRequest:
		s := proto.Size(x.CiSlaveDeleteRequest)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DnsCheckRequest:
		s := proto.Size(x.DnsCheckRequest)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DataBucketDeleteRequest:
		s := proto.Size(x.DataBucketDeleteRequest)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DbBackupScheduleRequest:
		s := proto.Size(x.DbBackupScheduleRequest)
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DataBucketDeleteRequest struct {
	DataType  string `protobuf:"bytes,1,opt,name=data_type,json=dataType" json:"data_type,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Prefix    string `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *DataBucketDeleteRequest) Reset()                    { *m = DataBucketDeleteRequest{} }
func (m *DataBucketDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DataBucketDeleteRequest) ProtoMessage()               {}
func (*DataBucketDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Auth)(nil), "gearman.Auth")
	proto.RegisterType((*Metadata)(nil), "gearman.Metadata")
	proto.RegisterType((*Operation)(nil), "gearman.Operation")
	proto.RegisterType((*DataBucketDeleteRequest)(nil), "gearman.DataBucketDeleteRequest")
	proto.RegisterEnum("gearman.OperationType", OperationType_name, OperationType_value)
}

func init() { proto.RegisterFile("gearman.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0xed, 0x6e, 0xda, 0x4a,
	0x10, 0x0d, 0x81, 0x00, 0x1e, 0x42, 0x42, 0x36, 0xb9, 0xc0, 0xe5, 0xe6, 0xea, 0xe6, 0x52, 0xa9,
	0xaa, 0x22, 0xd5, 0x69, 0xd3, 0x27, 0x30, 0xc6, 0x6a, 0x5a, 0x08, 0x41, 0x36, 0x56, 0xfb, 0x6f,
	0xb5, 0xb6, 0x37, 0xc5, 0x82, 0x80, 0xeb, 0x8f, 0xa8, 0x7d, 0xd4, 0x3e, 0x44, 0xdf, 0xa1, 0xbb,
	0xeb, 0x35, 0x98, 0x8f, 0xfc, 0xb3, 0xcf, 0x9c, 0x39, 0x7b, 0x66, 0x76, 0x67, 0xa0, 0xfe, 0x8d,
	0x92, 0xf0, 0x89, 0x2c, 0xd4, 0x20, 0x5c, 0xc6, 0x4b, 0x54, 0x91, 0xbf, 0x9d, 0xd7, 0x24, 0xf0,
	0x6f, 0x5c, 0x1a, 0xc6, 0xfe, 0xa3, 0xef, 0x92, 0x98, 0xde, 0x3c, 0xbf, 0x53, 0xdf, 0xe7, 0x81,
	0x34, 0xa1, 0xd3, 0x12, 0x3c, 0x3f, 0x0d, 0x47, 0x73, 0xf2, 0x9c, 0x05, 0xba, 0x3c, 0x30, 0x4b,
	0x1c, 0x1a, 0x2e, 0x68, 0x4c, 0x23, 0x99, 0x3f, 0x4f, 0xa2, 0x98, 0x86, 0x91, 0xe4, 0x74, 0x38,
	0xc7, 0x73, 0xd2, 0x98, 0x47, 0x62, 0xe2, 0x90, 0x48, 0xe6, 0x77, 0x13, 0x28, 0x69, 0x49, 0x3c,
	0x45, 0x97, 0xa0, 0x2c, 0xc8, 0x13, 0x8d, 0x02, 0xe2, 0xd2, 0x76, 0xe1, 0xaa, 0xf0, 0x46, 0x31,
	0xd7, 0x00, 0xea, 0x40, 0x35, 0x89, 0xd8, 0x11, 0x0c, 0x68, 0x1f, 0x8a, 0xe0, 0xea, 0x1f, 0x5d,
	0xc0, 0x51, 0xbc, 0x9c, 0xd1, 0x45, 0xbb, 0x28, 0x02, 0xe9, 0x0f, 0xfa, 0x17, 0x40, 0x7c, 0xe0,
	0xf8, 0x67, 0x40, 0xdb, 0xa5, 0x54, 0x50, 0x20, 0x13, 0x06, 0x74, 0x6f, 0xa0, 0x7a, 0x4f, 0x63,
	0xc2, 0xcd, 0xa0, 0x57, 0x50, 0x0f, 0x92, 0xd0, 0x9d, 0x32, 0x53, 0x38, 0x98, 0xfa, 0x1e, 0x3b,
	0xbe, 0xc8, 0xd8, 0xc7, 0x19, 0x38, 0x66, 0x58, 0xf7, 0x57, 0x05, 0x94, 0x87, 0x80, 0x86, 0x24,
	0xf6, 0x97, 0x0b, 0x84, 0xa0, 0x24, 0x99, 0x5c, 0x57, 0x7c, 0xa3, 0xff, 0xa1, 0x44, 0x58, 0x25,
	0xc2, 0x5f, 0xed, 0xb6, 0xae, 0x66, 0x1d, 0xe7, 0xe5, 0x99, 0x22, 0xc4, 0x28, 0xc7, 0x21, 0xfd,
	0x9e, 0xd0, 0x28, 0xc6, 0xb1, 0xcf, 0x4a, 0xe1, 0x8e, 0x8b, 0x66, 0x4d, 0x62, 0x13, 0x06, 0xa1,
	0xb7, 0x50, 0x7d, 0x92, 0xc6, 0x84, 0xeb, 0xda, 0xed, 0xd9, 0x4a, 0x29, 0x73, 0x6c, 0xae, 0x28,
	0xe8, 0x1a, 0x4a, 0xa2, 0xc0, 0x23, 0x46, 0x3d, 0xb9, 0x6d, 0xae, 0xa8, 0x2b, 0xab, 0xbc, 0x5a,
	0x53, 0x70, 0xd0, 0x57, 0x68, 0xca, 0x8b, 0xc1, 0x6e, 0x48, 0xd9, 0xdd, 0x62, 0x79, 0x70, 0xbb,
	0x2c, 0x0e, 0xba, 0x52, 0xd7, 0xf7, 0xa8, 0xea, 0x29, 0x53, 0x17, 0x44, 0x33, 0xe5, 0xdd, 0x1d,
	0x98, 0x17, 0xee, 0x1e, 0x1c, 0xd9, 0xf0, 0x57, 0xa6, 0x1c, 0xb9, 0x64, 0xbe, 0x16, 0xae, 0x08,
	0xe1, 0xff, 0xf6, 0x08, 0x5b, 0x9c, 0xb7, 0xd6, 0x3d, 0x77, 0x77, 0xe1, 0xbc, 0x61, 0x8f, 0xce,
	0x69, 0xce, 0x70, 0xf5, 0x45, 0xc3, 0x7d, 0x41, 0xdc, 0x35, 0xbc, 0x81, 0xe7, 0x95, 0x93, 0xc0,
	0xcb, 0xb7, 0x42, 0x79, 0x51, 0xd9, 0x16, 0xc4, 0x5d, 0xe5, 0x0d, 0x1c, 0x3d, 0x40, 0xcb, 0xf5,
	0xb1, 0x98, 0x90, 0xed, 0x2e, 0x83, 0x90, 0x6e, 0xaa, 0xae, 0xaf, 0x5a, 0x3c, 0xbe, 0xdb, 0x5b,
	0x7f, 0x17, 0xdf, 0x10, 0xdc, 0xea, 0x42, 0x6d, 0x4b, 0x70, 0xb7, 0x76, 0x7f, 0x17, 0x47, 0x9f,
	0xe1, 0xcc, 0x5b, 0x44, 0xd8, 0x9d, 0x52, 0x77, 0xb6, 0x92, 0x3a, 0x16, 0x52, 0x97, 0x6a, 0x7e,
	0xf2, 0xfb, 0x23, 0x4b, 0xe7, 0xa4, 0xb5, 0xe0, 0x29, 0x4b, 0xcc, 0x43, 0x08, 0x43, 0x87, 0x3f,
	0x43, 0xec, 0x24, 0xee, 0x8c, 0xc6, 0xdb, 0xfe, 0xea, 0xb2, 0x97, 0xd9, 0xa3, 0xec, 0x33, 0x6a,
	0x4f, 0x30, 0xb7, 0x9d, 0xb6, 0xbc, 0xfd, 0x21, 0x76, 0x51, 0x1d, 0xcf, 0xc1, 0x0e, 0x71, 0x67,
	0x49, 0xc0, 0xde, 0xd6, 0x94, 0x7a, 0x49, 0xee, 0x79, 0x9d, 0x88, 0x03, 0xfe, 0x56, 0x3d, 0x47,
	0xed, 0x09, 0x8a, 0x25, 0x19, 0x79, 0x65, 0x67, 0x6f, 0xa8, 0xa7, 0x40, 0x45, 0xca, 0x74, 0xe7,
	0xd0, 0x7a, 0xc1, 0x1a, 0xfa, 0x07, 0x14, 0x51, 0xa0, 0x18, 0xb2, 0x74, 0xda, 0xab, 0x1c, 0xe0,
	0x63, 0xb5, 0xb9, 0xb3, 0x0e, 0xb7, 0x77, 0x56, 0x13, 0xca, 0x41, 0x48, 0x1f, 0xfd, 0x1f, 0x72,
	0x31, 0xc9, 0xbf, 0xeb, 0xdf, 0x05, 0xa8, 0x6f, 0x8c, 0x27, 0xaa, 0x41, 0xc5, 0x1e, 0x0d, 0x46,
	0x0f, 0x5f, 0x46, 0x8d, 0x03, 0xb6, 0x5a, 0x4e, 0xf4, 0xa1, 0x6d, 0x4d, 0x0c, 0x13, 0xeb, 0xa6,
	0xa1, 0x4d, 0x8c, 0x46, 0x01, 0x9d, 0x41, 0x3d, 0xc3, 0x2c, 0x5d, 0x1b, 0x1a, 0x8d, 0xc3, 0x3c,
	0xad, 0x6f, 0x0c, 0x0d, 0x46, 0x2b, 0xe6, 0x31, 0x7b, 0xdc, 0xe7, 0xa9, 0x25, 0xe6, 0x02, 0xf5,
	0xec, 0x4f, 0xc3, 0x3e, 0xd6, 0x3e, 0x1a, 0xa3, 0x49, 0x26, 0x79, 0xb4, 0x8d, 0x4b, 0x8d, 0x32,
	0xaa, 0x83, 0xa2, 0xdf, 0x19, 0xfa, 0x00, 0xb3, 0xdb, 0x6f, 0x54, 0xd8, 0x72, 0x6d, 0x8c, 0xb4,
	0x7b, 0xc3, 0x1a, 0x6b, 0xba, 0x91, 0x25, 0x57, 0x79, 0x32, 0x93, 0xd7, 0x70, 0xcf, 0xd6, 0x07,
	0xc6, 0x2a, 0x59, 0x41, 0xe7, 0x70, 0xca, 0xf1, 0x9e, 0x66, 0x19, 0xb8, 0xa7, 0xe9, 0x03, 0x7b,
	0xdc, 0x00, 0xa7, 0x2c, 0x16, 0xfd, 0x87, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xec, 0xff,
	0x0e, 0x83, 0x06, 0x00, 0x00,
}
