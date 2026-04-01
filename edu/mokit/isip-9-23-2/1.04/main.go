package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// ! https://ip-api.com
// ! https://iplist.opencck.org/ru
func getSource(url string) []byte {
	file, _ := os.OpenFile("debag.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	defer file.Close()
	log.SetOutput(file)

	response, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return []byte{}
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Println(response.StatusCode)
		return []byte{}
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	log.Println(response.StatusCode)
	return body
}

func main() {
	body := getSource("https://iplist.opencck.org/?format=json&data=ip4&site=youtube.com")
	ips := struct {
		Servers []string `json:"youtube.com"`
	}{}

	json.Unmarshal(body, &ips)
	for _, ip := range ips.Servers {
		bodyIp := getSource("http://ip-api.com/json/" + ip)
		time.Sleep(1 * time.Second)
		fmt.Println(string(bodyIp))
	}
	//  Получить информацию со всех ip адресов
	//  Получить только timezone
	//  Каждую timezone нужно сохранить в отдельный массив
	//  В массиве не должны быть дубликаты
	// 	Необходимо в терминале писать о процессе:
	// 	Выполняется/СколькоВыполнять - готово
	// 	Например: 17/3100 - готово
}
