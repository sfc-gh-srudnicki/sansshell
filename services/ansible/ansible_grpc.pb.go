// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ansible

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

// PlaybookClient is the client API for Playbook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaybookClient interface {
	// Will run ansible-playbook only on the local host using the args passed.
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunReply, error)
}

type playbookClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaybookClient(cc grpc.ClientConnInterface) PlaybookClient {
	return &playbookClient{cc}
}

func (c *playbookClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunReply, error) {
	out := new(RunReply)
	err := c.cc.Invoke(ctx, "/Ansible.Playbook/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaybookServer is the server API for Playbook service.
// All implementations should embed UnimplementedPlaybookServer
// for forward compatibility
type PlaybookServer interface {
	// Will run ansible-playbook only on the local host using the args passed.
	Run(context.Context, *RunRequest) (*RunReply, error)
}

// UnimplementedPlaybookServer should be embedded to have forward compatible implementations.
type UnimplementedPlaybookServer struct {
}

func (UnimplementedPlaybookServer) Run(context.Context, *RunRequest) (*RunReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

// UnsafePlaybookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaybookServer will
// result in compilation errors.
type UnsafePlaybookServer interface {
	mustEmbedUnimplementedPlaybookServer()
}

func RegisterPlaybookServer(s grpc.ServiceRegistrar, srv PlaybookServer) {
	s.RegisterService(&Playbook_ServiceDesc, srv)
}

func _Playbook_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaybookServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ansible.Playbook/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaybookServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Playbook_ServiceDesc is the grpc.ServiceDesc for Playbook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Playbook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ansible.Playbook",
	HandlerType: (*PlaybookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Playbook_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ansible.proto",
}