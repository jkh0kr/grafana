// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: resource.proto

package api

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
	ResourceStore_GetResource_FullMethodName = "/api.ResourceStore/GetResource"
	ResourceStore_Create_FullMethodName      = "/api.ResourceStore/Create"
	ResourceStore_Update_FullMethodName      = "/api.ResourceStore/Update"
	ResourceStore_Delete_FullMethodName      = "/api.ResourceStore/Delete"
	ResourceStore_List_FullMethodName        = "/api.ResourceStore/List"
	ResourceStore_Watch_FullMethodName       = "/api.ResourceStore/Watch"
	ResourceStore_GetBlob_FullMethodName     = "/api.ResourceStore/GetBlob"
	ResourceStore_History_FullMethodName     = "/api.ResourceStore/History"
	ResourceStore_Origin_FullMethodName      = "/api.ResourceStore/Origin"
	ResourceStore_IsHealthy_FullMethodName   = "/api.ResourceStore/IsHealthy"
)

// ResourceStoreClient is the client API for ResourceStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The entity store provides a basic CRUD+watch for any resource
type ResourceStoreClient interface {
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ResourceStore_WatchClient, error)
	// Get the raw blob bytes and metadata
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error)
	// Show resource history (and trash)
	History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error)
	// Used for efficient provisioning
	Origin(ctx context.Context, in *OriginRequest, opts ...grpc.CallOption) (*OriginResponse, error)
	// Check if the service is healthy
	IsHealthy(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type resourceStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceStoreClient(cc grpc.ClientConnInterface) ResourceStoreClient {
	return &resourceStoreClient{cc}
}

func (c *resourceStoreClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResourceResponse)
	err := c.cc.Invoke(ctx, ResourceStore_GetResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, ResourceStore_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ResourceStore_WatchClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResourceStore_ServiceDesc.Streams[0], ResourceStore_Watch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &resourceStoreWatchClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResourceStore_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type resourceStoreWatchClient struct {
	grpc.ClientStream
}

func (x *resourceStoreWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resourceStoreClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlobResponse)
	err := c.cc.Invoke(ctx, ResourceStore_GetBlob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryResponse)
	err := c.cc.Invoke(ctx, ResourceStore_History_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Origin(ctx context.Context, in *OriginRequest, opts ...grpc.CallOption) (*OriginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OriginResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Origin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) IsHealthy(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, ResourceStore_IsHealthy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceStoreServer is the server API for ResourceStore service.
// All implementations should embed UnimplementedResourceStoreServer
// for forward compatibility
//
// The entity store provides a basic CRUD+watch for any resource
type ResourceStoreServer interface {
	GetResource(context.Context, *GetResourceRequest) (*GetResourceResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Watch(*WatchRequest, ResourceStore_WatchServer) error
	// Get the raw blob bytes and metadata
	GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error)
	// Show resource history (and trash)
	History(context.Context, *HistoryRequest) (*HistoryResponse, error)
	// Used for efficient provisioning
	Origin(context.Context, *OriginRequest) (*OriginResponse, error)
	// Check if the service is healthy
	IsHealthy(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// UnimplementedResourceStoreServer should be embedded to have forward compatible implementations.
type UnimplementedResourceStoreServer struct {
}

func (UnimplementedResourceStoreServer) GetResource(context.Context, *GetResourceRequest) (*GetResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (UnimplementedResourceStoreServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResourceStoreServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResourceStoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResourceStoreServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedResourceStoreServer) Watch(*WatchRequest, ResourceStore_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedResourceStoreServer) GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlob not implemented")
}
func (UnimplementedResourceStoreServer) History(context.Context, *HistoryRequest) (*HistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (UnimplementedResourceStoreServer) Origin(context.Context, *OriginRequest) (*OriginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Origin not implemented")
}
func (UnimplementedResourceStoreServer) IsHealthy(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHealthy not implemented")
}

// UnsafeResourceStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceStoreServer will
// result in compilation errors.
type UnsafeResourceStoreServer interface {
	mustEmbedUnimplementedResourceStoreServer()
}

func RegisterResourceStoreServer(s grpc.ServiceRegistrar, srv ResourceStoreServer) {
	s.RegisterService(&ResourceStore_ServiceDesc, srv)
}

func _ResourceStore_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_GetResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceStoreServer).Watch(m, &resourceStoreWatchServer{ServerStream: stream})
}

type ResourceStore_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type resourceStoreWatchServer struct {
	grpc.ServerStream
}

func (x *resourceStoreWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ResourceStore_GetBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).GetBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_GetBlob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).GetBlob(ctx, req.(*GetBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_History_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).History(ctx, req.(*HistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Origin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Origin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Origin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Origin(ctx, req.(*OriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_IsHealthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).IsHealthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_IsHealthy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).IsHealthy(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceStore_ServiceDesc is the grpc.ServiceDesc for ResourceStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ResourceStore",
	HandlerType: (*ResourceStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _ResourceStore_GetResource_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResourceStore_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceStore_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceStore_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResourceStore_List_Handler,
		},
		{
			MethodName: "GetBlob",
			Handler:    _ResourceStore_GetBlob_Handler,
		},
		{
			MethodName: "History",
			Handler:    _ResourceStore_History_Handler,
		},
		{
			MethodName: "Origin",
			Handler:    _ResourceStore_Origin_Handler,
		},
		{
			MethodName: "IsHealthy",
			Handler:    _ResourceStore_IsHealthy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ResourceStore_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "resource.proto",
}
