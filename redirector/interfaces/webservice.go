package interfaces

import (
	"github.com/gdg-belfast/HowToGoLab/redirector/usecases"

	"log"
	"net/http"
	"strings"
)

type WebService struct {
	MappingInteractor usecases.HostMapInteractor
}

func (ws *WebService) Redirect(w http.ResponseWriter, r *http.Request) {
	hostname := strings.Split(r.Host, ":")[0]
	if redirectUrl, err := ws.MappingInteractor.GetHostMapping(hostname); err != nil {
		log.Println(err)
		http.Error(w, "Glitch in the matrix", http.StatusBadRequest)
	} else {
		log.Println(redirectUrl)
		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
	}
}
