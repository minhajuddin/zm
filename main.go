package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	dir         string
	port        string
	fileHandler http.Handler
)

func main() {
	var err error
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	port = ":2909"
	fileHandler = http.FileServer(http.Dir(dir))
	log.Printf("starting server at http://localhost%s/", port)

	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(port, nil)
}

var rootHandlers = map[string]http.Handler{"GET": http.HandlerFunc(indexHandler), "POST": http.HandlerFunc(uploadHandler)}

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
	// we serve only the root path
	if r.URL.Path != "/" {
		fileHandler.ServeHTTP(w, r)
		return
	}

	listFilesHandler(w, r)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Awesome")
}

var uploadForm = `
<form enctype="multipart/form-data" method="post">
<input type=file name=file />
<input type=submit value="Upload" />
</form>
`

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, uploadForm)
	fmt.Fprintln(w, "<pre>")

	fmt.Fprintln(w, "<pre>")
	for _, f := range files {
		url := url.URL{Path: f.Name()}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", url.String(), f.Name())
	}
	fmt.Fprintln(w, "</pre>")
}
