package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var db string = "./data.json"

// Save json data to our file
func Save(redirects map[string]string) error {
	fmt.Println("saving...")
	// convert map to string
	if jstr, err := json.Marshal(redirects); err != nil {
		return err
	} else {
		// open our storage file so that we can append to it
		if err = ioutil.WriteFile(db, jstr, 0644); err != nil {
			return err
		}
	}
	return nil
}

// Read json data from the file
func Read() (map[string]string, error) {
	redirects := make(map[string]string, 0)
	if content, err := ioutil.ReadFile(db); err != nil {
		return redirects, err
	} else if string(content) != "" {
		if err := json.Unmarshal(content, &redirects); err != nil {
			return redirects, err
		}
	}
	return redirects, nil
}
