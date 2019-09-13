package dbcon

import (
	"database/sql"
	"fmt"
	"mysqlsample/config"
	_ "github.com/go-sql-driver/mysql"
)

// GetConnection will open db connection
// return DB connection
func GetConnection(conProperty *config.ConnectionProperty) (*sql.DB, string, error) {
	
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", conProperty.User, conProperty.Password, conProperty.Protocol, conProperty.Host, conProperty.Port, conProperty.Database)
	//db, err := sql.Open(conProperty.Driver, "user1:Mtvh11MySql@tcp(127.0.0.1:3306)/sakila?parseTime=true")
	db, err := sql.Open(conProperty.Driver, connectionString)
	
	return db, connectionString, err
}

// CloseConnection will close DB connection
func CloseConnection(db *sql.DB) {
	fmt.Println("Close Db Connection ...")
	db.Close()
}
