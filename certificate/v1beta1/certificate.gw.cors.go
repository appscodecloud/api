// Code generated by protoc-gen-grpc-gateway-cors
// source: certificate.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportCertificatesCorsPatterns returns an array of grpc gatway mux patterns for
// Certificates service to enable CORS.
func ExportCertificatesCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Certificates_List_0,
		pattern_Certificates_Describe_0,
		pattern_Certificates_Create_0,
		pattern_Certificates_Obtain_0,
		pattern_Certificates_Delete_0,
		pattern_Certificates_Renew_0,
		pattern_Certificates_Revoke_0,
		pattern_Certificates_Deploy_0,
	}
}
