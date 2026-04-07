package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getSource(url string) []byte {
	file, err := os.OpenFile("debag.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err == nil {
		defer file.Close()
		log.SetOutput(file)
	}
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	if response.StatusCode != 200 {
		log.Println(response.StatusCode)
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
func main() {
	config := struct {
		UrlApi string
	}{
		UrlApi: "https://iplist.opencck.org/?format=json&data=ip4&site=jut.su",
	}
	body := getSource(config.UrlApi)

	portal := struct {
		IpList []string `json:"jut.su"`
	}{}

	json.Unmarshal(body, &portal)

	for _, ip := range portal.IpList {
		// TODO: получить сведения из ip используя api http://ip-api.com
		// TODO: получить: timezone, city
		// TODO: В терминале отображать:
		// TODO: Выполняется(номер)/СколькоВыполнить(номер) - готово
		// TODO: Пример:
		// TODO: 17/143 - готово
		fmt.Println(ip)
	}
}
