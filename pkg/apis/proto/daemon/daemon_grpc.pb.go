//
//Copyright 2022 The Numaproj Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: pkg/apis/proto/daemon/daemon.proto

package daemon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DaemonService_ListBuffers_FullMethodName           = "/daemon.DaemonService/ListBuffers"
	DaemonService_GetBuffer_FullMethodName             = "/daemon.DaemonService/GetBuffer"
	DaemonService_GetVertexMetrics_FullMethodName      = "/daemon.DaemonService/GetVertexMetrics"
	DaemonService_GetPipelineWatermarks_FullMethodName = "/daemon.DaemonService/GetPipelineWatermarks"
	DaemonService_GetPipelineStatus_FullMethodName     = "/daemon.DaemonService/GetPipelineStatus"
	DaemonService_GetVertexErrors_FullMethodName       = "/daemon.DaemonService/GetVertexErrors"
)

// DaemonServiceClient is the client API for DaemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// DaemonService is a grpc service that is used to provide APIs for giving any pipeline information.
type DaemonServiceClient interface {
	ListBuffers(ctx context.Context, in *ListBuffersRequest, opts ...grpc.CallOption) (*ListBuffersResponse, error)
	GetBuffer(ctx context.Context, in *GetBufferRequest, opts ...grpc.CallOption) (*GetBufferResponse, error)
	GetVertexMetrics(ctx context.Context, in *GetVertexMetricsRequest, opts ...grpc.CallOption) (*GetVertexMetricsResponse, error)
	// GetPipelineWatermarks return the watermark of the given pipeline
	GetPipelineWatermarks(ctx context.Context, in *GetPipelineWatermarksRequest, opts ...grpc.CallOption) (*GetPipelineWatermarksResponse, error)
	GetPipelineStatus(ctx context.Context, in *GetPipelineStatusRequest, opts ...grpc.CallOption) (*GetPipelineStatusResponse, error)
	GetVertexErrors(ctx context.Context, in *GetVertexErrorsRequest, opts ...grpc.CallOption) (*GetVertexErrorsResponse, error)
}

type daemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonServiceClient(cc grpc.ClientConnInterface) DaemonServiceClient {
	return &daemonServiceClient{cc}
}

