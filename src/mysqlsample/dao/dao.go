package dao

import (
	"database/sql"
	"fmt"
	"mysqlsample/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type daoActor interface {
	FindbyID() *models.Actor
	FindAll() []models.Actor
}

// ActionActor will be used to tie actor dao functions
type ActionActor struct{}

// FindbyID function to look for
// id is primary to search for
// db is handle to db connection
// return Actor handle
func (ActionActor) FindbyID(id int, db *sql.DB) *models.Actor {

	var ID int
	var fname, lname string
	var lastupdate time.Time

	actor := models.Actor{}
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
func (ActionActor) FindAll(db *sql.DB) []models.Actor {

	var id int
	var fname, lname string
	var lastupdate time.Time

	selDB, err := db.Query("SELECT actor_id, first_name, last_name, last_update FROM actor ORDER BY actor_id")

	if err != nil {
		panic(err.Error())
	}

	actor := models.Actor{}
	actors := []models.Actor{}

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

//
type daoAddress interface {
	FindbyID() *models.Address
	FindAll() []models.Address
}

// ActionAddress will be used to tie actor dao functions
type ActionAddress struct{}

// FindbyID is to return address struct
func (ActionAddress) FindbyID(id int, db *sql.DB) *models.Address {

	var ID int
	var address string
	var location []byte

	addressRec := models.Address{}
	selDB, err := db.Query("SELECT address_id, address, location FROM address where address_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {

		err = selDB.Scan(&ID, &address, &location)
		if err != nil {
			fmt.Println(err.Error())
			panic(err.Error())
		}
		addressRec.ID = ID
		addressRec.Address = address
		addressRec.Location = location

	}

	return &addressRec
}
