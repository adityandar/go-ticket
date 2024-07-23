package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbInit() *sql.DB {

	err := godotenv.Load("config/.env")
	if err != nil {
		panic("Failed to load file environment")
	} else {
		fmt.Println("Success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("Connection failed when open DB: " + err.Error())
	}
	err = DB.Ping()
	if err != nil {
		panic("Connection failed when ping DB: " + err.Error())
	} else {
		fmt.Println("Connection Success")
	}
	dbMigrate(DB)

	return DB
}

func dbMigrate(dbParam *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/sql_migration",
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam
	fmt.Println("Applied", n, "migrations")
}
