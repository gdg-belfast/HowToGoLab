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

	// start proxy
	proxy.Start(mux)
	// start admin
	admin.Start(mux)

	http.ListenAndServe(":8080", mux)
}
