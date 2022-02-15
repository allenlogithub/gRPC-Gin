package databases

import (
	"database/sql"
	"fmt"
	"log"

	mysql "github.com/go-sql-driver/mysql"

	config "user-post-server/config"
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
	c := config.GetConfig()
	cfg := mysql.Config{
		User:   c.Get("databases.mysql.user").(string),
		Passwd: c.Get("databases.mysql.password").(string),
		Net:    c.Get("databases.mysql.net").(string),
		Addr:   c.Get("databases.mysql.domain").(string) + ":" + c.Get("databases.mysql.port").(string),
	}

	connMysql = connectMysql(&cfg, c.Get("databases.mysql.databaseName").(string))

	q := `
		CREATE TABLE IF NOT EXISTS friendrequest (
			id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			requestor_user_id BIGINT NOT NULL,
			receiver_user_id BIGINT NOT NULL,
			FOREIGN KEY(requestor_user_id)
				REFERENCES register(id)
				ON DELETE CASCADE,
			FOREIGN KEY(receiver_user_id)
				REFERENCES register(id)
				ON DELETE CASCADE
		)
	`
	_, err := connMysql.Exec(q)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.table:friendrequest failed.")
	}

	q = `
		CREATE TABLE IF NOT EXISTS friendlist (
			id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			user_id BIGINT NOT NULL,
			friend_user_id BIGINT NOT NULL,
			FOREIGN KEY(user_id)
				REFERENCES register(id)
				ON DELETE CASCADE
		)
	`
	_, err = connMysql.Exec(q)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Create Mysql.table:friendlist failed.")
	}
}

func GetMysql() *sql.DB {
	return connMysql
}
