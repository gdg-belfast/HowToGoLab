package proxy

import (
	"fmt"
	"net/http"
)

func Start(mux *http.ServeMux) {
	fmt.Println("Adding proxy handler")
	mux.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	path := r.Host + r.URL.Path
	fmt.Println(fmt.Sprintf("requested %s", path))

	redirects := map[string]string{
		"127.0.0.1:8080/": "https://google.com",
		"localhost:8080/": "https://rehabstudio.com",
	}

	if _, ok := redirects[path]; ok {
		url := redirects[path]
		fmt.Println(fmt.Sprintf(" ~> redirecting %s to %s", path, url))
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		w.Write(nil)
		return
	}
	fmt.Println("request doesnt match any redirects")

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("{ StatusCode: %d }", http.StatusNotFound)))
}
