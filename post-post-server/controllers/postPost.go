package controllers

import (
	"context"

	databases "post-post-server/databases"
	proto "post-post-server/proto"
)

type (
	Server struct {
		proto.UnimplementedPostPostServiceServer
	}
)

func (s *Server) PostPost(ctx context.Context, in *proto.PostPostRequest) (*proto.PostPostReply, error) {
	d := databases.PostArticleRequest{
		UserId:     in.GetUserId(),
		Content:    in.GetContent(),
		Visibility: in.GetVisibility(),
	}
	_, err := databases.AddArticle(&d)
	if err != nil {
		return &proto.PostPostReply{
			Success: false,
		}, err
	}

	return &proto.PostPostReply{
		Success: true,
	}, nil
}
