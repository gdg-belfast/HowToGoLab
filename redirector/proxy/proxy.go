package proxy

import (
	"fmt"
	"github.com/gdg-belfast/HowToGoLab/redirector/db"
	"net/http"
)

// handler catches all requests to our http.ListenAndServe. redirects if
// defined, 404s if not
func handler(w http.ResponseWriter, r *http.Request) {

	// gather the redirects every time a request is made - not ideal, but
	// ensures the data is up to date for now
	redirects, err := db.Read()
	if err != nil {
		panic(err)
	}

	// if the request is known, redirect
	path := r.Host + r.URL.Path
	fmt.Println(fmt.Sprintf("requested %s", path))
	if _, ok := redirects[path]; ok {
		url := redirects[path]
		fmt.Println(fmt.Sprintf(" ~> redirecting %s to %s", path, url))
		// redirect
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write(nil)
		return
	}
	fmt.Println("request doesnt match any redirects")
	// 404
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("{ StatusCode: %d }", http.StatusNotFound)))
}

// Start our server
func Start(mux *http.ServeMux) {

	fmt.Println("Adding proxy handler")

	mux.HandleFunc("/", handler)
}
