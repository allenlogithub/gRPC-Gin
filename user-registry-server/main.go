package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	proto "user-registry-server/proto"
)

type (
	server struct {
		// proto.RegisterRegistryServiceServer
		proto.UnimplementedRegistryServiceServer
	}
)

func newServer() *server {
	s := &server{}
	// s.loadFeatures(*jsonDBFile)
	return s
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	// databaseImplementation := os.Args[1]
	// db, err = database.Factory(databaseImplementation)
	// if err != nil {
	// 	panic(err)
	// }
	proto.RegisterRegistryServiceServer(srv, newServer())
	fmt.Println("Prepare to serve")
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) SetRegistry(ctx context.Context, in *proto.RegistryRequest) (*proto.RegistryReply, error) {
	// value, err := db.Set(in.GetKey(), in.GetValue())
	// return generateResponse(value, err)
	fmt.Println("Received: account", in.GetAccount())
	fmt.Println("Received: name", in.GetName())
	fmt.Println("Received: email", in.GetEmail())
	fmt.Println("Received: password", in.GetPassword())
	fmt.Println("Received: confirmPassword", in.GetConfirmPassword())
	return &proto.RegistryReply{Success: true}, nil
}
