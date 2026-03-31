package main

import (
	"log"
)

type Location struct {
	Country  string
	TimeZone string
	Cities   []string
}
type Config struct {
	FontSize    uint8
	ModeTheme   string
	LocationSys Location
}

func (config *Config) appendCity(newCity string) bool {
	for _, city := range config.LocationSys.Cities {
		if city == newCity {
			return false
		}
	}
	config.LocationSys.Cities = append(
		config.LocationSys.Cities, newCity,
	)
	return true
}

func (config *Config) replaceTheme(newTheme string) bool {
	themes := [3]string{"auto", "dark", "light"}
	for _, value := range themes {
		if value == newTheme {
			config.ModeTheme = newTheme
			return true
		}
	}
	return false
}

func main() {
	config := Config{ModeTheme: "auto"}
	config.LocationSys.TimeZone = "Europe/Moscow"
	config.replaceTheme("dark")
	config.appendCity("dark")
	log.Println(config)
}
