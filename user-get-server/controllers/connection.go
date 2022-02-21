package controllers

import (
	"context"

	databases "user-get-server/databases"
	proto "user-get-server/proto"
)

type (
	Server struct {
		proto.UnimplementedUserFriendServiceServer
	}
)

func (s *Server) GetFriendList(ctx context.Context, in *proto.GetFriendListRequest) (*proto.GetFriendListReply, error) {
	d := databases.GetFriendListRequest{
		UserId: in.GetUserId(),
	}
	result, err := databases.GetFriendList(&d)
	if err != nil {
		return &proto.GetFriendListReply{
			Items: nil,
		}, err
	}

	return &proto.GetFriendListReply{
		Items: result.Items,
	}, nil
}

func (s *Server) SearchUser(ctx context.Context, in *proto.SearchUserRequest) (*proto.SearchUserReply, error) {
	d := databases.SearchUserRequest{
		UserId:       in.GetUserId(),
		SearchString: in.GetSearchString(),
	}
	result, err := databases.SearchUser(&d)
	if err != nil {
		return &proto.SearchUserReply{
			Items: nil,
		}, err
	}

	return &proto.SearchUserReply{
		Items: result.Items,
	}, nil
}

func (s *Server) GetFriendRequestList(ctx context.Context, in *proto.GetFriendRequestListRequest) (*proto.GetFriendRequestListReply, error) {
	d := databases.GetFriendRequestListRequest{
		UserId: in.GetUserId(),
	}
	result, err := databases.GetFriendRequestList(&d)
	if err != nil {
		return &proto.GetFriendRequestListReply{
			Items: nil,
		}, err
	}

	return &proto.GetFriendRequestListReply{
		Items: result.Items,
	}, nil
}
