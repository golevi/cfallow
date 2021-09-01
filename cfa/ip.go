package cfa

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const ipURL = "https://levi.lol/ip"

type ipResponse struct {
	IP string `json:"ip"`
}

// GetIP gets your IP
func GetIP() string {

	resp, err := http.Get(ipURL)
	if err != nil {
		log.Println(err)

		// Allow skipping TLS verfication if there was an error
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}

		resp, err = http.Get(ipURL)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	ipr := &ipResponse{}
	err = json.Unmarshal(body, ipr)

	return ipr.IP
}
