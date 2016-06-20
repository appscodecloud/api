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

func init() { proto.RegisterFile("job.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x55, 0x12, 0xc7, 0x69, 0x6f, 0xbe, 0x2f, 0xa0, 0x49, 0x94, 0x5a, 0xa6, 0x12, 0xa9, 0x11,
	0xa1, 0x54, 0x95, 0x4d, 0x83, 0xc4, 0x02, 0x90, 0x90, 0x08, 0x62, 0x51, 0x21, 0x16, 0x06, 0x75,
	0x1b, 0xf9, 0x2f, 0xed, 0x20, 0xc7, 0x63, 0x3c, 0x0e, 0x52, 0x54, 0x75, 0xc3, 0x2b, 0xf0, 0x68,
	0x3c, 0x01, 0x12, 0x6f, 0xc1, 0x86, 0x99, 0x3b, 0x63, 0x70, 0x40, 0x49, 0xc4, 0xa6, 0xea, 0x9c,
	0xb9, 0xf7, 0x1c, 0xdf, 0x73, 0xcf, 0x04, 0xf6, 0x3f, 0xb0, 0xd0, 0xcd, 0x0b, 0x56, 0x32, 0xd2,
	0x8c, 0xa8, 0x7d, 0x78, 0xc9, 0xd8, 0x65, 0x9a, 0x78, 0x41, 0x4e, 0xbd, 0x20, 0xcb, 0x58, 0x19,
	0x94, 0x94, 0x65, 0x5c, 0x55, 0xd8, 0x43, 0x09, 0xc7, 0xe5, 0x2a, 0x4f, 0xb8, 0x87, 0x7f, 0x15,
	0xee, 0x5c, 0xc0, 0xad, 0x73, 0x16, 0xbe, 0xa1, 0xbc, 0xf4, 0x13, 0x9e, 0x8b, 0xfa, 0x84, 0x8c,
	0xc1, 0xe4, 0xa2, 0x79, 0xc9, 0xad, 0xc6, 0xa8, 0x71, 0xdc, 0x9d, 0xf4, 0x5c, 0xd5, 0xe7, 0xbe,
	0x43, 0xd4, 0xd7, 0xb7, 0xe4, 0x0e, 0x18, 0xe2, 0x0b, 0xb8, 0xd5, 0x1c, 0xb5, 0x44, 0x55, 0xc7,
	0x8d, 0xa8, 0x2b, 0xa8, 0x7c, 0x04, 0x9d, 0x67, 0xc8, 0xfb, 0x72, 0x49, 0xd3, 0xd8, 0x4f, 0x3e,
	0x2e, 0x13, 0x5e, 0x12, 0x02, 0x46, 0x16, 0x2c, 0x12, 0x64, 0xdd, 0xf7, 0xf1, 0x7f, 0x32, 0x80,
	0x76, 0x1e, 0x14, 0xc1, 0x42, 0x90, 0x48, 0x50, 0x1d, 0x9c, 0x63, 0x20, 0xa2, 0xf9, 0x55, 0xc2,
	0xa3, 0x82, 0x86, 0xc9, 0x96, 0x7e, 0xe7, 0x47, 0x03, 0xfa, 0x6b, 0xa5, 0xff, 0x38, 0x43, 0xc5,
	0xd9, 0xac, 0x7d, 0xd3, 0x08, 0xba, 0x31, 0xf2, 0xe5, 0xd2, 0x40, 0xab, 0x85, 0x57, 0x75, 0x88,
	0x1c, 0xc1, 0x7f, 0x31, 0xe5, 0x79, 0x1a, 0xac, 0x66, 0xd8, 0x6d, 0xe8, 0x12, 0x85, 0xbd, 0x95,
	0x24, 0x77, 0xa1, 0x1b, 0xca, 0xe1, 0x67, 0x11, 0x5b, 0x66, 0xa5, 0xd5, 0x16, 0x15, 0x2d, 0x1f,
	0x10, 0x9a, 0x4a, 0x44, 0xb8, 0x27, 0xf7, 0x27, 0xae, 0x53, 0x56, 0x58, 0x26, 0x12, 0xec, 0x09,
	0x60, 0x2a, 0xcf, 0xe4, 0x3e, 0xf4, 0xb0, 0x94, 0x8b, 0xfb, 0xc5, 0x22, 0xc8, 0x62, 0xab, 0x83,
	0x15, 0xff, 0x2b, 0x74, 0xaa, 0x40, 0x67, 0x0c, 0xb7, 0x71, 0xf8, 0x34, 0x29, 0xb7, 0xba, 0xf4,
	0x02, 0xeb, 0xa6, 0x45, 0x12, 0x6c, 0xad, 0x23, 0x07, 0xd0, 0xe1, 0x57, 0xb3, 0x39, 0x4d, 0x2b,
	0x43, 0x4c, 0x7e, 0xf5, 0x5a, 0x9c, 0x9c, 0x73, 0xe8, 0x49, 0x02, 0x96, 0xaf, 0xaa, 0xf6, 0xa1,
	0x30, 0x98, 0x2d, 0x8b, 0xa8, 0x22, 0xd0, 0x27, 0x6d, 0x5e, 0x49, 0x33, 0x4c, 0x9f, 0xa6, 0xa9,
	0x43, 0x8e, 0x07, 0x2d, 0xc1, 0xb5, 0x29, 0x0d, 0xca, 0x0f, 0x9d, 0x06, 0x3c, 0x4c, 0xbe, 0x19,
	0x60, 0x88, 0x0e, 0x4e, 0xde, 0x83, 0x21, 0x83, 0x4a, 0xfa, 0xd5, 0x32, 0x2f, 0x18, 0xad, 0xd2,
	0x65, 0xf7, 0x75, 0xfe, 0xea, 0x51, 0x76, 0x9c, 0xcf, 0x5f, 0xbf, 0x7f, 0x69, 0x1e, 0x12, 0x5b,
	0xbc, 0x8a, 0x9c, 0x47, 0x2c, 0x56, 0xcf, 0x23, 0xa2, 0xde, 0xa7, 0x47, 0xee, 0x99, 0x27, 0x93,
	0x4a, 0xe6, 0xb0, 0x57, 0xc5, 0x87, 0x0c, 0x35, 0xc9, 0x1f, 0xd1, 0xb3, 0x0f, 0xfe, 0xc2, 0xb5,
	0xc0, 0x43, 0x14, 0xb8, 0x47, 0x8e, 0x36, 0x0b, 0x78, 0xd7, 0x72, 0xb6, 0x1b, 0x12, 0x82, 0xa9,
	0x36, 0x40, 0x06, 0x9a, 0x6d, 0x6d, 0x21, 0xf6, 0x60, 0x7d, 0x2a, 0x2d, 0x70, 0x8a, 0x02, 0x63,
	0x7b, 0xb7, 0xc0, 0xd3, 0xc6, 0x09, 0xc9, 0xc0, 0x90, 0x4b, 0x22, 0xa4, 0x52, 0xf8, 0xbd, 0xb1,
	0x0d, 0xfc, 0xcf, 0x91, 0xff, 0x89, 0x73, 0xb6, 0x8d, 0x5f, 0xed, 0xf6, 0xc6, 0xbb, 0xae, 0xed,
	0x11, 0xf5, 0xe6, 0xd0, 0xc6, 0x27, 0x4e, 0x2a, 0xf7, 0xeb, 0x0f, 0x7e, 0x83, 0xe2, 0x04, 0x15,
	0x4f, 0xed, 0x07, 0x3b, 0x27, 0xf2, 0x30, 0xed, 0x52, 0x67, 0x06, 0xa6, 0x4a, 0xf9, 0x2f, 0xef,
	0xd6, 0x42, 0xbf, 0x41, 0x49, 0x2f, 0xe7, 0x64, 0xb7, 0x77, 0xa1, 0x89, 0xbf, 0x86, 0x8f, 0x7f,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x09, 0xd8, 0x49, 0x54, 0x05, 0x00, 0x00,
}
