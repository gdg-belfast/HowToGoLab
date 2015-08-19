package server

import (
	"fmt"
	"log"
	"net/http"
)

// Global map of desired redirects
var Redirects map[string]string

// Set the packages redirects
func SetRedirects(routes map[string]string) {
	Redirects = routes
}

func handler(w http.ResponseWriter, r *http.Request) {
	if _, ok := Redirects[r.Host+r.URL.Path]; ok {
		url := Redirects[r.Host+r.URL.Path]
		w.Header().Set("Location", url)
		fmt.Println(url)
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write(nil)
		return
	}
}

// Start our server
func Start(port int) error {

	http.HandleFunc("/", handler)
	// start listening
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}

	return nil
}

// Performs redirects
func redirectHandler(w http.ResponseWriter, r *http.Request,
	from string, to string, httpcode int) {
	log.Printf("-> redirecting to %s", to)
	http.Redirect(w, r, to, httpcode)
}
