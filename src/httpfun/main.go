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

	myrouter := controllers.MyServer{http.NewServeMux()}
	myrouter.InitializeRoutes()
	log.Fatal(http.ListenAndServe(":8080", myrouter.Myrouters))

}

func cleanup(sig os.Signal) {
	fmt.Println("some one hit ctrl-c")
	fmt.Println(sig)
}
