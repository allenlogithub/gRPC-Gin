package main

import (
	"flag"
	"fmt"
	"os"

	config "user-get-server/config"
	databases "user-get-server/databases"
	server "user-get-server/server"
)

func main() {
	environment := flag.String("e", "dev", "environment")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Env:", *environment)

	fmt.Println("Starting user-get-server")
	config.Init(*environment)
	databases.InitMysql()
	server.Init()
}
