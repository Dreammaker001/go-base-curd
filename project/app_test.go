package main

import (
	"database/sql"
	"testing"

	"github.com/Dreammaker001/go-base-curd/config"
	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	config := config.NewConfig("env", "json", "../../")
	db, err := sql.Open("mysql", config.GetString("database.dsn"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
