package main

import (
	"flag"
	"fmt"
	"os"

	client "user-api-gateway/client"
	config "user-api-gateway/config"
	server "user-api-gateway/server"
)

func main() {
	environment := flag.String("e", "dev", "environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Env:", *environment)

	fmt.Println("Starting user-api-gateway")
	config.Init(*environment)
	client.InitGrpcRegisterClient()
	client.InitGrpcAuthClient()
	client.InitGrpcUserPostClient()
	server.Init()
}
