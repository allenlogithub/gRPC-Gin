package main

import (
	"flag"
	"fmt"
	"os"

	client "post-api-gateway/client"
	config "post-api-gateway/config"
	server "post-api-gateway/server"
)

func main() {
	environment := flag.String("e", "dev", "environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Env:", *environment)

	fmt.Println("Starting post-api-gateway")
	config.Init(*environment)
	client.InitGrpcAuthClient()
	client.InitGrpcPostPostClient()
	server.Init()
}
