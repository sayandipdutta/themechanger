/*
@Author: Sayandip Dutta.

Change theme in OneCommander, Spyder, PythonIDLE, WindowsTerminal
from light to dark or dark to light, based on the theme name passed as argument.

Usage:
	.\themeChange.exe --theme light			// Windows
	./themeChange --theme dark				// Linux

	go run themeChange.go -theme=light -program=oneCommander {CURRENTLY NOT SUPPORTED}
*/

package main

import (
	"encoding/json"
	"flag"
	"log"
	"reflect"

	"fmt"

	"io/ioutil"
	"os"

	"github.com/sayandipdutta/themechanger/themeable"
)

var themeFlag string

func helpMessage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
	fmt.Println("Accepted value of theme flag: \n\tlight\n\tdark")

	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}

func init() {
	/*
		Provide command line argument -theme [dark/light] to set theme
		Default: dark
	*/
	flag.Usage = helpMessage
	flag.StringVar(&themeFlag, "theme", "dark", "Type of theme to be set")
	flag.Parse()
}

func LoadConfig() (map[string]themeable.ThemeConfig, error) {
	// read config file (JSON) and set theme
	jsonFile, err := os.Open("config\\config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	result := map[string]themeable.ThemeConfig{}
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	var config map[string]themeable.ThemeConfig
	config, err := LoadConfig()
	if err != nil {
		log.Println("ERROR: Whie loading config ->", err)
		return
	}

	var oneCom themeable.Themeable = themeable.OneCommander{
		ThemeConfig: config["OneCommander"],
	}
	var pyIDLE themeable.Themeable = themeable.PythonIDLE{
		ThemeConfig: config["PythonIDLE"],
	}
	var spyder themeable.Themeable = themeable.Spyder{
		ThemeConfig: config["Spyder"],
	}
	var winterm themeable.Themeable = themeable.WindowsTerminal{
		ThemeConfig: config["WindowsTerminal"],
	}

	var programs = []themeable.Themeable{oneCom, pyIDLE, spyder, winterm}
	for _, program := range programs {
		if err := themeable.SetTheme(program, themeFlag); err != nil {
			log.Println("ERROR:", reflect.TypeOf(program), "->", err)
		}
		// fmt.Println(program)
	}
}
