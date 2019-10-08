package main

import (
	"fmt"
	"mysqlsample/config"
	"mysqlsample/dao"
	"mysqlsample/dbcon"
	"mysqlsample/utils"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go MySQL Tutorial")
	// load DB connection properties
	conProperties := utils.LoadProperties()
	var conProperty config.ConnectionProperty
	fmt.Println(conProperties)
	if len(conProperties.ConnectionProperties) == 1 {
		conProperty = conProperties.ConnectionProperties[0]
	}

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, connectionString, err := dbcon.GetConnection(&conProperty)
	defer dbcon.CloseConnection(db)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("db connection is a GO ...")
	}

	actorAction := dao.ActionActor{}

	actor := actorAction.FindbyID(1, db)

	fmt.Println(actor.FirstName, actor.LastName, actor.LastUpdate)

	actors := actorAction.FindAll(db)

	jobs := make(chan string, 5)
	fp, _ := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)
	go writeToConsole(jobs)
	for _, actor := range actors {
		value := fmt.Sprintf("%d\t%s %s\t\t%s\n", actor.ID, actor.FirstName, actor.LastName, actor.LastUpdate)
		jobs <- value
		fp.WriteString(fmt.Sprintf("%d\t%s %s\t\t%s\n", actor.ID, actor.FirstName, actor.LastName, actor.LastUpdate))
	}
	close(jobs)
	
	addressAction := dao.ActionAddress{}
	address := addressAction.FindbyID(1, db)
	fmt.Println(address.ID, address.Address, address.Location)
	fmt.Printf("ConnectionString = %s\n", connectionString)

}

func writeToConsole(ch <-chan string) {
	for {
		value, more  := <- ch
		if more {
			fmt.Print(value)
		} else {
			return
		}
	}
}
