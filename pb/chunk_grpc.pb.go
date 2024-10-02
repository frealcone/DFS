// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc1
// source: chunk.proto

package pb

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
	Chunk_Read_FullMethodName         = "/pb.Chunk/Read"
	Chunk_PrimaryWrite_FullMethodName = "/pb.Chunk/PrimaryWrite"
	Chunk_Sync_FullMethodName         = "/pb.Chunk/Sync"
)

// ChunkClient is the client API for Chunk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunkClient interface {
	Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadResp, error)
	PrimaryWrite(ctx context.Context, in *PrimaryWriteReq, opts ...grpc.CallOption) (*PrimaryWriteResp, error)
	Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*SyncResp, error)
}

type chunkClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkClient(cc grpc.ClientConnInterface) ChunkClient {
	return &chunkClient{cc}
}

func (c *chunkClient) Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadResp, error) {
	out := new(ReadResp)
	err := c.cc.Invoke(ctx, Chunk_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkClient) PrimaryWrite(ctx context.Context, in *PrimaryWriteReq, opts ...grpc.CallOption) (*PrimaryWriteResp, error) {
	out := new(PrimaryWriteResp)
	err := c.cc.Invoke(ctx, Chunk_PrimaryWrite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkClient) Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, Chunk_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServer is the server API for Chunk service.
// All implementations must embed UnimplementedChunkServer
// for forward compatibility
type ChunkServer interface {
	Read(context.Context, *ReadReq) (*ReadResp, error)
	PrimaryWrite(context.Context, *PrimaryWriteReq) (*PrimaryWriteResp, error)
	Sync(context.Context, *SyncReq) (*SyncResp, error)
	mustEmbedUnimplementedChunkServer()
}

// UnimplementedChunkServer must be embedded to have forward compatible implementations.
type UnimplementedChunkServer struct {
}

func (UnimplementedChunkServer) Read(context.Context, *ReadReq) (*ReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedChunkServer) PrimaryWrite(context.Context, *PrimaryWriteReq) (*PrimaryWriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrimaryWrite not implemented")
}
func (UnimplementedChunkServer) Sync(context.Context, *SyncReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedChunkServer) mustEmbedUnimplementedChunkServer() {}

// UnsafeChunkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkServer will
// result in compilation errors.
type UnsafeChunkServer interface {
	mustEmbedUnimplementedChunkServer()
}

func RegisterChunkServer(s grpc.ServiceRegistrar, srv ChunkServer) {
	s.RegisterService(&Chunk_ServiceDesc, srv)
}

func _Chunk_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chunk_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServer).Read(ctx, req.(*ReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chunk_PrimaryWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryWriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServer).PrimaryWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chunk_PrimaryWrite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServer).PrimaryWrite(ctx, req.(*PrimaryWriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chunk_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chunk_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServer).Sync(ctx, req.(*SyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Chunk_ServiceDesc is the grpc.ServiceDesc for Chunk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chunk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chunk",
	HandlerType: (*ChunkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Chunk_Read_Handler,
		},
		{
			MethodName: "PrimaryWrite",
			Handler:    _Chunk_PrimaryWrite_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Chunk_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chunk.proto",
}
