// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package dtypes is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	Status
	Help
	ErrorDetails
	LongRunningResponse
	VoidRequest
	VoidResponse
*/
package dtypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Available Next Id 20.
type StatusCode int32

const (
	// Not an error. Returned on success.
	// Similar to HTTP Status code 2**.
	StatusCode_OK StatusCode = 0
	// Request Unsucesful
	StatusCode_FAILED StatusCode = 1
	// The request does not have valid authentication credentials for the
	// operation. Similar to HTTP status code 401.
	StatusCode_UNAUTHORIZED StatusCode = 2
	// The request contains invalid arguments.
	// Similar to HTTP status code 400
	StatusCode_BADREQUEST StatusCode = 3
	// The caller does not have permission to execute the specified
	// operation. Reserved For further use.
	// Similar to HTTP Status code 403
	StatusCode_PERMISSION_DENIED StatusCode = 4
	// Requested entity not found.
	// Similar to HTTP status code 404
	StatusCode_NOT_FOUND StatusCode = 5
	// The operation is not implemented or is not supported/enabled
	// Similar to HTTP status code 501
	StatusCode_UNIMPLEMENTED StatusCode = 6
	// Internal errors. This means that some invariants expected by the
	// underlying system have been broken.
	// Similar to HTTP status code 500
	StatusCode_INTERNAL StatusCode = 7
	// External Server Refuses Connection or Sends Invalid Data.
	StatusCode_EXTERNAL StatusCode = 8
	// The response contains invalid arguments.
	StatusCode_BAD_RESPONSE StatusCode = 9
	// Unknown error.
	// Errors raised by APIs that do not return enough error information
	// may be converted to this error.
	// Similar to HTTP status code 500
	StatusCode_UNKNOWN_ERROR StatusCode = 10
	// Quota and payment related error codes.
	// client request crossses the exisitng quta limits.
	StatusCode_QUOTA_LIMIT_EXCEED StatusCode = 11
	// no valid quta available
	StatusCode_INVALID_QUOTA StatusCode = 12
	// no payment informtaion avaible
	StatusCode_PAYMENT_INFORMATION_UNAVAILABLE StatusCode = 13
	// payment information not valid.
	StatusCode_INVALID_PAYMENT_INFORMATION StatusCode = 14
	// failed to transect the payment information
	StatusCode_TRANSACTION_FAILED StatusCode = 15
	// to give ARE_YOU_SURE option
	StatusCode_ARE_YOU_SURE StatusCode = 16
	// invalied data
	StatusCode_INVALIED_DATA StatusCode = 17
)

