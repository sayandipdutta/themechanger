/*
@Author: Sayandip Dutta.

Change theme in OneCommander, Spyder, PythonIDLE, WindowsTerminal
from light to dark or dark to light, based on the theme name passed as argument.

Usage:
	.\themeChange.exe --theme light			// Windows
	./themeChange --theme dark				// Linux

	go run themeChange.go -theme=light -program="oneCommander spyder" {CURRENTLY NOT SUPPORTED}
*/

package main

import (
	"flag"
	"log"
	"reflect"

	"fmt"

	"os"

	"github.com/sayandipdutta/themechanger/config"
	"github.com/sayandipdutta/themechanger/themeable"
)

var Logger *log.Logger
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

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	var conf map[string]themeable.ThemeConfig
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println("ERROR: Whie loading config ->", err)
		return
	}

	var programs = map[string]themeable.Themeable{
		"OneCommander": themeable.OneCommander{
			ThemeConfig: conf["OneCommander"],
		},
		"PythonIDLE": themeable.PythonIDLE{
			ThemeConfig: conf["PythonIDLE"],
		},
		"Spyder": themeable.Spyder{
			ThemeConfig: conf["Spyder"],
		},
		"WindowsTerminal": themeable.WindowsTerminal{
			ThemeConfig: conf["WindowsTerminal"],
		},
	}

	for _, program := range programs {
		if err := themeable.SetTheme(program, themeFlag); err != nil {
			Logger.Println(reflect.TypeOf(program), "->", err)
		}
		// fmt.Println(program)
	}
}
