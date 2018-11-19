// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

type ProjectListRequest struct {
	WithMember           bool     `protobuf:"varint,1,opt,name=with_member,json=withMember,proto3" json:"with_member,omitempty"`
	Members              []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectListRequest) Reset()         { *m = ProjectListRequest{} }
func (m *ProjectListRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectListRequest) ProtoMessage()    {}
func (*ProjectListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{0}
}

func (m *ProjectListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectListRequest.Unmarshal(m, b)
}
func (m *ProjectListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectListRequest.Marshal(b, m, deterministic)
}
func (m *ProjectListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectListRequest.Merge(m, src)
}
func (m *ProjectListRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectListRequest.Size(m)
}
func (m *ProjectListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectListRequest proto.InternalMessageInfo

func (m *ProjectListRequest) GetWithMember() bool {
	if m != nil {
		return m.WithMember
	}
	return false
}

func (m *ProjectListRequest) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ProjectListResponse struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProjectListResponse) Reset()         { *m = ProjectListResponse{} }
func (m *ProjectListResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectListResponse) ProtoMessage()    {}
func (*ProjectListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{1}
}

func (m *ProjectListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectListResponse.Unmarshal(m, b)
}
func (m *ProjectListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectListResponse.Marshal(b, m, deterministic)
}
func (m *ProjectListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectListResponse.Merge(m, src)
}
func (m *ProjectListResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectListResponse.Size(m)
}
func (m *ProjectListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectListResponse proto.InternalMessageInfo

func (m *ProjectListResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ProjectMemberListRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectMemberListRequest) Reset()         { *m = ProjectMemberListRequest{} }
func (m *ProjectMemberListRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectMemberListRequest) ProtoMessage()    {}
func (*ProjectMemberListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{2}
}

func (m *ProjectMemberListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectMemberListRequest.Unmarshal(m, b)
}
func (m *ProjectMemberListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectMemberListRequest.Marshal(b, m, deterministic)
}
func (m *ProjectMemberListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectMemberListRequest.Merge(m, src)
}
func (m *ProjectMemberListRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectMemberListRequest.Size(m)
}
func (m *ProjectMemberListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectMemberListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectMemberListRequest proto.InternalMessageInfo

func (m *ProjectMemberListRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ProjectMemberListResponse struct {
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectMemberListResponse) Reset()         { *m = ProjectMemberListResponse{} }
func (m *ProjectMemberListResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectMemberListResponse) ProtoMessage()    {}
func (*ProjectMemberListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{3}
}

func (m *ProjectMemberListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectMemberListResponse.Unmarshal(m, b)
}
func (m *ProjectMemberListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectMemberListResponse.Marshal(b, m, deterministic)
}
func (m *ProjectMemberListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectMemberListResponse.Merge(m, src)
}
func (m *ProjectMemberListResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectMemberListResponse.Size(m)
}
func (m *ProjectMemberListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectMemberListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectMemberListResponse proto.InternalMessageInfo

func (m *ProjectMemberListResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type Project struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phid                 string    `protobuf:"bytes,2,opt,name=phid,proto3" json:"phid,omitempty"`
	Type                 string    `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Status               string    `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	ViewPolicy           string    `protobuf:"bytes,5,opt,name=view_policy,json=viewPolicy,proto3" json:"view_policy,omitempty"`
	EditPolicy           string    `protobuf:"bytes,6,opt,name=edit_policy,json=editPolicy,proto3" json:"edit_policy,omitempty"`
	JoinPolicy           string    `protobuf:"bytes,7,opt,name=join_policy,json=joinPolicy,proto3" json:"join_policy,omitempty"`
	MembershipLocked     bool      `protobuf:"varint,8,opt,name=membership_locked,json=membershipLocked,proto3" json:"membership_locked,omitempty"`
	HasSubprojects       bool      `protobuf:"varint,9,opt,name=has_subprojects,json=hasSubprojects,proto3" json:"has_subprojects,omitempty"`
	Members              []*Member `protobuf:"bytes,10,rep,name=members,proto3" json:"members,omitempty"`
	CreatedAt            int64     `protobuf:"varint,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{4}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Project) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Project) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Project) GetViewPolicy() string {
	if m != nil {
		return m.ViewPolicy
	}
	return ""
}

func (m *Project) GetEditPolicy() string {
	if m != nil {
		return m.EditPolicy
	}
	return ""
}

func (m *Project) GetJoinPolicy() string {
	if m != nil {
		return m.JoinPolicy
	}
	return ""
}

func (m *Project) GetMembershipLocked() bool {
	if m != nil {
		return m.MembershipLocked
	}
	return false
}

func (m *Project) GetHasSubprojects() bool {
	if m != nil {
		return m.HasSubprojects
	}
	return false
}

func (m *Project) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Project) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Member struct {
	Phid                 string   `protobuf:"bytes,1,opt,name=phid,proto3" json:"phid,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	IsAdmin              bool     `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{5}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Member) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Member) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *Member) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func init() {
	proto.RegisterType((*ProjectListRequest)(nil), "appscode.auth.v1alpha1.ProjectListRequest")
	proto.RegisterType((*ProjectListResponse)(nil), "appscode.auth.v1alpha1.ProjectListResponse")
	proto.RegisterType((*ProjectMemberListRequest)(nil), "appscode.auth.v1alpha1.ProjectMemberListRequest")
	proto.RegisterType((*ProjectMemberListResponse)(nil), "appscode.auth.v1alpha1.ProjectMemberListResponse")
	proto.RegisterType((*Project)(nil), "appscode.auth.v1alpha1.Project")
	proto.RegisterType((*Member)(nil), "appscode.auth.v1alpha1.Member")
}

func init() { proto.RegisterFile("project.proto", fileDescriptor_8340e6318dfdfac2) }

var fileDescriptor_8340e6318dfdfac2 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x52, 0xd4, 0x4e,
	0x10, 0xc7, 0x2b, 0x59, 0x7e, 0x9b, 0xa4, 0xf9, 0xa9, 0x38, 0x56, 0x51, 0xc3, 0xa2, 0xb0, 0x95,
	0x8b, 0x5b, 0x40, 0x6d, 0x04, 0x2f, 0x88, 0x27, 0x38, 0x78, 0x02, 0xdd, 0x8a, 0x55, 0x1e, 0xbc,
	0xa4, 0x86, 0x64, 0x8a, 0x0c, 0xee, 0x66, 0x86, 0xcc, 0x04, 0x8a, 0xb2, 0xbc, 0xf0, 0x0a, 0x9e,
	0x7c, 0x10, 0x7d, 0x11, 0x1f, 0xc0, 0x8b, 0x77, 0x5f, 0xc1, 0x9a, 0x3f, 0x61, 0x97, 0x12, 0x8b,
	0xf5, 0x36, 0xf9, 0xf4, 0xb7, 0xd3, 0xdf, 0xe9, 0xee, 0x04, 0xee, 0x89, 0x9a, 0x9f, 0xd2, 0x5c,
	0x0d, 0x45, 0xcd, 0x15, 0x47, 0xcb, 0x44, 0x08, 0x99, 0xf3, 0x82, 0x0e, 0x49, 0xa3, 0xca, 0xe1,
	0xf9, 0x36, 0x19, 0x8b, 0x92, 0x6c, 0xf7, 0x1e, 0x9f, 0x70, 0x7e, 0x32, 0xa6, 0x09, 0x11, 0x2c,
	0x21, 0x55, 0xc5, 0x15, 0x51, 0x8c, 0x57, 0xd2, 0x66, 0xf5, 0xd6, 0xda, 0xac, 0xdb, 0xe3, 0xf1,
	0x1b, 0x40, 0x23, 0x5b, 0xe6, 0x90, 0x49, 0x95, 0xd2, 0xb3, 0x86, 0x4a, 0x85, 0xd6, 0x61, 0xf1,
	0x82, 0xa9, 0x32, 0x9b, 0xd0, 0xc9, 0x31, 0xad, 0xb1, 0xd7, 0xf7, 0x06, 0x61, 0x0a, 0x1a, 0x1d,
	0x19, 0x82, 0x30, 0x04, 0x36, 0x26, 0xb1, 0xdf, 0xef, 0x0c, 0xa2, 0xb4, 0x7d, 0x8c, 0x53, 0x78,
	0x74, 0xe3, 0x85, 0x52, 0xf0, 0x4a, 0x52, 0xf4, 0x12, 0x42, 0x77, 0x1d, 0x89, 0xbd, 0x7e, 0x67,
	0xb0, 0xb8, 0xb3, 0x3e, 0xbc, 0xfd, 0x42, 0x43, 0x97, 0x9e, 0x5e, 0x27, 0xc4, 0x5b, 0x80, 0x1d,
	0xb4, 0xe5, 0x67, 0xad, 0x2e, 0x41, 0xa7, 0x61, 0x85, 0xb1, 0x18, 0xa5, 0xfa, 0x18, 0xbf, 0x83,
	0x95, 0x5b, 0xd4, 0xce, 0xc7, 0x0b, 0x08, 0xdc, 0x6b, 0x4d, 0xca, 0x1c, 0x36, 0x5a, 0x7d, 0xfc,
	0xcb, 0x87, 0xc0, 0x41, 0x84, 0x60, 0xa1, 0x22, 0x13, 0xea, 0xca, 0x9a, 0xb3, 0x66, 0xa2, 0x64,
	0x05, 0xf6, 0x2d, 0xd3, 0x67, 0xcd, 0xd4, 0xa5, 0xa0, 0xb8, 0x63, 0x99, 0x3e, 0xa3, 0x65, 0xe8,
	0x4a, 0x45, 0x54, 0x23, 0xf1, 0x82, 0xa1, 0xee, 0x49, 0x37, 0xfd, 0x9c, 0xd1, 0x8b, 0x4c, 0xf0,
	0x31, 0xcb, 0x2f, 0xf1, 0x7f, 0x26, 0x08, 0x1a, 0x8d, 0x0c, 0xd1, 0x02, 0x5a, 0x30, 0xd5, 0x0a,
	0xba, 0x56, 0xa0, 0xd1, 0x54, 0x70, 0xca, 0x59, 0xd5, 0x0a, 0x02, 0x2b, 0xd0, 0xc8, 0x09, 0x36,
	0xe1, 0xa1, 0x9b, 0x53, 0xc9, 0x44, 0x36, 0xe6, 0xf9, 0x07, 0x5a, 0xe0, 0xd0, 0x4c, 0x77, 0x69,
	0x1a, 0x38, 0x34, 0x1c, 0x3d, 0x85, 0x07, 0x25, 0x91, 0x99, 0x6c, 0x8e, 0xaf, 0x27, 0x17, 0x19,
	0xe9, 0xfd, 0x92, 0xc8, 0xb7, 0x53, 0x8a, 0x76, 0xa7, 0xcb, 0x00, 0x66, 0xb4, 0x6b, 0x7f, 0xeb,
	0xa9, 0x1d, 0xc8, 0xf5, 0xb2, 0xa0, 0x27, 0x00, 0x79, 0x4d, 0x89, 0xa2, 0x45, 0x46, 0x14, 0x5e,
	0xec, 0x7b, 0x83, 0x4e, 0x1a, 0x39, 0xb2, 0xaf, 0xe2, 0x33, 0xe8, 0xba, 0x7d, 0x6b, 0x7b, 0xeb,
	0xcd, 0xf4, 0x76, 0x15, 0xa2, 0x46, 0xd2, 0x3a, 0x33, 0x83, 0xb0, 0x4d, 0x0f, 0x35, 0x78, 0xad,
	0x87, 0xb1, 0x0a, 0x51, 0x4d, 0xc9, 0xd8, 0x06, 0x6d, 0xf7, 0x43, 0x0d, 0x4c, 0x70, 0x05, 0x42,
	0x26, 0x33, 0x52, 0x4c, 0x58, 0x65, 0x66, 0x10, 0xa6, 0x01, 0x93, 0xfb, 0xfa, 0x71, 0xe7, 0x87,
	0x0f, 0xe1, 0xa8, 0xbd, 0xd8, 0x17, 0x0f, 0x16, 0xf4, 0xf6, 0xa0, 0x8d, 0x3b, 0x96, 0x64, 0x66,
	0x21, 0x7b, 0x9b, 0x73, 0x69, 0xed, 0x3a, 0xc6, 0xbb, 0x57, 0x5f, 0xb1, 0x1f, 0x7a, 0x57, 0xdf,
	0x7f, 0x7e, 0xf6, 0xb7, 0xd0, 0x46, 0x92, 0xdd, 0xfc, 0x5c, 0x1b, 0x55, 0x26, 0x6d, 0x7e, 0xd2,
	0xf6, 0x3b, 0x39, 0x95, 0xbc, 0x42, 0xdf, 0x3c, 0x08, 0x8e, 0x5c, 0x1b, 0x9f, 0xdd, 0x51, 0xf2,
	0x8f, 0xaf, 0xa6, 0xb7, 0xfd, 0x0f, 0x19, 0xce, 0xea, 0xab, 0x19, 0xab, 0x7b, 0x68, 0x77, 0x3e,
	0xab, 0x1f, 0x1b, 0x56, 0x7c, 0x4a, 0xdc, 0xb8, 0x8d, 0xf1, 0x83, 0x3d, 0x58, 0xcb, 0xf9, 0x64,
	0xa6, 0xbe, 0x60, 0x37, 0x3d, 0x1c, 0xfc, 0xef, 0x4c, 0x8c, 0xf4, 0x1f, 0x6a, 0xe4, 0xbd, 0x0f,
	0xdb, 0xc8, 0x71, 0xd7, 0xfc, 0xb4, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xda, 0xc8, 0xee,
	0xb8, 0x1b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectsClient interface {
	List(ctx context.Context, in *ProjectListRequest, opts ...grpc.CallOption) (*ProjectListResponse, error)
	Members(ctx context.Context, in *ProjectMemberListRequest, opts ...grpc.CallOption) (*ProjectMemberListResponse, error)
}

type projectsClient struct {
	cc *grpc.ClientConn
}

func NewProjectsClient(cc *grpc.ClientConn) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) List(ctx context.Context, in *ProjectListRequest, opts ...grpc.CallOption) (*ProjectListResponse, error) {
	out := new(ProjectListResponse)
	err := c.cc.Invoke(ctx, "/appscode.auth.v1alpha1.Projects/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) Members(ctx context.Context, in *ProjectMemberListRequest, opts ...grpc.CallOption) (*ProjectMemberListResponse, error) {
	out := new(ProjectMemberListResponse)
	err := c.cc.Invoke(ctx, "/appscode.auth.v1alpha1.Projects/Members", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
type ProjectsServer interface {
	List(context.Context, *ProjectListRequest) (*ProjectListResponse, error)
	Members(context.Context, *ProjectMemberListRequest) (*ProjectMemberListResponse, error)
}

func RegisterProjectsServer(s *grpc.Server, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1alpha1.Projects/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).List(ctx, req.(*ProjectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_Members_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectMemberListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).Members(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1alpha1.Projects/Members",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).Members(ctx, req.(*ProjectMemberListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.auth.v1alpha1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Projects_List_Handler,
		},
		{
			MethodName: "Members",
			Handler:    _Projects_Members_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}
