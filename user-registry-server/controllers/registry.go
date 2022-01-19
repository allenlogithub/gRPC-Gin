package controllers

import (
	"context"

	hs "user-registry-server/crypto"
	databases "user-registry-server/databases"
	proto "user-registry-server/proto"
)

type (
	Server struct {
		proto.UnimplementedRegisterServiceServer
	}
)

func (s *Server) SetRegister(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterReply, error) {
	hash, err := hs.HashAndSalt(in.GetPassword())
	if err != nil {
		return &proto.RegisterReply{Success: false}, err
	}
	r := databases.Register{
		Account:        in.GetAccount(),
		HashedPassword: hash,
		Name:           in.GetName(),
	}
	_, err = databases.AddRegister(&r)
	if err != nil {
		return &proto.RegisterReply{Success: false}, err
	}

	return &proto.RegisterReply{Success: true}, nil
}
