package main

import (
	"fmt"
	"net/http"
)

type MyServer struct {
	myrouters *http.ServeMux
}

func Preparedata(s string) string {
	return "Preparedata for " + s
}
func (s *MyServer) MyRoutes() {
	s.myrouters.HandleFunc("/api/", s.handleAPI())
	s.myrouters.HandleFunc("/about", s.handleAbout())
	s.myrouters.HandleFunc("/", s.handleIndex())
}

func (s *MyServer) handleAPI() http.HandlerFunc {
	thing := Preparedata("handleAPI")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(thing))
	}
}

func (s *MyServer) handleAbout() http.HandlerFunc {
	thing := Preparedata("handleAbout")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(thing))
	}
}

func (s *MyServer) handleIndex() http.HandlerFunc {
	thing := Preparedata("handleIndex")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fmt.Println("I got request as " + r.URL.Path)
			http.NotFound(w, r)
			return
		}

		w.Write([]byte(thing))
	}
}
