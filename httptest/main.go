package main

import "example/go-gin/httptest/restclient"

func main() {
	body := `{"foo":"bar"}`
	restclient.Post("", body, nil)
}
