package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetConnection will open db connection
// return DB connection
func GetConnection(url string, dbName string) (*sql.DB, error) {

	db, err := sql.Open("mysql", "user1:Mtvh11MySql@tcp(127.0.0.1:3306)/sakila?parseTime=true")

	return db, err
}

// CloseConnection will close DB connection
func CloseConnection(db *sql.DB) {
	fmt.Println("Close Db Connection ...")
	db.Close()
}
