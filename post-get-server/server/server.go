package server

import (
	"fmt"
	"net"
	"os"

	grpc "google.golang.org/grpc"

	config "post-get-server/config"
	controllers "post-get-server/controllers"
	proto "post-get-server/proto"
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
	proto.RegisterGetArticleServiceServer(srv, newServer())
	fmt.Println("Prepare to serve")
	if err := srv.Serve(listener); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
