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
	Spec
	AlertSpec
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
	StateType     string `protobuf:"bytes,1,opt,name=state_type,json=stateType" json:"state_type,omitempty"`
	State         string `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Time          string `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
	ServiceOutput string `protobuf:"bytes,4,opt,name=service_output,json=serviceOutput" json:"service_output,omitempty"`
	AlertPhid     string `protobuf:"bytes,5,opt,name=alert_phid,json=alertPhid" json:"alert_phid,omitempty"`
	HostName      string `protobuf:"bytes,6,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
}

func (m *NotificationRequest) Reset()                    { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()               {}
func (*NotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateReuquest struct {
	Name          string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Spec          *Spec      `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	MatrixPhid    string     `protobuf:"bytes,3,opt,name=matrix_phid,json=matrixPhid" json:"matrix_phid,omitempty"`
	Matrix        *Matrix    `protobuf:"bytes,4,opt,name=matrix" json:"matrix,omitempty"`
	AlertSpec     *AlertSpec `protobuf:"bytes,5,opt,name=alert_spec,json=alertSpec" json:"alert_spec,omitempty"`
	CheckInterval int64      `protobuf:"varint,6,opt,name=check_interval,json=checkInterval" json:"check_interval,omitempty"`
}

func (m *CreateReuquest) Reset()                    { *m = CreateReuquest{} }
func (m *CreateReuquest) String() string            { return proto.CompactTextString(m) }
func (*CreateReuquest) ProtoMessage()               {}
func (*CreateReuquest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateReuquest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CreateReuquest) GetMatrix() *Matrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

func (m *CreateReuquest) GetAlertSpec() *AlertSpec {
	if m != nil {
		return m.AlertSpec
	}
	return nil
}

type UpdateReuquest struct {
	Phid          string     `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name          string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Spec          *Spec      `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	MatrixPhid    string     `protobuf:"bytes,4,opt,name=matrix_phid,json=matrixPhid" json:"matrix_phid,omitempty"`
	Matrix        *Matrix    `protobuf:"bytes,5,opt,name=matrix" json:"matrix,omitempty"`
	AlertSpec     *AlertSpec `protobuf:"bytes,6,opt,name=alert_spec,json=alertSpec" json:"alert_spec,omitempty"`
	CheckInterval int64      `protobuf:"varint,7,opt,name=check_interval,json=checkInterval" json:"check_interval,omitempty"`
}

func (m *UpdateReuquest) Reset()                    { *m = UpdateReuquest{} }
func (m *UpdateReuquest) String() string            { return proto.CompactTextString(m) }
func (*UpdateReuquest) ProtoMessage()               {}
func (*UpdateReuquest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateReuquest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *UpdateReuquest) GetMatrix() *Matrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

func (m *UpdateReuquest) GetAlertSpec() *AlertSpec {
	if m != nil {
		return m.AlertSpec
	}
	return nil
}

type Spec struct {
	Cluster    string `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Namespace  string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	ObjectType string `protobuf:"bytes,4,opt,name=object_type,json=objectType" json:"object_type,omitempty"`
	ObjectName string `protobuf:"bytes,5,opt,name=object_name,json=objectName" json:"object_name,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type AlertSpec struct {
	WarningCondition  string `protobuf:"bytes,1,opt,name=warning_condition,json=warningCondition" json:"warning_condition,omitempty"`
	CriticalCondition string `protobuf:"bytes,2,opt,name=critical_condition,json=criticalCondition" json:"critical_condition,omitempty"`
	WarningUser       string `protobuf:"bytes,4,opt,name=warning_user,json=warningUser" json:"warning_user,omitempty"`
	WarningMethod     int32  `protobuf:"varint,5,opt,name=warning_method,json=warningMethod" json:"warning_method,omitempty"`
	CriticalUser      string `protobuf:"bytes,6,opt,name=critical_user,json=criticalUser" json:"critical_user,omitempty"`
	CriticalMethod    int32  `protobuf:"varint,7,opt,name=critical_method,json=criticalMethod" json:"critical_method,omitempty"`
	AlertInterval     int64  `protobuf:"varint,8,opt,name=alert_interval,json=alertInterval" json:"alert_interval,omitempty"`
}

