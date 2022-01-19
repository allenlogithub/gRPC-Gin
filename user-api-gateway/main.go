package main

import (
	"flag"
	"fmt"
	"os"

	client "user-api-gateway/client"
	server "user-api-gateway/server"
)

func main() {
	// environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	// config.Init(*environment)
	fmt.Println("Starting user-api-gateway")
	client.InitGrpcRegisterClient()
	client.InitGrpcAuthClient()
	server.Init()
}
