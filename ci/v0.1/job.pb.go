// Code generated by protoc-gen-go.
// source: job.proto
// DO NOT EDIT!

package ci

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

type JobListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Jobs   []*Job         `protobuf:"bytes,2,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *JobListResponse) Reset()                    { *m = JobListResponse{} }
func (m *JobListResponse) String() string            { return proto.CompactTextString(m) }
func (*JobListResponse) ProtoMessage()               {}
func (*JobListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *JobListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JobListResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type JobBuildRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Param string `protobuf:"bytes,2,opt,name=param" json:"param,omitempty"`
}

func (m *JobBuildRequest) Reset()                    { *m = JobBuildRequest{} }
func (m *JobBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*JobBuildRequest) ProtoMessage()               {}
func (*JobBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type JobDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *JobDescribeRequest) Reset()                    { *m = JobDescribeRequest{} }
func (m *JobDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*JobDescribeRequest) ProtoMessage()               {}
func (*JobDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type JobDescribeResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name          string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description   string         `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	DisplayName   string         `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	BuildCount    int64          `protobuf:"varint,5,opt,name=build_count,json=buildCount" json:"build_count,omitempty"`
	JobColor      string         `protobuf:"bytes,6,opt,name=job_color,json=jobColor" json:"job_color,omitempty"`
	BuildsCommand string         `protobuf:"bytes,7,opt,name=builds_command,json=buildsCommand" json:"builds_command,omitempty"`
}

func (m *JobDescribeResponse) Reset()                    { *m = JobDescribeResponse{} }
func (m *JobDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*JobDescribeResponse) ProtoMessage()               {}
func (*JobDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *JobDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type JobDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *JobDeleteRequest) Reset()                    { *m = JobDeleteRequest{} }
func (m *JobDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*JobDeleteRequest) ProtoMessage()               {}
func (*JobDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type JobCreateRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ShFile string `protobuf:"bytes,2,opt,name=sh_file,json=shFile" json:"sh_file,omitempty"`
}

func (m *JobCreateRequest) Reset()                    { *m = JobCreateRequest{} }
func (m *JobCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*JobCreateRequest) ProtoMessage()               {}
func (*JobCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type JobCopyRequest struct {
	Source      string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
}

func (m *JobCopyRequest) Reset()                    { *m = JobCopyRequest{} }
func (m *JobCopyRequest) String() string            { return proto.CompactTextString(m) }
func (*JobCopyRequest) ProtoMessage()               {}
func (*JobCopyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type Job struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color" json:"color,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*JobListResponse)(nil), "ci.JobListResponse")
	proto.RegisterType((*JobBuildRequest)(nil), "ci.JobBuildRequest")
	proto.RegisterType((*JobDescribeRequest)(nil), "ci.JobDescribeRequest")
	proto.RegisterType((*JobDescribeResponse)(nil), "ci.JobDescribeResponse")
	proto.RegisterType((*JobDeleteRequest)(nil), "ci.JobDeleteRequest")
	proto.RegisterType((*JobCreateRequest)(nil), "ci.JobCreateRequest")
	proto.RegisterType((*JobCopyRequest)(nil), "ci.JobCopyRequest")
	proto.RegisterType((*Job)(nil), "ci.Job")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Jobs service

type JobsClient interface {
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*JobListResponse, error)
	Describe(ctx context.Context, in *JobDescribeRequest, opts ...grpc.CallOption) (*JobDescribeResponse, error)
	Create(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Copy(ctx context.Context, in *JobCopyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Build(ctx context.Context, in *JobBuildRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *JobDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type jobsClient struct {
	cc *grpc.ClientConn
}

func NewJobsClient(cc *grpc.ClientConn) JobsClient {
	return &jobsClient{cc}
}

func (c *jobsClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*JobListResponse, error) {
	out := new(JobListResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Describe(ctx context.Context, in *JobDescribeRequest, opts ...grpc.CallOption) (*JobDescribeResponse, error) {
	out := new(JobDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Create(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Copy(ctx context.Context, in *JobCopyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Copy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Build(ctx context.Context, in *JobBuildRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Delete(ctx context.Context, in *JobDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.Jobs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jobs service

type JobsServer interface {
	List(context.Context, *dtypes.VoidRequest) (*JobListResponse, error)
	Describe(context.Context, *JobDescribeRequest) (*JobDescribeResponse, error)
	Create(context.Context, *JobCreateRequest) (*dtypes.VoidResponse, error)
	Copy(context.Context, *JobCopyRequest) (*dtypes.VoidResponse, error)
	Build(context.Context, *JobBuildRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *JobDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterJobsServer(s *grpc.Server, srv JobsServer) {
	s.RegisterService(&_Jobs_serviceDesc, srv)
}

func _Jobs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).List(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Describe(ctx, req.(*JobDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Create(ctx, req.(*JobCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Copy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Copy(ctx, req.(*JobCopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Build(ctx, req.(*JobBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Jobs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Delete(ctx, req.(*JobDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Jobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.Jobs",
	HandlerType: (*JobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Jobs_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Jobs_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Jobs_Create_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _Jobs_Copy_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _Jobs_Build_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Jobs_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0x36, 0x4b, 0xb7, 0x53, 0x28, 0xe0, 0x56, 0x5d, 0x08, 0xff, 0x4a, 0x26, 0xa6,
	0x32, 0x89, 0x04, 0xb6, 0x3b, 0xb8, 0x40, 0xa2, 0x88, 0x8b, 0x81, 0xb8, 0x08, 0x52, 0x25, 0x2e,
	0xa6, 0x2a, 0x4d, 0xcd, 0x66, 0x94, 0xc6, 0x21, 0x4e, 0x91, 0xaa, 0x69, 0x37, 0xbc, 0x02, 0x8f,
	0xc1, 0xe3, 0xf0, 0x0a, 0xbc, 0x05, 0x37, 0xd8, 0xc7, 0x0e, 0xa4, 0x8c, 0x54, 0xe2, 0x66, 0x9a,
	0x3f, 0x9f, 0xf3, 0x73, 0xfc, 0x9d, 0xcf, 0x85, 0x9d, 0x8f, 0x7c, 0xe6, 0x67, 0x39, 0x2f, 0x38,
	0x69, 0xc6, 0xcc, 0xbd, 0x7d, 0xca, 0xf9, 0x69, 0x42, 0x83, 0x28, 0x63, 0x41, 0x94, 0xa6, 0xbc,
	0x88, 0x0a, 0xc6, 0x53, 0xa1, 0x2b, 0xdc, 0x81, 0x92, 0xe7, 0xc5, 0x2a, 0xa3, 0x22, 0xc0, 0xbf,
	0x5a, 0xf7, 0x26, 0x70, 0xed, 0x98, 0xcf, 0xde, 0x30, 0x51, 0x84, 0x54, 0x64, 0xb2, 0x9e, 0x92,
	0x7d, 0xb0, 0x85, 0x6c, 0x5e, 0x0a, 0xa7, 0x31, 0x6c, 0x8c, 0x3a, 0x87, 0x5d, 0x5f, 0xf7, 0xf9,
	0xef, 0x50, 0x0d, 0xcd, 0x2e, 0xb9, 0x05, 0x96, 0xfc, 0x02, 0xe1, 0x34, 0x87, 0x2d, 0x59, 0xd5,
	0xf6, 0x63, 0xe6, 0x4b, 0x54, 0x88, 0xa2, 0xf7, 0x0c, 0xb9, 0x2f, 0x96, 0x2c, 0x99, 0x87, 0xf4,
	0xd3, 0x92, 0x8a, 0x82, 0x10, 0xb0, 0xd2, 0x68, 0x41, 0x91, 0xba, 0x13, 0xe2, 0xff, 0xa4, 0x0f,
	0x5b, 0x59, 0x94, 0x47, 0x0b, 0x09, 0x51, 0xa2, 0x5e, 0x78, 0x23, 0x20, 0xb2, 0xf9, 0x25, 0x15,
	0x71, 0xce, 0x66, 0x74, 0x43, 0xbf, 0xf7, 0xb3, 0x01, 0xbd, 0xb5, 0xd2, 0xff, 0xbc, 0x43, 0xc9,
	0x6c, 0x56, 0xbe, 0x69, 0x08, 0x9d, 0x39, 0xf2, 0x32, 0x65, 0xa0, 0xd3, 0xc2, 0xad, 0xaa, 0x44,
	0xee, 0xc3, 0x95, 0x39, 0x13, 0x59, 0x12, 0xad, 0xa6, 0xd8, 0x6d, 0x99, 0x12, 0xad, 0xbd, 0x55,
	0x90, 0x7b, 0xd0, 0x99, 0xa9, 0xcb, 0x4f, 0x63, 0xbe, 0x4c, 0x0b, 0x67, 0x4b, 0x56, 0xb4, 0x42,
	0x40, 0x69, 0xac, 0x14, 0xe9, 0x9e, 0x9a, 0x9f, 0xdc, 0x4e, 0x78, 0xee, 0xd8, 0x08, 0xd8, 0x96,
	0xc2, 0x58, 0xad, 0xc9, 0x03, 0xe8, 0x62, 0xa9, 0x90, 0xfb, 0x8b, 0x45, 0x94, 0xce, 0x9d, 0x36,
	0x56, 0x5c, 0xd5, 0xea, 0x58, 0x8b, 0xde, 0x3e, 0x5c, 0xc7, 0xcb, 0x27, 0xb4, 0xd8, 0xe8, 0xd2,
	0x73, 0xac, 0x1b, 0xe7, 0x34, 0xda, 0x58, 0x47, 0x76, 0xa1, 0x2d, 0xce, 0xa6, 0x1f, 0x58, 0x52,
	0x1a, 0x62, 0x8b, 0xb3, 0x57, 0x72, 0xe5, 0x1d, 0x43, 0x57, 0x01, 0x78, 0xb6, 0x2a, 0xdb, 0x07,
	0xd2, 0x60, 0xbe, 0xcc, 0xe3, 0x12, 0x60, 0x56, 0xc6, 0xbc, 0x82, 0xa5, 0x98, 0x3e, 0x83, 0xa9,
	0x4a, 0x5e, 0x00, 0x2d, 0xc9, 0xaa, 0x4b, 0x83, 0xf6, 0xc3, 0xa4, 0x01, 0x17, 0x87, 0xdf, 0x2c,
	0xb0, 0x64, 0x87, 0x20, 0xaf, 0xc1, 0x52, 0x41, 0x25, 0xbd, 0x72, 0x98, 0x13, 0xce, 0xca, 0x74,
	0xb9, 0x3d, 0x93, 0xbf, 0x6a, 0x94, 0xbd, 0x9b, 0x5f, 0xbe, 0xff, 0xf8, 0xda, 0xec, 0x91, 0x1b,
	0xf8, 0x2a, 0x62, 0x16, 0x7c, 0x7e, 0xec, 0x3f, 0x09, 0x54, 0x40, 0xc9, 0x09, 0x6c, 0x97, 0xa9,
	0x21, 0x03, 0xd3, 0xfb, 0x57, 0xe2, 0xdc, 0xdd, 0x4b, 0xba, 0xe1, 0x0e, 0x91, 0xeb, 0x12, 0xe7,
	0x12, 0x37, 0x38, 0x57, 0x37, 0xb9, 0x20, 0xef, 0xc1, 0xd6, 0x7e, 0x93, 0xbe, 0x81, 0xac, 0xd9,
	0xef, 0xf6, 0xd7, 0xef, 0x60, 0xb8, 0x7b, 0xc8, 0xbd, 0xe3, 0xd6, 0x72, 0x9f, 0x36, 0x0e, 0x08,
	0x05, 0x4b, 0x4d, 0x82, 0x90, 0x12, 0xfc, 0x67, 0x2c, 0x35, 0xd8, 0x23, 0xc4, 0x3e, 0xf2, 0x46,
	0xff, 0xc0, 0xea, 0xb9, 0x5d, 0x04, 0xe7, 0x95, 0x19, 0xe1, 0x31, 0x27, 0xb0, 0x85, 0xcf, 0x97,
	0x94, 0xce, 0x56, 0x1f, 0x73, 0xcd, 0x41, 0x0f, 0xf1, 0xa0, 0x3d, 0xf7, 0x6e, 0xdd, 0xf7, 0x07,
	0x18, 0x60, 0x85, 0x9f, 0x80, 0xad, 0x83, 0xfb, 0xdb, 0xa0, 0xb5, 0x1c, 0xd7, 0x1c, 0x60, 0x8c,
	0x3f, 0xa8, 0x35, 0x68, 0x66, 0xe3, 0xef, 0xda, 0xd1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e,
	0xad, 0xd3, 0xca, 0x1e, 0x05, 0x00, 0x00,
}