func (m *AlertSpec) Reset()                    { *m = AlertSpec{} }
func (m *AlertSpec) String() string            { return proto.CompactTextString(m) }
func (*AlertSpec) ProtoMessage()               {}
func (*AlertSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListRequest struct {
	Spec *Spec `protobuf:"bytes,1,opt,name=spec" json:"spec,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type DeleteRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Spec *Spec  `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts []*Alert       `protobuf:"bytes,2,rep,name=alerts" json:"alerts,omitempty"`
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

func (m *ListResponse) GetAlerts() []*Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

type Alert struct {
	Phid            string     `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name            string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IcingaUrl       string     `protobuf:"bytes,3,opt,name=icinga_url,json=icingaUrl" json:"icinga_url,omitempty"`
	Matrix          *Matrix    `protobuf:"bytes,4,opt,name=matrix" json:"matrix,omitempty"`
	Spec            *AlertSpec `protobuf:"bytes,5,opt,name=spec" json:"spec,omitempty"`
	RefreshInterval int64      `protobuf:"varint,6,opt,name=refresh_interval,json=refreshInterval" json:"refresh_interval,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Alert) GetMatrix() *Matrix {
	if m != nil {
		return m.Matrix
	}
	return nil
}

func (m *Alert) GetSpec() *AlertSpec {
	if m != nil {
		return m.Spec
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
func (*Matrix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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
	proto.RegisterType((*Spec)(nil), "alert.Spec")
	proto.RegisterType((*AlertSpec)(nil), "alert.AlertSpec")
	proto.RegisterType((*ListRequest)(nil), "alert.ListRequest")
	proto.RegisterType((*DeleteRequest)(nil), "alert.DeleteRequest")
	proto.RegisterType((*ListResponse)(nil), "alert.ListResponse")
	proto.RegisterType((*Alert)(nil), "alert.Alert")
	proto.RegisterType((*Matrix)(nil), "alert.Matrix")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

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
	// 931 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x96, 0x93, 0xd8, 0xbd, 0x39, 0xf9, 0xb9, 0xb9, 0x73, 0xcb, 0x55, 0x94, 0x16, 0x41, 0x4d,
	0xcb, 0x4f, 0x11, 0x0e, 0x94, 0x4d, 0xd5, 0x1d, 0x6a, 0x59, 0x20, 0xd1, 0x02, 0x2e, 0x01, 0x09,
	0xb5, 0x04, 0xd7, 0x99, 0x36, 0xa6, 0x8e, 0x6d, 0xec, 0x71, 0x21, 0xaa, 0x2a, 0x21, 0x1e, 0x80,
	0x0d, 0x0b, 0x24, 0x1e, 0x80, 0x87, 0x60, 0xc7, 0x06, 0x89, 0x35, 0xaf, 0x00, 0x0f, 0xc1, 0x8e,
	0x99, 0x33, 0x33, 0xce, 0x8f, 0x92, 0x12, 0x5d, 0xa9, 0x9b, 0xc8, 0xf3, 0x9d, 0xd3, 0x33, 0xdf,
	0x37, 0xe7, 0xaf, 0x50, 0xf3, 0x42, 0x9a, 0x32, 0x27, 0x49, 0x63, 0x16, 0x13, 0x13, 0x0f, 0x9d,
	0xcd, 0xab, 0x38, 0xbe, 0x0a, 0x69, 0xd7, 0x4b, 0x82, 0xae, 0x17, 0x45, 0x31, 0xf3, 0x58, 0x10,
	0x47, 0x99, 0x74, 0xea, 0x3c, 0x13, 0xf0, 0x80, 0x8d, 0x13, 0x9a, 0x75, 0xf1, 0x57, 0xe2, 0xf6,
	0xef, 0x06, 0x3c, 0x3d, 0x89, 0x59, 0x70, 0x19, 0xf8, 0xe8, 0xef, 0xd2, 0x6f, 0x72, 0x9a, 0x31,
	0xf2, 0x22, 0x40, 0xc6, 0x23, 0xd0, 0xbe, 0x70, 0x6e, 0x1b, 0x2f, 0x1b, 0xaf, 0x57, 0xdd, 0x2a,
	0x22, 0x9f, 0x72, 0x80, 0xac, 0x83, 0x89, 0x87, 0x76, 0x09, 0x2d, 0xf2, 0x40, 0x08, 0x54, 0x58,
	0x30, 0xa2, 0xed, 0x32, 0x82, 0xf8, 0x4d, 0x76, 0xa0, 0x99, 0xd1, 0xf4, 0x26, 0xf0, 0x69, 0x3f,
	0xce, 0x59, 0x92, 0xb3, 0x76, 0x05, 0xad, 0x0d, 0x85, 0x7e, 0x84, 0xa0, 0xb8, 0x0f, 0x65, 0xf4,
	0x93, 0x61, 0x30, 0x68, 0x9b, 0xf2, 0x3e, 0x44, 0x3e, 0xe6, 0x00, 0xd9, 0x80, 0xea, 0x30, 0xce,
	0x58, 0x3f, 0xf2, 0x78, 0x78, 0x0b, 0xad, 0x8f, 0x04, 0x70, 0xc2, 0xcf, 0xf6, 0x3f, 0x06, 0x34,
	0x0f, 0x53, 0xca, 0x19, 0xb8, 0x34, 0x97, 0xf4, 0x39, 0x13, 0x74, 0x95, 0xc4, 0xf1, 0x9b, 0xbc,
	0x04, 0x95, 0x2c, 0xa1, 0x3e, 0x52, 0xae, 0xed, 0xd5, 0x1c, 0xf9, 0x86, 0xa7, 0x1c, 0x72, 0xd1,
	0xc0, 0x1d, 0x6a, 0x23, 0x8f, 0xa5, 0xc1, 0x77, 0x92, 0x84, 0x54, 0x01, 0x12, 0x42, 0x16, 0x3b,
	0x60, 0xc9, 0x13, 0x6a, 0xa8, 0xed, 0x35, 0x54, 0x8c, 0x63, 0x04, 0x5d, 0x65, 0x24, 0x5d, 0xad,
	0x05, 0xaf, 0x33, 0xd1, 0xb5, 0xa5, 0x5c, 0xdf, 0x13, 0xbf, 0x78, 0xa7, 0x54, 0x27, 0x3e, 0xc5,
	0x1b, 0xf9, 0x43, 0xea, 0x5f, 0xf7, 0x83, 0x88, 0xf1, 0x57, 0xf1, 0x42, 0x94, 0x58, 0x76, 0x1b,
	0x88, 0x7e, 0xa0, 0x40, 0xfb, 0x5f, 0xae, 0xb3, 0x97, 0x0c, 0xe6, 0x74, 0x22, 0x57, 0xa5, 0x53,
	0x7c, 0x17, 0xda, 0x4b, 0x0b, 0xb4, 0x97, 0x57, 0xd4, 0x5e, 0xb9, 0x47, 0xbb, 0xb9, 0xba, 0x76,
	0xeb, 0x79, 0xb4, 0xaf, 0x2d, 0xd2, 0xfe, 0xbd, 0x01, 0x15, 0xf4, 0x6f, 0xc3, 0x9a, 0x1f, 0xe6,
	0x19, 0x87, 0x95, 0x40, 0x7d, 0x24, 0x9b, 0x50, 0x15, 0x5a, 0xb3, 0xc4, 0xf3, 0x75, 0x09, 0x4e,
	0x00, 0x21, 0x30, 0xbe, 0xf8, 0x9a, 0xfa, 0x4c, 0x56, 0xb4, 0x12, 0x28, 0x21, 0x2c, 0xe9, 0x89,
	0x03, 0xbe, 0x9e, 0x39, 0xed, 0x80, 0x65, 0xf6, 0x6b, 0x09, 0xaa, 0x85, 0x04, 0xf2, 0x26, 0x3c,
	0xf9, 0xd6, 0x4b, 0xa3, 0x20, 0xba, 0xea, 0xfb, 0x71, 0x34, 0x08, 0x44, 0xf3, 0xa8, 0x34, 0xb4,
	0x94, 0xe1, 0x50, 0xe3, 0xe4, 0x2d, 0x20, 0x7e, 0xca, 0x3f, 0x7d, 0x2f, 0x9c, 0xf2, 0x96, 0xfc,
	0x9f, 0x68, 0xcb, 0xc4, 0x7d, 0x0b, 0xea, 0x3a, 0x76, 0xce, 0xdb, 0x44, 0x91, 0xad, 0x29, 0xac,
	0xc7, 0x21, 0xf1, 0x6c, 0xda, 0x65, 0x44, 0xd9, 0x30, 0x96, 0x3d, 0x63, 0xba, 0x0d, 0x85, 0x1e,
	0x23, 0x48, 0x5e, 0x81, 0x46, 0x71, 0x31, 0x86, 0x92, 0xbd, 0x53, 0xd7, 0x20, 0xc6, 0x7a, 0x0d,
	0x1e, 0x17, 0x4e, 0x2a, 0xd8, 0x1a, 0x06, 0x6b, 0x6a, 0x58, 0x45, 0xe3, 0x97, 0xca, 0xe4, 0x16,
	0xb9, 0x7a, 0x24, 0x73, 0x85, 0x68, 0x91, 0x2b, 0x07, 0x6a, 0x1f, 0x06, 0x19, 0xd3, 0xa3, 0x44,
	0xd7, 0x9e, 0xb1, 0xa4, 0xf6, 0xec, 0x23, 0x68, 0x1c, 0xd1, 0x90, 0x8a, 0xb2, 0x5e, 0x5e, 0xd5,
	0xff, 0xd7, 0xbd, 0xf6, 0x19, 0xd4, 0xe5, 0xad, 0x59, 0xc2, 0xc7, 0x1e, 0x25, 0xaf, 0x82, 0x25,
	0xa6, 0x52, 0x9e, 0xa9, 0x8b, 0x9b, 0x8e, 0x1c, 0x7f, 0xce, 0x29, 0xa2, 0xae, 0xb2, 0x92, 0x6d,
	0xb0, 0x30, 0x56, 0xc6, 0x43, 0x97, 0xb9, 0x5f, 0x7d, 0xba, 0x5a, 0x5d, 0x65, 0xb3, 0xff, 0x30,
	0xc0, 0x44, 0x64, 0xe5, 0x96, 0xe3, 0x13, 0x2d, 0xf0, 0x79, 0x26, 0xbc, 0x7e, 0x9e, 0x86, 0xba,
	0x1e, 0x25, 0xd2, 0x4b, 0xc3, 0x55, 0x67, 0xc9, 0xb6, 0x92, 0xbd, 0x6c, 0x8a, 0xc8, 0xee, 0x7d,
	0x03, 0x5a, 0x29, 0xbd, 0x4c, 0x69, 0x36, 0x9c, 0x1f, 0x21, 0x8f, 0x15, 0x5e, 0x24, 0xe7, 0x37,
	0x03, 0x2c, 0x79, 0xc7, 0xc2, 0x21, 0x29, 0xda, 0x2b, 0x1e, 0x8d, 0xbc, 0x68, 0x50, 0xb4, 0x97,
	0x3c, 0x12, 0x07, 0x4c, 0x9e, 0x9d, 0x74, 0xcc, 0xa5, 0x88, 0x67, 0x6a, 0xcf, 0xf0, 0x75, 0x3e,
	0x11, 0xa6, 0xf7, 0x23, 0x96, 0x8e, 0x5d, 0xe9, 0x26, 0x22, 0x5d, 0xc6, 0xe9, 0x28, 0x0f, 0x3d,
	0x55, 0xbf, 0xfa, 0xd8, 0xd9, 0x07, 0x98, 0xb8, 0x93, 0x16, 0x94, 0xaf, 0xe9, 0x58, 0x91, 0x10,
	0x9f, 0x62, 0xb9, 0x70, 0xa6, 0x79, 0xb1, 0x5c, 0xf0, 0x70, 0x50, 0xda, 0x37, 0xf6, 0xfe, 0x34,
	0xc1, 0x42, 0xed, 0x19, 0x39, 0x07, 0x0b, 0xf7, 0xd6, 0x98, 0x74, 0x14, 0x93, 0x05, 0x6b, 0xac,
	0xb3, 0xae, 0x93, 0xfe, 0x59, 0x1c, 0x0c, 0x74, 0x69, 0xd8, 0x5b, 0x3f, 0xfc, 0xf5, 0xf7, 0x4f,
	0xa5, 0x0d, 0xfb, 0x99, 0x5c, 0x96, 0xe2, 0xaf, 0xbb, 0x37, 0x6f, 0x3b, 0xef, 0x74, 0x23, 0x8c,
	0x78, 0x60, 0xec, 0x92, 0x9f, 0xf9, 0x33, 0xc9, 0x9d, 0x42, 0x5e, 0x50, 0xf1, 0x67, 0x57, 0xcc,
	0x92, 0xd0, 0x5f, 0x61, 0xe8, 0x2f, 0x3a, 0xbd, 0xf9, 0xd0, 0xb7, 0x22, 0x51, 0x8e, 0x9a, 0x55,
	0x77, 0xea, 0x58, 0x8c, 0x27, 0x0d, 0x4c, 0x8d, 0xa8, 0x39, 0x48, 0xb8, 0xde, 0x69, 0x66, 0x72,
	0x0b, 0x14, 0xcc, 0x66, 0x97, 0xc2, 0xfd, 0xcc, 0xec, 0x87, 0x63, 0xf6, 0x0b, 0x67, 0x26, 0x1b,
	0x99, 0xac, 0x2b, 0x66, 0x33, 0x7d, 0xbd, 0x84, 0xd8, 0x00, 0x89, 0x7d, 0xb9, 0x7b, 0xf6, 0x20,
	0xc4, 0xba, 0xb7, 0xa2, 0x43, 0xef, 0xc8, 0x8f, 0x7c, 0x81, 0x88, 0xf9, 0x40, 0x88, 0xa2, 0x36,
	0x35, 0xa2, 0x3a, 0x4f, 0x67, 0x30, 0xc5, 0xeb, 0x1c, 0x79, 0x7d, 0x4e, 0x1e, 0xe6, 0xc1, 0x2e,
	0x2c, 0xfc, 0x07, 0xec, 0xdd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xab, 0xfa, 0xac, 0x4f, 0xcc,
	0x09, 0x00, 0x00,
}