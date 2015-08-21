package admin

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the admin panel"))
}

// start a webserver that will allow us to administer our redirector
func Start(port int) error {

	http.HandleFunc("/admin", handler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}
	return nil
}
