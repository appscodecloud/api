// Code generated by protoc-gen-go.
// source: gearman.proto
// DO NOT EDIT!

/*
Package gearman is a generated protocol buffer package.

It is generated from these files:
	gearman.proto

It has these top-level messages:
	Operation
	Auth
	Metadata
*/
package gearman

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ci "github.com/appscode/api/ci"
import kubernetes "github.com/appscode/api/kubernetes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Next Id: 11
type OperationType int32

const (
	OperationType_UNKNOWN          OperationType = 0
	OperationType_CLUSTER_CREATE   OperationType = 1
	OperationType_CLUSTER_SCALE    OperationType = 2
	OperationType_CLUSTER_DELETE   OperationType = 3
	OperationType_NAMESPACE_CREATE OperationType = 4
	OperationType_SLAVE_CREATE     OperationType = 5
	OperationType_SLAVE_DELETE     OperationType = 6
)

var OperationType_name = map[int32]string{
	0: "UNKNOWN",
	1: "CLUSTER_CREATE",
	2: "CLUSTER_SCALE",
	3: "CLUSTER_DELETE",
	4: "NAMESPACE_CREATE",
	5: "SLAVE_CREATE",
	6: "SLAVE_DELETE",
}
var OperationType_value = map[string]int32{
	"UNKNOWN":          0,
	"CLUSTER_CREATE":   1,
	"CLUSTER_SCALE":    2,
	"CLUSTER_DELETE":   3,
	"NAMESPACE_CREATE": 4,
	"SLAVE_CREATE":     5,
	"SLAVE_DELETE":     6,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Next Id: 10
type Operation struct {
	Phid string        `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Type OperationType `protobuf:"varint,2,opt,name=type,enum=gearman.OperationType" json:"type,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*Operation_ClusterCreateRequest
	//	*Operation_ClusterScaleRequest
	//	*Operation_ClusterDeleteRequest
	//	*Operation_CiSlaveCreateRequest
	//	*Operation_CiSlaveDeleteRequest
	Request     isOperation_Request `protobuf_oneof:"request"`
	Auth        *Auth               `protobuf:"bytes,8,opt,name=auth" json:"auth,omitempty"`
	RequestTime int64               `protobuf:"varint,9,opt,name=request_time" json:"request_time,omitempty"`
	Metadata    *Metadata           `protobuf:"bytes,10,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isOperation_Request interface {
	isOperation_Request()
}

type Operation_ClusterCreateRequest struct {
	ClusterCreateRequest *kubernetes.ClusterCreateRequest `protobuf:"bytes,3,opt,name=cluster_create_request,oneof"`
}
type Operation_ClusterScaleRequest struct {
	ClusterScaleRequest *kubernetes.ClusterScaleRequest `protobuf:"bytes,4,opt,name=cluster_scale_request,oneof"`
}
type Operation_ClusterDeleteRequest struct {
	ClusterDeleteRequest *kubernetes.ClusterDeleteRequest `protobuf:"bytes,5,opt,name=cluster_delete_request,oneof"`
}
type Operation_CiSlaveCreateRequest struct {
	CiSlaveCreateRequest *ci.SlaveCreateRequest `protobuf:"bytes,6,opt,name=ci_slave_create_request,oneof"`
}
type Operation_CiSlaveDeleteRequest struct {
	CiSlaveDeleteRequest *ci.SlaveDeleteRequest `protobuf:"bytes,7,opt,name=ci_slave_delete_request,oneof"`
}

func (*Operation_ClusterCreateRequest) isOperation_Request() {}
func (*Operation_ClusterScaleRequest) isOperation_Request()  {}
func (*Operation_ClusterDeleteRequest) isOperation_Request() {}
func (*Operation_CiSlaveCreateRequest) isOperation_Request() {}
func (*Operation_CiSlaveDeleteRequest) isOperation_Request() {}

func (m *Operation) GetRequest() isOperation_Request {
	if m != nil {
		return m.Request
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

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_ClusterCreateRequest)(nil),
		(*Operation_ClusterScaleRequest)(nil),
		(*Operation_ClusterDeleteRequest)(nil),
		(*Operation_CiSlaveCreateRequest)(nil),
		(*Operation_CiSlaveDeleteRequest)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterCreateRequest); err != nil {
			return err
		}
	case *Operation_ClusterScaleRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterScaleRequest); err != nil {
			return err
		}
	case *Operation_ClusterDeleteRequest:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterDeleteRequest); err != nil {
			return err
		}
	case *Operation_CiSlaveCreateRequest:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiSlaveCreateRequest); err != nil {
			return err
		}
	case *Operation_CiSlaveDeleteRequest:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiSlaveDeleteRequest); err != nil {
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
	case 3: // request.cluster_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterCreateRequest{msg}
		return true, err
	case 4: // request.cluster_scale_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterScaleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterScaleRequest{msg}
		return true, err
	case 5: // request.cluster_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterDeleteRequest{msg}
		return true, err
	case 6: // request.ci_slave_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.SlaveCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiSlaveCreateRequest{msg}
		return true, err
	case 7: // request.ci_slave_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.SlaveDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiSlaveDeleteRequest{msg}
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
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterScaleRequest:
		s := proto.Size(x.ClusterScaleRequest)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterDeleteRequest:
		s := proto.Size(x.ClusterDeleteRequest)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiSlaveCreateRequest:
		s := proto.Size(x.CiSlaveCreateRequest)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiSlaveDeleteRequest:
		s := proto.Size(x.CiSlaveDeleteRequest)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Auth struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Metadata struct {
	SubscriptionPhid string `protobuf:"bytes,1,opt,name=subscription_phid" json:"subscription_phid,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Operation)(nil), "gearman.Operation")
	proto.RegisterType((*Auth)(nil), "gearman.Auth")
	proto.RegisterType((*Metadata)(nil), "gearman.Metadata")
	proto.RegisterEnum("gearman.OperationType", OperationType_name, OperationType_value)
}

var fileDescriptor0 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x9b, 0xc6, 0xf9, 0xe3, 0x69, 0x5c, 0x39, 0x2b, 0x08, 0xa6, 0x08, 0x51, 0x05, 0x90,
	0x10, 0x07, 0x47, 0x2a, 0x27, 0x38, 0xe1, 0xba, 0x96, 0x90, 0x48, 0x53, 0x14, 0xa7, 0x70, 0xb4,
	0x36, 0x9b, 0x11, 0x5d, 0x35, 0x71, 0xcc, 0xee, 0x1a, 0x89, 0xc7, 0xe0, 0x15, 0x79, 0x12, 0x36,
	0x1b, 0x3b, 0x36, 0x16, 0x3d, 0xee, 0xcc, 0xf7, 0xfd, 0x66, 0x76, 0x66, 0xc0, 0xf9, 0x8e, 0x54,
	0x6c, 0x68, 0xea, 0x67, 0x62, 0xab, 0xb6, 0xa4, 0x57, 0x3c, 0xcf, 0x08, 0xcd, 0xf8, 0x84, 0xf1,
	0x89, 0x5c, 0xd3, 0x9f, 0xb8, 0x4f, 0x9e, 0x3d, 0xdf, 0xc5, 0xee, 0xf3, 0x25, 0x8a, 0x14, 0x15,
	0xca, 0x09, 0x5b, 0xe7, 0x52, 0xa1, 0x90, 0xfb, 0xf4, 0xf8, 0x4f, 0x1b, 0xec, 0x9b, 0x0c, 0x05,
	0x55, 0x7c, 0x9b, 0x92, 0x01, 0x58, 0xd9, 0x1d, 0x5f, 0x79, 0xad, 0xf3, 0xd6, 0x1b, 0x9b, 0xbc,
	0x02, 0x4b, 0xfd, 0xca, 0xd0, 0x3b, 0xd6, 0xaf, 0xd3, 0x8b, 0x91, 0x5f, 0x56, 0x3d, 0xe8, 0x17,
	0x3a, 0x4b, 0x2e, 0x61, 0x54, 0x30, 0x13, 0x26, 0x90, 0x2a, 0x4c, 0x04, 0xfe, 0xc8, 0x51, 0x2a,
	0xaf, 0xad, 0x7d, 0x27, 0x17, 0xe7, 0x7e, 0x55, 0xdd, 0x0f, 0xf7, 0xca, 0xd0, 0x08, 0xe7, 0x7b,
	0xdd, 0xa7, 0x23, 0xf2, 0x11, 0x1e, 0x97, 0x0c, 0xc9, 0xe8, 0xba, 0x42, 0x58, 0x06, 0xf1, 0xe2,
	0x3f, 0x88, 0x78, 0xa7, 0xab, 0x08, 0xb5, 0x2e, 0x56, 0xb8, 0xc6, 0x5a, 0x17, 0x9d, 0x07, 0xbb,
	0xb8, 0x32, 0xc2, 0x8a, 0xf1, 0x1e, 0x9e, 0x30, 0x9e, 0x98, 0xe1, 0x35, 0xbf, 0xd2, 0x35, 0x90,
	0x91, 0xcf, 0xb8, 0x1f, 0xef, 0xf2, 0xcd, 0x0f, 0xd4, 0xad, 0x8d, 0xfa, 0xbd, 0x86, 0xb5, 0x59,
	0xf5, 0x19, 0x58, 0x34, 0x57, 0x77, 0x5e, 0xdf, 0xe8, 0x9c, 0xc3, 0x94, 0x03, 0x1d, 0x24, 0x8f,
	0x60, 0x50, 0x70, 0x12, 0xc5, 0x37, 0xe8, 0xd9, 0x5a, 0xd4, 0x26, 0x2f, 0xa1, 0xbf, 0x41, 0x45,
	0x57, 0x54, 0x51, 0x0f, 0x8c, 0x6d, 0x78, 0xb0, 0x5d, 0x17, 0x89, 0x4b, 0x1b, 0x7a, 0x85, 0x75,
	0xfc, 0x01, 0x2c, 0x43, 0x1b, 0x82, 0x9d, 0xd2, 0x0d, 0xca, 0x8c, 0x32, 0x2c, 0x76, 0xec, 0x42,
	0x3f, 0x97, 0x7a, 0x2e, 0x3a, 0x6c, 0xf6, 0x6c, 0x13, 0x07, 0x3a, 0x6a, 0x7b, 0x8f, 0xa9, 0x59,
	0x9f, 0x3d, 0x7e, 0x0d, 0xfd, 0x12, 0x49, 0x9e, 0xc2, 0x50, 0xe6, 0x4b, 0xc9, 0x04, 0xcf, 0x76,
	0xeb, 0x4f, 0xaa, 0x5b, 0x79, 0xfb, 0xbb, 0x05, 0xce, 0xbf, 0x77, 0x71, 0x02, 0xbd, 0xdb, 0xd9,
	0xe7, 0xd9, 0xcd, 0xb7, 0x99, 0x7b, 0x44, 0x08, 0x9c, 0x86, 0xd3, 0xdb, 0x78, 0x11, 0xcd, 0x93,
	0x70, 0x1e, 0x05, 0x8b, 0xc8, 0x6d, 0xe9, 0x6e, 0x9c, 0x32, 0x16, 0x87, 0xc1, 0x34, 0x72, 0x8f,
	0xeb, 0xb2, 0xab, 0x68, 0x1a, 0x69, 0x59, 0x5b, 0x8f, 0xc0, 0x9d, 0x05, 0xd7, 0x51, 0xfc, 0x25,
	0x08, 0xa3, 0xd2, 0x6c, 0xe9, 0xbe, 0x07, 0xf1, 0x34, 0xf8, 0x7a, 0x88, 0x74, 0xaa, 0x48, 0xe1,
	0xec, 0x2e, 0xbb, 0xe6, 0xc4, 0xdf, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x48, 0x97, 0x58,
	0x2f, 0x03, 0x00, 0x00,
}
