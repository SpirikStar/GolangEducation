package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func getAvatar(id int) []byte {
	url := fmt.Sprintf("%s%d", baseUrl, id)
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
func downloadAvatar(id int) {
	avatar := getAvatar(id)
	if len(avatar) == 0 {
		return
	}
	path := fmt.Sprintf("%s/%d.png", folder, id)
	err := os.WriteFile(path, avatar, 0644)
	if err != nil {
		return
	}
}
