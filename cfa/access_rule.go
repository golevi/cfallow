package cfa

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// AccessRule represents a Cloudflare access rule.
type AccessRule struct {
	ID            string           `json:"id"`
	Notes         string           `json:"notes"`
	Mode          string           `json:"mode"`
	Configuration AccessRuleConfig `json:"configuration"`
}

// AccessRuleResponse is a response.
type AccessRuleResponse struct {
	Success bool         `json:"success"`
	Result  []AccessRule `json:"result"`
}

// AccessRuleRequest represents a Cloudflare access rule.
type AccessRuleRequest struct {
	ID            string           `json:"-"`
	Notes         string           `json:"notes"`
	Mode          string           `json:"mode"`
	Configuration AccessRuleConfig `json:"configuration"`
}

// AccessRuleConfig is a key/value rule.
type AccessRuleConfig struct {
	Target string `json:"target"`
	Value  string `json:"value"`
}

// GetRulesByNote _
func GetRulesByNote(account string, note string) *AccessRuleResponse {
	req, err := http.NewRequest("GET", url+"accounts/"+account+"/firewall/access_rules/rules?notes="+note, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var ar = &AccessRuleResponse{}
	err = json.Unmarshal(body, &ar)

	return ar
}

// ListAccessRules search, sort, and filter IP/country access rules
// GET accounts/:account_identifier/firewall/access_rules/rules
func ListAccessRules(account string) *AccessRuleResponse {
	req, err := http.NewRequest("GET", url+"accounts/"+account+"/firewall/access_rules/rules", nil)
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

	var ar = &AccessRuleResponse{}
	err = json.Unmarshal(body, &ar)

	return ar
}

// CreateAccessRule Make a new IP, IP range, or country access rule for all
// zones owned by the account.
// POST accounts/:account_identifier/firewall/access_rules/rules
func CreateAccessRule(account string, ip string, note string) bool {
	ipa := net.ParseIP(ip)
	target := "ip6"
	if ipa.To4() != nil {
		target = "ip"
	}

	config := AccessRuleConfig{
		Target: target,
		Value:  ip,
	}

	accessRule := &AccessRuleRequest{
		Mode:          "whitelist",
		Configuration: config,
		Notes:         note,
	}

	data, err := json.Marshal(accessRule)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", url+"accounts/"+account+"/firewall/access_rules/rules", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var ar = &AccessRuleResponse{}
	err = json.Unmarshal(body, &ar)

	return ar.Success
}

// // UpdateAccessRule update rule state and/or configuration. This will be applied
// // across all zones owned by the account.
// // PATCH accounts/:account_identifier/firewall/access_rules/rules/:identifier
// func UpdateAccessRule() {

// }

// DeleteAccessRule remove an access rule so it is no longer evaluated during
// requests. This will apply to all zones owned by the account.
// DELETE accounts/:account_identifier/firewall/access_rules/rules/:identifier
func DeleteAccessRule(account string, id string) bool {
	req, err := http.NewRequest("DELETE", url+"accounts/"+account+"/firewall/access_rules/rules/"+id, nil)
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

	var ar = &AccessRuleResponse{}
	err = json.Unmarshal(body, &ar)

	return ar.Success
}

// NewIPAccessRule _
func NewIPAccessRule(ip string) *AccessRuleConfig {
	return &AccessRuleConfig{
		Target: "ip",
		Value:  ip,
	}
}

// NewIP6AccessRule _
func NewIP6AccessRule(ip string) *AccessRuleConfig {
	return &AccessRuleConfig{
		Target: "ip6",
		Value:  ip,
	}
}
