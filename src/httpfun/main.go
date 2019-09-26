package main

import (
	"fmt"
	"httpfun/controllers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		cleanup(sig)
		os.Exit(1)
	}()

	//myrouter := controllers.MyServer{http.NewServeMux()}
	myrouter := controllers.MyServer{mux.NewRouter()}
	myrouter.InitializeRoutes()
	fmt.Printf("http server is up on http://localhost:8080\n")
	http.Handle("/", myrouter.Myrouters)
	log.Fatal(http.ListenAndServe(":8080", myrouter.Myrouters))

}

func cleanup(sig os.Signal) {
	fmt.Println("some one hit ctrl-c")
	fmt.Println(sig)
}
