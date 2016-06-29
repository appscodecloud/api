// Code generated by protoc-gen-go.
// source: clusters.proto
// DO NOT EDIT!

package kubernetes

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

type Cluster struct {
	Phid      string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Provider  string            `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	Zone      string            `protobuf:"bytes,4,opt,name=zone" json:"zone,omitempty"`
	Region    string            `protobuf:"bytes,5,opt,name=region" json:"region,omitempty"`
	Os        string            `protobuf:"bytes,6,opt,name=os" json:"os,omitempty"`
	Links     map[string]string `protobuf:"bytes,7,rep,name=links" json:"links,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nodes     int32             `protobuf:"varint,8,opt,name=nodes" json:"nodes,omitempty"`
	CreatedAt string            `protobuf:"bytes,9,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Cluster) GetLinks() map[string]string {
	if m != nil {
		return m.Links
	}
	return nil
}

type ClusterDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDescribeRequest) Reset()                    { *m = ClusterDescribeRequest{} }
func (m *ClusterDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeRequest) ProtoMessage()               {}
func (*ClusterDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type ClusterDescribeResponse struct {
	Status  *dtypes.Status                 `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Cluster *Cluster                       `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Specs   *ClusterDescribeResponse_Specs `protobuf:"bytes,3,opt,name=specs" json:"specs,omitempty"`
}

func (m *ClusterDescribeResponse) Reset()                    { *m = ClusterDescribeResponse{} }
func (m *ClusterDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse) ProtoMessage()               {}
func (*ClusterDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ClusterDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterDescribeResponse) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClusterDescribeResponse) GetSpecs() *ClusterDescribeResponse_Specs {
	if m != nil {
		return m.Specs
	}
	return nil
}

type ClusterDescribeResponse_Specs struct {
	Containers  int32 `protobuf:"varint,1,opt,name=containers" json:"containers,omitempty"`
	Pods        int32 `protobuf:"varint,2,opt,name=pods" json:"pods,omitempty"`
	Services    int32 `protobuf:"varint,3,opt,name=services" json:"services,omitempty"`
	Rcs         int32 `protobuf:"varint,4,opt,name=rcs" json:"rcs,omitempty"`
	CpuCore     int64 `protobuf:"varint,5,opt,name=cpu_core" json:"cpu_core,omitempty"`
	TotalMemory int64 `protobuf:"varint,6,opt,name=total_memory" json:"total_memory,omitempty"`
}

func (m *ClusterDescribeResponse_Specs) Reset()         { *m = ClusterDescribeResponse_Specs{} }
func (m *ClusterDescribeResponse_Specs) String() string { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse_Specs) ProtoMessage()    {}
func (*ClusterDescribeResponse_Specs) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{2, 0}
}

type ClusterListResponse struct {
	Status   *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Clusters []*Cluster     `protobuf:"bytes,2,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ClusterListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type ClusterCreateRequest struct {
	Name                string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider            string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Zone                string            `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	CloudCredential     string            `protobuf:"bytes,4,opt,name=cloud_credential" json:"cloud_credential,omitempty"`
	CloudCredentialData map[string]string `protobuf:"bytes,5,rep,name=cloud_credential_data" json:"cloud_credential_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeSet             map[string]int64  `protobuf:"bytes,6,rep,name=node_set" json:"node_set,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	SaltbaseVersion     string            `protobuf:"bytes,7,opt,name=saltbase_version" json:"saltbase_version,omitempty"`
	KubeStarterVersion  string            `protobuf:"bytes,8,opt,name=kube_starter_version" json:"kube_starter_version,omitempty"`
	KubeVersion         string            `protobuf:"bytes,9,opt,name=kube_version" json:"kube_version,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ClusterCreateRequest) GetCloudCredentialData() map[string]string {
	if m != nil {
		return m.CloudCredentialData
	}
	return nil
}

func (m *ClusterCreateRequest) GetNodeSet() map[string]int64 {
	if m != nil {
		return m.NodeSet
	}
	return nil
}

type ClusterScaleRequest struct {
	Name        string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NodeChanges map[string]int64 `protobuf:"bytes,2,rep,name=node_changes" json:"node_changes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ClusterScaleRequest) Reset()                    { *m = ClusterScaleRequest{} }
func (m *ClusterScaleRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterScaleRequest) ProtoMessage()               {}
func (*ClusterScaleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ClusterScaleRequest) GetNodeChanges() map[string]int64 {
	if m != nil {
		return m.NodeChanges
	}
	return nil
}

type ClusterDeleteRequest struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ReleaseReservedIp bool   `protobuf:"varint,2,opt,name=release_reserved_ip" json:"release_reserved_ip,omitempty"`
}

