package main

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
