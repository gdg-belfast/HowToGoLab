package server

import (
	"fmt"
	"net/http"
)

// Global map of desired redirects
var Redirects map[string]string

// Set the packages redirects
func SetRedirects(routes map[string]string) {
	Redirects = routes
}

// proxyLogger is just a placeholder. we could spit this out into I/O later
func proxyLogger(s string) {
	// just a simple stdout log for now
	fmt.Println(s)
}

// handler catches all requests to our http.ListenAndServe. redirects if
// defined, 404s if not
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.Host + r.URL.Path
	proxyLogger(fmt.Sprintf("requested %s", path))
	if _, ok := Redirects[path]; ok {
		url := Redirects[path]
		proxyLogger(fmt.Sprintf(" ~> redirecting %s to %s", path, url))
		// redirect
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write(nil)
		return
	}
	// 404
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("{ StatusCode: %d }", http.StatusNotFound)))
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
