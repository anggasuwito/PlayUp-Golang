package database

import (
	"database/sql"
	"fmt"
	"log"
	"playUp/main/master/utils/helper"
	_ "github.com/go-sql-driver/mysql"
)

// ConnectionDB -> initilize connection to the database
func ConnectionDB() *sql.DB {
	DB_USER := helper.ReadEnv("dbUser")

	DB_PASSWORD := helper.ReadEnv("dbPassword")

	DB_HOST := helper.ReadEnv("dbHost")

	DB_PORT := helper.ReadEnv("dbPort")

	DB_SCHEMA := helper.ReadEnv("dbName")

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_SCHEMA)
	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Println(err)
	}
	log.Println("-> DB has started")
	return db
}
