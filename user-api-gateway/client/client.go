package client

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	config "user-api-gateway/config"
	proto "user-api-gateway/proto"
)

var (
	AuthCli          proto.AuthServiceClient
	RegisterCli      proto.RegisterServiceClient
	UserPostCli      proto.UserPostServiceClient
	UserGetFriendCli proto.UserFriendServiceClient
)

func InitGrpcRegisterClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.user.register.ip").(string)+":"+config.GetConfig().Get("grpcServer.user.register.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("GrpcRegisterClient did not connect: %v", err)
	}
	RegisterCli = proto.NewRegisterServiceClient(conn)
}

func InitGrpcAuthClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.user.auth.ip").(string)+":"+config.GetConfig().Get("grpcServer.user.auth.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("GrpcAuthClient did not connect: %v", err)
	}
	AuthCli = proto.NewAuthServiceClient(conn)
}

func InitGrpcUserPostClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.user.post.ip").(string)+":"+config.GetConfig().Get("grpcServer.user.post.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("conn: ", conn)
	if err != nil {
		log.Fatalf("GrpcUserPostClient did not connect: %v", err)
	}
	UserPostCli = proto.NewUserPostServiceClient(conn)
}

func InitGrpcUserGetFriendClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.user.get.ip").(string)+":"+config.GetConfig().Get("grpcServer.user.get.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("conn: ", conn)
	if err != nil {
		log.Fatalf("GrpcUserGetFriendClient did not connect: %v", err)
	}
	UserGetFriendCli = proto.NewUserFriendServiceClient(conn)
}

func GetRegisterCli() proto.RegisterServiceClient {
	return RegisterCli
}

func GetAuthCli() proto.AuthServiceClient {
	return AuthCli
}

func GetUserPostCli() proto.UserPostServiceClient {
	return UserPostCli
}

func GetUserGetFriendCli() proto.UserFriendServiceClient {
	return UserGetFriendCli
}
