package main

import (
	"io"
	"net/http"
)

var baseUrl string = "https://avatars.githubusercontent.com/u/"

func getAvatar(id int) []byte {
	url := baseUrl + string(id)
	response, err := http.Get(url)
	if err != nil {
		return []byte{}
	}
	if response.StatusCode != http.StatusOK {
		return []byte{}
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}
	}
	return data
}
