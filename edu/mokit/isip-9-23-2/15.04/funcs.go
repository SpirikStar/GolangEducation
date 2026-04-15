package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadAvatar(id int, isCreated bool) any {
	code := getAvatar(id)
	if !isCreated {
		return code
	}
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
