package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

//MyServer is structure
//type MyServer struct {
//	Myrouters *http.ServeMux
//}
// MyServer is struct data type for mux.Router
type MyServer struct {
	Myrouters *mux.Router
}

// Preparedata function is to return "Prepareddata for + input string"
func Preparedata(s string) string {
	return "Preparedata for " + s
}

// InitializeRoutes is main function to initialize all http routes
func (s *MyServer) InitializeRoutes() {
	s.Myrouters.HandleFunc("/api/", s.handleAPI())
	s.Myrouters.HandleFunc("/about/", s.handleAbout())
	s.Myrouters.HandleFunc("/", s.handleIndex())
	s.Myrouters.HandleFunc("/index/", s.serveTemplate())
	s.Myrouters.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	
}

func (s *MyServer) handleAPI() http.HandlerFunc {
	thing := Preparedata("handleAPI")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(thing))
	}
}

func (s *MyServer) handleAbout() http.HandlerFunc {
	thing := Preparedata("handleAbout")
	return func(wr http.ResponseWriter, r *http.Request) {
		wr.Write([]byte(thing))
	}
}

func (s *MyServer) handleIndex() http.HandlerFunc {

	return func(wr http.ResponseWriter, r *http.Request) {
		lp := filepath.Join("templates", "index.html")
		fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

		//indexTmpl, err := template.New("index").ParseFiles(lp)

		tmpl, err := template.ParseFiles(lp)
		if err != nil {
			fmt.Printf("Error = %s\t template = %s filePath = %s URL.Path = %s\n", err.Error(), lp, fp, r.URL.Path)
			fmt.Fprintf(wr, "Page not found")
			return
		}

		err = tmpl.Execute(wr, nil)
		//err = indexTmpl.Execute(wr, nil)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Fprintf(wr, "Error ExecuteTemplate")
		}
		return
	}
}
func (s *MyServer) serveTemplate() http.HandlerFunc {

	return func(wr http.ResponseWriter, r *http.Request) {
		lp := filepath.Join("./templates", "index.html")
		fp := filepath.Join("./templates", filepath.Clean(r.URL.Path))

		// Return a 404 if the template doesn't exist
		info, err := os.Stat(fp)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(wr, r)
				return
			}
		}

		// Return a 404 if the request is for a directory
		if info.IsDir() {
			http.NotFound(wr, r)
			return
		}

		tmpl, err := template.ParseFiles(lp, fp)
		if err != nil {
			// Log the detailed error
			log.Println(err.Error())
			// Return a generic "Internal Server Error" message
			http.Error(wr, http.StatusText(500), 500)
			return
		}

		if err := tmpl.ExecuteTemplate(wr, "index", nil); err != nil {
			log.Println(err.Error())
			http.Error(wr, http.StatusText(500), 500)
		}
		return
	}
}
