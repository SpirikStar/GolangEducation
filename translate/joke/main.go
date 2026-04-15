package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Joke struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func getReponse(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}
	}
	if response.StatusCode != http.StatusOK {
		return []byte{}
	}
	data, _ := io.ReadAll(response.Body)
	return data
}
func main() {
	result := getReponse("https://official-joke-api.appspot.com/random_joke")
	var joke Joke
	json.Unmarshal(result, &joke)
}
