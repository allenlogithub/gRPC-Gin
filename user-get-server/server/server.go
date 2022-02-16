package server

import (
	"fmt"
	"net"
	"os"

	grpc "google.golang.org/grpc"

	config "user-get-server/config"
	controllers "user-get-server/controllers"
	proto "user-get-server/proto"
)

func newServer() *controllers.Server {
	s := &controllers.Server{}

	return s
}

func Init() {
	listener, err := net.Listen("tcp", ":"+config.GetConfig().Get("server.port").(string))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	srv := grpc.NewServer()
	proto.RegisterUserFriendServiceServer(srv, newServer())
	fmt.Println("Prepare to serve")
	if err := srv.Serve(listener); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
