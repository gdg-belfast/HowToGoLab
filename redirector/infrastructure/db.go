package infrastructure

import (
	"github.com/gdg-belfast/HowToGoLab/redirector/domain"

	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

const (
	dbPath string = "./data.json"
)

type JsonStorage struct {
	filepath string
}

func NewJsonStorage() *JsonStorage {
	return &JsonStorage{
		filepath: dbPath,
	}
}

func (jstore *JsonStorage) FromJson() (map[string]*domain.HostMap, error) {
	scratch := &tempJsonStorage{}
	redirects := make(map[string]*domain.HostMap, 0)

	if content, err := ioutil.ReadFile(jstore.filepath); err != nil || len(content) == 0 {
		return redirects, err
	} else {
		if err := json.Unmarshal(content, &scratch); err != nil {
			return redirects, err
		}
	}
	log.Println(scratch)
	for _, mapping := range scratch.Mappings {
		redirects[mapping.From] = mapping
	}
	return redirects, nil
}

func (jstore *JsonStorage) ToJson(mappings map[string]*domain.HostMap) error {
	scratch := &tempJsonStorage{}
	for _, val := range mappings {
		scratch.Mappings = append(scratch.Mappings, val)
	}
	if jstr, err := json.Marshal(scratch); err != nil {
		return err
	} else {
		// open our storage file so that we can append to it
		if err = ioutil.WriteFile(jstore.filepath, jstr, 0644); err != nil {
			return err
		}
	}
	return nil
}

type tempJsonStorage struct {
	Mappings []*domain.HostMap `json:"Maps"`
}

// Create saves a new hostmap to the JSON file
func (jstore *JsonStorage) Create(hostmap *domain.HostMap) error {
	hosts, err := jstore.FromJson()
	if err != nil {
		return err
	}

	if _, found := hosts[hostmap.From]; !found {
		hosts[hostmap.From] = hostmap
		if err = jstore.ToJson(hosts); err != nil {
			return errors.New("Could not create mapping")
		}
	}
	return nil
}

// Read looks for a value inside the JSON
func (jstore *JsonStorage) Read(host string) (*domain.HostMap, error) {
	hosts, err := jstore.FromJson()
	if err != nil {
		return &domain.HostMap{}, err
	}

	if foundHost, found := hosts[host]; !found {
		return &domain.HostMap{}, errors.New("Host not found")
	} else {
		return foundHost, nil
	}
}

func (jstore *JsonStorage) Update(hostmap *domain.HostMap) error {
	hosts, err := jstore.FromJson()
	if err != nil {
		return err
	}

	if _, found := hosts[hostmap.From]; !found {
		return errors.New("Original does not exist")
	}

	hosts[hostmap.From] = hostmap
	if err = jstore.ToJson(hosts); err != nil {
		return errors.New("Could not update mapping")
	}
	return nil
}

func (jstore *JsonStorage) Delete(hostmap *domain.HostMap) error {
	hosts, err := jstore.FromJson()
	if err != nil {
		return err
	}

	if _, found := hosts[hostmap.From]; !found {
		return errors.New("Original does not exist")
	}

	delete(hosts, hostmap.From)
	if err = jstore.ToJson(hosts); err != nil {
		return errors.New("Could not delete mapping")
	}
	return nil
}
