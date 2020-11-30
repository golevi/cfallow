package cfa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Account is a Cloudflare account.
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AccountResponse is a response.
type AccountResponse struct {
	Success bool      `json:"success"`
	Result  []Account `json:"result"`
}

// ListAccounts list all accounts you have ownership or verified access to.
// GET accounts
func ListAccounts() *AccountResponse {
	req, err := http.NewRequest("GET", url+"accounts", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var ar = &AccountResponse{}
	err = json.Unmarshal(body, &ar)

	return ar
}
