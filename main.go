/*
@Author: Sayandip Dutta.

Change theme of applications based on config from light to dark or vice versa
to light, when theme is passed as argument

Usage:
	.\themeChange.exe --theme light
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var themeFlag *bool

// Help message
func helpMessage() {
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Println(".\\themeChange.exe")
	fmt.Println(".\\themeChange.exe -light")
}

func main() {
	flag.Usage = helpMessage
	themeFlag = flag.Bool("light", false, "If given, `light` theme will be set.")
	flag.Parse()

	file, err := os.OpenFile("themechanger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Printf("INFO: Application started successfully!")

	var conf map[string]themeConfig
	conf, err = loadConfig()
	if err != nil {
		log.Fatalln("ERROR: Whie loading config ->", err)
	}

	theme := "dark"
	if *themeFlag {
		theme = "light"
	}

	for programName, config := range conf {
		err := config.setTheme(*themeFlag)
		if err != nil {
			log.Println(programName, "->", err)
		} else {
			log.Printf("INFO: successfully changed theme of %s to %s\n", programName, theme)
		}
	}
}