var StatusCode_name = map[int32]string{
	0:  "OK",
	1:  "FAILED",
	2:  "UNAUTHORIZED",
	3:  "BADREQUEST",
	4:  "PERMISSION_DENIED",
	5:  "NOT_FOUND",
	6:  "UNIMPLEMENTED",
	7:  "INTERNAL",
	8:  "EXTERNAL",
	9:  "BAD_RESPONSE",
	10: "UNKNOWN_ERROR",
	11: "QUOTA_LIMIT_EXCEED",
	12: "INVALID_QUOTA",
	13: "PAYMENT_INFORMATION_UNAVAILABLE",
	14: "INVALID_PAYMENT_INFORMATION",
	15: "TRANSACTION_FAILED",
	16: "ARE_YOU_SURE",
	17: "INVALIED_DATA",
}
var StatusCode_value = map[string]int32{
	"OK":                              0,
	"FAILED":                          1,
	"UNAUTHORIZED":                    2,
	"BADREQUEST":                      3,
	"PERMISSION_DENIED":               4,
	"NOT_FOUND":                       5,
	"UNIMPLEMENTED":                   6,
	"INTERNAL":                        7,
	"EXTERNAL":                        8,
	"BAD_RESPONSE":                    9,
	"UNKNOWN_ERROR":                   10,
	"QUOTA_LIMIT_EXCEED":              11,
	"INVALID_QUOTA":                   12,
	"PAYMENT_INFORMATION_UNAVAILABLE": 13,
	"INVALID_PAYMENT_INFORMATION":     14,
	"TRANSACTION_FAILED":              15,
	"ARE_YOU_SURE":                    16,
	"INVALIED_DATA":                   17,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Next Id 5.
type Status struct {
	// Response status code
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	// Response status code string.
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	// User facing message.
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	// Optional. Help link if there is an error.
	Help *Help `protobuf:"bytes,4,opt,name=help" json:"help,omitempty"`
	// A list of messages that carry the error details.  There will be a
	// common set of message types for APIs to use.
	Details []*google_protobuf.Any `protobuf:"bytes,5,rep,name=details" json:"details,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetHelp() *Help {
	if m != nil {
		return m.Help
	}
	return nil
}

func (m *Status) GetDetails() []*google_protobuf.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// Provides links to documentation or for performing an out of band action.
// Next Id 3;
type Help struct {
	// Describe what link offers
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// The URL of The link.
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *Help) Reset()                    { *m = Help{} }
func (m *Help) String() string            { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()               {}
func (*Help) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Basic Error details message to send in response. Application specific
// error messages can be provided.
type ErrorDetails struct {
	RequestedResource string `protobuf:"bytes,1,opt,name=requested_resource,json=requestedResource" json:"requested_resource,omitempty"`
	Stacktrace        string `protobuf:"bytes,2,opt,name=stacktrace" json:"stacktrace,omitempty"`
}

func (m *ErrorDetails) Reset()                    { *m = ErrorDetails{} }
func (m *ErrorDetails) String() string            { return proto.CompactTextString(m) }
func (*ErrorDetails) ProtoMessage()               {}
func (*ErrorDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Available Next ID: 6
type LongRunningResponse struct {
	Status  *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	TypeId  string  `protobuf:"bytes,2,opt,name=type_id,json=typeId" json:"type_id,omitempty"`
	JobId   string  `protobuf:"bytes,3,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	JobType string  `protobuf:"bytes,4,opt,name=job_type,json=jobType" json:"job_type,omitempty"`
}

func (m *LongRunningResponse) Reset()                    { *m = LongRunningResponse{} }
func (m *LongRunningResponse) String() string            { return proto.CompactTextString(m) }
func (*LongRunningResponse) ProtoMessage()               {}
func (*LongRunningResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LongRunningResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Void Requests and response to use with other types.
type VoidRequest struct {
}

func (m *VoidRequest) Reset()                    { *m = VoidRequest{} }
func (m *VoidRequest) String() string            { return proto.CompactTextString(m) }
func (*VoidRequest) ProtoMessage()               {}
func (*VoidRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type VoidResponse struct {
	Status *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *VoidResponse) Reset()                    { *m = VoidResponse{} }
func (m *VoidResponse) String() string            { return proto.CompactTextString(m) }
func (*VoidResponse) ProtoMessage()               {}
func (*VoidResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VoidResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "dtypes.Status")
	proto.RegisterType((*Help)(nil), "dtypes.Help")
	proto.RegisterType((*ErrorDetails)(nil), "dtypes.ErrorDetails")
	proto.RegisterType((*LongRunningResponse)(nil), "dtypes.LongRunningResponse")
	proto.RegisterType((*VoidRequest)(nil), "dtypes.VoidRequest")
	proto.RegisterType((*VoidResponse)(nil), "dtypes.VoidResponse")
	proto.RegisterEnum("dtypes.StatusCode", StatusCode_name, StatusCode_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0xf5, 0xcf, 0xd2, 0xed, 0xb6, 0xdd, 0xcf, 0x35, 0xff, 0x3a, 0x90, 0xd8, 0x54, 0x24,
	0x84, 0x90, 0xc8, 0x24, 0x90, 0x78, 0xe0, 0xcd, 0x5b, 0x3c, 0xcd, 0x5a, 0xea, 0x74, 0x4e, 0x32,
	0x36, 0x24, 0x64, 0x75, 0x8d, 0x29, 0x85, 0x92, 0x94, 0x24, 0x7d, 0xd8, 0x37, 0xe0, 0x63, 0xf0,
	0xf9, 0xf8, 0x14, 0xd8, 0x4e, 0x33, 0xf6, 0xc0, 0x0b, 0x4f, 0xf1, 0x39, 0xf7, 0xf8, 0xde, 0x73,
	0x4f, 0x0c, 0xdd, 0xf2, 0x66, 0xa5, 0x0a, 0x77, 0x95, 0x67, 0x65, 0x86, 0x9d, 0xc4, 0xa2, 0xc7,
	0x7b, 0xf3, 0x2c, 0x9b, 0x2f, 0xd5, 0xa1, 0x65, 0xaf, 0xd7, 0x9f, 0x0e, 0xa7, 0xe9, 0x4d, 0x25,
	0x19, 0xfd, 0x6c, 0x80, 0x13, 0x96, 0xd3, 0x72, 0x5d, 0x60, 0x0c, 0xed, 0x59, 0x96, 0xa8, 0x61,
	0xe3, 0xa0, 0xf1, 0x62, 0x47, 0xd8, 0x33, 0x7e, 0x08, 0x4e, 0x61, 0xab, 0xc3, 0xa6, 0x65, 0x37,
	0x08, 0x0f, 0xa1, 0xf3, 0x4d, 0x15, 0xc5, 0x74, 0xae, 0x86, 0x2d, 0x5b, 0xa8, 0x21, 0x3e, 0x80,
	0xf6, 0x67, 0xb5, 0x5c, 0x0d, 0xdb, 0x9a, 0xee, 0xbe, 0xee, 0xb9, 0x95, 0x05, 0xf7, 0x54, 0x73,
	0xc2, 0x56, 0xb0, 0x0b, 0x9d, 0x44, 0x95, 0xd3, 0xc5, 0xb2, 0x18, 0x6e, 0x1d, 0xb4, 0xb4, 0xe8,
	0xbe, 0x5b, 0xf9, 0x73, 0x6b, 0x7f, 0x2e, 0x49, 0x6f, 0x44, 0x2d, 0x1a, 0xbd, 0x83, 0xb6, 0xb9,
	0xad, 0x3b, 0x77, 0x13, 0x55, 0xcc, 0xf2, 0xc5, 0xaa, 0x5c, 0x64, 0xe9, 0xc6, 0xe6, 0x5d, 0x0a,
	0x23, 0x68, 0xad, 0xf3, 0xe5, 0xc6, 0xaa, 0x39, 0x8e, 0x3e, 0x42, 0x8f, 0xe6, 0x79, 0x96, 0x7b,
	0x55, 0x2f, 0xfc, 0x0a, 0x70, 0xae, 0xbe, 0xaf, 0x55, 0x51, 0xaa, 0x44, 0xe6, 0xaa, 0xc8, 0xd6,
	0xf9, 0xac, 0xde, 0x78, 0x70, 0x5b, 0x11, 0x9b, 0x02, 0x7e, 0x0a, 0xa0, 0x17, 0x9e, 0x7d, 0x2d,
	0xf3, 0xa9, 0x96, 0x55, 0x7d, 0xef, 0x30, 0xa3, 0x1f, 0x0d, 0xb8, 0xe7, 0x67, 0xe9, 0x5c, 0xac,
	0xd3, 0x74, 0xa1, 0x3f, 0xaa, 0x58, 0x65, 0x69, 0xa1, 0xf0, 0xf3, 0xdb, 0xd8, 0x1a, 0x36, 0x86,
	0xdd, 0x3a, 0x86, 0x2a, 0xea, 0xdb, 0x18, 0x1f, 0x41, 0xc7, 0xf0, 0x72, 0x91, 0xd4, 0xf9, 0x1a,
	0xc8, 0x12, 0xfc, 0x00, 0x9c, 0x2f, 0xd9, 0xb5, 0xe1, 0xab, 0x78, 0xb7, 0x34, 0xd2, 0xf4, 0x1e,
	0x6c, 0x1b, 0xda, 0x88, 0x6c, 0xc0, 0x3a, 0x77, 0x8d, 0x23, 0x0d, 0x47, 0x7d, 0xe8, 0x5e, 0x64,
	0x0b, 0x6d, 0xdd, 0xee, 0x30, 0x7a, 0x0b, 0xbd, 0x0a, 0xfe, 0x9b, 0xa3, 0x97, 0xbf, 0x9a, 0x00,
	0x15, 0x75, 0x6c, 0xfe, 0xbf, 0x03, 0xcd, 0xe0, 0x0c, 0xfd, 0x87, 0x01, 0x9c, 0x13, 0xc2, 0x7c,
	0xea, 0xa1, 0x86, 0x4e, 0xb9, 0x17, 0x73, 0x12, 0x47, 0xa7, 0x81, 0x60, 0x1f, 0x34, 0xd3, 0xc4,
	0xbb, 0x00, 0x47, 0xc4, 0x13, 0xf4, 0x3c, 0xa6, 0x61, 0x84, 0x5a, 0xda, 0xfd, 0x60, 0x42, 0xc5,
	0x98, 0x85, 0x21, 0x0b, 0xb8, 0xf4, 0x28, 0x67, 0x5a, 0xd6, 0xc6, 0x7d, 0xd8, 0xe1, 0x41, 0x24,
	0x4f, 0x82, 0x98, 0x7b, 0x68, 0x0b, 0x0f, 0xa0, 0x1f, 0x73, 0x36, 0x9e, 0xf8, 0x74, 0x4c, 0x79,
	0xa4, 0x15, 0x0e, 0xee, 0xc1, 0x36, 0xd3, 0x47, 0xc1, 0x89, 0x8f, 0x3a, 0x06, 0xd1, 0xcb, 0x0d,
	0xda, 0x36, 0x63, 0xf5, 0x10, 0x29, 0x68, 0x38, 0x09, 0x78, 0x48, 0xd1, 0x4e, 0xd5, 0xe0, 0x8c,
	0x07, 0xef, 0xb9, 0xa4, 0x42, 0x04, 0x02, 0x81, 0x7e, 0xaf, 0xf8, 0x3c, 0x0e, 0x22, 0x22, 0x7d,
	0x36, 0x66, 0x91, 0xa4, 0x97, 0xc7, 0x54, 0x37, 0xee, 0x1a, 0x29, 0xe3, 0x17, 0xc4, 0x67, 0x9e,
	0xb4, 0x75, 0xd4, 0xc3, 0xcf, 0x60, 0x7f, 0x42, 0xae, 0xcc, 0x68, 0xc9, 0xf8, 0x49, 0x20, 0xc6,
	0x24, 0x32, 0x6e, 0xf5, 0x6a, 0x17, 0x7a, 0x53, 0x72, 0xe4, 0x53, 0xd4, 0xc7, 0xfb, 0xf0, 0xa4,
	0xbe, 0xf7, 0x17, 0x31, 0xda, 0x35, 0x03, 0x23, 0x41, 0x78, 0x48, 0x8e, 0xed, 0xed, 0x4d, 0x48,
	0xff, 0x1b, 0xb7, 0x44, 0x50, 0x79, 0x15, 0xc4, 0x32, 0x8c, 0x05, 0x45, 0xe8, 0x8f, 0x05, 0xea,
	0x49, 0x8f, 0x68, 0x0b, 0x83, 0x6b, 0xc7, 0x3e, 0xf8, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x8c, 0xe8, 0x5a, 0xac, 0xb5, 0x03, 0x00, 0x00,
}
