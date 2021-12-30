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
	"flag"
	"log"
	"reflect"

	"fmt"

	"os"

	"github.com/sayandipdutta/themechanger/config"
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

func main() {
	var conf map[string]themeable.ThemeConfig
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println("ERROR: Whie loading config ->", err)
		return
	}

	var oneCom themeable.Themeable = themeable.OneCommander{
		ThemeConfig: conf["OneCommander"],
	}
	var pyIDLE themeable.Themeable = themeable.PythonIDLE{
		ThemeConfig: conf["PythonIDLE"],
	}
	var spyder themeable.Themeable = themeable.Spyder{
		ThemeConfig: conf["Spyder"],
	}
	var winterm themeable.Themeable = themeable.WindowsTerminal{
		ThemeConfig: conf["WindowsTerminal"],
	}

	var programs = map[string]themeable.Themeable{
		"OneCommander":    oneCom,
		"PythonIDLE":      pyIDLE,
		"Spyder":          spyder,
		"WindowsTerminal": winterm,
	}
	for _, program := range programs {
		if err := themeable.SetTheme(program, themeFlag); err != nil {
			log.Println("ERROR:", reflect.TypeOf(program), "->", err)
		}
		// fmt.Println(program)
	}
}
