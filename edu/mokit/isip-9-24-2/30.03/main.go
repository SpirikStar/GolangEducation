package main

import "log"

type ConfSys struct {
	Theme string
}

func (conf *ConfSys) replaceTheme(theme string) {
	conf.Theme = theme
}

func replTheme(theme string, conf *ConfSys) {
	conf.Theme = theme
}

func main() {
	usConfig := ConfSys{}
	usConfig.replaceTheme("auto")
	replTheme("auto", &usConfig)
	
	log.Println(usConfig)

	config := struct {
		FontSize  uint8
		ModeTheme string
		TimeZone  struct {
			Zone, Local string
		}
	}{
		FontSize: 14,
		TimeZone: struct {
			Zone, Local string
		}{Local: "Russian"},
	}
	config.TimeZone.Zone = "EuropeMoscow"
	config.ModeTheme = "auto"

	log.Println(config)
}
