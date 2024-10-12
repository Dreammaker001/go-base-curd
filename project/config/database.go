package config

import (
	"database/sql"
	"time"

	"github.com/Dreammaker001/go-base-curd/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	config := NewConfig("env", "json", "../../")
	db, err := sql.Open("mysql", config.GetString("database.dsn"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
