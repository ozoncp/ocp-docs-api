// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_docs_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OcpDocsApiClient is the client API for OcpDocsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpDocsApiClient interface {
	ListDocsV1(ctx context.Context, in *ListDocsV1Request, opts ...grpc.CallOption) (*ListDocsV1Response, error)
	DescribeDocV1(ctx context.Context, in *DescribeDocV1Request, opts ...grpc.CallOption) (*DescribeDocV1Response, error)
	CreateDocV1(ctx context.Context, in *CreateDocV1Request, opts ...grpc.CallOption) (*CreateDocV1Response, error)
	RemoveDocV1(ctx context.Context, in *RemoveDocV1Request, opts ...grpc.CallOption) (*RemoveDocV1Response, error)
}

type ocpDocsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpDocsApiClient(cc grpc.ClientConnInterface) OcpDocsApiClient {
	return &ocpDocsApiClient{cc}
}

func (c *ocpDocsApiClient) ListDocsV1(ctx context.Context, in *ListDocsV1Request, opts ...grpc.CallOption) (*ListDocsV1Response, error) {
	out := new(ListDocsV1Response)
	err := c.cc.Invoke(ctx, "/ocp.docs.api.OcpDocsApi/ListDocsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpDocsApiClient) DescribeDocV1(ctx context.Context, in *DescribeDocV1Request, opts ...grpc.CallOption) (*DescribeDocV1Response, error) {
	out := new(DescribeDocV1Response)
	err := c.cc.Invoke(ctx, "/ocp.docs.api.OcpDocsApi/DescribeDocV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpDocsApiClient) CreateDocV1(ctx context.Context, in *CreateDocV1Request, opts ...grpc.CallOption) (*CreateDocV1Response, error) {
	out := new(CreateDocV1Response)
	err := c.cc.Invoke(ctx, "/ocp.docs.api.OcpDocsApi/CreateDocV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpDocsApiClient) RemoveDocV1(ctx context.Context, in *RemoveDocV1Request, opts ...grpc.CallOption) (*RemoveDocV1Response, error) {
	out := new(RemoveDocV1Response)
	err := c.cc.Invoke(ctx, "/ocp.docs.api.OcpDocsApi/RemoveDocV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpDocsApiServer is the server API for OcpDocsApi service.
// All implementations must embed UnimplementedOcpDocsApiServer
// for forward compatibility
type OcpDocsApiServer interface {
	ListDocsV1(context.Context, *ListDocsV1Request) (*ListDocsV1Response, error)
	DescribeDocV1(context.Context, *DescribeDocV1Request) (*DescribeDocV1Response, error)
	CreateDocV1(context.Context, *CreateDocV1Request) (*CreateDocV1Response, error)
	RemoveDocV1(context.Context, *RemoveDocV1Request) (*RemoveDocV1Response, error)
	mustEmbedUnimplementedOcpDocsApiServer()
}

// UnimplementedOcpDocsApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpDocsApiServer struct {
}

func (UnimplementedOcpDocsApiServer) ListDocsV1(context.Context, *ListDocsV1Request) (*ListDocsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocsV1 not implemented")
}
func (UnimplementedOcpDocsApiServer) DescribeDocV1(context.Context, *DescribeDocV1Request) (*DescribeDocV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDocV1 not implemented")
}
func (UnimplementedOcpDocsApiServer) CreateDocV1(context.Context, *CreateDocV1Request) (*CreateDocV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocV1 not implemented")
}
func (UnimplementedOcpDocsApiServer) RemoveDocV1(context.Context, *RemoveDocV1Request) (*RemoveDocV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDocV1 not implemented")
}
func (UnimplementedOcpDocsApiServer) mustEmbedUnimplementedOcpDocsApiServer() {}

// UnsafeOcpDocsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpDocsApiServer will
// result in compilation errors.
type UnsafeOcpDocsApiServer interface {
	mustEmbedUnimplementedOcpDocsApiServer()
}

func RegisterOcpDocsApiServer(s grpc.ServiceRegistrar, srv OcpDocsApiServer) {
	s.RegisterService(&OcpDocsApi_ServiceDesc, srv)
}

func _OcpDocsApi_ListDocsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpDocsApiServer).ListDocsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.docs.api.OcpDocsApi/ListDocsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpDocsApiServer).ListDocsV1(ctx, req.(*ListDocsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpDocsApi_DescribeDocV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDocV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpDocsApiServer).DescribeDocV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.docs.api.OcpDocsApi/DescribeDocV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpDocsApiServer).DescribeDocV1(ctx, req.(*DescribeDocV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpDocsApi_CreateDocV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpDocsApiServer).CreateDocV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.docs.api.OcpDocsApi/CreateDocV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpDocsApiServer).CreateDocV1(ctx, req.(*CreateDocV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpDocsApi_RemoveDocV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDocV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpDocsApiServer).RemoveDocV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.docs.api.OcpDocsApi/RemoveDocV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpDocsApiServer).RemoveDocV1(ctx, req.(*RemoveDocV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpDocsApi_ServiceDesc is the grpc.ServiceDesc for OcpDocsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpDocsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.docs.api.OcpDocsApi",
	HandlerType: (*OcpDocsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDocsV1",
			Handler:    _OcpDocsApi_ListDocsV1_Handler,
		},
		{
			MethodName: "DescribeDocV1",
			Handler:    _OcpDocsApi_DescribeDocV1_Handler,
		},
		{
			MethodName: "CreateDocV1",
			Handler:    _OcpDocsApi_CreateDocV1_Handler,
		},
		{
			MethodName: "RemoveDocV1",
			Handler:    _OcpDocsApi_RemoveDocV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-docs-api/ocp-docs-api.proto",
}
