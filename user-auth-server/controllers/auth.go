package controllers

import (
	"context"
	"errors"

	"github.com/google/uuid"

	hs "user-auth-server/crypto"
	databases "user-auth-server/databases"
	proto "user-auth-server/proto"
)

type (
	Server struct {
		proto.UnimplementedLoginServiceServer
	}
)

func (s *Server) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginReply, error) {
	r := databases.RegisterInfoRequest{
		Account: in.GetAccount(),
	}
	res, err := databases.GetRegisterInfo(&r)
	if err != nil {
		return &proto.LoginReply{
			AccessToken:  "",
			DeviceNumber: "",
		}, err
	}
	isMatch := hs.ComparePassword(res.HashedPassword, in.GetPassword())
	if isMatch {
		return &proto.LoginReply{
			AccessToken:  "12312312312313",
			DeviceNumber: uuid.New().String(),
		}, nil
	}

	return &proto.LoginReply{
		AccessToken:  "",
		DeviceNumber: "",
	}, errors.New("InvalidLoginInfo")
}
