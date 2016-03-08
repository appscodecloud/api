// Code generated by protoc-gen-grpc-gateway
// source: cloud.proto
// DO NOT EDIT!

/*
Package credential is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package credential

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/appscode/api/dtypes"
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

func request_Cloud_List_0(ctx context.Context, client CloudClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq dtypes.VoidRequest
	var metadata runtime.ServerMetadata

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Cloud_Create_0(ctx context.Context, client CloudClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CloudCredentialCreateRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Create(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Cloud_Update_0(ctx context.Context, client CloudClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CloudCredentialUpdateRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Update(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Cloud_Delete_0(ctx context.Context, client CloudClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CloudCredentialDeleteRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.Delete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterCloudHandlerFromEndpoint is same as RegisterCloudHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCloudHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCloudHandler(ctx, mux, conn)
}

// RegisterCloudHandler registers the http handlers for service Cloud to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCloudHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewCloudClient(conn)

	mux.Handle("GET", pattern_Cloud_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Cloud_List_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Cloud_List_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_Cloud_Create_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Cloud_Create_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Cloud_Create_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Cloud_Update_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Cloud_Update_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Cloud_Update_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_Cloud_Delete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Cloud_Delete_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Cloud_Delete_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Cloud_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "credentials", "v0", "cloud"}, ""))

	pattern_Cloud_Create_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "credentials", "v0", "cloud"}, ""))

	pattern_Cloud_Update_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "credentials", "v0", "cloud"}, ""))

	pattern_Cloud_Delete_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "credentials", "v0", "cloud", "name"}, ""))
)

var (
	forward_Cloud_List_0 = runtime.ForwardResponseMessage

	forward_Cloud_Create_0 = runtime.ForwardResponseMessage

	forward_Cloud_Update_0 = runtime.ForwardResponseMessage

	forward_Cloud_Delete_0 = runtime.ForwardResponseMessage
)