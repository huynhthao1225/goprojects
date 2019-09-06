package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := GetConnection("mySql", "connectionString")
	defer CloseConnection(db)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("db connection is a GO ...")
	}

	actor := FindbyID(1, db)

	fmt.Println(actor.FirstName, actor.LastName, actor.LastUpdate)

	actors := FindAll(db)

	for _, actor := range actors {
		fmt.Printf("%d\t%s %s\t\t%s\n", actor.ID, actor.FirstName, actor.LastName, actor.LastUpdate)
	}

}
