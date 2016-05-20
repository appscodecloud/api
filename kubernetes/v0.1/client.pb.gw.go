// Code generated by protoc-gen-grpc-gateway
// source: client.proto
// DO NOT EDIT!

/*
Package kubernetes is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package kubernetes

import (
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_Clients_Nodes_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.Nodes(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_Secrets_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.Secrets(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_Jobs_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.Jobs(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_Namespaces_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.Namespaces(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_Pods_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.Pods(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_Services_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.Services(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_ReplicationControllers_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.ReplicationControllers(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_Apps_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.Apps(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_App_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AppDescribeRequest
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

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

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

	msg, err := client.App(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_AppPods_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AppDescribeRequest
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

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

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

	msg, err := client.AppPods(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_ConfigMaps_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClientRequest
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

	msg, err := client.ConfigMaps(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_ConfigMap_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ConfigMapDescribeRequest
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

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

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

	msg, err := client.ConfigMap(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clients_Secret_0(ctx context.Context, marshaler runtime.Marshaler, client ClientsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SecretDescribeRequest
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

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

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

	msg, err := client.Secret(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterClientsHandlerFromEndpoint is same as RegisterClientsHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterClientsHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterClientsHandler(ctx, mux, conn)
}

// RegisterClientsHandler registers the http handlers for service Clients to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterClientsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewClientsClient(conn)

	mux.Handle("GET", pattern_Clients_Nodes_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Nodes_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Nodes_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_Secrets_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Secrets_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Secrets_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_Jobs_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Jobs_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Jobs_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_Namespaces_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Namespaces_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Namespaces_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_Pods_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Pods_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Pods_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_Services_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Services_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Services_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_ReplicationControllers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_ReplicationControllers_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_ReplicationControllers_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_Apps_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Apps_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Apps_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_App_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_App_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_App_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_AppPods_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_AppPods_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_AppPods_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_ConfigMaps_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_ConfigMaps_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_ConfigMaps_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_ConfigMap_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_ConfigMap_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_ConfigMap_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clients_Secret_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		resp, md, err := request_Clients_Secret_0(runtime.AnnotateContext(ctx, req), inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_Clients_Secret_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Clients_Nodes_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "nodes"}, ""))

	pattern_Clients_Secrets_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "secrets"}, ""))

	pattern_Clients_Jobs_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "jobs"}, ""))

	pattern_Clients_Namespaces_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "namespaces"}, ""))

	pattern_Clients_Pods_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "pods"}, ""))

	pattern_Clients_Services_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "services"}, ""))

	pattern_Clients_ReplicationControllers_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "rcs"}, ""))

	pattern_Clients_Apps_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "apps"}, ""))

	pattern_Clients_App_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6}, []string{"kubernetes", "v0.1", "client", "cluster", "apps", "namespace", "name"}, ""))

	pattern_Clients_AppPods_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6, 2, 7}, []string{"kubernetes", "v0.1", "client", "cluster", "apps", "namespace", "name", "pods"}, ""))

	pattern_Clients_ConfigMaps_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"kubernetes", "v0.1", "client", "cluster", "configmaps"}, ""))

	pattern_Clients_ConfigMap_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6}, []string{"kubernetes", "v0.1", "client", "cluster", "configmaps", "namespace", "name"}, ""))

	pattern_Clients_Secret_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6}, []string{"kubernetes", "v0.1", "client", "cluster", "secrets", "namespace", "name"}, ""))
)

var (
	forward_Clients_Nodes_0 = runtime.ForwardResponseMessage

	forward_Clients_Secrets_0 = runtime.ForwardResponseMessage

	forward_Clients_Jobs_0 = runtime.ForwardResponseMessage

	forward_Clients_Namespaces_0 = runtime.ForwardResponseMessage

	forward_Clients_Pods_0 = runtime.ForwardResponseMessage

	forward_Clients_Services_0 = runtime.ForwardResponseMessage

	forward_Clients_ReplicationControllers_0 = runtime.ForwardResponseMessage

	forward_Clients_Apps_0 = runtime.ForwardResponseMessage

	forward_Clients_App_0 = runtime.ForwardResponseMessage

	forward_Clients_AppPods_0 = runtime.ForwardResponseMessage

	forward_Clients_ConfigMaps_0 = runtime.ForwardResponseMessage

	forward_Clients_ConfigMap_0 = runtime.ForwardResponseMessage

	forward_Clients_Secret_0 = runtime.ForwardResponseMessage
)
