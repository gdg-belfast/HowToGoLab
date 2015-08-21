package main

import (
	"fmt"
	"github.com/gdg-belfast/HowToGoLab/redirector/admin"
	"github.com/gdg-belfast/HowToGoLab/redirector/proxy"
)

func main() {

	finish := make(chan bool)

	fmt.Println("[ Proxy demo ]")

	// start proxy
	if err := proxy.Start(8080); err != nil {
		panic(err)
	}

	// start admin
	if err := admin.Start(8000); err != nil {
		panic(err)
	}

	<-finish
}
