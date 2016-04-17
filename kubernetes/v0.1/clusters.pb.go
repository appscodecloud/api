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
	Phid        string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Provider    string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	Zone        string `protobuf:"bytes,4,opt,name=zone" json:"zone,omitempty"`
	Region      string `protobuf:"bytes,5,opt,name=region" json:"region,omitempty"`
	Os          string `protobuf:"bytes,6,opt,name=os" json:"os,omitempty"`
	GraphanaUrl string `protobuf:"bytes,7,opt,name=graphana_url,json=graphanaUrl" json:"graphana_url,omitempty"`
	KibanaUrl   string `protobuf:"bytes,8,opt,name=kibana_url,json=kibanaUrl" json:"kibana_url,omitempty"`
	IcingaUrl   string `protobuf:"bytes,9,opt,name=icinga_url,json=icingaUrl" json:"icinga_url,omitempty"`
	Nodes       int32  `protobuf:"varint,10,opt,name=nodes" json:"nodes,omitempty"`
	CreatedAt   string `protobuf:"bytes,11,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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
	CpuCore     int64 `protobuf:"varint,5,opt,name=cpu_core,json=cpuCore" json:"cpu_core,omitempty"`
	TotalMemory int64 `protobuf:"varint,6,opt,name=total_memory,json=totalMemory" json:"total_memory,omitempty"`
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
	CloudCredential     string            `protobuf:"bytes,4,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	CloudCredentialData map[string]string `protobuf:"bytes,5,rep,name=cloud_credential_data,json=cloudCredentialData" json:"cloud_credential_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeSet             map[string]int64  `protobuf:"bytes,6,rep,name=node_set,json=nodeSet" json:"node_set,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	SaltbaseVersion     string            `protobuf:"bytes,7,opt,name=saltbase_version,json=saltbaseVersion" json:"saltbase_version,omitempty"`
	KubeStarterVersion  string            `protobuf:"bytes,8,opt,name=kube_starter_version,json=kubeStarterVersion" json:"kube_starter_version,omitempty"`
	KubeVersion         string            `protobuf:"bytes,9,opt,name=kube_version,json=kubeVersion" json:"kube_version,omitempty"`
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
	NodeChanges map[string]int64 `protobuf:"bytes,2,rep,name=node_changes,json=nodeChanges" json:"node_changes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
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
	ReleaseReservedIp bool   `protobuf:"varint,2,opt,name=release_reserved_ip,json=releaseReservedIp" json:"release_reserved_ip,omitempty"`
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
	DiskName string `protobuf:"bytes,2,opt,name=disk_name,json=diskName" json:"disk_name,omitempty"`
}

func (m *ClusterClientContainerRequest) Reset()                    { *m = ClusterClientContainerRequest{} }
func (m *ClusterClientContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterClientContainerRequest) ProtoMessage()               {}
func (*ClusterClientContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type ClusterInstanceListRequest struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
}

