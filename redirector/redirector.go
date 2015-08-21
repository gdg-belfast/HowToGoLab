package main

import (
	"github.com/gdg-belfast/HowToGoLab/redirector/admin"
	"github.com/gdg-belfast/HowToGoLab/redirector/logger"
	"github.com/gdg-belfast/HowToGoLab/redirector/proxy"
)

func main() {

	// set the redirects
	proxy.SetRedirects(map[string]string{
		"127.0.0.1:8080/": "https://google.com",
		"localhost:8080/": "https://rehabstudio.com",
	})

	// start proxy
	proxy.Start(8080)

	// start admin
	admin.Start(8000)
}
