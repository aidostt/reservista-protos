// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: reservista/reservation/reservation.proto

package proto_reservation

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

// ReservationClient is the client API for Reservation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationClient interface {
	MakeReservation(ctx context.Context, in *ReservationSQLRequest, opts ...grpc.CallOption) (*IDRequest, error)
	GetReservation(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ReservationObject, error)
	DeleteReservationById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetAllReservationByUserId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ReservationListResponse, error)
	GetAllReservationByRestaurantId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ReservationListResponse, error)
	UpdateReservation(ctx context.Context, in *ReservationSQLRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	GetRestaurantByReservationId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*RestaurantObject, error)
	GetTableByReservationId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableObject, error)
}

type reservationClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationClient(cc grpc.ClientConnInterface) ReservationClient {
	return &reservationClient{cc}
}

func (c *reservationClient) MakeReservation(ctx context.Context, in *ReservationSQLRequest, opts ...grpc.CallOption) (*IDRequest, error) {
	out := new(IDRequest)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/MakeReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) GetReservation(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ReservationObject, error) {
	out := new(ReservationObject)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/GetReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) DeleteReservationById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/DeleteReservationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) GetAllReservationByUserId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ReservationListResponse, error) {
	out := new(ReservationListResponse)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/GetAllReservationByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) GetAllReservationByRestaurantId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ReservationListResponse, error) {
	out := new(ReservationListResponse)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/GetAllReservationByRestaurantId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) UpdateReservation(ctx context.Context, in *ReservationSQLRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/UpdateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) GetRestaurantByReservationId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*RestaurantObject, error) {
	out := new(RestaurantObject)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/GetRestaurantByReservationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) GetTableByReservationId(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*TableObject, error) {
	out := new(TableObject)
	err := c.cc.Invoke(ctx, "/reservation.Reservation/GetTableByReservationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServer is the server API for Reservation service.
// All implementations must embed UnimplementedReservationServer
// for forward compatibility
type ReservationServer interface {
	MakeReservation(context.Context, *ReservationSQLRequest) (*IDRequest, error)
	GetReservation(context.Context, *IDRequest) (*ReservationObject, error)
	DeleteReservationById(context.Context, *IDRequest) (*StatusResponse, error)
	GetAllReservationByUserId(context.Context, *IDRequest) (*ReservationListResponse, error)
	GetAllReservationByRestaurantId(context.Context, *IDRequest) (*ReservationListResponse, error)
	UpdateReservation(context.Context, *ReservationSQLRequest) (*StatusResponse, error)
	GetRestaurantByReservationId(context.Context, *IDRequest) (*RestaurantObject, error)
	GetTableByReservationId(context.Context, *IDRequest) (*TableObject, error)
	mustEmbedUnimplementedReservationServer()
}

// UnimplementedReservationServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServer struct {
}

func (UnimplementedReservationServer) MakeReservation(context.Context, *ReservationSQLRequest) (*IDRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeReservation not implemented")
}
func (UnimplementedReservationServer) GetReservation(context.Context, *IDRequest) (*ReservationObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservation not implemented")
}
func (UnimplementedReservationServer) DeleteReservationById(context.Context, *IDRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservationById not implemented")
}
func (UnimplementedReservationServer) GetAllReservationByUserId(context.Context, *IDRequest) (*ReservationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservationByUserId not implemented")
}
func (UnimplementedReservationServer) GetAllReservationByRestaurantId(context.Context, *IDRequest) (*ReservationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservationByRestaurantId not implemented")
}
func (UnimplementedReservationServer) UpdateReservation(context.Context, *ReservationSQLRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (UnimplementedReservationServer) GetRestaurantByReservationId(context.Context, *IDRequest) (*RestaurantObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantByReservationId not implemented")
}
func (UnimplementedReservationServer) GetTableByReservationId(context.Context, *IDRequest) (*TableObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTableByReservationId not implemented")
}
func (UnimplementedReservationServer) mustEmbedUnimplementedReservationServer() {}

// UnsafeReservationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServer will
// result in compilation errors.
type UnsafeReservationServer interface {
	mustEmbedUnimplementedReservationServer()
}

func RegisterReservationServer(s grpc.ServiceRegistrar, srv ReservationServer) {
	s.RegisterService(&Reservation_ServiceDesc, srv)
}

func _Reservation_MakeReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationSQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).MakeReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/MakeReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).MakeReservation(ctx, req.(*ReservationSQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_GetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).GetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/GetReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).GetReservation(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_DeleteReservationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).DeleteReservationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/DeleteReservationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).DeleteReservationById(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_GetAllReservationByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).GetAllReservationByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/GetAllReservationByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).GetAllReservationByUserId(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_GetAllReservationByRestaurantId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).GetAllReservationByRestaurantId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/GetAllReservationByRestaurantId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).GetAllReservationByRestaurantId(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationSQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/UpdateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).UpdateReservation(ctx, req.(*ReservationSQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_GetRestaurantByReservationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).GetRestaurantByReservationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/GetRestaurantByReservationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).GetRestaurantByReservationId(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_GetTableByReservationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).GetTableByReservationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/GetTableByReservationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).GetTableByReservationId(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reservation_ServiceDesc is the grpc.ServiceDesc for Reservation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reservation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.Reservation",
	HandlerType: (*ReservationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeReservation",
			Handler:    _Reservation_MakeReservation_Handler,
		},
		{
			MethodName: "GetReservation",
			Handler:    _Reservation_GetReservation_Handler,
		},
		{
			MethodName: "DeleteReservationById",
			Handler:    _Reservation_DeleteReservationById_Handler,
		},
		{
			MethodName: "GetAllReservationByUserId",
			Handler:    _Reservation_GetAllReservationByUserId_Handler,
		},
		{
			MethodName: "GetAllReservationByRestaurantId",
			Handler:    _Reservation_GetAllReservationByRestaurantId_Handler,
		},
		{
			MethodName: "UpdateReservation",
			Handler:    _Reservation_UpdateReservation_Handler,
		},
		{
			MethodName: "GetRestaurantByReservationId",
			Handler:    _Reservation_GetRestaurantByReservationId_Handler,
		},
		{
			MethodName: "GetTableByReservationId",
			Handler:    _Reservation_GetTableByReservationId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservista/reservation/reservation.proto",
}
