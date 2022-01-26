package controllers

import (
	"context"

	databases "post-post-server/databases"
	proto "post-post-server/proto"
)

func (s *Server) AddArticleComment(ctx context.Context, in *proto.AddArticleCommentRequest) (*proto.AddArticleCommentReply, error) {
	d := databases.PostArticleCommentRequest{
		UserId:    in.GetUserId(),
		Content:   in.GetContent(),
		ArticleId: in.GetArticleId(),
	}
	_, err := databases.AddArticleComment(&d)
	if err != nil {
		return &proto.AddArticleCommentReply{
			Success: false,
		}, err
	}

	return &proto.AddArticleCommentReply{
		Success: true,
	}, nil
}

func (s *Server) DelArticleComment(ctx context.Context, in *proto.DelArticleCommentRequest) (*proto.DelArticleCommentReply, error) {
	d := databases.DelArticleCommentRequest{
		UserId:           in.GetUserId(),
		ArticleCommentId: in.GetArticleCommentId(),
	}
	err := databases.DelArticleComment(&d)
	if err != nil {
		return &proto.DelArticleCommentReply{
			Success: false,
		}, err
	}

	return &proto.DelArticleCommentReply{
		Success: true,
	}, nil
}