func (m *ClusterInstanceListRequest) Reset()                    { *m = ClusterInstanceListRequest{} }
func (m *ClusterInstanceListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterInstanceListRequest) ProtoMessage()               {}
func (*ClusterInstanceListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

type ClusterInstance struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId" json:"external_id,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	ExternalIp string `protobuf:"bytes,4,opt,name=external_ip,json=externalIp" json:"external_ip,omitempty"`
	InternalIp string `protobuf:"bytes,5,opt,name=internal_ip,json=internalIp" json:"internal_ip,omitempty"`
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
	SaltbaseVersion    string `protobuf:"bytes,2,opt,name=saltbase_version,json=saltbaseVersion" json:"saltbase_version,omitempty"`
	KubeStarterVersion string `protobuf:"bytes,3,opt,name=kube_starter_version,json=kubeStarterVersion" json:"kube_starter_version,omitempty"`
	KubeVersion        string `protobuf:"bytes,4,opt,name=kube_version,json=kubeVersion" json:"kube_version,omitempty"`
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
const _ = grpc.SupportPackageIsVersion1

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

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Scale(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_StartupScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterStartupScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).StartupScript(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_ClientConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterClientConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).ClientConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Instances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterInstanceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Instances(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
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

var fileDescriptor1 = []byte{
	// 1273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x57, 0xcd, 0x6e, 0xdc, 0x54,
	0x14, 0x96, 0xe7, 0x2f, 0x33, 0x67, 0xd2, 0x36, 0xdc, 0x84, 0x30, 0x38, 0x0d, 0x29, 0xa6, 0xa4,
	0x69, 0xd5, 0xcc, 0xa4, 0x01, 0xa1, 0x36, 0x8b, 0x56, 0x68, 0x02, 0x22, 0x52, 0xa9, 0x90, 0x47,
	0xed, 0x82, 0x8d, 0xe5, 0xd8, 0x97, 0xa9, 0x35, 0xce, 0xb5, 0xeb, 0x6b, 0x47, 0x84, 0x2a, 0x12,
	0xf0, 0x0a, 0x2c, 0x41, 0xb0, 0x60, 0x81, 0xc4, 0x43, 0xb0, 0xe8, 0x2b, 0xb0, 0x64, 0xcb, 0x83,
	0x70, 0x7f, 0x3d, 0x9e, 0xc4, 0x99, 0x9f, 0x05, 0x9b, 0xd1, 0xf5, 0x77, 0xce, 0xb9, 0xe7, 0xff,
	0x9c, 0x3b, 0x70, 0xdd, 0x0b, 0x33, 0x9a, 0xe2, 0x84, 0x76, 0xe3, 0x24, 0x4a, 0x23, 0x04, 0xa3,
	0xec, 0x18, 0x27, 0x04, 0xa7, 0x98, 0x9a, 0x37, 0x87, 0x51, 0x34, 0x0c, 0x71, 0xcf, 0x8d, 0x83,
	0x9e, 0x4b, 0x48, 0x94, 0xba, 0x69, 0x10, 0x11, 0xc5, 0x69, 0xae, 0x73, 0xd8, 0x4f, 0xcf, 0x62,
	0x4c, 0x7b, 0xe2, 0x57, 0xe2, 0xd6, 0x6f, 0x15, 0x58, 0xea, 0xcb, 0x4b, 0x11, 0x82, 0x5a, 0xfc,
	0x32, 0xf0, 0x3b, 0xc6, 0x2d, 0x63, 0xa7, 0x65, 0x8b, 0x33, 0xc7, 0x88, 0x7b, 0x82, 0x3b, 0x15,
	0x89, 0xf1, 0x33, 0x32, 0xa1, 0xc9, 0x84, 0x4f, 0x03, 0x1f, 0x27, 0x9d, 0xaa, 0xc0, 0xf3, 0x6f,
	0xce, 0xff, 0x5d, 0x44, 0x70, 0xa7, 0x26, 0xf9, 0xf9, 0x19, 0xad, 0x43, 0x23, 0xc1, 0x43, 0x66,
	0x4c, 0xa7, 0x2e, 0x50, 0xf5, 0x85, 0xae, 0x43, 0x25, 0xa2, 0x9d, 0x86, 0xc0, 0xd8, 0x09, 0xbd,
	0x0f, 0xcb, 0xc3, 0xc4, 0x8d, 0x5f, 0xba, 0xc4, 0x75, 0xb2, 0x24, 0xec, 0x2c, 0x09, 0x4a, 0x5b,
	0x63, 0xcf, 0x93, 0x10, 0x6d, 0x02, 0x8c, 0x82, 0x63, 0xcd, 0xd0, 0x14, 0x0c, 0x2d, 0x89, 0x28,
	0x72, 0xe0, 0x05, 0x64, 0x28, 0xc9, 0x2d, 0x49, 0x96, 0x08, 0x27, 0xaf, 0x41, 0x9d, 0x44, 0x3e,
	0xa6, 0x1d, 0x60, 0x94, 0xba, 0x2d, 0x3f, 0xb8, 0x90, 0x97, 0x60, 0x37, 0xc5, 0xbe, 0xe3, 0xa6,
	0x9d, 0xb6, 0x14, 0x52, 0xc8, 0xa7, 0xa9, 0x75, 0x1f, 0xd6, 0x55, 0x80, 0x0e, 0x31, 0xf5, 0x92,
	0xe0, 0x18, 0xdb, 0xf8, 0x55, 0x86, 0x69, 0x9a, 0xc7, 0xc6, 0x18, 0xc7, 0xc6, 0xfa, 0xa7, 0x02,
	0xef, 0x5c, 0x62, 0xa7, 0x31, 0x4b, 0x04, 0x46, 0xdb, 0xd0, 0xa0, 0x2c, 0x2b, 0x19, 0x15, 0x12,
	0xed, 0xfd, 0xeb, 0x5d, 0x99, 0x90, 0xee, 0x40, 0xa0, 0xb6, 0xa2, 0xa2, 0x5d, 0x58, 0x52, 0x79,
	0x16, 0x61, 0x6f, 0xef, 0xaf, 0x76, 0xc7, 0x79, 0xee, 0xaa, 0xdb, 0x6d, 0xcd, 0x83, 0x9e, 0x40,
	0x9d, 0xc6, 0xd8, 0xa3, 0x22, 0x17, 0xed, 0xfd, 0xbb, 0x25, 0xcc, 0x17, 0x4d, 0xe9, 0x0e, 0xb8,
	0x80, 0x2d, 0xe5, 0xcc, 0x3f, 0x0c, 0xa8, 0x0b, 0x00, 0xbd, 0xc7, 0x42, 0x11, 0x91, 0xd4, 0x0d,
	0x08, 0xab, 0x31, 0x61, 0x65, 0xdd, 0x2e, 0x20, 0xa2, 0x42, 0x22, 0x9f, 0x0a, 0xb3, 0xea, 0xb6,
	0x38, 0xf3, 0x6a, 0xa0, 0x38, 0x39, 0x0d, 0x3c, 0x2c, 0x2d, 0xa8, 0xdb, 0xf9, 0x37, 0x5a, 0x81,
	0x6a, 0xc2, 0x0c, 0xab, 0x09, 0x98, 0x1f, 0xd1, 0xbb, 0xd0, 0xf4, 0xe2, 0xcc, 0xf1, 0xa2, 0x04,
	0x8b, 0x6a, 0xa8, 0x32, 0x3f, 0xe2, 0xac, 0xcf, 0x3e, 0x79, 0xfa, 0x53, 0x56, 0xb5, 0xa1, 0x73,
	0x82, 0x4f, 0xa2, 0xe4, 0x4c, 0x14, 0x46, 0xd5, 0x6e, 0x0b, 0xec, 0x4b, 0x01, 0x59, 0x04, 0x56,
	0x95, 0x47, 0x4f, 0x03, 0x9a, 0x2e, 0x1c, 0xd8, 0x1e, 0x53, 0xae, 0x1a, 0x88, 0xb9, 0x50, 0xbd,
	0x2a, 0xb2, 0x39, 0x93, 0xf5, 0xa6, 0x06, 0x6b, 0x0a, 0xed, 0x8b, 0x82, 0x98, 0x92, 0xfa, 0x89,
	0xb6, 0xa8, 0x5c, 0xd1, 0x16, 0xd5, 0x42, 0x5b, 0xdc, 0x85, 0x15, 0x2f, 0x8c, 0x32, 0xdf, 0x61,
	0xb5, 0xe6, 0x63, 0x92, 0x06, 0x6e, 0xa8, 0xda, 0xe6, 0x86, 0xc0, 0xfb, 0x39, 0x8c, 0x4e, 0xe0,
	0xed, 0x8b, 0xac, 0x8e, 0xef, 0xa6, 0x2e, 0x0b, 0x21, 0xf7, 0xe2, 0x51, 0x89, 0x17, 0x13, 0xf6,
	0x32, 0x70, 0xe2, 0xc2, 0x43, 0x26, 0xfb, 0x19, 0x49, 0x93, 0x33, 0x7b, 0xd5, 0xbb, 0x4c, 0x41,
	0x5f, 0x40, 0x93, 0xb7, 0x86, 0x43, 0x71, 0xca, 0xb2, 0xc0, 0x35, 0xec, 0xce, 0xd4, 0xf0, 0x8c,
	0x09, 0x0c, 0x70, 0x2a, 0x6f, 0x5d, 0x22, 0xf2, 0x8b, 0xfb, 0x48, 0xdd, 0x30, 0x3d, 0x76, 0x29,
	0x76, 0x4e, 0x59, 0x44, 0xf9, 0x10, 0x90, 0x6d, 0x7d, 0x43, 0xe3, 0x2f, 0x24, 0x8c, 0xf6, 0x60,
	0x8d, 0xeb, 0x70, 0x58, 0xae, 0x12, 0x76, 0x7b, 0xce, 0x2e, 0x9b, 0x1c, 0x71, 0xda, 0x40, 0x92,
	0xb4, 0x04, 0x2b, 0x18, 0x21, 0xa1, 0x39, 0x65, 0xbf, 0xb7, 0x39, 0xa6, 0x58, 0xcc, 0xcf, 0xa1,
	0x73, 0x95, 0xeb, 0xbc, 0x38, 0x47, 0xf8, 0x4c, 0xa5, 0x90, 0x1f, 0xf9, 0x7c, 0x38, 0x75, 0xc3,
	0x4c, 0x4f, 0x3b, 0xf9, 0x71, 0x50, 0x79, 0x68, 0x98, 0x07, 0xb0, 0x5c, 0x74, 0x70, 0x96, 0x6c,
	0xb5, 0x20, 0x6b, 0xfd, 0x65, 0xe4, 0x55, 0x3b, 0xf0, 0xdc, 0x70, 0x6a, 0x0d, 0x0d, 0x60, 0x59,
	0x44, 0xde, 0x63, 0x03, 0x6f, 0x88, 0x75, 0x95, 0xee, 0x95, 0x44, 0xbf, 0x78, 0x95, 0x08, 0x7e,
	0x5f, 0x8a, 0xc8, 0x04, 0xb4, 0xc9, 0x18, 0x31, 0x1f, 0xc3, 0xca, 0x45, 0x86, 0x85, 0x1c, 0xf8,
	0x3a, 0x6f, 0x82, 0x43, 0x1c, 0xe2, 0xe9, 0x4d, 0xd0, 0x85, 0xd5, 0x84, 0x31, 0xf1, 0x7c, 0x27,
	0x98, 0xcf, 0x01, 0x36, 0x55, 0x83, 0x58, 0xdc, 0xd9, 0xb4, 0xdf, 0x52, 0x24, 0x5b, 0x51, 0x8e,
	0x62, 0xeb, 0x01, 0x6c, 0x68, 0x87, 0x78, 0x72, 0xb3, 0x78, 0xc0, 0x06, 0x55, 0x9c, 0x16, 0x54,
	0x24, 0x51, 0x98, 0xab, 0xe0, 0x67, 0x2b, 0x84, 0x9b, 0xe5, 0x22, 0x0b, 0x4e, 0x83, 0xdb, 0x70,
	0x8d, 0x8d, 0xb6, 0x6f, 0x82, 0x61, 0x96, 0x88, 0x55, 0xa9, 0xb2, 0x3e, 0x09, 0x5a, 0x7b, 0x60,
	0xea, 0x7a, 0x0f, 0x03, 0x56, 0x42, 0x7d, 0x41, 0x9d, 0xb6, 0x02, 0x46, 0xb9, 0x4b, 0x93, 0x12,
	0xff, 0x8b, 0x79, 0x5f, 0xc1, 0xe6, 0x45, 0x65, 0x72, 0x58, 0x4f, 0x4b, 0xd2, 0x06, 0xb4, 0xfc,
	0x80, 0x8e, 0x9c, 0xc2, 0x66, 0x6f, 0x72, 0xe0, 0x19, 0x37, 0xff, 0x49, 0xee, 0xf0, 0x11, 0x61,
	0xa6, 0x10, 0x0f, 0xcb, 0x59, 0x2b, 0xaf, 0x63, 0x3d, 0xa7, 0xa6, 0xa3, 0x53, 0xb8, 0xb6, 0xad,
	0x30, 0x71, 0xc1, 0x1b, 0x03, 0x6e, 0x5c, 0xb8, 0xa1, 0xf4, 0x69, 0xb1, 0x05, 0x6d, 0xfc, 0x2d,
	0x63, 0x22, 0x6c, 0x98, 0x31, 0x92, 0xb4, 0x03, 0x34, 0x74, 0x34, 0x7e, 0x7b, 0x54, 0x0b, 0xa6,
	0x4f, 0x08, 0xc5, 0x6a, 0x5e, 0x8e, 0x85, 0x62, 0xce, 0x10, 0x90, 0x31, 0x83, 0x7c, 0x71, 0x80,
	0x86, 0x18, 0x83, 0x2e, 0xa9, 0xc6, 0xb8, 0xa4, 0x78, 0x37, 0xd0, 0x51, 0xa6, 0x26, 0x13, 0x3f,
	0x5a, 0xdf, 0x1b, 0x79, 0x16, 0x27, 0xc3, 0xb0, 0x60, 0x16, 0x1f, 0x41, 0x2b, 0x50, 0xf2, 0xba,
	0x9b, 0x37, 0x4a, 0xba, 0x59, 0xeb, 0xb0, 0xc7, 0xdc, 0xd6, 0x9f, 0x46, 0xde, 0x77, 0xcf, 0x63,
	0x7f, 0xc6, 0xf2, 0x29, 0x1b, 0xb4, 0x95, 0xc5, 0x06, 0x6d, 0x75, 0xee, 0x41, 0x5b, 0xbb, 0x34,
	0x68, 0xf7, 0x7f, 0x6f, 0x41, 0x53, 0x19, 0x4b, 0xd1, 0x0f, 0x06, 0x34, 0xf5, 0x93, 0x03, 0x59,
	0x53, 0xdf, 0x23, 0xc2, 0x23, 0xf3, 0x83, 0x39, 0xde, 0x2c, 0xd6, 0xfd, 0x1f, 0xff, 0xfe, 0xf7,
	0xa7, 0xca, 0x36, 0xba, 0x2d, 0x9e, 0xb8, 0x63, 0x81, 0xde, 0xe9, 0x5e, 0xf7, 0x41, 0x4f, 0xef,
	0xec, 0xde, 0x6b, 0x1e, 0x8f, 0x73, 0xf4, 0x0a, 0x1a, 0x72, 0x41, 0xa1, 0x5b, 0xb3, 0x76, 0x97,
	0xb9, 0xa1, 0x93, 0xf7, 0x34, 0x22, 0x43, 0x3b, 0x23, 0x84, 0x3d, 0x18, 0x73, 0xb5, 0x3b, 0x42,
	0xad, 0x65, 0x6e, 0x4e, 0x55, 0x7b, 0x60, 0xdc, 0x43, 0x11, 0x7b, 0x46, 0xf1, 0xa9, 0x8c, 0xb6,
	0x66, 0xcc, 0xeb, 0xb9, 0x14, 0x5a, 0xb3, 0x15, 0x66, 0xd0, 0x90, 0x13, 0xb9, 0xd4, 0xc7, 0x89,
	0x61, 0x3d, 0x5d, 0xa5, 0x0a, 0xed, 0xbd, 0xf9, 0x42, 0xcb, 0xd4, 0xca, 0x82, 0x2c, 0x55, 0x3b,
	0x51, 0xab, 0x73, 0xa9, 0x9d, 0x33, 0xa3, 0x1e, 0xd4, 0x78, 0x0b, 0xa2, 0x55, 0x7d, 0xe5, 0x8b,
	0x28, 0xf0, 0xb5, 0x9e, 0xb2, 0x90, 0x17, 0x1b, 0xd6, 0xfa, 0x50, 0xe8, 0xda, 0x42, 0xd3, 0xa3,
	0x8a, 0x7e, 0x35, 0xe0, 0xda, 0xc4, 0x5a, 0x41, 0x77, 0xca, 0x92, 0x59, 0xb2, 0xab, 0xcc, 0x9d,
	0xd9, 0x8c, 0xca, 0x96, 0x03, 0x61, 0xcb, 0xc7, 0x68, 0x7f, 0x9a, 0x2d, 0xbb, 0x54, 0xca, 0xee,
	0x52, 0x21, 0xdc, 0x7b, 0xcd, 0x27, 0xd5, 0x39, 0xfa, 0xd9, 0x80, 0xe5, 0xe2, 0x5e, 0x41, 0xdb,
	0x65, 0xe5, 0x7d, 0x79, 0x55, 0x99, 0x77, 0x66, 0xf2, 0x2d, 0x62, 0x9d, 0xce, 0x0a, 0xfb, 0xe6,
	0x57, 0xec, 0xca, 0xad, 0x84, 0x7e, 0x31, 0xa0, 0xa5, 0x67, 0x19, 0x2d, 0x35, 0xad, 0x64, 0xa9,
	0x94, 0x9a, 0x56, 0x36, 0x75, 0xad, 0xc7, 0xc2, 0xb4, 0x87, 0xe8, 0x93, 0x19, 0xa6, 0x15, 0x57,
	0xd4, 0x79, 0x2f, 0x1f, 0xa9, 0xc7, 0x0d, 0xf1, 0xa7, 0xf7, 0xa3, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x18, 0xc7, 0x98, 0x48, 0x0f, 0x00, 0x00,
}