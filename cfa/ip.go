package cfa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ipResponse struct {
	IP string `json:"ip"`
}

// GetIP gets your IP
func GetIP() string {
	resp, err := http.Get("https://levi.lol/ip")
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	ipr := &ipResponse{}
	err = json.Unmarshal(body, ipr)

	return ipr.IP
}
