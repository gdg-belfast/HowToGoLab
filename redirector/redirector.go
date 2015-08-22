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
	if err := proxy.Start(mux); err != nil {
		panic(err)
	}

	// start admin
	if err := admin.Start(mux); err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", mux)
}
