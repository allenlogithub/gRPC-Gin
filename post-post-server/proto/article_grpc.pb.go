// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// PostArticleServiceClient is the client API for PostArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostArticleServiceClient interface {
	AddArticle(ctx context.Context, in *AddArticleRequest, opts ...grpc.CallOption) (*AddArticleReply, error)
	DelArticle(ctx context.Context, in *DelArticleRequest, opts ...grpc.CallOption) (*DelArticleReply, error)
}

type postArticleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostArticleServiceClient(cc grpc.ClientConnInterface) PostArticleServiceClient {
	return &postArticleServiceClient{cc}
}

func (c *postArticleServiceClient) AddArticle(ctx context.Context, in *AddArticleRequest, opts ...grpc.CallOption) (*AddArticleReply, error) {
	out := new(AddArticleReply)
	err := c.cc.Invoke(ctx, "/proto.PostArticleService/AddArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postArticleServiceClient) DelArticle(ctx context.Context, in *DelArticleRequest, opts ...grpc.CallOption) (*DelArticleReply, error) {
	out := new(DelArticleReply)
	err := c.cc.Invoke(ctx, "/proto.PostArticleService/DelArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostArticleServiceServer is the server API for PostArticleService service.
// All implementations must embed UnimplementedPostArticleServiceServer
// for forward compatibility
type PostArticleServiceServer interface {
	AddArticle(context.Context, *AddArticleRequest) (*AddArticleReply, error)
	DelArticle(context.Context, *DelArticleRequest) (*DelArticleReply, error)
	mustEmbedUnimplementedPostArticleServiceServer()
}

// UnimplementedPostArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostArticleServiceServer struct {
}

func (UnimplementedPostArticleServiceServer) AddArticle(context.Context, *AddArticleRequest) (*AddArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddArticle not implemented")
}
func (UnimplementedPostArticleServiceServer) DelArticle(context.Context, *DelArticleRequest) (*DelArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelArticle not implemented")
}
func (UnimplementedPostArticleServiceServer) mustEmbedUnimplementedPostArticleServiceServer() {}

// UnsafePostArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostArticleServiceServer will
// result in compilation errors.
type UnsafePostArticleServiceServer interface {
	mustEmbedUnimplementedPostArticleServiceServer()
}

func RegisterPostArticleServiceServer(s grpc.ServiceRegistrar, srv PostArticleServiceServer) {
	s.RegisterService(&PostArticleService_ServiceDesc, srv)
}

func _PostArticleService_AddArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostArticleServiceServer).AddArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostArticleService/AddArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostArticleServiceServer).AddArticle(ctx, req.(*AddArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostArticleService_DelArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostArticleServiceServer).DelArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostArticleService/DelArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostArticleServiceServer).DelArticle(ctx, req.(*DelArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostArticleService_ServiceDesc is the grpc.ServiceDesc for PostArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PostArticleService",
	HandlerType: (*PostArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddArticle",
			Handler:    _PostArticleService_AddArticle_Handler,
		},
		{
			MethodName: "DelArticle",
			Handler:    _PostArticleService_DelArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
