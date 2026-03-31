package main

import (
	"io"
	"log"
	"net/http"
)

// TODO: https://iplist.opencck.org/ru
// TODO: http://ip-api.com

func request(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("error:", err)
		return []byte{}
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("error:", err)
		return []byte{}
	}
	return body
}

func main() {
	request("https://iplist.opencck.org/?format=json&data=ip4&site=youtube.com")
}
