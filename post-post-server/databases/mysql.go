package databases

import (
	"database/sql"
	"fmt"
	"log"

	mysql "github.com/go-sql-driver/mysql"

	config "post-post-server/config"
)

var (
	connMysql *sql.DB
	err       error
)

func connectMysql(cfg *mysql.Config, dbName string) *sql.DB {
	if dbName != "" {
		cfg.DBName = dbName
	}
	connMysql, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		if dbName != "" {
			log.Fatal("Connect to Mysql.database:" + dbName + "failed.")
		}
		log.Fatal("Connect to Mysql failed.")
	}

	return connMysql
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
	connMysql = connectMysql(&cfg, c.Get("databases.mysql.databaseName").(string))

	// create table:article if not exists
	q := `
		CREATE TABLE IF NOT EXISTS article (
			id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			user_id BIGINT NOT NULL,
			content TEXT NOT NULL,
			visibility VARCHAR(10) NOT NULL
		)
	`
	_, err := connMysql.Exec(q)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.table:article failed.")
	}
}

func GetMysql() *sql.DB {
	return connMysql
}
