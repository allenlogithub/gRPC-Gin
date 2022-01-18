package databases

import (
	"database/sql"
	"fmt"
	"log"

	mysql "github.com/go-sql-driver/mysql"

	config "user-registry-server/config"
)

var (
	mysqlDb *sql.DB
	err     error
)

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
	mysqlDb, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		log.Fatal("Connect to Mysql failed.")
	}
	defer mysqlDb.Close()
	// create database:user if not exists
	_, err = mysqlDb.Exec("CREATE DATABASE IF NOT EXISTS " + c.Get("databases.mysql.databaseName").(string))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.database:user failed.")
	}
	// connect to the database:user
	cfg.DBName = c.Get("databases.mysql.databaseName").(string)
	mysqlDb, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		log.Fatal("Init Mysql.database:user failed.")
	}
	// create table:register if not exists
	// register
	q := `
		CREATE TABLE IF NOT EXISTS register (
			id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			account VARCHAR(50) NOT NULL,			
			hashed_password VARCHAR(128) NOT NULL,
			email VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL
		)
	`
	_, err := mysqlDb.Exec(q)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.table:register failed.")
	}
}

func GetMysql() *sql.DB {
	return mysqlDb
}
