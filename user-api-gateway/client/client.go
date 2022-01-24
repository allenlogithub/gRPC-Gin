package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	config "user-api-gateway/config"
	proto "user-api-gateway/proto"
)

var (
	AuthCli     proto.LoginServiceClient
	RegisterCli proto.RegisterServiceClient
)

func InitGrpcRegisterClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.user.register.ip").(string)+":"+config.GetConfig().Get("grpcServer.user.register.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	RegisterCli = proto.NewRegisterServiceClient(conn)
}

func InitGrpcAuthClient() {
	conn, err := grpc.Dial(config.GetConfig().Get("grpcServer.user.auth.ip").(string)+":"+config.GetConfig().Get("grpcServer.user.auth.port").(string), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	AuthCli = proto.NewLoginServiceClient(conn)
}

func GetRegisterCli() proto.RegisterServiceClient {
	return RegisterCli
}

func GetAuthCli() proto.LoginServiceClient {
	return AuthCli
}
