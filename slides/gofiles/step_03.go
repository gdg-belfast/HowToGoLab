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

	w.Write([]byte("{ StatusCode: 0 }"))
}
