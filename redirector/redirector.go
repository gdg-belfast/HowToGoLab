package main

import (
	//"github.com/gdg-belfast/HowToGoLab/redirector/admin"
	"github.com/gdg-belfast/HowToGoLab/redirector/service"

	"log"
	"os"
	"strconv"
)

func main() {

	log.Println("[ Proxy demo ]")

	var port int64
	var err error

	if port, err = getPort(); err != nil {
		log.Fatalln("Could not get listening port: ", err)
	} else {
		log.Println("Port set to", port)
	}

	redirectorService := service.NewRedirectorService()
	// add admin handlers
	//admin.Start(mux)
	err = redirectorService.Start(port)
	if err != nil {
		log.Fatalln(err)
	}
}

func getPort() (int64, error) {
	port := os.Getenv("REDIRECTOR_PORT")
	if port == "" {
		return 8080, nil
	}
	return strconv.ParseInt(port, 10, 0)
}
