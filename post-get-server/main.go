package main

import (
	"flag"
	"fmt"
	"os"

	config "post-get-server/config"
	databases "post-get-server/databases"
	server "post-get-server/server"
)

func main() {
	environment := flag.String("e", "dev", "environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Env:", *environment)

	fmt.Println("Starting post-get-server")
	config.Init(*environment)
	databases.InitMysql()
	server.Init()
}
