package main

import (
	"flag"
	"fmt"
	"os"

	config "user-post-server/config"
	databases "user-post-server/databases"
	server "user-post-server/server"
)

func main() {
	environment := flag.String("e", "dev", "environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Env:", *environment)

	fmt.Println("Starting user-post-server")
	config.Init(*environment)
	databases.InitMysql()
	server.Init()
}
