package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	config "post-api-gateway/config"
	proto "post-api-gateway/proto"
)

var (
	AuthCli     proto.AuthServiceClient
	PostPostCli proto.PostPostServiceClient
)

func InitGrpcAuthClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.user.auth.ip").(string)+":"+config.GetConfig().Get("grpcServer.user.auth.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	AuthCli = proto.NewAuthServiceClient(conn)
}

func GetAuthCli() proto.AuthServiceClient {
	return AuthCli
}

func InitGrpcPostPostClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.post.post.ip").(string)+":"+config.GetConfig().Get("grpcServer.post.post.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	PostPostCli = proto.NewPostPostServiceClient(conn)
}

func GetPostPostCli() proto.PostPostServiceClient {
	return PostPostCli
}
