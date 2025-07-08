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
	"io"
	"log"
	"os"
	"path/filepath"
)

var themeFlag *bool

func helpMessage() {
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Println(".\\themeChange.exe")
	fmt.Println(".\\themeChange.exe -light")
}

func getLogFolder() (string, bool) {
	appDataDir := os.Getenv("LOCALAPPDATA")
	if appDataDir == "" {
		return "", false
	}
	logfolder := filepath.Join(appDataDir, "themechanger")
	if info, err := os.Stat(logfolder); err != nil || !info.IsDir() {
		err = os.Mkdir(logfolder, 0666)
		if err != nil {
			return "", false
		}
	}
	return logfolder, true
}

func main() {
	flag.Usage = helpMessage
	themeFlag = flag.Bool("light", false, "If given, `light` theme will be set, otherwise `dark`.")
	flag.Parse()

	logfolder, ok := getLogFolder()
	if ok {
		logfile := filepath.Join(logfolder, "themechanger.log")
		logwriter, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer logwriter.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(io.MultiWriter(os.Stderr, logwriter))
	} else {
		log.SetOutput(os.Stderr)
	}
	log.Printf("INFO: Application started successfully!")

	var conf map[string]themeConfig
	conf, err := loadConfig()
	if err != nil {
		log.Fatalln("WARNING: Error while loading config ->", err, " Skipping!")
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
