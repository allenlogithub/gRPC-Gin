package controllers

import (
	"context"

	databases "user-post-server/databases"
	proto "user-post-server/proto"
)

type (
	Server struct {
		proto.UnimplementedUserPostServiceServer
	}
)

func (s *Server) SendFriendRequest(ctx context.Context, in *proto.SendFriendRequestRequest) (*proto.SendFriendRequestReply, error) {
	d := databases.AddFriendRequestRequest{
		RequestorUserId: in.GetRequestorUserId(),
		ReceiverUserId:  in.GetReceiverUserId(),
	}
	if err := databases.AddFriendRequest(&d); err != nil {
		return &proto.SendFriendRequestReply{
			IsSent: false,
		}, err
	}

	return &proto.SendFriendRequestReply{
		IsSent: true,
	}, nil
}

func (s *Server) AcceptFriendRequest(ctx context.Context, in *proto.AcceptFriendRequestRequest) (*proto.AcceptFriendRequestReply, error) {
	d := databases.AddFriendListRequest{
		RequestorUserId: in.GetRequestorUserId(),
		ReceiverUserId:  in.GetReceiverUserId(),
	}
	if err := databases.AddFriendList(&d); err != nil {
		return &proto.AcceptFriendRequestReply{
			IsAccepted: false,
		}, err
	}

	return &proto.AcceptFriendRequestReply{
		IsAccepted: true,
	}, nil
}

func (s *Server) RejectFriendRequest(ctx context.Context, in *proto.RejectFriendRequestRequest) (*proto.RejectFriendRequestReply, error) {
	d := databases.DelFriendRequestRequest{
		RequestorUserId: in.GetRequestorUserId(),
		ReceiverUserId:  in.GetReceiverUserId(),
	}
	if err := databases.DelFriendRequest(&d); err != nil {
		return &proto.RejectFriendRequestReply{
			IsRejected: false,
		}, err
	}

	return &proto.RejectFriendRequestReply{
		IsRejected: true,
	}, nil
}
