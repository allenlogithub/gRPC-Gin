package controllers

import (
	"context"
	// "errors"
	// "strconv"

	// config "post-post-server/config"
	// hs "post-post-server/crypto"
	// databases "post-post-server/databases"
	// jwt "post-post-server/jwt"
	proto "post-post-server/proto"
)

type (
	Server struct {
		proto.UnimplementedPostPostServiceServer
	}
)

func (s *Server) PostPost(ctx context.Context, in *proto.PostPostRequest) (*proto.PostPostReply, error) {
	// d := databases.UserOwnedToken{
	// 	AccessToken: in.GetAccessToken(),
	// }
	// if err := databases.DeleteAuthToken(&d); err != nil {
	// 	return &proto.LogoutReply{
	// 		Success: false,
	// 	}, err
	// }

	return &proto.PostPostReply{
		Success: true,
	}, nil
}
