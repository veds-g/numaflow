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
// source: pkg/apis/proto/mvtxdaemon/mvtxdaemon.proto

package mvtxdaemon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MonoVertexDaemonService_GetMonoVertexMetrics_FullMethodName = "/mvtxdaemon.MonoVertexDaemonService/GetMonoVertexMetrics"
	MonoVertexDaemonService_GetMonoVertexStatus_FullMethodName  = "/mvtxdaemon.MonoVertexDaemonService/GetMonoVertexStatus"
	MonoVertexDaemonService_GetMonoVertexErrors_FullMethodName  = "/mvtxdaemon.MonoVertexDaemonService/GetMonoVertexErrors"
)

// MonoVertexDaemonServiceClient is the client API for MonoVertexDaemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MonoVertexDaemonService is a grpc service that is used to provide APIs for giving any MonoVertex information.
type MonoVertexDaemonServiceClient interface {
	GetMonoVertexMetrics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMonoVertexMetricsResponse, error)
	GetMonoVertexStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMonoVertexStatusResponse, error)
	GetMonoVertexErrors(ctx context.Context, in *GetMonoVertexErrorsRequest, opts ...grpc.CallOption) (*GetMonoVertexErrorsResponse, error)
}

type monoVertexDaemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMonoVertexDaemonServiceClient(cc grpc.ClientConnInterface) MonoVertexDaemonServiceClient {
	return &monoVertexDaemonServiceClient{cc}
}

func (c *monoVertexDaemonServiceClient) GetMonoVertexMetrics(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMonoVertexMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMonoVertexMetricsResponse)
	err := c.cc.Invoke(ctx, MonoVertexDaemonService_GetMonoVertexMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monoVertexDaemonServiceClient) GetMonoVertexStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMonoVertexStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMonoVertexStatusResponse)
	err := c.cc.Invoke(ctx, MonoVertexDaemonService_GetMonoVertexStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monoVertexDaemonServiceClient) GetMonoVertexErrors(ctx context.Context, in *GetMonoVertexErrorsRequest, opts ...grpc.CallOption) (*GetMonoVertexErrorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMonoVertexErrorsResponse)
	err := c.cc.Invoke(ctx, MonoVertexDaemonService_GetMonoVertexErrors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonoVertexDaemonServiceServer is the server API for MonoVertexDaemonService service.
// All implementations must embed UnimplementedMonoVertexDaemonServiceServer
// for forward compatibility
//
// MonoVertexDaemonService is a grpc service that is used to provide APIs for giving any MonoVertex information.
type MonoVertexDaemonServiceServer interface {
	GetMonoVertexMetrics(context.Context, *emptypb.Empty) (*GetMonoVertexMetricsResponse, error)
	GetMonoVertexStatus(context.Context, *emptypb.Empty) (*GetMonoVertexStatusResponse, error)
	GetMonoVertexErrors(context.Context, *GetMonoVertexErrorsRequest) (*GetMonoVertexErrorsResponse, error)
	mustEmbedUnimplementedMonoVertexDaemonServiceServer()
}

// UnimplementedMonoVertexDaemonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMonoVertexDaemonServiceServer struct {
}

func (UnimplementedMonoVertexDaemonServiceServer) GetMonoVertexMetrics(context.Context, *emptypb.Empty) (*GetMonoVertexMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonoVertexMetrics not implemented")
}
func (UnimplementedMonoVertexDaemonServiceServer) GetMonoVertexStatus(context.Context, *emptypb.Empty) (*GetMonoVertexStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonoVertexStatus not implemented")
}
func (UnimplementedMonoVertexDaemonServiceServer) GetMonoVertexErrors(context.Context, *GetMonoVertexErrorsRequest) (*GetMonoVertexErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonoVertexErrors not implemented")
}
func (UnimplementedMonoVertexDaemonServiceServer) mustEmbedUnimplementedMonoVertexDaemonServiceServer() {
}

// UnsafeMonoVertexDaemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonoVertexDaemonServiceServer will
// result in compilation errors.
type UnsafeMonoVertexDaemonServiceServer interface {
	mustEmbedUnimplementedMonoVertexDaemonServiceServer()
}

func RegisterMonoVertexDaemonServiceServer(s grpc.ServiceRegistrar, srv MonoVertexDaemonServiceServer) {
	s.RegisterService(&MonoVertexDaemonService_ServiceDesc, srv)
}

func _MonoVertexDaemonService_GetMonoVertexMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonoVertexDaemonServiceServer).GetMonoVertexMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonoVertexDaemonService_GetMonoVertexMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonoVertexDaemonServiceServer).GetMonoVertexMetrics(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonoVertexDaemonService_GetMonoVertexStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonoVertexDaemonServiceServer).GetMonoVertexStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonoVertexDaemonService_GetMonoVertexStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonoVertexDaemonServiceServer).GetMonoVertexStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonoVertexDaemonService_GetMonoVertexErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonoVertexErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonoVertexDaemonServiceServer).GetMonoVertexErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonoVertexDaemonService_GetMonoVertexErrors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonoVertexDaemonServiceServer).GetMonoVertexErrors(ctx, req.(*GetMonoVertexErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MonoVertexDaemonService_ServiceDesc is the grpc.ServiceDesc for MonoVertexDaemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonoVertexDaemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mvtxdaemon.MonoVertexDaemonService",
	HandlerType: (*MonoVertexDaemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMonoVertexMetrics",
			Handler:    _MonoVertexDaemonService_GetMonoVertexMetrics_Handler,
		},
		{
			MethodName: "GetMonoVertexStatus",
			Handler:    _MonoVertexDaemonService_GetMonoVertexStatus_Handler,
		},
		{
			MethodName: "GetMonoVertexErrors",
			Handler:    _MonoVertexDaemonService_GetMonoVertexErrors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apis/proto/mvtxdaemon/mvtxdaemon.proto",
}
