package main

import (
	"database/sql"
	"go-ticket/database"
	"go-ticket/router"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	DB = database.DbInit()
	defer DB.Close()

	router.RunServer()

}
