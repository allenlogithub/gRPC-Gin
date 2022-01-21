package controllers

import (
	"context"
	"errors"

	config "user-auth-server/config"
	hs "user-auth-server/crypto"
	databases "user-auth-server/databases"
	jwt "user-auth-server/jwt"
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
			AccessToken: "",
		}, err
	}
	isMatch := hs.ComparePassword(res.HashedPassword, in.GetPassword())
	if isMatch {
		j := jwt.JwtInfo{
			Account: in.GetAccount(),
			UserId:  res.Id,
		}
		tk, err := jwt.CreateToken(&j)
		if err != nil {
			return &proto.LoginReply{
				AccessToken: "",
			}, errors.New("CreateTokenFailed")
		}
		d := databases.UserAccessToken{
			UserId:      res.Id,
			AccessToken: tk,
			TTL:         config.GetConfig().Get("jwt.ttl").(int),
		}
		if err := databases.SetAuthToken(&d); err != nil {
			return &proto.LoginReply{
				AccessToken: "",
			}, errors.New("SetTokenToRedisFailed")
		}
		return &proto.LoginReply{
			AccessToken: tk,
		}, nil
	}

	return &proto.LoginReply{
		AccessToken: "",
	}, errors.New("InvalidLoginInfo")
}