func (m *ClusterDeleteRequest) Reset()                    { *m = ClusterDeleteRequest{} }
func (m *ClusterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDeleteRequest) ProtoMessage()               {}
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type ClusterStartupScriptRequest struct {
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *ClusterStartupScriptRequest) Reset()                    { *m = ClusterStartupScriptRequest{} }
func (m *ClusterStartupScriptRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterStartupScriptRequest) ProtoMessage()               {}
func (*ClusterStartupScriptRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type ClusterStartupScriptResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Configuration string         `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ClusterStartupScriptResponse) Reset()                    { *m = ClusterStartupScriptResponse{} }
func (m *ClusterStartupScriptResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterStartupScriptResponse) ProtoMessage()               {}
func (*ClusterStartupScriptResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ClusterStartupScriptResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterClientConfigRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterClientConfigRequest) Reset()                    { *m = ClusterClientConfigRequest{} }
func (m *ClusterClientConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterClientConfigRequest) ProtoMessage()               {}
func (*ClusterClientConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type ClusterClientConfigResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Configuration string         `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ClusterClientConfigResponse) Reset()                    { *m = ClusterClientConfigResponse{} }
func (m *ClusterClientConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterClientConfigResponse) ProtoMessage()               {}
func (*ClusterClientConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *ClusterClientConfigResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterClientContainerRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DiskName string `protobuf:"bytes,2,opt,name=disk_name" json:"disk_name,omitempty"`
}

func (m *ClusterClientContainerRequest) Reset()                    { *m = ClusterClientContainerRequest{} }
func (m *ClusterClientContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterClientContainerRequest) ProtoMessage()               {}
func (*ClusterClientContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type ClusterInstanceListRequest struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name" json:"cluster_name,omitempty"`
}

