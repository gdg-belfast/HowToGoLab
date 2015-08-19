package main

import (
	"github.com/gdg-belfast/HowToGoLab/redirector/server"
)

func main() {

	// set the redirects
	server.SetRedirects(map[string]string{
		"127.0.0.1:8080/": "https://google.com",
		"localhost:8080/": "https://rehabstudio.com",
	})

	// start listening
	server.Start(8080)
}
