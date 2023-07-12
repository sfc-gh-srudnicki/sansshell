// Copyright (c) 2022 Snowflake Inc. All rights reserved.
//
//Licensed under the Apache License, Version 2.0 (the
//"License"); you may not use this file except in compliance
//with the License.  You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing,
//software distributed under the License is distributed on an
//"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//KIND, either express or implied.  See the License for the
//specific language governing permissions and limitations
//under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: httpoverrpc.proto

package httpoverrpc

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

const (
	HTTPOverRPC_Localhost_FullMethodName = "/HTTPOverRPC.HTTPOverRPC/Localhost"
)

// HTTPOverRPCClient is the client API for HTTPOverRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPOverRPCClient interface {
	// Make an HTTP call to localhost
	Localhost(ctx context.Context, in *LocalhostHTTPRequest, opts ...grpc.CallOption) (*HTTPReply, error)
}

type hTTPOverRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPOverRPCClient(cc grpc.ClientConnInterface) HTTPOverRPCClient {
	return &hTTPOverRPCClient{cc}
}

func (c *hTTPOverRPCClient) Localhost(ctx context.Context, in *LocalhostHTTPRequest, opts ...grpc.CallOption) (*HTTPReply, error) {
	out := new(HTTPReply)
	err := c.cc.Invoke(ctx, HTTPOverRPC_Localhost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPOverRPCServer is the server API for HTTPOverRPC service.
// All implementations should embed UnimplementedHTTPOverRPCServer
// for forward compatibility
type HTTPOverRPCServer interface {
	// Make an HTTP call to localhost
	Localhost(context.Context, *LocalhostHTTPRequest) (*HTTPReply, error)
}

// UnimplementedHTTPOverRPCServer should be embedded to have forward compatible implementations.
type UnimplementedHTTPOverRPCServer struct {
}

func (UnimplementedHTTPOverRPCServer) Localhost(context.Context, *LocalhostHTTPRequest) (*HTTPReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Localhost not implemented")
}

// UnsafeHTTPOverRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPOverRPCServer will
// result in compilation errors.
type UnsafeHTTPOverRPCServer interface {
	mustEmbedUnimplementedHTTPOverRPCServer()
}

func RegisterHTTPOverRPCServer(s grpc.ServiceRegistrar, srv HTTPOverRPCServer) {
	s.RegisterService(&HTTPOverRPC_ServiceDesc, srv)
}

func _HTTPOverRPC_Localhost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalhostHTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPOverRPCServer).Localhost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPOverRPC_Localhost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPOverRPCServer).Localhost(ctx, req.(*LocalhostHTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPOverRPC_ServiceDesc is the grpc.ServiceDesc for HTTPOverRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPOverRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HTTPOverRPC.HTTPOverRPC",
	HandlerType: (*HTTPOverRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Localhost",
			Handler:    _HTTPOverRPC_Localhost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httpoverrpc.proto",
}
