package main

import (
	"encoding/json"
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

type YouTube struct {
	ArrayIp []string `json:"youtube.com"`
}

func main() {
	body := request("https://iplist.opencck.org/?format=json&data=ip4&site=youtube.com")
	var ips YouTube
	json.Unmarshal(body, &ips)
	for _, ip := range ips.ArrayIp {
		url := "http://ip-api.com/json/" + ip
		info := request(url)
		log.Println(string(info))
	}
}
