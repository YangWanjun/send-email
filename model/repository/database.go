package repository

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

var Db *sql.DB

func Init() {
	var err error
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USERNAME"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("MYSQL_ADDRESS"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
		ParseTime: true,
	}
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
}
