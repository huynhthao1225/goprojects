package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func (p Person) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
func main() {
	router := http.NewServeMux()
	thandler := &timeHandler{format: time.RFC1123}
	router.Handle("/time/", thandler)

	person := &Person{"Thao", "Huynht"}
	router.Handle("/me/", person)

	book := new(Book)
	book.Author = "Die for you"
	book.Title = "Thao Manh Huynh"
	router.Handle("/book/", book)

	router.HandleFunc("/", hello)
	router.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "welcome to the home page")
}
func ping(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "pong Pong ...")
}
