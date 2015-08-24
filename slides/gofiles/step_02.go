package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("[ Proxy demo ]")

	mux := http.NewServeMux()

	// start listening & serving
	http.ListenAndServe(":8080", mux)
}
