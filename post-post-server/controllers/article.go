package controllers

import (
	"context"

	databases "post-post-server/databases"
	proto "post-post-server/proto"
)

type (
	Server struct {
		proto.UnimplementedPostArticleServiceServer
	}
)

func (s *Server) AddArticle(ctx context.Context, in *proto.AddArticleRequest) (*proto.AddArticleReply, error) {
	d := databases.PostArticleRequest{
		UserId:     in.GetUserId(),
		Content:    in.GetContent(),
		Visibility: in.GetVisibility(),
	}
	_, err := databases.AddArticle(&d)
	if err != nil {
		return &proto.AddArticleReply{
			Success: false,
		}, err
	}

	return &proto.AddArticleReply{
		Success: true,
	}, nil
}

func (s *Server) DelArticle(ctx context.Context, in *proto.DelArticleRequest) (*proto.DelArticleReply, error) {
	d := databases.DelArticleRequest{
		UserId:    in.GetUserId(),
		ArticleId: in.GetArticleId(),
	}
	err := databases.DelArticle(&d)
	if err != nil {
		return &proto.DelArticleReply{
			Success: false,
		}, err
	}

	return &proto.DelArticleReply{
		Success: true,
	}, nil
}
