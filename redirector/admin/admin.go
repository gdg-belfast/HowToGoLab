package admin

import (
	"fmt"
	"github.com/gdg-belfast/HowToGoLab/redirector/db"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("from:", r.Form["from[]"])
		fmt.Println("to:", r.Form["to[]"])

		data := make(map[string]string)
		for i, v := range r.Form["from[]"] {
			data[v] = r.Form["to[]"][i]
		}
		if err := db.Save(data); err != nil {
			panic(err)
		}
	}

	// read data
	data, err := db.Read()
	if err != nil {
		panic(err)
	}
	type templateData struct {
		Redirects map[string]string
	}

	var templates = template.Must(template.ParseFiles("admin/view/form.html"))
	if err := templates.ExecuteTemplate(w, "adminform", templateData{Redirects: data}); err != nil {
		panic(err)
	}
}

// start a webserver that will allow us to administer our redirector
func Start(port int) error {

	adminMux := http.NewServeMux()

	listenAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Admin started: %s", listenAddr)

	adminMux.HandleFunc("/", handler)

	go func() {
		http.ListenAndServe(listenAddr, adminMux)
	}()

	return nil
}
