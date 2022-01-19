package server

import (
	"fmt"
	"net"
	"os"

	grpc "google.golang.org/grpc"

	controllers "user-auth-server/controllers"
	proto "user-auth-server/proto"
)

func newServer() *controllers.Server {
	s := &controllers.Server{}

	return s
}

func Init() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	srv := grpc.NewServer()
	proto.RegisterLoginServiceServer(srv, newServer())
	fmt.Println("Prepare to serve")
	if err := srv.Serve(listener); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}