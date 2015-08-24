package service

import (
	"github.com/gorilla/mux"

	"github.com/gdg-belfast/HowToGoLab/redirector/infrastructure"
	"github.com/gdg-belfast/HowToGoLab/redirector/interfaces"
	"github.com/gdg-belfast/HowToGoLab/redirector/usecases"

	"fmt"
	"net/http"
)

type RedirectorService struct {
	Router *mux.Router
}

func NewRedirectorService() *RedirectorService {
	service := &RedirectorService{
		Router: mux.NewRouter(),
	}

	mapper := &usecases.HostMapper{
		Repo: infrastructure.NewJsonStorage(),
	}

	redirectHandler := &interfaces.WebService{
		MappingInteractor: mapper,
	}

	service.Router.HandleFunc("/", redirectHandler.Redirect)
	return service
}

func (service *RedirectorService) Start(port int64) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), service.Router)
}
