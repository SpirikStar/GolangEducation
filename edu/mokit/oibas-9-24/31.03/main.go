package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type IpInfo struct {
	Status   string `json:"status"`
	City     string `json:"city"`
	TimeZone string `json:"timezone"`
	Query    string `json:"query"`
}

// TODO: https://iplist.opencck.org
// TODO: Тип данных ipv4
// TODO: youtube
// TODO: Сохранить результат в struct
// TODO: Создать

func requestUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("error:", err)
		return []byte{}
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("error:", err)
		return []byte{}
	}
	return body
}
func main() {
	body := requestUrl("http://ip-api.com/json")

	var info IpInfo
	json.Unmarshal(body, &info)

	log.Println(info.Status)
}
