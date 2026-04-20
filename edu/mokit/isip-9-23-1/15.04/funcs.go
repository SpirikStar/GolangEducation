package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func downloadAvatar(id int, isCreated bool, wg *sync.WaitGroup) any {
	defer wg.Done()
	code := getAvatar(id)
	if !isCreated {
		return code
	}
	if len(code) == 0 {
		return false
	}
	defaultPhoto, _ := os.ReadFile("default/default.jpg")
	if string(defaultPhoto) == string(code) {
		return false
	}
	test, _ := os.ReadFile("avatars/100088.jpg")
	fmt.Println(test)
	path := folderAvatars + fmt.Sprintf("/%d.jpg", id)
	err := os.WriteFile(path, code, 0644)
	if err != nil {
		return false
	}
	return true
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