func (c *daemonServiceClient) ListBuffers(ctx context.Context, in *ListBuffersRequest, opts ...grpc.CallOption) (*ListBuffersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBuffersResponse)
	err := c.cc.Invoke(ctx, DaemonService_ListBuffers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetBuffer(ctx context.Context, in *GetBufferRequest, opts ...grpc.CallOption) (*GetBufferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBufferResponse)
	err := c.cc.Invoke(ctx, DaemonService_GetBuffer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetVertexMetrics(ctx context.Context, in *GetVertexMetricsRequest, opts ...grpc.CallOption) (*GetVertexMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVertexMetricsResponse)
	err := c.cc.Invoke(ctx, DaemonService_GetVertexMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetPipelineWatermarks(ctx context.Context, in *GetPipelineWatermarksRequest, opts ...grpc.CallOption) (*GetPipelineWatermarksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPipelineWatermarksResponse)
	err := c.cc.Invoke(ctx, DaemonService_GetPipelineWatermarks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetPipelineStatus(ctx context.Context, in *GetPipelineStatusRequest, opts ...grpc.CallOption) (*GetPipelineStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPipelineStatusResponse)
	err := c.cc.Invoke(ctx, DaemonService_GetPipelineStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetVertexErrors(ctx context.Context, in *GetVertexErrorsRequest, opts ...grpc.CallOption) (*GetVertexErrorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVertexErrorsResponse)
	err := c.cc.Invoke(ctx, DaemonService_GetVertexErrors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServiceServer is the server API for DaemonService service.
// All implementations must embed UnimplementedDaemonServiceServer
// for forward compatibility
//
// DaemonService is a grpc service that is used to provide APIs for giving any pipeline information.
type DaemonServiceServer interface {
	ListBuffers(context.Context, *ListBuffersRequest) (*ListBuffersResponse, error)
	GetBuffer(context.Context, *GetBufferRequest) (*GetBufferResponse, error)
	GetVertexMetrics(context.Context, *GetVertexMetricsRequest) (*GetVertexMetricsResponse, error)
	// GetPipelineWatermarks return the watermark of the given pipeline
	GetPipelineWatermarks(context.Context, *GetPipelineWatermarksRequest) (*GetPipelineWatermarksResponse, error)
	GetPipelineStatus(context.Context, *GetPipelineStatusRequest) (*GetPipelineStatusResponse, error)
	GetVertexErrors(context.Context, *GetVertexErrorsRequest) (*GetVertexErrorsResponse, error)
	mustEmbedUnimplementedDaemonServiceServer()
}

// UnimplementedDaemonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServiceServer struct {
}

func (UnimplementedDaemonServiceServer) ListBuffers(context.Context, *ListBuffersRequest) (*ListBuffersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuffers not implemented")
}
func (UnimplementedDaemonServiceServer) GetBuffer(context.Context, *GetBufferRequest) (*GetBufferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuffer not implemented")
}
func (UnimplementedDaemonServiceServer) GetVertexMetrics(context.Context, *GetVertexMetricsRequest) (*GetVertexMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVertexMetrics not implemented")
}
func (UnimplementedDaemonServiceServer) GetPipelineWatermarks(context.Context, *GetPipelineWatermarksRequest) (*GetPipelineWatermarksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineWatermarks not implemented")
}
func (UnimplementedDaemonServiceServer) GetPipelineStatus(context.Context, *GetPipelineStatusRequest) (*GetPipelineStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineStatus not implemented")
}
func (UnimplementedDaemonServiceServer) GetVertexErrors(context.Context, *GetVertexErrorsRequest) (*GetVertexErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVertexErrors not implemented")
}
func (UnimplementedDaemonServiceServer) mustEmbedUnimplementedDaemonServiceServer() {}

// UnsafeDaemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServiceServer will
// result in compilation errors.
type UnsafeDaemonServiceServer interface {
	mustEmbedUnimplementedDaemonServiceServer()
}

func RegisterDaemonServiceServer(s grpc.ServiceRegistrar, srv DaemonServiceServer) {
	s.RegisterService(&DaemonService_ServiceDesc, srv)
}

func _DaemonService_ListBuffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuffersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListBuffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DaemonService_ListBuffers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListBuffers(ctx, req.(*ListBuffersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBufferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DaemonService_GetBuffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetBuffer(ctx, req.(*GetBufferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetVertexMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVertexMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetVertexMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DaemonService_GetVertexMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetVertexMetrics(ctx, req.(*GetVertexMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetPipelineWatermarks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineWatermarksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetPipelineWatermarks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DaemonService_GetPipelineWatermarks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetPipelineWatermarks(ctx, req.(*GetPipelineWatermarksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetPipelineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetPipelineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DaemonService_GetPipelineStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetPipelineStatus(ctx, req.(*GetPipelineStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetVertexErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVertexErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetVertexErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DaemonService_GetVertexErrors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetVertexErrors(ctx, req.(*GetVertexErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DaemonService_ServiceDesc is the grpc.ServiceDesc for DaemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DaemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daemon.DaemonService",
	HandlerType: (*DaemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBuffers",
			Handler:    _DaemonService_ListBuffers_Handler,
		},
		{
			MethodName: "GetBuffer",
			Handler:    _DaemonService_GetBuffer_Handler,
		},
		{
			MethodName: "GetVertexMetrics",
			Handler:    _DaemonService_GetVertexMetrics_Handler,
		},
		{
			MethodName: "GetPipelineWatermarks",
			Handler:    _DaemonService_GetPipelineWatermarks_Handler,
		},
		{
			MethodName: "GetPipelineStatus",
			Handler:    _DaemonService_GetPipelineStatus_Handler,
		},
		{
			MethodName: "GetVertexErrors",
			Handler:    _DaemonService_GetVertexErrors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apis/proto/daemon/daemon.proto",
}
