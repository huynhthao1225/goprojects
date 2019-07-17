package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type timeHandler struct {
	format string
}

type jsonString struct {
	json string
}

func (th timeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is : " + tm))
}

// Person Firstname is person first name
// Person Lastname is person last name
type Person struct {
	Firstname string
	Lastname  string
}

//Book has title and author
type Book struct {
	Title  string
	Author string
}

func (b Book) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	js, err := json.Marshal(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

var c chan os.Signal
var c1 chan http.ResponseWriter

func main() {
	c = make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go iamdying()
	go keepFeeding()

	router := http.NewServeMux()
	thandler := &timeHandler{format: time.RFC1123}
	router.Handle("/time/", thandler)

	//person := &Person{"Thao", "Huynht"}
	person := new(Person)
	person.Firstname = "Thao"
	person.Lastname = "Huynh"

	router.Handle("/me/", person)

	book := new(Book)
	book.Author = "Die for you"
	book.Title = "Thao Manh Huynh"
	router.Handle("/book/", book)

	router.HandleFunc("/", hello)
	router.HandleFunc("/ping", ping)

	router.HandleFunc("/stopme", stopme)
	router.HandleFunc("/funSlice", funSlice)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "welcome to the home page")
}
func ping(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "pong Pong ...")
}

func (p Person) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func stopme(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(rw, "I am about to be terminated")
	c1 <- rw
	//	time.Sleep(1000 * time.Millisecond)
	c <- os.Interrupt
	close(c1)
	close(c)
}

func keepFeeding() {
	rw := <-c1

	if rw != nil {
		fmt.Fprintln(rw, "I am here ...")
	}
	fmt.Println("I am here")

}
func iamdying() {
	sig := <-c
	cleanup(sig)
	os.Exit(1)
}

func cleanup(sig os.Signal) {
	fmt.Printf("some one ask me to terminate with signal %s\n", sig)
}

func funSlice(rw http.ResponseWriter, req *http.Request) {
	names := []string{"value1", "value2", "value3"}
	writeSlice(rw, "values", names)

}

func writeSlice(rw http.ResponseWriter, firstone string, values []string) {
	fmt.Fprintln(rw, firstone)
	for _, value := range values {
		rw.Write([]byte(value))
	}
}
