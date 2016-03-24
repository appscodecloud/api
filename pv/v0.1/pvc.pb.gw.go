// Code generated by protoc-gen-grpc-gateway
// source: pvc.proto
// DO NOT EDIT!

/*
Package pv is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pv

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.NewDoubleArray

func request_PersistentVolumeClaim_PVCList_0(ctx context.Context, client PersistentVolumeClaimClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PVCListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["cluster"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "cluster")
	}

	protoReq.Cluster, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.PVCList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_PersistentVolumeClaim_PVCRegister_0(ctx context.Context, client PersistentVolumeClaimClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PVCRegisterRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["cluster"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "cluster")
	}

	protoReq.Cluster, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.PVCRegister(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_PersistentVolumeClaim_PVCUnregister_0 = &utilities.DoubleArray{Encoding: map[string]int{"cluster": 0, "name": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}
)

func request_PersistentVolumeClaim_PVCUnregister_0(ctx context.Context, client PersistentVolumeClaimClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PVCUnregisterRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["cluster"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "cluster")
	}

	protoReq.Cluster, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_PersistentVolumeClaim_PVCUnregister_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.PVCUnregister(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterPersistentVolumeClaimHandlerFromEndpoint is same as RegisterPersistentVolumeClaimHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPersistentVolumeClaimHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPersistentVolumeClaimHandler(ctx, mux, conn)
}

// RegisterPersistentVolumeClaimHandler registers the http handlers for service PersistentVolumeClaim to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPersistentVolumeClaimHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewPersistentVolumeClaimClient(conn)

	mux.Handle("GET", pattern_PersistentVolumeClaim_PVCList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_PersistentVolumeClaim_PVCList_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_PersistentVolumeClaim_PVCList_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_PersistentVolumeClaim_PVCRegister_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_PersistentVolumeClaim_PVCRegister_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_PersistentVolumeClaim_PVCRegister_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_PersistentVolumeClaim_PVCUnregister_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_PersistentVolumeClaim_PVCUnregister_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_PersistentVolumeClaim_PVCUnregister_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PersistentVolumeClaim_PVCList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "pv", "v0", "pvc", "cluster"}, ""))

	pattern_PersistentVolumeClaim_PVCRegister_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "pv", "v0", "pvc", "cluster", "name"}, ""))

	pattern_PersistentVolumeClaim_PVCUnregister_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "pv", "v0", "pvc", "cluster", "name"}, ""))
)

var (
	forward_PersistentVolumeClaim_PVCList_0 = runtime.ForwardResponseMessage

	forward_PersistentVolumeClaim_PVCRegister_0 = runtime.ForwardResponseMessage

	forward_PersistentVolumeClaim_PVCUnregister_0 = runtime.ForwardResponseMessage
)
