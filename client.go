package namecheck

import (
	"net/http"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var Client HTTPClient = http.DefaultClient
