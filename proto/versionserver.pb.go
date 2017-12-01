// Code generated by protoc-gen-go. DO NOT EDIT.
// source: versionserver.proto

/*
Package versionserver is a generated protocol buffer package.

It is generated from these files:
	versionserver.proto

It has these top-level messages:
	Version
	GetVersionRequest
	GetVersionResponse
	SetVersionRequest
	SetVersionResponse
*/
package versionserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Version struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Version) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Version) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetVersionRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetVersionRequest) Reset()                    { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()               {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetVersionRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetVersionResponse struct {
	Version *Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *GetVersionResponse) Reset()                    { *m = GetVersionResponse{} }
func (m *GetVersionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetVersionResponse) ProtoMessage()               {}
func (*GetVersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetVersionResponse) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type SetVersionRequest struct {
	Set *Version `protobuf:"bytes,1,opt,name=set" json:"set,omitempty"`
}

func (m *SetVersionRequest) Reset()                    { *m = SetVersionRequest{} }
func (m *SetVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*SetVersionRequest) ProtoMessage()               {}
func (*SetVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SetVersionRequest) GetSet() *Version {
	if m != nil {
		return m.Set
	}
	return nil
}

type SetVersionResponse struct {
	Response *Version `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *SetVersionResponse) Reset()                    { *m = SetVersionResponse{} }
func (m *SetVersionResponse) String() string            { return proto.CompactTextString(m) }
func (*SetVersionResponse) ProtoMessage()               {}
func (*SetVersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SetVersionResponse) GetResponse() *Version {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Version)(nil), "versionserver.Version")
	proto.RegisterType((*GetVersionRequest)(nil), "versionserver.GetVersionRequest")
	proto.RegisterType((*GetVersionResponse)(nil), "versionserver.GetVersionResponse")
	proto.RegisterType((*SetVersionRequest)(nil), "versionserver.SetVersionRequest")
	proto.RegisterType((*SetVersionResponse)(nil), "versionserver.SetVersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VersionServer service

type VersionServerClient interface {
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
	SetVersion(ctx context.Context, in *SetVersionRequest, opts ...grpc.CallOption) (*SetVersionResponse, error)
}

type versionServerClient struct {
	cc *grpc.ClientConn
}

func NewVersionServerClient(cc *grpc.ClientConn) VersionServerClient {
	return &versionServerClient{cc}
}

func (c *versionServerClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := grpc.Invoke(ctx, "/versionserver.VersionServer/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionServerClient) SetVersion(ctx context.Context, in *SetVersionRequest, opts ...grpc.CallOption) (*SetVersionResponse, error) {
	out := new(SetVersionResponse)
	err := grpc.Invoke(ctx, "/versionserver.VersionServer/SetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VersionServer service

type VersionServerServer interface {
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	SetVersion(context.Context, *SetVersionRequest) (*SetVersionResponse, error)
}

func RegisterVersionServerServer(s *grpc.Server, srv VersionServerServer) {
	s.RegisterService(&_VersionServer_serviceDesc, srv)
}

func _VersionServer_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServerServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versionserver.VersionServer/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServerServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionServer_SetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServerServer).SetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versionserver.VersionServer/SetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServerServer).SetVersion(ctx, req.(*SetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VersionServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "versionserver.VersionServer",
	HandlerType: (*VersionServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _VersionServer_GetVersion_Handler,
		},
		{
			MethodName: "SetVersion",
			Handler:    _VersionServer_SetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "versionserver.proto",
}

func init() { proto.RegisterFile("versionserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4b, 0x2d, 0x2a,
	0xce, 0xcc, 0xcf, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0x32, 0xe4, 0x62, 0x0f, 0x83, 0x08, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa7,
	0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89,
	0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x92, 0x2a, 0x97, 0xa0,
	0x7b, 0x6a, 0x09, 0x54, 0x57, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0xa6, 0x66, 0x25, 0x37,
	0x2e, 0x21, 0x64, 0x65, 0xc5, 0x05, 0x20, 0x3b, 0x85, 0x0c, 0xb8, 0xd8, 0xa1, 0x0e, 0x00, 0x1b,
	0xca, 0x6d, 0x24, 0xa6, 0x87, 0xea, 0x4a, 0x98, 0x06, 0x98, 0x32, 0x25, 0x5b, 0x2e, 0xc1, 0x60,
	0x0c, 0xeb, 0x34, 0xb8, 0x98, 0x8b, 0x53, 0x4b, 0xc0, 0xd6, 0xe1, 0x36, 0x02, 0xa4, 0x44, 0xc9,
	0x83, 0x4b, 0x28, 0x18, 0xd3, 0x19, 0x46, 0x5c, 0x1c, 0x45, 0x50, 0x36, 0x01, 0x43, 0xe0, 0xea,
	0x8c, 0x76, 0x32, 0x72, 0xf1, 0x42, 0x45, 0x83, 0xc1, 0x6a, 0x84, 0x82, 0xb9, 0xb8, 0x10, 0x5e,
	0x14, 0x52, 0x40, 0x33, 0x01, 0x23, 0x90, 0xa4, 0x14, 0xf1, 0xa8, 0x80, 0x58, 0xa2, 0xc4, 0x00,
	0x32, 0x34, 0x18, 0xb7, 0xa1, 0xc1, 0x04, 0x0d, 0x0d, 0xc6, 0x62, 0x68, 0x12, 0x1b, 0x38, 0xf2,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x75, 0x2c, 0xdb, 0x13, 0x02, 0x00, 0x00,
}