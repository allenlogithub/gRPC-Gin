package controllers

import (
	"context"
	"errors"
	"strconv"

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

func (s *Server) Logout(ctx context.Context, in *proto.LogoutRequest) (*proto.LogoutReply, error) {
	d := databases.UserOwnedToken{
		AccessToken: in.GetAccessToken(),
	}
	if err := databases.DeleteAuthToken(&d); err != nil {
		return &proto.LogoutReply{
			Success: false,
		}, err
	}

	return &proto.LogoutReply{
		Success: true,
	}, nil
}

func (s *Server) ValidateJWT(ctx context.Context, in *proto.JWTValidationRequest) (*proto.JWTValidationReply, error) {
	// JWT verification
	jwtInfo, err1 := jwt.ExtractTokenData(in.GetAccessToken())
	if err1 != nil {
		return &proto.JWTValidationReply{
			Valid: false,
		}, err1
	}
	// check redis for JWT TTL
	q := databases.UserOwnedToken{
		AccessToken: in.GetAccessToken(),
	}
	res, err2 := databases.GetUserId(&q)
	i64, err3 := strconv.Atoi(res)
	if err2 != nil || err3 != nil || int64(i64) != jwtInfo.UserId {
		return &proto.JWTValidationReply{
			Valid: false,
		}, err2
	}

	return &proto.JWTValidationReply{
		Valid: true,
	}, nil
}
