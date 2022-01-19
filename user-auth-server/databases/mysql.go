package databases

import (
	"database/sql"
	"fmt"
	"log"

	mysql "github.com/go-sql-driver/mysql"

	config "user-auth-server/config"
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

func InitMysql() {
	// set database connection config
	c := config.GetConfig()
	cfg := mysql.Config{
		User:   c.Get("databases.mysql.user").(string),
		Passwd: c.Get("databases.mysql.password").(string),
		Net:    c.Get("databases.mysql.net").(string),
		Addr:   c.Get("databases.mysql.domain").(string) + ":" + c.Get("databases.mysql.port").(string),
	}

	// connect to the database:user
	conn = connectMysql(&cfg, c.Get("databases.mysql.databaseName").(string))
}

func GetMysql() *sql.DB {
	return conn
}
