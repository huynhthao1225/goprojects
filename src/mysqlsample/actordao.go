package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// FindbyID function to look for
// id is primary to search for
// db is handle to db connection
// return Actor handle
func FindbyID(id int, db *sql.DB) *Actor {

	var ID int
	var fname, lname string
	var lastupdate time.Time

	actor := Actor{}
	selDB, err := db.Query("SELECT actor_id, first_name, last_name, last_update FROM actor where actor_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {

		err = selDB.Scan(&ID, &fname, &lname, &lastupdate)
		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
		actor.ID = ID
		actor.FirstName = fname
		actor.LastName = lname
		actor.LastUpdate = lastupdate

	}

	return &actor

}

// FindAll get all records from actor table
func FindAll(db *sql.DB) []Actor {

	var id int
	var fname, lname string
	var lastupdate time.Time

	selDB, err := db.Query("SELECT actor_id, first_name, last_name, last_update FROM actor ORDER BY actor_id")

	if err != nil {
		panic(err.Error())
	}

	actor := Actor{}
	actors := []Actor{}

	for selDB.Next() {
		err = selDB.Scan(&id, &fname, &lname, &lastupdate)
		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
		actor.ID = id
		actor.FirstName = fname
		actor.LastName = lname
		actor.LastUpdate = lastupdate
		actors = append(actors, actor)

	}
	return actors
}
