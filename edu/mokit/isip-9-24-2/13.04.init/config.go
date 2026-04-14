package main

var Config = make(map[string]string)

func init() {
	Config["protocol"] = "http"
	Config["host"] = "localhost"
	Config["port"] = "8000"
}
