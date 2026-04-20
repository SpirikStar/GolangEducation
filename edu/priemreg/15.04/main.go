// TODO: go mod init github
// TODO: go run .
package main

var baseUrl string = "https://avatars.githubusercontent.com/u/"
var folder string = "avatars"

func main() {
	for id := 1; id < 10; id++ {
		downloadAvatar(id)
	}
}
