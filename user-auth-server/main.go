package main

import (
	"flag"
	"fmt"
	"os"

	config "user-auth-server/config"
	databases "user-auth-server/databases"
	server "user-auth-server/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	fmt.Println("Starting user-auth-server")
	config.Init(*environment)
	databases.InitMysql()
	databases.InitRedis()
	server.Init()
}
