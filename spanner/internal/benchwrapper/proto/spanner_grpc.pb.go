// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spanner_bench

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

// SpannerBenchWrapperClient is the client API for SpannerBenchWrapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpannerBenchWrapperClient interface {
	// Read represents operations like Go's ReadOnlyTransaction.Query, Java's
	// ReadOnlyTransaction.executeQuery, Python's snapshot.read, and Node's
	// Transaction.Read.
	//
	// It will typically be used to read many items.
	Read(ctx context.Context, in *ReadQuery, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Insert represents operations like Go's Client.Apply, Java's
	// DatabaseClient.writeAtLeastOnce, Python's transaction.commit, and Node's
	// Transaction.Commit.
	//
	// It will typically be used to insert many items.
	Insert(ctx context.Context, in *InsertQuery, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Update represents operations like Go's ReadWriteTransaction.BatchUpdate,
	// Java's TransactionRunner.run, Python's Batch.update, and Node's
	// Transaction.BatchUpdate.
	//
	// It will typically be used to update many items.
	Update(ctx context.Context, in *UpdateQuery, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type spannerBenchWrapperClient struct {
	cc grpc.ClientConnInterface
}

func NewSpannerBenchWrapperClient(cc grpc.ClientConnInterface) SpannerBenchWrapperClient {
	return &spannerBenchWrapperClient{cc}
}

func (c *spannerBenchWrapperClient) Read(ctx context.Context, in *ReadQuery, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/spanner_bench.SpannerBenchWrapper/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spannerBenchWrapperClient) Insert(ctx context.Context, in *InsertQuery, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/spanner_bench.SpannerBenchWrapper/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spannerBenchWrapperClient) Update(ctx context.Context, in *UpdateQuery, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/spanner_bench.SpannerBenchWrapper/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpannerBenchWrapperServer is the server API for SpannerBenchWrapper service.
// All implementations must embed UnimplementedSpannerBenchWrapperServer
// for forward compatibility
type SpannerBenchWrapperServer interface {
	// Read represents operations like Go's ReadOnlyTransaction.Query, Java's
	// ReadOnlyTransaction.executeQuery, Python's snapshot.read, and Node's
	// Transaction.Read.
	//
	// It will typically be used to read many items.
	Read(context.Context, *ReadQuery) (*EmptyResponse, error)
	// Insert represents operations like Go's Client.Apply, Java's
	// DatabaseClient.writeAtLeastOnce, Python's transaction.commit, and Node's
	// Transaction.Commit.
	//
	// It will typically be used to insert many items.
	Insert(context.Context, *InsertQuery) (*EmptyResponse, error)
	// Update represents operations like Go's ReadWriteTransaction.BatchUpdate,
	// Java's TransactionRunner.run, Python's Batch.update, and Node's
	// Transaction.BatchUpdate.
	//
	// It will typically be used to update many items.
	Update(context.Context, *UpdateQuery) (*EmptyResponse, error)
	mustEmbedUnimplementedSpannerBenchWrapperServer()
}

// UnimplementedSpannerBenchWrapperServer must be embedded to have forward compatible implementations.
type UnimplementedSpannerBenchWrapperServer struct {
}

func (UnimplementedSpannerBenchWrapperServer) Read(context.Context, *ReadQuery) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedSpannerBenchWrapperServer) Insert(context.Context, *InsertQuery) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedSpannerBenchWrapperServer) Update(context.Context, *UpdateQuery) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSpannerBenchWrapperServer) mustEmbedUnimplementedSpannerBenchWrapperServer() {}

// UnsafeSpannerBenchWrapperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpannerBenchWrapperServer will
// result in compilation errors.
type UnsafeSpannerBenchWrapperServer interface {
	mustEmbedUnimplementedSpannerBenchWrapperServer()
}

func RegisterSpannerBenchWrapperServer(s grpc.ServiceRegistrar, srv SpannerBenchWrapperServer) {
	s.RegisterService(&SpannerBenchWrapper_ServiceDesc, srv)
}

func _SpannerBenchWrapper_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpannerBenchWrapperServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spanner_bench.SpannerBenchWrapper/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpannerBenchWrapperServer).Read(ctx, req.(*ReadQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpannerBenchWrapper_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpannerBenchWrapperServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spanner_bench.SpannerBenchWrapper/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpannerBenchWrapperServer).Insert(ctx, req.(*InsertQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpannerBenchWrapper_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpannerBenchWrapperServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spanner_bench.SpannerBenchWrapper/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpannerBenchWrapperServer).Update(ctx, req.(*UpdateQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// SpannerBenchWrapper_ServiceDesc is the grpc.ServiceDesc for SpannerBenchWrapper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpannerBenchWrapper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spanner_bench.SpannerBenchWrapper",
	HandlerType: (*SpannerBenchWrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _SpannerBenchWrapper_Read_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _SpannerBenchWrapper_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SpannerBenchWrapper_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spanner.proto",
}