func (m *ClusterInstanceListRequest) Reset()                    { *m = ClusterInstanceListRequest{} }
func (m *ClusterInstanceListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterInstanceListRequest) ProtoMessage()               {}
func (*ClusterInstanceListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

type ClusterInstance struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	ExternalId string `protobuf:"bytes,2,opt,name=external_id" json:"external_id,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	ExternalIp string `protobuf:"bytes,4,opt,name=external_ip" json:"external_ip,omitempty"`
	InternalIp string `protobuf:"bytes,5,opt,name=internal_ip" json:"internal_ip,omitempty"`
	Role       string `protobuf:"bytes,6,opt,name=role" json:"role,omitempty"`
	Sku        string `protobuf:"bytes,7,opt,name=sku" json:"sku,omitempty"`
}

func (m *ClusterInstance) Reset()                    { *m = ClusterInstance{} }
func (m *ClusterInstance) String() string            { return proto.CompactTextString(m) }
func (*ClusterInstance) ProtoMessage()               {}
func (*ClusterInstance) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

type ClusterInstanceListResponse struct {
	Status    *dtypes.Status     `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Instances []*ClusterInstance `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
}

func (m *ClusterInstanceListResponse) Reset()                    { *m = ClusterInstanceListResponse{} }
func (m *ClusterInstanceListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterInstanceListResponse) ProtoMessage()               {}
func (*ClusterInstanceListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *ClusterInstanceListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterInstanceListResponse) GetInstances() []*ClusterInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type ClusterUpdateRequest struct {
	Name               string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SaltbaseVersion    string `protobuf:"bytes,2,opt,name=saltbase_version" json:"saltbase_version,omitempty"`
	KubeStarterVersion string `protobuf:"bytes,3,opt,name=kube_starter_version" json:"kube_starter_version,omitempty"`
	KubeVersion        string `protobuf:"bytes,4,opt,name=kube_version" json:"kube_version,omitempty"`
}

func (m *ClusterUpdateRequest) Reset()                    { *m = ClusterUpdateRequest{} }
func (m *ClusterUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterUpdateRequest) ProtoMessage()               {}
func (*ClusterUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func init() {
	proto.RegisterType((*Cluster)(nil), "kubernetes.Cluster")
	proto.RegisterType((*ClusterDescribeRequest)(nil), "kubernetes.ClusterDescribeRequest")
	proto.RegisterType((*ClusterDescribeResponse)(nil), "kubernetes.ClusterDescribeResponse")
	proto.RegisterType((*ClusterDescribeResponse_Specs)(nil), "kubernetes.ClusterDescribeResponse.Specs")
	proto.RegisterType((*ClusterListResponse)(nil), "kubernetes.ClusterListResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "kubernetes.ClusterCreateRequest")
	proto.RegisterType((*ClusterScaleRequest)(nil), "kubernetes.ClusterScaleRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "kubernetes.ClusterDeleteRequest")
	proto.RegisterType((*ClusterStartupScriptRequest)(nil), "kubernetes.ClusterStartupScriptRequest")
	proto.RegisterType((*ClusterStartupScriptResponse)(nil), "kubernetes.ClusterStartupScriptResponse")
	proto.RegisterType((*ClusterClientConfigRequest)(nil), "kubernetes.ClusterClientConfigRequest")
	proto.RegisterType((*ClusterClientConfigResponse)(nil), "kubernetes.ClusterClientConfigResponse")
	proto.RegisterType((*ClusterClientContainerRequest)(nil), "kubernetes.ClusterClientContainerRequest")
	proto.RegisterType((*ClusterInstanceListRequest)(nil), "kubernetes.ClusterInstanceListRequest")
	proto.RegisterType((*ClusterInstance)(nil), "kubernetes.ClusterInstance")
	proto.RegisterType((*ClusterInstanceListResponse)(nil), "kubernetes.ClusterInstanceListResponse")
	proto.RegisterType((*ClusterUpdateRequest)(nil), "kubernetes.ClusterUpdateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Clusters service

type ClustersClient interface {
	Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error)
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Scale(ctx context.Context, in *ClusterScaleRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Update(ctx context.Context, in *ClusterUpdateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	StartupScript(ctx context.Context, in *ClusterStartupScriptRequest, opts ...grpc.CallOption) (*ClusterStartupScriptResponse, error)
	ClientConfig(ctx context.Context, in *ClusterClientConfigRequest, opts ...grpc.CallOption) (*ClusterClientConfigResponse, error)
	Instances(ctx context.Context, in *ClusterInstanceListRequest, opts ...grpc.CallOption) (*ClusterInstanceListResponse, error)
}

type clustersClient struct {
	cc *grpc.ClientConn
}

func NewClustersClient(cc *grpc.ClientConn) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error) {
	out := new(ClusterDescribeResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Scale(ctx context.Context, in *ClusterScaleRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Scale", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Update(ctx context.Context, in *ClusterUpdateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) StartupScript(ctx context.Context, in *ClusterStartupScriptRequest, opts ...grpc.CallOption) (*ClusterStartupScriptResponse, error) {
	out := new(ClusterStartupScriptResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/StartupScript", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) ClientConfig(ctx context.Context, in *ClusterClientConfigRequest, opts ...grpc.CallOption) (*ClusterClientConfigResponse, error) {
	out := new(ClusterClientConfigResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/ClientConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Instances(ctx context.Context, in *ClusterInstanceListRequest, opts ...grpc.CallOption) (*ClusterInstanceListResponse, error) {
	out := new(ClusterInstanceListResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Instances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clusters service

type ClustersServer interface {
	Describe(context.Context, *ClusterDescribeRequest) (*ClusterDescribeResponse, error)
	Create(context.Context, *ClusterCreateRequest) (*dtypes.LongRunningResponse, error)
	Scale(context.Context, *ClusterScaleRequest) (*dtypes.LongRunningResponse, error)
	Delete(context.Context, *ClusterDeleteRequest) (*dtypes.LongRunningResponse, error)
	Update(context.Context, *ClusterUpdateRequest) (*dtypes.LongRunningResponse, error)
	List(context.Context, *dtypes.VoidRequest) (*ClusterListResponse, error)
	StartupScript(context.Context, *ClusterStartupScriptRequest) (*ClusterStartupScriptResponse, error)
	ClientConfig(context.Context, *ClusterClientConfigRequest) (*ClusterClientConfigResponse, error)
	Instances(context.Context, *ClusterInstanceListRequest) (*ClusterInstanceListResponse, error)
}

func RegisterClustersServer(s *grpc.Server, srv ClustersServer) {
	s.RegisterService(&_Clusters_serviceDesc, srv)
}

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Describe(ctx, req.(*ClusterDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Create(ctx, req.(*ClusterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Scale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/Scale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Scale(ctx, req.(*ClusterScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Delete(ctx, req.(*ClusterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Update(ctx, req.(*ClusterUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).List(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_StartupScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterStartupScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).StartupScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/StartupScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).StartupScript(ctx, req.(*ClusterStartupScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_ClientConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterClientConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).ClientConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/ClientConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).ClientConfig(ctx, req.(*ClusterClientConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Instances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterInstanceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Instances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetes.Clusters/Instances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Instances(ctx, req.(*ClusterInstanceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clusters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetes.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Clusters_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Clusters_Create_Handler,
		},
		{
			MethodName: "Scale",
			Handler:    _Clusters_Scale_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Clusters_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Clusters_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Clusters_List_Handler,
		},
		{
			MethodName: "StartupScript",
			Handler:    _Clusters_StartupScript_Handler,
		},
		{
			MethodName: "ClientConfig",
			Handler:    _Clusters_ClientConfig_Handler,
		},
		{
			MethodName: "Instances",
			Handler:    _Clusters_Instances_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("clusters.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x57, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0x96, 0x93, 0x26, 0x4d, 0x4e, 0xda, 0xfe, 0xfa, 0x9b, 0x76, 0x17, 0xcb, 0x2d, 0x4b, 0x65,
	0xa0, 0x5b, 0xba, 0xc4, 0xd9, 0x0d, 0x42, 0x5a, 0x8a, 0xd0, 0xb2, 0x4a, 0x01, 0x21, 0x55, 0x5c,
	0x6c, 0x59, 0xc4, 0x05, 0x52, 0xe4, 0xda, 0x43, 0xd6, 0xaa, 0x3b, 0x36, 0x1e, 0x3b, 0xda, 0xee,
	0xaa, 0x5c, 0x70, 0x83, 0x04, 0x97, 0x3c, 0x00, 0x37, 0x20, 0x1e, 0x80, 0x47, 0xe1, 0x09, 0x90,
	0x78, 0x04, 0x1e, 0x80, 0x33, 0x33, 0xb6, 0x9b, 0xa4, 0x93, 0xc6, 0x85, 0x9b, 0x55, 0x3c, 0x73,
	0xfe, 0x7c, 0xf3, 0x9d, 0x73, 0xbe, 0xd3, 0x85, 0x35, 0x2f, 0xcc, 0x78, 0x4a, 0x13, 0xee, 0xc4,
	0x49, 0x94, 0x46, 0x04, 0x4e, 0xb3, 0x13, 0x9a, 0x30, 0x9a, 0x52, 0x6e, 0x6d, 0x8f, 0xa2, 0x68,
	0x14, 0xd2, 0x9e, 0x1b, 0x07, 0x3d, 0x97, 0xb1, 0x28, 0x75, 0xd3, 0x20, 0x62, 0xb9, 0xa5, 0x75,
	0x5b, 0x1c, 0xfb, 0xe9, 0x79, 0x4c, 0x79, 0x4f, 0xfe, 0xab, 0xce, 0xed, 0x3f, 0x0d, 0x58, 0x1e,
	0xa8, 0xa0, 0x64, 0x05, 0x96, 0xe2, 0x67, 0x81, 0x6f, 0x1a, 0x3b, 0xc6, 0x5e, 0x5b, 0x7c, 0x31,
	0xf7, 0x8c, 0x9a, 0x35, 0xf9, 0xb5, 0x0e, 0x2d, 0x74, 0x18, 0x07, 0x3e, 0x4d, 0xcc, 0x7a, 0x71,
	0xff, 0x22, 0x62, 0xd4, 0x5c, 0x92, 0x5f, 0x6b, 0xd0, 0x4c, 0xe8, 0x08, 0x13, 0x9a, 0x0d, 0xf9,
	0x0d, 0x50, 0x8b, 0xb8, 0xd9, 0x94, 0xbf, 0xbb, 0xd0, 0x08, 0x03, 0x76, 0xca, 0xcd, 0xe5, 0x9d,
	0xfa, 0x5e, 0xa7, 0x7f, 0xc7, 0xb9, 0x44, 0xed, 0xe4, 0xb9, 0x9d, 0x23, 0x61, 0xf0, 0x11, 0x4b,
	0x93, 0x73, 0xb2, 0x0a, 0x0d, 0x16, 0xf9, 0x94, 0x9b, 0x2d, 0xf4, 0x6e, 0x10, 0x0c, 0xe5, 0x25,
	0xd4, 0x4d, 0xa9, 0x3f, 0x74, 0x53, 0xb3, 0x2d, 0x22, 0x5a, 0x6f, 0x03, 0x4c, 0x38, 0x74, 0xa0,
	0x7e, 0x4a, 0xcf, 0x73, 0xd8, 0xe8, 0x3d, 0x76, 0xc3, 0x2c, 0xc7, 0x7d, 0x50, 0x7b, 0x68, 0xd8,
	0xbb, 0x70, 0x3b, 0x4f, 0x73, 0x48, 0xb9, 0x97, 0x04, 0x27, 0xf4, 0x09, 0xfd, 0x26, 0xa3, 0x3c,
	0x2d, 0xdf, 0x28, 0x5d, 0xed, 0x1f, 0x6b, 0xf0, 0xca, 0x15, 0x43, 0x1e, 0x23, 0x89, 0x94, 0xdc,
	0x81, 0x26, 0x47, 0x46, 0x33, 0x2e, 0x6d, 0x3b, 0xfd, 0x35, 0x47, 0x91, 0xe9, 0x1c, 0xcb, 0x53,
	0xf2, 0x06, 0x2c, 0xe7, 0xb5, 0x91, 0x89, 0x3b, 0xfd, 0x0d, 0xcd, 0x2b, 0xc9, 0x43, 0x68, 0xf0,
	0x98, 0x7a, 0x5c, 0x52, 0xd8, 0xe9, 0xbf, 0xa5, 0xb1, 0x99, 0xcd, 0xec, 0x1c, 0x0b, 0x07, 0x2b,
	0x86, 0x86, 0xfc, 0x21, 0xe9, 0x88, 0x58, 0xea, 0x06, 0x0c, 0xdb, 0x40, 0x82, 0x69, 0xc8, 0xc2,
	0x45, 0x3e, 0x97, 0x99, 0x1b, 0xa2, 0x54, 0x9c, 0x26, 0xe3, 0xc0, 0xa3, 0x2a, 0x4f, 0x43, 0x10,
	0x94, 0x60, 0xd2, 0xa5, 0xe2, 0xda, 0x8b, 0xb3, 0xa1, 0x17, 0x25, 0x54, 0xd6, 0xaa, 0x4e, 0x36,
	0x61, 0x25, 0xc5, 0x76, 0x09, 0x87, 0x67, 0xf4, 0x2c, 0x4a, 0xce, 0x65, 0xd5, 0xea, 0xf6, 0x57,
	0xb0, 0x91, 0x43, 0x3a, 0x0a, 0x78, 0x5a, 0x99, 0x88, 0x37, 0x31, 0x7c, 0xde, 0xa4, 0x88, 0xa7,
	0x3e, 0x87, 0x09, 0xfb, 0xf7, 0x3a, 0x6c, 0xe6, 0xbf, 0x07, 0xb2, 0xba, 0xda, 0x92, 0x4c, 0xb5,
	0x5d, 0x6d, 0xaa, 0xed, 0x54, 0x13, 0x9a, 0xb0, 0xee, 0x85, 0x51, 0xe6, 0x0f, 0xb1, 0x45, 0x7c,
	0xca, 0xd2, 0xc0, 0x0d, 0xf3, 0x86, 0xfc, 0x12, 0x6e, 0xcd, 0xde, 0x0c, 0x7d, 0x37, 0x75, 0xf1,
	0xcd, 0x02, 0xd4, 0x7b, 0x1a, 0x50, 0x53, 0x40, 0xf0, 0x10, 0xbd, 0x07, 0xa5, 0xf3, 0x21, 0xfa,
	0xaa, 0x76, 0x7b, 0x04, 0x2d, 0xd1, 0x9f, 0x43, 0x4e, 0x53, 0xa4, 0x4a, 0x04, 0xeb, 0x2e, 0x0c,
	0xf6, 0x19, 0x3a, 0x1c, 0xd3, 0x54, 0x05, 0x40, 0xd0, 0xdc, 0x0d, 0xd3, 0x13, 0x97, 0xd3, 0xe1,
	0x18, 0x79, 0x12, 0x53, 0xb3, 0x2c, 0x41, 0x6f, 0xc3, 0xa6, 0x88, 0x34, 0x44, 0x86, 0x13, 0x8c,
	0x51, 0xde, 0xb6, 0xe4, 0x2d, 0xd6, 0x49, 0xde, 0x16, 0xa7, 0x6a, 0x16, 0x0e, 0xc0, 0x9c, 0x0b,
	0x75, 0xc1, 0x64, 0x58, 0x0e, 0xac, 0x4c, 0x21, 0x9b, 0x6f, 0x5f, 0x97, 0x93, 0xf4, 0xb3, 0x51,
	0x36, 0xc5, 0xb1, 0xe7, 0x86, 0x73, 0x8a, 0xf6, 0x31, 0xac, 0x48, 0x82, 0xbc, 0x67, 0x2e, 0x1b,
	0xd1, 0xa2, 0x0d, 0xee, 0x6b, 0x48, 0x9a, 0x0c, 0x22, 0x39, 0x1a, 0x28, 0x17, 0x89, 0xc6, 0xea,
	0xc3, 0xfa, 0xec, 0xd9, 0x42, 0x84, 0x8f, 0xcb, 0xb6, 0x3a, 0xa4, 0x21, 0x9d, 0xd7, 0x56, 0x5b,
	0xb0, 0x91, 0xe0, 0xb5, 0x28, 0x40, 0x42, 0xc5, 0xb0, 0xa0, 0xb8, 0x04, 0xb1, 0x0c, 0xd3, 0xb2,
	0xef, 0xc1, 0x56, 0x01, 0x4f, 0x94, 0x21, 0x8b, 0x8f, 0x71, 0x20, 0xe3, 0x74, 0x22, 0x52, 0x12,
	0x85, 0x85, 0x66, 0x3c, 0x85, 0x6d, 0xbd, 0x71, 0xc5, 0x71, 0xb9, 0x05, 0xab, 0x38, 0xce, 0x5f,
	0x07, 0xa3, 0x2c, 0x91, 0x7a, 0xad, 0x8a, 0x63, 0xef, 0x83, 0x55, 0xf4, 0x51, 0x18, 0x60, 0x55,
	0x07, 0xd2, 0x46, 0x2f, 0x5b, 0x9f, 0x97, 0x78, 0xa7, 0x6d, 0xff, 0x1b, 0x82, 0x0f, 0xe1, 0xd5,
	0xd9, 0xa8, 0x4a, 0x74, 0xf4, 0x8c, 0xfe, 0x1f, 0xda, 0x7e, 0xc0, 0x4f, 0x87, 0x97, 0x2b, 0xc3,
	0xee, 0x97, 0x6f, 0xf8, 0x94, 0x21, 0x02, 0xe6, 0x51, 0x25, 0x24, 0xca, 0x1d, 0x9b, 0x39, 0xd7,
	0x89, 0xe1, 0xc4, 0x5b, 0xbe, 0x37, 0xe0, 0x7f, 0x33, 0x4e, 0x33, 0x6b, 0x69, 0x03, 0x3a, 0xf4,
	0x39, 0x5e, 0x33, 0x9c, 0x67, 0x3c, 0xac, 0x4d, 0xed, 0xaa, 0xfa, 0x55, 0x93, 0x38, 0xd7, 0x03,
	0x3c, 0x0c, 0xd8, 0xe5, 0x61, 0xa3, 0xf0, 0x93, 0xb5, 0x54, 0x7b, 0x0a, 0x7b, 0x8b, 0x9f, 0x66,
	0x6a, 0x14, 0xed, 0xb3, 0x92, 0xd5, 0x69, 0xf4, 0x15, 0x59, 0x75, 0xa0, 0x1d, 0xe4, 0x7e, 0xc5,
	0x00, 0x6c, 0x69, 0x06, 0xa0, 0x88, 0x6d, 0x27, 0x65, 0xdf, 0x3e, 0x8d, 0xfd, 0xb9, 0x72, 0xa8,
	0x53, 0x8e, 0xda, 0xb5, 0xca, 0x51, 0xd7, 0x2a, 0x87, 0xa4, 0xa4, 0xff, 0x77, 0x1b, 0x5a, 0x79,
	0x52, 0x4e, 0x7e, 0x30, 0xa0, 0x55, 0xec, 0x1e, 0x62, 0x5f, 0xbb, 0x98, 0x24, 0x32, 0xeb, 0xf5,
	0x0a, 0xcb, 0xcb, 0x7e, 0xf7, 0xbb, 0x3f, 0xfe, 0xfa, 0xa9, 0xd6, 0x23, 0x5d, 0xfc, 0xb3, 0x24,
	0xe6, 0x1e, 0x4e, 0xb3, 0xfc, 0xfb, 0xe4, 0xd2, 0xb3, 0x37, 0xbe, 0xef, 0x3c, 0xe8, 0x15, 0x1b,
	0xa3, 0xf7, 0x52, 0x3c, 0xf6, 0x82, 0xbc, 0x80, 0xa6, 0xd2, 0x4f, 0xb2, 0xb3, 0x48, 0x5a, 0xad,
	0xad, 0x82, 0xf9, 0xa3, 0x88, 0x8d, 0x9e, 0x64, 0x8c, 0x05, 0xac, 0x6c, 0x7e, 0xfb, 0x81, 0xcc,
	0x7f, 0xcf, 0xda, 0xad, 0x96, 0xff, 0xc0, 0xd8, 0x27, 0xcf, 0x71, 0xd3, 0x0a, 0x59, 0x22, 0xaf,
	0x2d, 0x10, 0xac, 0x4a, 0x99, 0xed, 0x1b, 0x64, 0xfe, 0x16, 0x9a, 0x4a, 0xb4, 0xb4, 0xaf, 0x9e,
	0xd2, 0xb3, 0xeb, 0x73, 0xe7, 0xac, 0xef, 0xdf, 0x90, 0x75, 0xcc, 0xaf, 0x9a, 0x4f, 0x9b, 0x7f,
	0xaa, 0x2f, 0x2b, 0xe5, 0xb7, 0x6f, 0x98, 0x3f, 0x84, 0x25, 0x31, 0x63, 0x64, 0xa3, 0x88, 0xfd,
	0x45, 0x14, 0xf8, 0x45, 0x42, 0x5d, 0x35, 0x26, 0x27, 0xd2, 0x76, 0x64, 0xd2, 0x3d, 0x52, 0x91,
	0x70, 0xf2, 0x9b, 0x01, 0xab, 0x53, 0x9a, 0x4d, 0xee, 0xea, 0x0a, 0xae, 0x59, 0x01, 0xd6, 0xde,
	0x62, 0xc3, 0x1c, 0xd4, 0x40, 0x82, 0xfa, 0x80, 0xbc, 0x5f, 0x09, 0x54, 0x97, 0xab, 0x20, 0x5d,
	0x2e, 0xa3, 0xf4, 0x5e, 0x0a, 0x71, 0xba, 0x20, 0xbf, 0x18, 0xb0, 0x32, 0x29, 0xed, 0x64, 0x57,
	0x37, 0x14, 0x57, 0xf7, 0x84, 0x75, 0x77, 0xa1, 0xdd, 0xbf, 0x82, 0x59, 0x14, 0x0c, 0xbf, 0x45,
	0xac, 0xae, 0x5a, 0x1f, 0xe4, 0x57, 0x03, 0xda, 0x85, 0x9e, 0x71, 0x2d, 0x46, 0xcd, 0x1e, 0xd0,
	0x62, 0xd4, 0x29, 0xae, 0xfd, 0x89, 0xc4, 0xf8, 0x98, 0x3c, 0xaa, 0x8a, 0x71, 0x72, 0xbd, 0x5c,
	0xf4, 0x4a, 0x35, 0x3e, 0x69, 0xca, 0xff, 0xf9, 0xbc, 0xf3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xda, 0x38, 0xa8, 0x02, 0x4d, 0x0d, 0x00, 0x00,
}
