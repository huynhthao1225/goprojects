package main

import (
	"fmt"
	"httpfun/controllers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		cleanup(sig)
		os.Exit(1)
	}()

	myServer := controllers.InitializeRoutes()
	fmt.Printf("http server is up on http://localhost:8080\n")

	log.Fatal(http.ListenAndServe(":8080", myServer))

}

func cleanup(sig os.Signal) {
	fmt.Println("some one hit ctrl-c")
	fmt.Println(sig)
}