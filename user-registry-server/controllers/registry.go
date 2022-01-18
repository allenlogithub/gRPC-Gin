package controllers

import (
	"context"

	hs "user-registry-server/crypto"
	databases "user-registry-server/databases"
	proto "user-registry-server/proto"
)

type (
	Server struct {
		proto.UnimplementedRegistryServiceServer
	}
)

func (s *Server) SetRegistry(ctx context.Context, in *proto.RegistryRequest) (*proto.RegistryReply, error) {
	hash, err := hs.HashAndSalt(in.GetPassword())
	if err != nil {
		return &proto.RegistryReply{Success: false}, err
	}
	r := databases.Register{
		Account:        in.GetAccount(),
		HashedPassword: hash,
		Email:          in.GetEmail(),
		Name:           in.GetName(),
	}
	_, err = databases.AddRegister(&r)
	if err != nil {
		return &proto.RegistryReply{Success: false}, err
	}

	return &proto.RegistryReply{Success: true}, nil
}
