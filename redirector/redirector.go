package main

import (
	"fmt"
	"github.com/gdg-belfast/HowToGoLab/redirector/admin"
	"github.com/gdg-belfast/HowToGoLab/redirector/proxy"
	"net/http"
)

func main() {

	fmt.Println("[ Proxy demo ]")

	mux := http.NewServeMux()

	// add proxy handlers
	proxy.Start(mux)
	// add admin handlers
	admin.Start(mux)

	// start listening & serving
	http.ListenAndServe(":8080", mux)
}
