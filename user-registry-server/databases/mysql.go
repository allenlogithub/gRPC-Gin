package databases

import (
	"database/sql"
	"fmt"
	"log"

	mysql "github.com/go-sql-driver/mysql"

	config "user-registry-server/config"
)

var (
	conn *sql.DB
	err  error
)

func connectMysql(cfg *mysql.Config, dbName string) *sql.DB {
	if dbName != "" {
		cfg.DBName = dbName
	}
	conn, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		if dbName != "" {
			log.Fatal("Connect to Mysql.database:" + dbName + "failed.")
		}
		log.Fatal("Connect to Mysql failed.")
	}

	return conn
}

func createMysqlDB(conn *sql.DB, dbName string) {
	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.database:" + dbName + "failed.")
	}
}

func InitMysql() {
	// set database connection config
	c := config.GetConfig()
	cfg := mysql.Config{
		User:   c.Get("databases.mysql.user").(string),
		Passwd: c.Get("databases.mysql.password").(string),
		Net:    c.Get("databases.mysql.net").(string),
		Addr:   c.Get("databases.mysql.domain").(string) + ":" + c.Get("databases.mysql.port").(string),
	}

	// connect to the Mysql
	conn = connectMysql(&cfg, "")
	defer conn.Close()

	// create database:user if not exists
	createMysqlDB(conn, c.Get("databases.mysql.databaseName").(string))

	// connect to the database:user
	conn = connectMysql(&cfg, c.Get("databases.mysql.databaseName").(string))

	// create table:register if not exists
	q := `
		CREATE TABLE IF NOT EXISTS register (
			id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			account VARCHAR(50) NOT NULL,
			hashed_password VARCHAR(128) NOT NULL,
			email VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			UNIQUE (account)
		)
	`
	_, err := conn.Exec(q)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.table:register failed.")
	}
}

func GetMysql() *sql.DB {
	return conn
}
