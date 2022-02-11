package controllers

import (
	"context"

	databases "post-get-server/databases"
	proto "post-get-server/proto"
)

type (
	Server struct {
		proto.UnimplementedGetArticleServiceServer
	}
)

func (s *Server) GetPersonalArticle(ctx context.Context, in *proto.GetPersonalArticleRequest) (*proto.GetPersonalArticleReply, error) {
	d := databases.GetPersonalArticleRequest{
		UserId: in.GetUserId(),
	}
	res, err := databases.GetPersonalArticle(&d)
	if err != nil {
		return &proto.GetPersonalArticleReply{
			Items: nil,
		}, err
	}

	return &proto.GetPersonalArticleReply{
		Items: res.Items,
	}, nil
}

func (s *Server) GetArticleComment(ctx context.Context, in *proto.GetArticleCommentRequest) (*proto.GetArticleCommentReply, error) {
	d := databases.GetArticleCommentRequest{
		UserId:    in.GetUserId(),
		ArticleId: in.GetArticleId(),
	}
	res, err := databases.GetArticleComment(&d)
	if err != nil {
		return &proto.GetArticleCommentReply{
			Items: nil,
		}, err
	}

	return &proto.GetArticleCommentReply{
		Items: res.Items,
	}, nil
}
