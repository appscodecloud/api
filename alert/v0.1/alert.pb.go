// Code generated by protoc-gen-go.
// source: alert.proto
// DO NOT EDIT!

/*
Package alert is a generated protocol buffer package.

It is generated from these files:
	alert.proto

It has these top-level messages:
	NotificationRequest
	AcknowledgeRequest
	CreateRequest
	UpdateRequest
	Spec
	AlertSpec
	ListRequest
	DeleteRequest
	ListResponse
	Alert
	IcingaService
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

type AcknowledgeRequest struct {
	Phid    string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
}

func (m *AcknowledgeRequest) Reset()                    { *m = AcknowledgeRequest{} }
func (m *AcknowledgeRequest) String() string            { return proto.CompactTextString(m) }
func (*AcknowledgeRequest) ProtoMessage()               {}
func (*AcknowledgeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateRequest struct {
	Name              string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Spec              *Spec          `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	IcingaServicePhid string         `protobuf:"bytes,3,opt,name=icinga_service_phid,json=icingaServicePhid" json:"icinga_service_phid,omitempty"`
	IcingaService     *IcingaService `protobuf:"bytes,4,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	AlertSpec         *AlertSpec     `protobuf:"bytes,5,opt,name=alert_spec,json=alertSpec" json:"alert_spec,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CreateRequest) GetIcingaService() *IcingaService {
	if m != nil {
		return m.IcingaService
	}
	return nil
}

func (m *CreateRequest) GetAlertSpec() *AlertSpec {
	if m != nil {
		return m.AlertSpec
	}
	return nil
}

type UpdateRequest struct {
	Phid              string         `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name              string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Spec              *Spec          `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	IcingaServicePhid string         `protobuf:"bytes,4,opt,name=icinga_service_phid,json=icingaServicePhid" json:"icinga_service_phid,omitempty"`
	IcingaService     *IcingaService `protobuf:"bytes,5,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	AlertSpec         *AlertSpec     `protobuf:"bytes,6,opt,name=alert_spec,json=alertSpec" json:"alert_spec,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *UpdateRequest) GetIcingaService() *IcingaService {
	if m != nil {
		return m.IcingaService
	}
	return nil
}

func (m *UpdateRequest) GetAlertSpec() *AlertSpec {
	if m != nil {
		return m.AlertSpec
	}
	return nil
}

type Spec struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Namespace  string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	ObjectType string `protobuf:"bytes,3,opt,name=object_type,json=objectType" json:"object_type,omitempty"`
	ObjectName string `protobuf:"bytes,4,opt,name=object_name,json=objectName" json:"object_name,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AlertSpec struct {
	WarningUser    string `protobuf:"bytes,1,opt,name=warning_user,json=warningUser" json:"warning_user,omitempty"`
	WarningMethod  int32  `protobuf:"varint,2,opt,name=warning_method,json=warningMethod" json:"warning_method,omitempty"`
	CriticalUser   string `protobuf:"bytes,3,opt,name=critical_user,json=criticalUser" json:"critical_user,omitempty"`
	CriticalMethod int32  `protobuf:"varint,4,opt,name=critical_method,json=criticalMethod" json:"critical_method,omitempty"`
	AlertInterval  int64  `protobuf:"varint,5,opt,name=alert_interval,json=alertInterval" json:"alert_interval,omitempty"`
}

func (m *AlertSpec) Reset()                    { *m = AlertSpec{} }
func (m *AlertSpec) String() string            { return proto.CompactTextString(m) }
func (*AlertSpec) ProtoMessage()               {}
func (*AlertSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListRequest struct {
	Spec *Spec `protobuf:"bytes,1,opt,name=spec" json:"spec,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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
	Phid          string         `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name          string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IcingaUrl     string         `protobuf:"bytes,3,opt,name=icinga_url,json=icingaUrl" json:"icinga_url,omitempty"`
	IcingaService *IcingaService `protobuf:"bytes,4,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	AlertSpec     *AlertSpec     `protobuf:"bytes,5,opt,name=alert_spec,json=alertSpec" json:"alert_spec,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Alert) GetIcingaService() *IcingaService {
	if m != nil {
		return m.IcingaService
	}
	return nil
}

func (m *Alert) GetAlertSpec() *AlertSpec {
	if m != nil {
		return m.AlertSpec
	}
	return nil
}

type IcingaService struct {
	Name              string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CheckCommand      string            `protobuf:"bytes,2,opt,name=check_command,json=checkCommand" json:"check_command,omitempty"`
	Query             map[string]string `protobuf:"bytes,3,rep,name=query" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FormulaR          string            `protobuf:"bytes,4,opt,name=formula_r,json=formulaR" json:"formula_r,omitempty"`
	WarningCondition  string            `protobuf:"bytes,5,opt,name=warning_condition,json=warningCondition" json:"warning_condition,omitempty"`
	CriticalCondition string            `protobuf:"bytes,6,opt,name=critical_condition,json=criticalCondition" json:"critical_condition,omitempty"`
	CheckInterval     int64             `protobuf:"varint,7,opt,name=check_interval,json=checkInterval" json:"check_interval,omitempty"`
}

func (m *IcingaService) Reset()                    { *m = IcingaService{} }
func (m *IcingaService) String() string            { return proto.CompactTextString(m) }
func (*IcingaService) ProtoMessage()               {}
func (*IcingaService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IcingaService) GetQuery() map[string]string {
	if m != nil {
		return m.Query
	}
	return nil
}

func init() {
	proto.RegisterType((*NotificationRequest)(nil), "alert.NotificationRequest")
	proto.RegisterType((*AcknowledgeRequest)(nil), "alert.AcknowledgeRequest")
	proto.RegisterType((*CreateRequest)(nil), "alert.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "alert.UpdateRequest")
	proto.RegisterType((*Spec)(nil), "alert.Spec")
	proto.RegisterType((*AlertSpec)(nil), "alert.AlertSpec")
	proto.RegisterType((*ListRequest)(nil), "alert.ListRequest")
	proto.RegisterType((*DeleteRequest)(nil), "alert.DeleteRequest")
	proto.RegisterType((*ListResponse)(nil), "alert.ListResponse")
	proto.RegisterType((*Alert)(nil), "alert.Alert")
	proto.RegisterType((*IcingaService)(nil), "alert.IcingaService")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Alerts service

type AlertsClient interface {
	Notify(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Acknowledge(ctx context.Context, in *AcknowledgeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
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

func (c *alertsClient) Acknowledge(ctx context.Context, in *AcknowledgeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Acknowledge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.Alerts/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
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
	Acknowledge(context.Context, *AcknowledgeRequest) (*dtypes.VoidResponse, error)
	Create(context.Context, *CreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *UpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*dtypes.VoidResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterAlertsServer(s *grpc.Server, srv AlertsServer) {
	s.RegisterService(&_Alerts_serviceDesc, srv)
}

func _Alerts_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Notify(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Acknowledge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Acknowledge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Acknowledge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Acknowledge(ctx, req.(*AcknowledgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.Alerts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "Acknowledge",
			Handler:    _Alerts_Acknowledge_Handler,
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
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0x97, 0xf7, 0x8f, 0xdb, 0x7d, 0xbb, 0x1b, 0x92, 0x49, 0x84, 0x96, 0xa5, 0x28, 0x60, 0xc2,
	0xbf, 0x20, 0xbc, 0x10, 0x84, 0x54, 0x95, 0x53, 0x49, 0x39, 0x54, 0x82, 0x02, 0x0e, 0xe1, 0x84,
	0x30, 0xae, 0x3d, 0xd9, 0x98, 0x78, 0x3d, 0xc6, 0x33, 0x4e, 0xb5, 0xaa, 0x22, 0x21, 0xbe, 0x02,
	0x07, 0x8e, 0x1c, 0xf8, 0x1a, 0x5c, 0xb8, 0x72, 0xe5, 0xc4, 0x15, 0xc1, 0x57, 0xe0, 0xcc, 0xcc,
	0x9b, 0x19, 0xef, 0x9a, 0xee, 0xb6, 0xa5, 0x52, 0x73, 0x59, 0xcd, 0xfc, 0xe6, 0xcd, 0x7b, 0xbf,
	0xdf, 0xf3, 0x9b, 0xf7, 0x16, 0xfa, 0x51, 0x46, 0x4b, 0xe1, 0x17, 0x25, 0x13, 0x8c, 0x74, 0x71,
	0x33, 0xbe, 0x36, 0x65, 0x6c, 0x9a, 0xd1, 0x49, 0x54, 0xa4, 0x93, 0x28, 0xcf, 0x99, 0x88, 0x44,
	0xca, 0x72, 0xae, 0x8d, 0xc6, 0xcf, 0x2a, 0x38, 0x11, 0xf3, 0x82, 0xf2, 0x09, 0xfe, 0x6a, 0xdc,
	0xfb, 0xd5, 0x81, 0xed, 0x3b, 0x4c, 0xa4, 0x27, 0x69, 0x8c, 0xf6, 0x01, 0xfd, 0xb6, 0xa2, 0x5c,
	0x90, 0x17, 0x00, 0xb8, 0xf4, 0x40, 0x43, 0x65, 0x3c, 0x72, 0x5e, 0x74, 0x5e, 0xef, 0x05, 0x3d,
	0x44, 0x3e, 0x97, 0x00, 0xd9, 0x81, 0x2e, 0x6e, 0x46, 0x2d, 0x3c, 0xd1, 0x1b, 0x42, 0xa0, 0x23,
	0xd2, 0x19, 0x1d, 0xb5, 0x11, 0xc4, 0x35, 0x79, 0x05, 0x36, 0x38, 0x2d, 0xcf, 0xd3, 0x98, 0x86,
	0xac, 0x12, 0x45, 0x25, 0x46, 0x1d, 0x3c, 0x1d, 0x1a, 0xf4, 0x13, 0x04, 0x55, 0x3c, 0x94, 0x11,
	0x16, 0xa7, 0x69, 0x32, 0xea, 0xea, 0x78, 0x88, 0x7c, 0x2a, 0x01, 0xf2, 0x3c, 0xf4, 0x4e, 0x19,
	0x17, 0x61, 0x1e, 0x49, 0xf7, 0x2e, 0x9e, 0x5e, 0x55, 0xc0, 0x1d, 0xb9, 0xf7, 0x3e, 0x00, 0x72,
	0x33, 0x3e, 0xcb, 0xd9, 0xbd, 0x8c, 0x26, 0x53, 0x6a, 0x15, 0x48, 0x32, 0xe8, 0x4b, 0x73, 0xc7,
	0x35, 0x19, 0xc1, 0x95, 0x98, 0xcd, 0x66, 0x34, 0x17, 0x86, 0xb8, 0xdd, 0x7a, 0x7f, 0x3a, 0x30,
	0x3c, 0x2c, 0xa9, 0x54, 0xb1, 0x74, 0x1f, 0xa3, 0x99, 0xfb, 0x6a, 0x4d, 0x76, 0xa1, 0xc3, 0x0b,
	0x1a, 0xe3, 0xe5, 0xfe, 0x41, 0xdf, 0xd7, 0x9f, 0xe1, 0x48, 0x42, 0x01, 0x1e, 0x10, 0x1f, 0xb6,
	0xd3, 0x38, 0xcd, 0xa7, 0x51, 0x68, 0x45, 0x23, 0x07, 0x9d, 0x90, 0x2d, 0x7d, 0x74, 0xa4, 0x4f,
	0x50, 0xd7, 0xfb, 0xb0, 0xd1, 0xb4, 0xc7, 0xec, 0xf4, 0x0f, 0x76, 0x8c, 0xeb, 0xdb, 0xcb, 0x37,
	0x82, 0x61, 0xc3, 0x01, 0x99, 0xd8, 0x9c, 0x21, 0xa7, 0x2e, 0x5e, 0xdc, 0x34, 0x17, 0x6f, 0xaa,
	0x5f, 0x24, 0xa6, 0xb3, 0xa8, 0x96, 0xde, 0x3f, 0x52, 0xe4, 0x71, 0x91, 0x34, 0x45, 0x3e, 0x90,
	0x24, 0x2b, 0xbc, 0xb5, 0x42, 0x78, 0xfb, 0x7f, 0x0a, 0xef, 0x3c, 0xbe, 0xf0, 0xee, 0x93, 0x0a,
	0x77, 0x1f, 0x2d, 0xfc, 0x3b, 0x07, 0x3a, 0x6a, 0x81, 0x05, 0x90, 0x55, 0x5c, 0xd0, 0xd2, 0x48,
	0xb6, 0x5b, 0x72, 0x0d, 0x7a, 0x4a, 0x29, 0x2f, 0xa2, 0xd8, 0x4a, 0x5f, 0x00, 0x52, 0x7f, 0x9f,
	0xdd, 0xfd, 0x86, 0xc6, 0x42, 0xbf, 0x07, 0xfd, 0x3d, 0x41, 0x43, 0xf8, 0x20, 0x16, 0x06, 0x98,
	0xbb, 0xce, 0xb2, 0x01, 0x16, 0xe9, 0x6f, 0x0e, 0xf4, 0x6a, 0x6e, 0xe4, 0x25, 0x18, 0xdc, 0x8b,
	0xca, 0x5c, 0x8a, 0x0a, 0x2b, 0x5e, 0x93, 0xe9, 0x1b, 0xec, 0x58, 0x42, 0xea, 0xe1, 0x58, 0x93,
	0x19, 0x15, 0xa7, 0x2c, 0x41, 0x56, 0xdd, 0x60, 0x68, 0xd0, 0x8f, 0x11, 0x24, 0x2f, 0xc3, 0x30,
	0x2e, 0x53, 0x21, 0x9f, 0x6f, 0xa6, 0x5d, 0x69, 0x6e, 0x03, 0x0b, 0xa2, 0xaf, 0xd7, 0xe0, 0x99,
	0xda, 0xc8, 0x38, 0xeb, 0xa0, 0xb3, 0x0d, 0x0b, 0x1b, 0x6f, 0x32, 0xa8, 0xce, 0x6c, 0x9a, 0xcb,
	0xa4, 0x9c, 0x47, 0x19, 0x7e, 0x96, 0x76, 0x30, 0x44, 0xf4, 0xb6, 0x01, 0x3d, 0x1f, 0xfa, 0x1f,
	0xa5, 0x5c, 0xd8, 0x2a, 0xb2, 0xd5, 0xe1, 0xac, 0xa9, 0x0e, 0xef, 0x16, 0x0c, 0x6f, 0xd1, 0x8c,
	0x3e, 0xbc, 0xee, 0x1e, 0xf5, 0xb8, 0xbc, 0x2f, 0x61, 0xa0, 0xa3, 0xf2, 0x42, 0x36, 0x36, 0x4a,
	0x5e, 0x05, 0x57, 0xf5, 0x9d, 0x8a, 0x9b, 0xc0, 0x1b, 0xbe, 0x6e, 0x70, 0xfe, 0x11, 0xa2, 0x81,
	0x39, 0x25, 0x7b, 0xe0, 0xa2, 0x2f, 0x2e, 0x5d, 0xb7, 0xa5, 0xdd, 0x60, 0xb9, 0x54, 0x02, 0x73,
	0xe6, 0xfd, 0xe2, 0x40, 0x17, 0x91, 0xc7, 0x7e, 0x14, 0xb2, 0x67, 0x99, 0x1a, 0xae, 0xca, 0xcc,
	0xe4, 0xbd, 0xa7, 0x91, 0xe3, 0x32, 0xbb, 0xe4, 0xb7, 0xfd, 0x47, 0x0b, 0x86, 0x0d, 0x8f, 0x2b,
	0x1b, 0x98, 0xaa, 0x96, 0x53, 0x1a, 0x9f, 0x85, 0xaa, 0xef, 0x45, 0x79, 0x62, 0xf4, 0x0c, 0x10,
	0x3c, 0xd4, 0x18, 0x79, 0x0f, 0xba, 0xf2, 0x2b, 0x95, 0x73, 0x29, 0x49, 0xa5, 0x6b, 0x77, 0x15,
	0x5f, 0xff, 0x33, 0x65, 0xf1, 0x61, 0x2e, 0xca, 0x79, 0xa0, 0xad, 0x55, 0x8f, 0x3e, 0x61, 0xe5,
	0xac, 0xca, 0xa2, 0xb0, 0x34, 0x0f, 0xe0, 0xaa, 0x01, 0x02, 0xf2, 0x26, 0x6c, 0xd9, 0x6a, 0x8e,
	0x59, 0x9e, 0xa4, 0x6a, 0xd6, 0x98, 0x36, 0xbf, 0x69, 0x0e, 0x0e, 0x2d, 0x4e, 0xde, 0x02, 0x52,
	0x97, 0xeb, 0xc2, 0x5a, 0xb7, 0xfd, 0x2d, 0x7b, 0xb2, 0x30, 0x97, 0x45, 0xab, 0x45, 0xd5, 0x45,
	0x7b, 0x45, 0x17, 0x2d, 0xa2, 0xb6, 0x68, 0xc7, 0xd7, 0x01, 0x16, 0xa4, 0xc9, 0x26, 0xb4, 0xcf,
	0xe8, 0xdc, 0x24, 0x47, 0x2d, 0xd5, 0x4c, 0x93, 0x66, 0x55, 0x3d, 0xd3, 0x70, 0x73, 0xa3, 0x75,
	0xdd, 0x39, 0xf8, 0xdb, 0x05, 0x17, 0x93, 0xce, 0xc9, 0x09, 0xb8, 0x38, 0x2e, 0xe7, 0x64, 0x6c,
	0xd2, 0xb2, 0x62, 0x7a, 0x8e, 0x77, 0x6c, 0x25, 0x7e, 0xc1, 0xd2, 0xc4, 0xd6, 0xab, 0xb7, 0xff,
	0xfd, 0xef, 0x7f, 0xfd, 0xd0, 0xda, 0xf3, 0x76, 0xe5, 0x8c, 0x2e, 0x78, 0xcc, 0x12, 0x33, 0xac,
	0x95, 0x9b, 0xc9, 0xf9, 0xdb, 0xfe, 0x3b, 0x93, 0x1c, 0x5d, 0xdf, 0x70, 0xf6, 0x09, 0x83, 0xfe,
	0xd2, 0x4c, 0x23, 0xcf, 0xd9, 0x4f, 0xff, 0xc0, 0x9c, 0x5b, 0x13, 0x6b, 0x82, 0xb1, 0xde, 0xf0,
	0xf6, 0xd6, 0xc6, 0x8a, 0x16, 0xae, 0x54, 0xc0, 0x9f, 0x1c, 0x70, 0xf5, 0x00, 0x24, 0xb6, 0x40,
	0x1b, 0xf3, 0x70, 0x4d, 0x9c, 0x14, 0xe3, 0xc4, 0xe3, 0xaf, 0xd6, 0xc6, 0xb9, 0xaf, 0x2a, 0xd8,
	0x37, 0x6d, 0xf6, 0xc2, 0x6c, 0xeb, 0xce, 0x6a, 0x81, 0xa5, 0xee, 0xfa, 0x1f, 0x48, 0x99, 0x5e,
	0x58, 0x86, 0x7a, 0x7a, 0xd5, 0x0c, 0x1b, 0xc3, 0xec, 0xe1, 0x0c, 0xbd, 0x4b, 0x60, 0xf8, 0xb3,
	0x64, 0xa8, 0xfb, 0x5c, 0xcd, 0xb0, 0xd1, 0xf6, 0xd6, 0x30, 0x9c, 0x21, 0xc3, 0xe9, 0x3e, 0x7d,
	0xba, 0x0c, 0x27, 0xf7, 0x55, 0x27, 0xbb, 0x20, 0x3f, 0xca, 0x61, 0xa8, 0xfa, 0x28, 0x21, 0x86,
	0xe3, 0x52, 0x2b, 0x1f, 0x6f, 0x37, 0x30, 0x43, 0xf0, 0x04, 0x09, 0x7e, 0x4d, 0x9e, 0x72, 0x0a,
	0xef, 0xba, 0xf8, 0x9f, 0xf4, 0xdd, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x17, 0x24, 0xed, 0xc1,
	0xdf, 0x0a, 0x00, 0x00,
}
