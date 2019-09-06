package main

import (
	"context"
	"fmt"
	"log"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {

	client := GetClient()

	// Check the connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	Disconnect(client)
}
