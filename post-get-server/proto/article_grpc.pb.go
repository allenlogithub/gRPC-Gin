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
	AddArticleComment(ctx context.Context, in *AddArticleCommentRequest, opts ...grpc.CallOption) (*AddArticleCommentReply, error)
	DelArticleComment(ctx context.Context, in *DelArticleCommentRequest, opts ...grpc.CallOption) (*DelArticleCommentReply, error)
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

func (c *postArticleServiceClient) AddArticleComment(ctx context.Context, in *AddArticleCommentRequest, opts ...grpc.CallOption) (*AddArticleCommentReply, error) {
	out := new(AddArticleCommentReply)
	err := c.cc.Invoke(ctx, "/proto.PostArticleService/AddArticleComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postArticleServiceClient) DelArticleComment(ctx context.Context, in *DelArticleCommentRequest, opts ...grpc.CallOption) (*DelArticleCommentReply, error) {
	out := new(DelArticleCommentReply)
	err := c.cc.Invoke(ctx, "/proto.PostArticleService/DelArticleComment", in, out, opts...)
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
	AddArticleComment(context.Context, *AddArticleCommentRequest) (*AddArticleCommentReply, error)
	DelArticleComment(context.Context, *DelArticleCommentRequest) (*DelArticleCommentReply, error)
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
func (UnimplementedPostArticleServiceServer) AddArticleComment(context.Context, *AddArticleCommentRequest) (*AddArticleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddArticleComment not implemented")
}
func (UnimplementedPostArticleServiceServer) DelArticleComment(context.Context, *DelArticleCommentRequest) (*DelArticleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelArticleComment not implemented")
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

func _PostArticleService_AddArticleComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddArticleCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostArticleServiceServer).AddArticleComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostArticleService/AddArticleComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostArticleServiceServer).AddArticleComment(ctx, req.(*AddArticleCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostArticleService_DelArticleComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelArticleCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostArticleServiceServer).DelArticleComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostArticleService/DelArticleComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostArticleServiceServer).DelArticleComment(ctx, req.(*DelArticleCommentRequest))
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
		{
			MethodName: "AddArticleComment",
			Handler:    _PostArticleService_AddArticleComment_Handler,
		},
		{
			MethodName: "DelArticleComment",
			Handler:    _PostArticleService_DelArticleComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}

// GetArticleServiceClient is the client API for GetArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetArticleServiceClient interface {
	GetPersonalArticle(ctx context.Context, in *GetPersonalArticleRequest, opts ...grpc.CallOption) (*GetPersonalArticleReply, error)
	GetFriendArticle(ctx context.Context, in *GetFriendArticleRequest, opts ...grpc.CallOption) (*GetFriendArticleReply, error)
}

type getArticleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGetArticleServiceClient(cc grpc.ClientConnInterface) GetArticleServiceClient {
	return &getArticleServiceClient{cc}
}

func (c *getArticleServiceClient) GetPersonalArticle(ctx context.Context, in *GetPersonalArticleRequest, opts ...grpc.CallOption) (*GetPersonalArticleReply, error) {
	out := new(GetPersonalArticleReply)
	err := c.cc.Invoke(ctx, "/proto.GetArticleService/GetPersonalArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getArticleServiceClient) GetFriendArticle(ctx context.Context, in *GetFriendArticleRequest, opts ...grpc.CallOption) (*GetFriendArticleReply, error) {
	out := new(GetFriendArticleReply)
	err := c.cc.Invoke(ctx, "/proto.GetArticleService/GetFriendArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetArticleServiceServer is the server API for GetArticleService service.
// All implementations must embed UnimplementedGetArticleServiceServer
// for forward compatibility
type GetArticleServiceServer interface {
	GetPersonalArticle(context.Context, *GetPersonalArticleRequest) (*GetPersonalArticleReply, error)
	GetFriendArticle(context.Context, *GetFriendArticleRequest) (*GetFriendArticleReply, error)
	mustEmbedUnimplementedGetArticleServiceServer()
}

// UnimplementedGetArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGetArticleServiceServer struct {
}

func (UnimplementedGetArticleServiceServer) GetPersonalArticle(context.Context, *GetPersonalArticleRequest) (*GetPersonalArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonalArticle not implemented")
}
func (UnimplementedGetArticleServiceServer) GetFriendArticle(context.Context, *GetFriendArticleRequest) (*GetFriendArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendArticle not implemented")
}
func (UnimplementedGetArticleServiceServer) mustEmbedUnimplementedGetArticleServiceServer() {}

// UnsafeGetArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetArticleServiceServer will
// result in compilation errors.
type UnsafeGetArticleServiceServer interface {
	mustEmbedUnimplementedGetArticleServiceServer()
}

func RegisterGetArticleServiceServer(s grpc.ServiceRegistrar, srv GetArticleServiceServer) {
	s.RegisterService(&GetArticleService_ServiceDesc, srv)
}

func _GetArticleService_GetPersonalArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonalArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetArticleServiceServer).GetPersonalArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GetArticleService/GetPersonalArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetArticleServiceServer).GetPersonalArticle(ctx, req.(*GetPersonalArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetArticleService_GetFriendArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetArticleServiceServer).GetFriendArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GetArticleService/GetFriendArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetArticleServiceServer).GetFriendArticle(ctx, req.(*GetFriendArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetArticleService_ServiceDesc is the grpc.ServiceDesc for GetArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GetArticleService",
	HandlerType: (*GetArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersonalArticle",
			Handler:    _GetArticleService_GetPersonalArticle_Handler,
		},
		{
			MethodName: "GetFriendArticle",
			Handler:    _GetArticleService_GetFriendArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
