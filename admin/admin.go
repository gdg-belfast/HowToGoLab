package admin

import (
	"fmt"
	"github.com/gdg-belfast/HowToGoLab/redirector/usecases"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// has form data been submitted?
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

	// gather redirects from our database
	data, err := db.Read()
	if err != nil {
		panic(err)
	}
	type templateData struct {
		Redirects map[string]string
	}

	// display something to the user
	var templates = template.Must(template.ParseFiles("admin/view/form.html"))
	if err := templates.ExecuteTemplate(w, "adminform", templateData{Redirects: data}); err != nil {
		panic(err)
	}
}

// start a webserver that will allow us to administer our redirector
func Start(mux *http.ServeMux) {

	fmt.Println("Adding admin handlers")

	mux.HandleFunc("/admin", handler)
}
