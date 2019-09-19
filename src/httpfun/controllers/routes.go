package controllers

import (
	"fmt"
	"net/http"
)

// MyServer is structure
type MyServer struct {
	Myrouters *http.ServeMux
}

// Preparedata function is to return "Prepareddata for + input string"
func Preparedata(s string) string {
	return "Preparedata for " + s
}

// InitializeRoutes is main function to initialize all http routes
func (s *MyServer) InitializeRoutes() {
	s.Myrouters.HandleFunc("/api/", s.handleAPI())
	s.Myrouters.HandleFunc("/about", s.handleAbout())
	s.Myrouters.HandleFunc("/", s.handleIndex())
	http.
	//	s.Myrouters.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
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
