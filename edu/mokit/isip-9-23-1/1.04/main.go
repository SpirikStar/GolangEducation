package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// TODO: https://iplist.opencck.org/ru
// TODO: http://ip-api.com

func request(url string) []byte {
	file, _ := os.OpenFile("debag.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	log.SetOutput(file)

	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	log.Println(response.Status)
	return body
}

type Ips struct {
	ArrayIp []string `json:"youtube.com"`
}
// TODO: go mod init project
// TODO: go get github.com/patrickmn/go-cache

func main() {
	cache := make(map[string]string)
	body := request("https://iplist.opencck.org/?format=json&data=ip4&site=youtube.com")
	var ips Ips
	json.Unmarshal(body, &ips)
	for _, ip := range ips.ArrayIp[0:15] {
		if cache[ip] == "" {
			info := request("http://ip-api.com/json/" + ip)
			cache[ip] = string(info)
			fmt.Println("request:", string(info))
		} else {
			fmt.Println("cache:", cache[ip])
		}
	}
}
