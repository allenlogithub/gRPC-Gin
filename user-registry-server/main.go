package main

import (
	"flag"
	"fmt"
	"os"

	config "user-registry-server/config"
	databases "user-registry-server/databases"
	server "user-registry-server/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	config.Init(*environment)
	databases.InitMysql()
	server.Init()
}
