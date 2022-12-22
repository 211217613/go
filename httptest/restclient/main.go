package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

// post to github api
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}

	request.Header = headers

	return Client.Do(request)
}
