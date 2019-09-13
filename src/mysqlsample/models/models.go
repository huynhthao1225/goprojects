package models

import (
	"time"
)

// Actor is data structure of MySql Actor table
type Actor struct {
	ID         int
	FirstName  string
	LastName   string
	LastUpdate time.Time
}

// Address is data structure of MySql Address table
type Address struct {
	ID         int
	Address    string
	Address2   string
	District   string
	CityID     int
	PostalCode string
	Phone      string
	Location   []byte
	LastUpdate time.Time
}
