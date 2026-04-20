package main

// TODO: go mod init avatars
// TODO: go run .

const (
	url = "https://avatars.githubusercontent.com/u/"
)

func main() {
	for id := 1; id < 10; id++ {
		getByteImage(id)
	}
}
