package usecases

import (
	"github.com/gdg-belfast/HowToGoLab/redirector/domain"
)

// HostMapInteractor defines the interface used to call the business
// logic from the interface layer
type HostMapInteractor interface {
	GetHostMapping(string) (string, error)
}

// HostMapper implements the HostMapInteractor interface
type HostMapper struct {
	Repo domain.HostMapRepository
}

// GetHostMapping takes a URL and looks it up to determine the redirect
// address. If it is not found it returns a blank address and an error
func (mapper *HostMapper) GetHostMapping(hostname string) (string, error) {
	var urlMapping *domain.HostMap
	var err error
	if urlMapping, err = mapper.Repo.Read(hostname); err != nil {
		return "", err
	}
	return urlMapping.To, err
}
