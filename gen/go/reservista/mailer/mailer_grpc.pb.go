// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: reservista/mailer/mailer.proto

package proto_mailer

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

// MailerClient is the client API for Mailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailerClient interface {
	SendWelcome(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error)
	SendQR(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error)
	SendAuthCode(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error)
	SendReminder(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error)
	SendAdvert(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error)
	SendResetCode(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error)
}

type mailerClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerClient(cc grpc.ClientConnInterface) MailerClient {
	return &mailerClient{cc}
}

func (c *mailerClient) SendWelcome(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/mailer.Mailer/SendWelcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerClient) SendQR(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/mailer.Mailer/SendQR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerClient) SendAuthCode(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/mailer.Mailer/SendAuthCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerClient) SendReminder(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/mailer.Mailer/SendReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerClient) SendAdvert(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/mailer.Mailer/SendAdvert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerClient) SendResetCode(ctx context.Context, in *EmailInput, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/mailer.Mailer/SendResetCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServer is the server API for Mailer service.
// All implementations must embed UnimplementedMailerServer
// for forward compatibility
type MailerServer interface {
	SendWelcome(context.Context, *EmailInput) (*StatusResponse, error)
	SendQR(context.Context, *EmailInput) (*StatusResponse, error)
	SendAuthCode(context.Context, *EmailInput) (*StatusResponse, error)
	SendReminder(context.Context, *EmailInput) (*StatusResponse, error)
	SendAdvert(context.Context, *EmailInput) (*StatusResponse, error)
	SendResetCode(context.Context, *EmailInput) (*StatusResponse, error)
	mustEmbedUnimplementedMailerServer()
}

// UnimplementedMailerServer must be embedded to have forward compatible implementations.
type UnimplementedMailerServer struct {
}

func (UnimplementedMailerServer) SendWelcome(context.Context, *EmailInput) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWelcome not implemented")
}
func (UnimplementedMailerServer) SendQR(context.Context, *EmailInput) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendQR not implemented")
}
func (UnimplementedMailerServer) SendAuthCode(context.Context, *EmailInput) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAuthCode not implemented")
}
func (UnimplementedMailerServer) SendReminder(context.Context, *EmailInput) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReminder not implemented")
}
func (UnimplementedMailerServer) SendAdvert(context.Context, *EmailInput) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAdvert not implemented")
}
func (UnimplementedMailerServer) SendResetCode(context.Context, *EmailInput) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResetCode not implemented")
}
func (UnimplementedMailerServer) mustEmbedUnimplementedMailerServer() {}

// UnsafeMailerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServer will
// result in compilation errors.
type UnsafeMailerServer interface {
	mustEmbedUnimplementedMailerServer()
}

func RegisterMailerServer(s grpc.ServiceRegistrar, srv MailerServer) {
	s.RegisterService(&Mailer_ServiceDesc, srv)
}

func _Mailer_SendWelcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendWelcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.Mailer/SendWelcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendWelcome(ctx, req.(*EmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailer_SendQR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendQR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.Mailer/SendQR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendQR(ctx, req.(*EmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailer_SendAuthCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendAuthCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.Mailer/SendAuthCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendAuthCode(ctx, req.(*EmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailer_SendReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.Mailer/SendReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendReminder(ctx, req.(*EmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailer_SendAdvert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendAdvert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.Mailer/SendAdvert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendAdvert(ctx, req.(*EmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailer_SendResetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendResetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.Mailer/SendResetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendResetCode(ctx, req.(*EmailInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Mailer_ServiceDesc is the grpc.ServiceDesc for Mailer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mailer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mailer.Mailer",
	HandlerType: (*MailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendWelcome",
			Handler:    _Mailer_SendWelcome_Handler,
		},
		{
			MethodName: "SendQR",
			Handler:    _Mailer_SendQR_Handler,
		},
		{
			MethodName: "SendAuthCode",
			Handler:    _Mailer_SendAuthCode_Handler,
		},
		{
			MethodName: "SendReminder",
			Handler:    _Mailer_SendReminder_Handler,
		},
		{
			MethodName: "SendAdvert",
			Handler:    _Mailer_SendAdvert_Handler,
		},
		{
			MethodName: "SendResetCode",
			Handler:    _Mailer_SendResetCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservista/mailer/mailer.proto",
}
