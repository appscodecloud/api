// Code generated by protoc-gen-go.
// source: alert.proto
// DO NOT EDIT!

/*
Package alert is a generated protocol buffer package.

It is generated from these files:
	alert.proto

It has these top-level messages:
	NotificationRequest
	CreateReuquest
	UpdateReuquest
	ListRequest
	DeleteRequest
	ListResponse
	Alert
	Matrix
*/
package alert

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

type NotificationRequest struct {
	StateType     string `protobuf:"bytes,1,opt,name=state_type" json:"state_type,omitempty"`
	State         string `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Time          string `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
	ServiceOutput string `protobuf:"bytes,4,opt,name=service_output" json:"service_output,omitempty"`
	AlertPhid     string `protobuf:"bytes,5,opt,name=alert_phid" json:"alert_phid,omitempty"`
	HostName      string `protobuf:"bytes,6,opt,name=host_name" json:"host_name,omitempty"`
}

func (m *NotificationRequest) Reset()                    { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()               {}
func (*NotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateReuquest struct {
	Name              string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cluster           string  `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Namespace         string  `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	ObjectType        string  `protobuf:"bytes,4,opt,name=object_type" json:"object_type,omitempty"`
	ObjectName        string  `protobuf:"bytes,5,opt,name=object_name" json:"object_name,omitempty"`
	MatrixPhid        string  `protobuf:"bytes,6,opt,name=matrix_phid" json:"matrix_phid,omitempty"`
	Matrix            *Matrix `protobuf:"bytes,7,opt,name=matrix" json:"matrix,omitempty"`
	WarningCondition  string  `protobuf:"bytes,8,opt,name=warning_condition" json:"warning_condition,omitempty"`
	CriticalCondition string  `protobuf:"bytes,9,opt,name=critical_condition" json:"critical_condition,omitempty"`
	CheckInterval     int64   `protobuf:"varint,10,opt,name=check_interval" json:"check_interval,omitempty"`
	WarningUser       string  `protobuf:"bytes,11,opt,name=warning_user" json:"warning_user,omitempty"`
	WarningMethod     int32   `protobuf:"varint,12,opt,name=warning_method" json:"warning_method,omitempty"`
	CriticalUser      string  `protobuf:"bytes,13,opt,name=critical_user" json:"critical_user,omitempty"`
	CriticalMethod    int32   `protobuf:"varint,14,opt,name=critical_method" json:"critical_method,omitempty"`
	AlertInterval     int64   `protobuf:"varint,15,opt,name=alert_interval" json:"alert_interval,omitempty"`
}

func (m *CreateReuquest) Reset()                    { *m = CreateReuquest{} }
func (m *CreateReuquest) String() string            { return proto.CompactTextString(m) }
func (*CreateReuquest) ProtoMessage()               {}
func (*CreateReuquest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateReuquest) GetMatrix() *Matrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

type UpdateReuquest struct {
	Phid              string  `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name              string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Cluster           string  `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	Namespace         string  `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	ObjectType        string  `protobuf:"bytes,5,opt,name=object_type" json:"object_type,omitempty"`
	ObjectName        string  `protobuf:"bytes,6,opt,name=object_name" json:"object_name,omitempty"`
	MatrixPhid        string  `protobuf:"bytes,7,opt,name=matrix_phid" json:"matrix_phid,omitempty"`
	Matrix            *Matrix `protobuf:"bytes,8,opt,name=matrix" json:"matrix,omitempty"`
	WarningCondition  string  `protobuf:"bytes,9,opt,name=warning_condition" json:"warning_condition,omitempty"`
	CriticalCondition string  `protobuf:"bytes,10,opt,name=critical_condition" json:"critical_condition,omitempty"`
	CheckInterval     int64   `protobuf:"varint,11,opt,name=check_interval" json:"check_interval,omitempty"`
	WarningUser       string  `protobuf:"bytes,12,opt,name=warning_user" json:"warning_user,omitempty"`
	WarningMethod     int32   `protobuf:"varint,13,opt,name=warning_method" json:"warning_method,omitempty"`
	CriticalUser      string  `protobuf:"bytes,14,opt,name=critical_user" json:"critical_user,omitempty"`
	CriticalMethod    int32   `protobuf:"varint,15,opt,name=critical_method" json:"critical_method,omitempty"`
	AlertInterval     int64   `protobuf:"varint,16,opt,name=alert_interval" json:"alert_interval,omitempty"`
}

func (m *UpdateReuquest) Reset()                    { *m = UpdateReuquest{} }
func (m *UpdateReuquest) String() string            { return proto.CompactTextString(m) }
func (*UpdateReuquest) ProtoMessage()               {}
func (*UpdateReuquest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateReuquest) GetMatrix() *Matrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

type ListRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	ObjectName string `protobuf:"bytes,2,opt,name=object_name" json:"object_name,omitempty"`
	Namespace  string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	ObjectType string `protobuf:"bytes,4,opt,name=object_type" json:"object_type,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	ObjectName string `protobuf:"bytes,2,opt,name=object_name" json:"object_name,omitempty"`
	Namespace  string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	ObjectType string `protobuf:"bytes,4,opt,name=object_type" json:"object_type,omitempty"`
	Phid       string `protobuf:"bytes,5,opt,name=phid" json:"phid,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts []*Alert       `protobuf:"bytes,2,rep,name=alerts" json:"alerts,omitempty"`
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

func (m *ListResponse) GetAlerts() []*Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

type Alert struct {
	Phid            string       `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name            string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IcingaUrl       string       `protobuf:"bytes,3,opt,name=icinga_url" json:"icinga_url,omitempty"`
	Spcs            *Alert_Specs `protobuf:"bytes,4,opt,name=spcs" json:"spcs,omitempty"`
	AlertUser       string       `protobuf:"bytes,5,opt,name=alert_user" json:"alert_user,omitempty"`
	WarningMethod   string       `protobuf:"bytes,6,opt,name=warning_method" json:"warning_method,omitempty"`
	CriticalUser    string       `protobuf:"bytes,7,opt,name=critical_user" json:"critical_user,omitempty"`
	CriticalMethod  string       `protobuf:"bytes,8,opt,name=critical_method" json:"critical_method,omitempty"`
	AlertInterval   int64        `protobuf:"varint,9,opt,name=alert_interval" json:"alert_interval,omitempty"`
	RefreshInterval int64        `protobuf:"varint,10,opt,name=refresh_interval" json:"refresh_interval,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Alert) GetSpcs() *Alert_Specs {
	if m != nil {
		return m.Spcs
	}
	return nil
}

type Alert_Specs struct {
	Matrix            *Matrix `protobuf:"bytes,1,opt,name=matrix" json:"matrix,omitempty"`
	WarningCondition  string  `protobuf:"bytes,2,opt,name=warning_condition" json:"warning_condition,omitempty"`
	CriticalCondition string  `protobuf:"bytes,3,opt,name=critical_condition" json:"critical_condition,omitempty"`
}

func (m *Alert_Specs) Reset()                    { *m = Alert_Specs{} }
func (m *Alert_Specs) String() string            { return proto.CompactTextString(m) }
func (*Alert_Specs) ProtoMessage()               {}
func (*Alert_Specs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *Alert_Specs) GetMatrix() *Matrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

type Matrix struct {
	Name    string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Command string            `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	Query   map[string]string `protobuf:"bytes,3,rep,name=query" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Formula string            `protobuf:"bytes,4,opt,name=formula" json:"formula,omitempty"`
}

func (m *Matrix) Reset()                    { *m = Matrix{} }
func (m *Matrix) String() string            { return proto.CompactTextString(m) }
func (*Matrix) ProtoMessage()               {}
func (*Matrix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Matrix) GetQuery() map[string]string {
	if m != nil {
		return m.Query
	}
	return nil
}

func init() {
	proto.RegisterType((*NotificationRequest)(nil), "alert.NotificationRequest")
	proto.RegisterType((*CreateReuquest)(nil), "alert.CreateReuquest")
	proto.RegisterType((*UpdateReuquest)(nil), "alert.UpdateReuquest")
	proto.RegisterType((*ListRequest)(nil), "alert.ListRequest")
	proto.RegisterType((*DeleteRequest)(nil), "alert.DeleteRequest")
	proto.RegisterType((*ListResponse)(nil), "alert.ListResponse")
	proto.RegisterType((*Alert)(nil), "alert.Alert")
	proto.RegisterType((*Alert_Specs)(nil), "alert.Alert.Specs")
	proto.RegisterType((*Matrix)(nil), "alert.Matrix")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Alerts service

type AlertsClient interface {
	Notify(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Create(ctx context.Context, in *CreateReuquest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *UpdateReuquest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type alertsClient struct {
	cc *grpc.ClientConn
}

func NewAlertsClient(cc *grpc.ClientConn) AlertsClient {
	return &alertsClient{cc}
}

func (c *alertsClient) Notify(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Create(ctx context.Context, in *CreateReuquest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Update(ctx context.Context, in *UpdateReuquest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Alerts service

type AlertsServer interface {
	Notify(context.Context, *NotificationRequest) (*dtypes.VoidResponse, error)
	Create(context.Context, *CreateReuquest) (*dtypes.VoidResponse, error)
	Update(context.Context, *UpdateReuquest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*dtypes.VoidResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterAlertsServer(s *grpc.Server, srv AlertsServer) {
	s.RegisterService(&_Alerts_serviceDesc, srv)
}

func _Alerts_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AlertsServer).Notify(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Alerts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateReuquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AlertsServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Alerts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateReuquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AlertsServer).Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Alerts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AlertsServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Alerts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AlertsServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Alerts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.Alerts",
	HandlerType: (*AlertsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Alerts_Notify_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Alerts_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Alerts_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Alerts_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Alerts_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x4e, 0xe3, 0x46,
	0x14, 0x96, 0x93, 0xd8, 0x21, 0xc7, 0xf9, 0x81, 0x21, 0x50, 0x63, 0xd1, 0x16, 0xf9, 0xa6, 0x15,
	0xaa, 0x92, 0x2a, 0xbd, 0xa9, 0xb8, 0xa9, 0x50, 0xe9, 0x8f, 0x2a, 0x8a, 0x54, 0x50, 0x7b, 0x83,
	0xaa, 0xc8, 0x38, 0x03, 0x71, 0x71, 0x6c, 0xd7, 0x33, 0xa6, 0x44, 0x88, 0x8b, 0x56, 0xed, 0x75,
	0x2f, 0xf6, 0x9d, 0xf6, 0x05, 0xf6, 0x15, 0x76, 0x6f, 0xf7, 0x66, 0x5f, 0x60, 0x67, 0xce, 0xd8,
	0xf9, 0x23, 0x8e, 0x58, 0xb1, 0x7b, 0x83, 0x32, 0xdf, 0x9c, 0xf9, 0xce, 0xe7, 0xef, 0x9c, 0x39,
	0x03, 0x98, 0x6e, 0x40, 0x13, 0xde, 0x89, 0x93, 0x88, 0x47, 0x44, 0xc7, 0x85, 0xbd, 0x7b, 0x15,
	0x45, 0x57, 0x01, 0xed, 0xba, 0xb1, 0xdf, 0x75, 0xc3, 0x30, 0xe2, 0x2e, 0xf7, 0xa3, 0x90, 0xa9,
	0x20, 0x7b, 0x5b, 0xc2, 0x03, 0x3e, 0x8e, 0x29, 0xeb, 0xe2, 0x5f, 0x85, 0x3b, 0xff, 0x69, 0xb0,
	0x79, 0x12, 0x71, 0xff, 0xd2, 0xf7, 0x30, 0xfe, 0x94, 0xfe, 0x99, 0x52, 0xc6, 0x09, 0x01, 0x60,
	0x82, 0x81, 0xf6, 0x65, 0xb0, 0xa5, 0xed, 0x69, 0x9f, 0xd7, 0x48, 0x03, 0x74, 0xc4, 0xac, 0x12,
	0x2e, 0xeb, 0x50, 0xe1, 0xfe, 0x88, 0x5a, 0x65, 0x5c, 0x6d, 0x43, 0x93, 0xd1, 0xe4, 0xc6, 0xf7,
	0x68, 0x3f, 0x4a, 0x79, 0x9c, 0x72, 0xab, 0x82, 0xb8, 0x20, 0x42, 0x7d, 0xfd, 0x78, 0xe8, 0x0f,
	0x2c, 0x1d, 0xb1, 0x0d, 0xa8, 0x0d, 0x23, 0xc6, 0xfb, 0xa1, 0x2b, 0x8e, 0x1b, 0x12, 0x72, 0x5e,
	0x95, 0xa0, 0xf9, 0x6d, 0x42, 0x05, 0xfb, 0x29, 0x4d, 0x95, 0x04, 0xc1, 0x8f, 0x01, 0x2a, 0x79,
	0x0b, 0xaa, 0x5e, 0x90, 0x32, 0x4e, 0x93, 0x2c, 0xbd, 0x20, 0x91, 0xdb, 0x2c, 0x76, 0xbd, 0x5c,
	0xc3, 0x26, 0x98, 0xd1, 0xc5, 0x1f, 0xd4, 0xe3, 0x4a, 0x75, 0x65, 0x01, 0x44, 0x36, 0x3d, 0x07,
	0x47, 0x2e, 0x4f, 0xfc, 0x5b, 0x25, 0x0b, 0x35, 0x90, 0x8f, 0xc1, 0x50, 0xa0, 0x55, 0x15, 0x6b,
	0xb3, 0xd7, 0xe8, 0x28, 0x9b, 0x7f, 0x46, 0x90, 0xec, 0xc0, 0xc6, 0x5f, 0x6e, 0x12, 0xfa, 0xe1,
	0x55, 0xdf, 0x8b, 0xc2, 0x81, 0x2f, 0xed, 0xb2, 0xd6, 0xf0, 0xa4, 0x0d, 0xc4, 0x4b, 0x04, 0xe0,
	0xb9, 0xc1, 0xcc, 0x5e, 0x2d, 0x37, 0xc6, 0x1b, 0x52, 0xef, 0xba, 0xef, 0x87, 0x42, 0xfc, 0x8d,
	0x1b, 0x58, 0x20, 0xf0, 0x32, 0x69, 0x43, 0x3d, 0xa7, 0x4b, 0x85, 0x73, 0x96, 0x99, 0x47, 0xe7,
	0xe8, 0x88, 0xf2, 0x61, 0x34, 0xb0, 0xea, 0x02, 0xd7, 0xc9, 0x16, 0x34, 0x26, 0x19, 0x30, 0xbc,
	0x81, 0xe1, 0x1f, 0x41, 0x6b, 0x02, 0x67, 0xf1, 0x4d, 0x8c, 0x17, 0x3c, 0xca, 0xf6, 0x49, 0xd6,
	0x96, 0xcc, 0xea, 0xbc, 0x11, 0x3e, 0xff, 0x1a, 0x0f, 0x16, 0x7c, 0x46, 0x13, 0xb4, 0xbc, 0xaa,
	0xe8, 0x53, 0x69, 0xd1, 0xf5, 0xf2, 0x43, 0xd7, 0x2b, 0xcb, 0x5c, 0xd7, 0x97, 0xb9, 0x6e, 0x2c,
	0x73, 0xbd, 0xba, 0xe0, 0xfa, 0xda, 0xa3, 0x5d, 0xaf, 0xad, 0x70, 0x1d, 0x0a, 0x5c, 0x37, 0x97,
	0xba, 0x5e, 0x2f, 0x70, 0xbd, 0xb1, 0xdc, 0xf5, 0x66, 0x91, 0xeb, 0xad, 0x02, 0xd7, 0xd7, 0xd1,
	0xf5, 0x73, 0x30, 0x8f, 0x7d, 0xc6, 0xf3, 0xcb, 0x35, 0xe3, 0xaa, 0xb6, 0xcc, 0xad, 0x77, 0x6a,
	0x70, 0xe7, 0x1a, 0x1a, 0x47, 0x34, 0xa0, 0xb2, 0xa2, 0x1f, 0x80, 0x7e, 0xd2, 0x1e, 0x58, 0x57,
	0xe7, 0x18, 0xea, 0xea, 0x4b, 0x58, 0x2c, 0x86, 0x0b, 0x25, 0x9f, 0x80, 0x21, 0x67, 0x42, 0xca,
	0x30, 0x95, 0xd9, 0x6b, 0x76, 0xd4, 0x90, 0xe9, 0x9c, 0x21, 0x4a, 0x76, 0xc1, 0x40, 0x47, 0x98,
	0xc8, 0x5a, 0x16, 0xfb, 0xf5, 0xac, 0xba, 0x87, 0xf2, 0xaf, 0xf3, 0xbc, 0x04, 0x3a, 0xfe, 0x5a,
	0xd9, 0x84, 0x62, 0x84, 0xf8, 0x9e, 0xa8, 0x8d, 0xdb, 0x4f, 0x93, 0x20, 0x93, 0xba, 0x07, 0x15,
	0x16, 0x7b, 0x0c, 0x35, 0x9a, 0x3d, 0x32, 0xcb, 0xda, 0x39, 0x8b, 0xa9, 0xc7, 0xa6, 0x83, 0x07,
	0x0b, 0xa7, 0x17, 0xd4, 0x59, 0x35, 0xe6, 0x83, 0x3a, 0x57, 0x8b, 0xea, 0xbc, 0x96, 0xf3, 0x2c,
	0xd4, 0xb9, 0x86, 0xdd, 0x65, 0xc1, 0x7a, 0x42, 0x2f, 0x13, 0xca, 0x86, 0x0b, 0xb7, 0xdd, 0xfe,
	0x1d, 0x74, 0x25, 0x6b, 0xda, 0xee, 0xda, 0xa3, 0xdb, 0xbd, 0xb4, 0xa2, 0xdd, 0xd1, 0x0e, 0xe7,
	0x7f, 0x0d, 0x8c, 0x8c, 0xe1, 0xe1, 0xd8, 0x8c, 0x46, 0x23, 0x37, 0x1c, 0x64, 0x2c, 0x9f, 0x81,
	0x2e, 0xba, 0x24, 0x19, 0x8b, 0x83, 0xb2, 0x1e, 0xd6, 0x5c, 0xfa, 0xce, 0x2f, 0x72, 0xeb, 0xbb,
	0x90, 0x27, 0x63, 0x79, 0xf2, 0x32, 0x4a, 0x46, 0x69, 0xe0, 0xaa, 0x46, 0xb0, 0xbf, 0x00, 0x98,
	0xd9, 0x36, 0xa1, 0x7c, 0x4d, 0xc7, 0xd3, 0x97, 0x41, 0x7c, 0x6a, 0x9a, 0x15, 0xec, 0xa0, 0xf4,
	0xb5, 0xd6, 0x7b, 0x5d, 0x01, 0x03, 0xcb, 0xc1, 0xc8, 0x39, 0x18, 0xf8, 0xc4, 0x8c, 0x89, 0x9d,
	0x65, 0x5b, 0xf2, 0xe2, 0xd8, 0xed, 0xbc, 0x73, 0x7e, 0x8b, 0xfc, 0x41, 0xde, 0x5f, 0xce, 0xa7,
	0xff, 0xbc, 0x78, 0xf9, 0xac, 0xb4, 0xe3, 0xb4, 0xd5, 0xbb, 0x26, 0x4f, 0x77, 0x6f, 0xbe, 0xec,
	0x86, 0xc8, 0x77, 0xa0, 0xed, 0x93, 0xbf, 0xc5, 0x97, 0xab, 0x87, 0x83, 0x6c, 0x65, 0xec, 0xf3,
	0xef, 0x48, 0x01, 0xf1, 0x4f, 0x48, 0x7c, 0x64, 0x7f, 0x33, 0x4f, 0x7c, 0x97, 0xdd, 0x9c, 0xfb,
	0xee, 0xdd, 0xe4, 0x76, 0x88, 0xdf, 0x33, 0xd7, 0x62, 0xba, 0x92, 0x01, 0xf7, 0xb9, 0x06, 0x35,
	0x54, 0x27, 0x1a, 0xe6, 0x67, 0xec, 0x6a, 0x0d, 0xce, 0xfb, 0xd0, 0xf0, 0xaf, 0xd0, 0xa0, 0xc6,
	0x00, 0x69, 0x67, 0x1a, 0xe6, 0xa6, 0x42, 0x81, 0x84, 0x13, 0x94, 0xf0, 0xe3, 0xfe, 0xf7, 0x4f,
	0x94, 0xd0, 0xbd, 0x93, 0xd7, 0xf7, 0x9e, 0xdc, 0x42, 0x45, 0xce, 0x07, 0x92, 0xdf, 0xc8, 0x99,
	0xb1, 0x67, 0x6f, 0xce, 0x61, 0x99, 0x80, 0x1f, 0x50, 0xc0, 0x21, 0x79, 0xaa, 0x07, 0x17, 0x06,
	0xfe, 0x43, 0xf3, 0xd5, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x67, 0xf1, 0x7e, 0x1c, 0x09,
	0x00, 0x00,
}