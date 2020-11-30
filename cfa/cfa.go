package cfa

import (
	"os"
)

var url = "https://api.cloudflare.com/client/v4/"
var token string

func init() {
	token = os.Getenv("CF_ALLOW")

	if len(token) == 0 {
		panic("CF_ALLOW environment variable required")
	}
}
