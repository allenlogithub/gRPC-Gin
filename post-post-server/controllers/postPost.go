package controllers

import (
	"context"
	"fmt"
	// "errors"
	// "strconv"

	// config "post-post-server/config"
	// hs "post-post-server/crypto"
	// databases "post-post-server/databases"
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
	// userId, err := databases.GetUserId(&d)
	// if err != nil {
	// 	return &proto.PostPostReply{
	// 		Success: false,
	// 	}, err
	// }
	// fmt.Println(userId)
	// if err := databases.DeleteAuthToken(&d); err != nil {
	// 	return &proto.LogoutReply{
	// 		Success: false,
	// 	}, err
	// }
	fmt.Println(in.GetUserId())
	fmt.Println(in.GetContent())
	fmt.Println(in.GetVisibility())

	return &proto.PostPostReply{
		Success: true,
	}, nil
}
