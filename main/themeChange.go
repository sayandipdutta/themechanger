/*
@Author: Sayandip Dutta.

Change theme in OneCommander, Spyder, PythonIDLE, WindowsTerminal
from light to dark or dark to light, based on the theme name passed as argument.

Usage:
	.\themeChange.exe --theme light			// Windows

	go run themeChange.go -theme=light -program="OneCommander Spyder"
*/

package main

import (
	"flag"
	"log"
	"reflect"
	"strings"

	"fmt"

	"os"

	"github.com/sayandipdutta/themechanger/config"
	"github.com/sayandipdutta/themechanger/themeable"
)

var Logger *log.Logger
var themeFlag string
var commandLineProgs string

// Help message
func helpMessage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
	fmt.Println("Accepted value of theme flag: \n\tlight\n\tdark\t//default value.\n ")
	fmt.Println("Accepted value of program flag: \n\tall\t\t\t\t\t//default value. Theme all programs.")
	fmt.Println("or,\t\"<program_name> <program_name2> ...\"\tTheme only specified programs.\n ")
	fmt.Println(`program should be provided in double quotes ("") if it contains spaces.`)
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Println("\n.\\themeChange.exe \t\t\t//default value of theme == dark, program == all")
	fmt.Println(".\\themeChange.exe --theme=dark\t\t//default value of program == all")
	fmt.Println(".\\themeChange.exe --theme=light --program=\"OneCommander Spyder\"")
}

// Parse flags, Initialize logger.
func init() {
	/*
		Provide command line argument -theme [dark/light] to set theme
		Default: dark
	*/
	flag.Usage = helpMessage
	flag.StringVar(&themeFlag, "theme", "dark", "Type of theme to be set")
	flag.StringVar(&commandLineProgs, "program", "all", "Program to be themed")
	flag.Parse()
	// Enable logging

	Logger = config.SetLogger()
}

// Validate flags
// Given theme and program flags, validate them.
// If theme or program is invalid, return error.
func validateFlags(theme string, program string) error {
	if theme != "light" && theme != "dark" {
		return fmt.Errorf("%s.SetTheme: invalid theme name: %s. Use -theme=light or -theme=dark", reflect.TypeOf(program), theme)
	}
	if program == "all" {
		return nil
	}
	if err := config.ValidateProgramFlag(strings.Split(program, " ")); err != nil {
		return err
	}
	return nil
}

func main() {
	var conf map[string]themeable.ThemeConfig
	conf, err := themeable.LoadConfig()
	if err != nil {
		Logger.Fatalln("ERROR: Whie loading config ->", err)
	}

	programs := make(map[string]themeable.Themeable, len(conf))
	for key, value := range conf {
		programs[key] = themeable.Registry[key](value)
	}
	_, p, err := config.GetListedPrograms()
	if err != nil {
		Logger.Fatalln("ERROR: While getting listed programs ->", err)
	}
	if err := validateFlags(themeFlag, commandLineProgs); err != nil {
		Logger.Fatalln("ERROR: While validating flags ->", err)
	}
	if commandLineProgs != "all" {
		p = strings.Split(strings.Trim(commandLineProgs, " "), " ")
	}

	for _, program := range p {
		if err := themeable.SetTheme(programs[program], themeFlag); err != nil {
			Logger.Fatalln(reflect.TypeOf(program), "->", err)
		}
	}
}
