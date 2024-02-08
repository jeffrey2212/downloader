// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: model_downloader.proto

package downloader_backend

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

// ModelDownloaderClient is the client API for ModelDownloader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelDownloaderClient interface {
	DownloadModel(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
}

type modelDownloaderClient struct {
	cc grpc.ClientConnInterface
}

func NewModelDownloaderClient(cc grpc.ClientConnInterface) ModelDownloaderClient {
	return &modelDownloaderClient{cc}
}

func (c *modelDownloaderClient) DownloadModel(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/modeldownloader.ModelDownloader/DownloadModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelDownloaderServer is the server API for ModelDownloader service.
// All implementations must embed UnimplementedModelDownloaderServer
// for forward compatibility
type ModelDownloaderServer interface {
	DownloadModel(context.Context, *DownloadRequest) (*DownloadResponse, error)
	mustEmbedUnimplementedModelDownloaderServer()
}

// UnimplementedModelDownloaderServer must be embedded to have forward compatible implementations.
type UnimplementedModelDownloaderServer struct {
}

func (UnimplementedModelDownloaderServer) DownloadModel(context.Context, *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadModel not implemented")
}
func (UnimplementedModelDownloaderServer) mustEmbedUnimplementedModelDownloaderServer() {}

// UnsafeModelDownloaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelDownloaderServer will
// result in compilation errors.
type UnsafeModelDownloaderServer interface {
	mustEmbedUnimplementedModelDownloaderServer()
}

func RegisterModelDownloaderServer(s grpc.ServiceRegistrar, srv ModelDownloaderServer) {
	s.RegisterService(&ModelDownloader_ServiceDesc, srv)
}

func _ModelDownloader_DownloadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelDownloaderServer).DownloadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modeldownloader.ModelDownloader/DownloadModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelDownloaderServer).DownloadModel(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelDownloader_ServiceDesc is the grpc.ServiceDesc for ModelDownloader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelDownloader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "modeldownloader.ModelDownloader",
	HandlerType: (*ModelDownloaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadModel",
			Handler:    _ModelDownloader_DownloadModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model_downloader.proto",
}