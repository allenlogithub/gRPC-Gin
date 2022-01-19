package controllers

import (
	"context"
	"errors"
	"strings"

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
	if in.GetPassword() != in.GetConfirmPassword() {
		return &proto.RegisterReply{Success: false}, errors.New("InvalidPassword")
	}
	hash, err := hs.HashAndSalt(in.GetPassword())
	if err != nil {
		return &proto.RegisterReply{Success: false}, errors.New("PasswordHashingFailed")
	}
	r := databases.Register{
		Account:        in.GetAccount(),
		HashedPassword: hash,
		Name:           in.GetName(),
	}
	_, err = databases.AddRegister(&r)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			return &proto.RegisterReply{Success: false}, errors.New("DuplicateEntry")
		}
		return &proto.RegisterReply{Success: false}, errors.New("AddRegisterUnknownFailed")
	}

	return &proto.RegisterReply{Success: true}, nil
}
