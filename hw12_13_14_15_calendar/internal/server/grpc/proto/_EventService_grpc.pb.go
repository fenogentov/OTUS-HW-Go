// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto_event

import (
	context "context"
//	"hw12_13_14_15_calendar/internal/storage"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarClient interface {
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	DeleteEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	GetEvents(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/proto_event.Calendar/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/proto_event.Calendar/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DeleteEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/proto_event.Calendar/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEvents(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/proto_event.Calendar/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
// All implementations must embed UnimplementedCalendarServer
// for forward compatibility
type CalendarServer interface {
	CreateEvent(context.Context, *Event) (*Event, error)
	UpdateEvent(context.Context, *Event) (*Event, error)
	DeleteEvent(context.Context, *Event) (*Event, error)
	GetEvents(context.Context, *Event) (*Event, error)
	mustEmbedUnimplementedCalendarServer()
}

// UnimplementedCalendarServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (UnimplementedCalendarServer) CreateEvent(context.Context, *Event) (*Event, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedCalendarServer) UpdateEvent(context.Context, *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedCalendarServer) DeleteEvent(context.Context, *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedCalendarServer) GetEvents(context.Context, *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedCalendarServer) mustEmbedUnimplementedCalendarServer() {}

// UnsafeCalendarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServer will
// result in compilation errors.
type UnsafeCalendarServer interface {
	mustEmbedUnimplementedCalendarServer()
}

func RegisterCalendarServer(s grpc.ServiceRegistrar, srv CalendarServer) {
	s.RegisterService(&Calendar_ServiceDesc, srv)
}

func _Calendar_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_event.Calendar/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_event.Calendar/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).UpdateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_event.Calendar/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DeleteEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_event.Calendar/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvents(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

// Calendar_ServiceDesc is the grpc.ServiceDesc for Calendar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calendar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_event.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Calendar_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Calendar_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Calendar_DeleteEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _Calendar_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/EventService.proto",
}
