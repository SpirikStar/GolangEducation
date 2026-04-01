package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ! https://ip-api.com
// ! https://iplist.opencck.org/ru
func getSource(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}
	}
	if response.StatusCode != 200 {
		return []byte{}
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}
	}
	return body
}
func main() {
	body := getSource("https://iplist.opencck.org/?format=json&data=ip4&site=youtube.com")

	ips := struct {
		Servers []string `json:"youtube.com"`
	}{}

	json.Unmarshal(body, &ips)
	log.Println(ips)
}
