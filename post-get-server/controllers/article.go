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
