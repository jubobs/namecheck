package mock

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jubobs/namecheck"
)

type clientFunc func(*http.Request) (*http.Response, error)

func (f clientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func ClientWithError(err error) namecheck.HTTPClient {
	do := func(_ *http.Request) (*http.Response, error) {
		return nil, err
	}
	return clientFunc(do)
}

func ClientWithStatusCode(code int) namecheck.HTTPClient {
	do := func(_ *http.Request) (*http.Response, error) {
		res := http.Response{
			StatusCode: code,
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}
		return &res, nil
	}
	return clientFunc(do)
}
