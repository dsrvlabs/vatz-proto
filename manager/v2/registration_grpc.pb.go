// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/vatz/manager/v2/registration.proto

package v2

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
	RegistrationService_RegisterPlugin_FullMethodName = "/vatz.manager.RegistrationService/RegisterPlugin"
)

// RegistrationServiceClient is the client API for RegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationServiceClient interface {
	RegisterPlugin(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationServiceClient(cc grpc.ClientConnInterface) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) RegisterPlugin(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, RegistrationService_RegisterPlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServiceServer is the server API for RegistrationService service.
// All implementations must embed UnimplementedRegistrationServiceServer
// for forward compatibility
type RegistrationServiceServer interface {
	RegisterPlugin(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedRegistrationServiceServer()
}

// UnimplementedRegistrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationServiceServer struct {
}

func (UnimplementedRegistrationServiceServer) RegisterPlugin(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPlugin not implemented")
}
func (UnimplementedRegistrationServiceServer) mustEmbedUnimplementedRegistrationServiceServer() {}

// UnsafeRegistrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServiceServer will
// result in compilation errors.
type UnsafeRegistrationServiceServer interface {
	mustEmbedUnimplementedRegistrationServiceServer()
}

func RegisterRegistrationServiceServer(s grpc.ServiceRegistrar, srv RegistrationServiceServer) {
	s.RegisterService(&RegistrationService_ServiceDesc, srv)
}

func _RegistrationService_RegisterPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).RegisterPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegistrationService_RegisterPlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).RegisterPlugin(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistrationService_ServiceDesc is the grpc.ServiceDesc for RegistrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vatz.manager.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPlugin",
			Handler:    _RegistrationService_RegisterPlugin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vatz/manager/v2/registration.proto",
}
