package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var dir string
var port string

func main() {
	var err error
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	port = ":2909"
	log.Printf("starting server at http://localhost%s/", port)

	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(port, nil)
}

func loadFiles() {
}

var rootHandlers = map[string]http.Handler{"GET": http.FileServer(http.Dir(dir)), "POST": http.HandlerFunc(uploadHandler)}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)

	h, ok := rootHandlers[r.Method]
	if !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	h.ServeHTTP(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Teri hasi")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Awesome")
}
