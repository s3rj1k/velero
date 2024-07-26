// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: BackupItemAction.proto

package generated

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
	BackupItemAction_AppliesTo_FullMethodName = "/generated.BackupItemAction/AppliesTo"
	BackupItemAction_Execute_FullMethodName   = "/generated.BackupItemAction/Execute"
)

// BackupItemActionClient is the client API for BackupItemAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupItemActionClient interface {
	AppliesTo(ctx context.Context, in *BackupItemActionAppliesToRequest, opts ...grpc.CallOption) (*BackupItemActionAppliesToResponse, error)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type backupItemActionClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupItemActionClient(cc grpc.ClientConnInterface) BackupItemActionClient {
	return &backupItemActionClient{cc}
}

func (c *backupItemActionClient) AppliesTo(ctx context.Context, in *BackupItemActionAppliesToRequest, opts ...grpc.CallOption) (*BackupItemActionAppliesToResponse, error) {
	out := new(BackupItemActionAppliesToResponse)
	err := c.cc.Invoke(ctx, BackupItemAction_AppliesTo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupItemActionClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, BackupItemAction_Execute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupItemActionServer is the server API for BackupItemAction service.
// All implementations should embed UnimplementedBackupItemActionServer
// for forward compatibility
type BackupItemActionServer interface {
	AppliesTo(context.Context, *BackupItemActionAppliesToRequest) (*BackupItemActionAppliesToResponse, error)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

// UnimplementedBackupItemActionServer should be embedded to have forward compatible implementations.
type UnimplementedBackupItemActionServer struct {
}

func (UnimplementedBackupItemActionServer) AppliesTo(context.Context, *BackupItemActionAppliesToRequest) (*BackupItemActionAppliesToResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliesTo not implemented")
}
func (UnimplementedBackupItemActionServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeBackupItemActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupItemActionServer will
// result in compilation errors.
type UnsafeBackupItemActionServer interface {
	mustEmbedUnimplementedBackupItemActionServer()
}

func RegisterBackupItemActionServer(s grpc.ServiceRegistrar, srv BackupItemActionServer) {
	s.RegisterService(&BackupItemAction_ServiceDesc, srv)
}

func _BackupItemAction_AppliesTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupItemActionAppliesToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).AppliesTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupItemAction_AppliesTo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).AppliesTo(ctx, req.(*BackupItemActionAppliesToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupItemAction_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupItemActionServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupItemAction_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupItemActionServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BackupItemAction_ServiceDesc is the grpc.ServiceDesc for BackupItemAction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackupItemAction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generated.BackupItemAction",
	HandlerType: (*BackupItemActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppliesTo",
			Handler:    _BackupItemAction_AppliesTo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _BackupItemAction_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BackupItemAction.proto",
}