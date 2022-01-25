package main

import (
	"flag"
	"fmt"
	"os"

	config "post-post-server/config"
	databases "post-post-server/databases"
	server "post-post-server/server"
)

func main() {
	environment := flag.String("e", "dev", "environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Env:", *environment)

	fmt.Println("Starting post-post-server")
	config.Init(*environment)
	databases.InitMysql()
	server.Init()
}
