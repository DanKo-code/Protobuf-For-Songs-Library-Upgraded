// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: song/songsLibrary.proto

package songsLibraryv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Song_GetSongs_FullMethodName   = "/songsLibrary.Song/GetSongs"
	Song_DeleteSong_FullMethodName = "/songsLibrary.Song/DeleteSong"
)

// SongClient is the client API for Song service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SongClient interface {
	GetSongs(ctx context.Context, in *GetSongsRequest, opts ...grpc.CallOption) (*GetSongsResponseList, error)
	DeleteSong(ctx context.Context, in *DeleteSongsRequest, opts ...grpc.CallOption) (*DeleteSongsResponse, error)
}

type songClient struct {
	cc grpc.ClientConnInterface
}

func NewSongClient(cc grpc.ClientConnInterface) SongClient {
	return &songClient{cc}
}

func (c *songClient) GetSongs(ctx context.Context, in *GetSongsRequest, opts ...grpc.CallOption) (*GetSongsResponseList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSongsResponseList)
	err := c.cc.Invoke(ctx, Song_GetSongs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songClient) DeleteSong(ctx context.Context, in *DeleteSongsRequest, opts ...grpc.CallOption) (*DeleteSongsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSongsResponse)
	err := c.cc.Invoke(ctx, Song_DeleteSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SongServer is the server API for Song service.
// All implementations must embed UnimplementedSongServer
// for forward compatibility.
type SongServer interface {
	GetSongs(context.Context, *GetSongsRequest) (*GetSongsResponseList, error)
	DeleteSong(context.Context, *DeleteSongsRequest) (*DeleteSongsResponse, error)
	mustEmbedUnimplementedSongServer()
}

// UnimplementedSongServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSongServer struct{}

func (UnimplementedSongServer) GetSongs(context.Context, *GetSongsRequest) (*GetSongsResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongs not implemented")
}
func (UnimplementedSongServer) DeleteSong(context.Context, *DeleteSongsRequest) (*DeleteSongsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedSongServer) mustEmbedUnimplementedSongServer() {}
func (UnimplementedSongServer) testEmbeddedByValue()              {}

// UnsafeSongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SongServer will
// result in compilation errors.
type UnsafeSongServer interface {
	mustEmbedUnimplementedSongServer()
}

func RegisterSongServer(s grpc.ServiceRegistrar, srv SongServer) {
	// If the following call pancis, it indicates UnimplementedSongServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Song_ServiceDesc, srv)
}

func _Song_GetSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServer).GetSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Song_GetSongs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServer).GetSongs(ctx, req.(*GetSongsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Song_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Song_DeleteSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServer).DeleteSong(ctx, req.(*DeleteSongsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Song_ServiceDesc is the grpc.ServiceDesc for Song service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Song_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "songsLibrary.Song",
	HandlerType: (*SongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSongs",
			Handler:    _Song_GetSongs_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _Song_DeleteSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "song/songsLibrary.proto",
}

const (
	SongData_GetSongData_FullMethodName = "/songsLibrary.SongData/GetSongData"
)

// SongDataClient is the client API for SongData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SongDataClient interface {
	GetSongData(ctx context.Context, in *GetSongDataRequest, opts ...grpc.CallOption) (*GetSongDataResponse, error)
}

type songDataClient struct {
	cc grpc.ClientConnInterface
}

func NewSongDataClient(cc grpc.ClientConnInterface) SongDataClient {
	return &songDataClient{cc}
}

func (c *songDataClient) GetSongData(ctx context.Context, in *GetSongDataRequest, opts ...grpc.CallOption) (*GetSongDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSongDataResponse)
	err := c.cc.Invoke(ctx, SongData_GetSongData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SongDataServer is the server API for SongData service.
// All implementations must embed UnimplementedSongDataServer
// for forward compatibility.
type SongDataServer interface {
	GetSongData(context.Context, *GetSongDataRequest) (*GetSongDataResponse, error)
	mustEmbedUnimplementedSongDataServer()
}

// UnimplementedSongDataServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSongDataServer struct{}

func (UnimplementedSongDataServer) GetSongData(context.Context, *GetSongDataRequest) (*GetSongDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongData not implemented")
}
func (UnimplementedSongDataServer) mustEmbedUnimplementedSongDataServer() {}
func (UnimplementedSongDataServer) testEmbeddedByValue()                  {}

// UnsafeSongDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SongDataServer will
// result in compilation errors.
type UnsafeSongDataServer interface {
	mustEmbedUnimplementedSongDataServer()
}

func RegisterSongDataServer(s grpc.ServiceRegistrar, srv SongDataServer) {
	// If the following call pancis, it indicates UnimplementedSongDataServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SongData_ServiceDesc, srv)
}

func _SongData_GetSongData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSongDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongDataServer).GetSongData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SongData_GetSongData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongDataServer).GetSongData(ctx, req.(*GetSongDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SongData_ServiceDesc is the grpc.ServiceDesc for SongData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SongData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "songsLibrary.SongData",
	HandlerType: (*SongDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSongData",
			Handler:    _SongData_GetSongData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "song/songsLibrary.proto",
}
