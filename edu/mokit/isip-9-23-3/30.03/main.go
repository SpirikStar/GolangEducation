package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	cache = make(map[string]string)
	mutex sync.Mutex
)
// TODO: http://ip-api.com/json/
// TODO: http://ip-api.com/json/178.176.73.187
// TODO: http://localhost:8000?ip=178.176.73.187
// ? {"status":"success","country":"The Netherlands","countryCode":"NL","region":"NH","regionName":"North Holland","city":"Amsterdam","zip":"1012","lat":52.3676,"lon":4.90414,"timezone":"Europe/Amsterdam","isp":"Timeweb, LLP","org":"Timeweb, LLP","as":"AS210976 Timeweb, LLP","query":"195.133.41.44"}
func handle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		return
	}
	key := "result:" + id
	
	mutex.Lock()
	if result, ok := cache[key]; ok {
		mutex.Unlock()
		fmt.Fprintf(w, "Result cache: %s\n", result)
	}
	mutex.Unlock()

	time.Sleep(800 * time.Millisecond)
	result := "Result from " + id

	mutex.Lock()
	cache[key] = result
	mutex.Unlock()

	fmt.Fprintf(w, "Вычисление: %s\n", result)
}
func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}
