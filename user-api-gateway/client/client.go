package client

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "user-api-gateway/proto"
)

var (
	AuthCli     proto.LoginServiceClient
	RegisterCli proto.RegisterServiceClient

	authAddr     = flag.String("authAddr", "172.17.0.8:4040", "the auth server address to connect to")
	registerAddr = flag.String("registerAddr", "172.17.0.6:4040", "the register server address to connect to")
)

func InitGrpcRegisterClient() {
	conn, err := grpc.Dial(*registerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	RegisterCli = proto.NewRegisterServiceClient(conn)
}

func InitGrpcAuthClient() {
	conn, err := grpc.Dial(*authAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
