package config

import (
	"fmt"
	"log"
	"os"

	"github.com/sayandipdutta/themechanger/setup"
	"github.com/sayandipdutta/themechanger/themeable"
)

// SetLogger sets the logger for the program
func SetLogger() *log.Logger {
	file, err := os.OpenFile(setup.Setup().LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// GetListedPrograms returns the list of programs and their themes
// If the config file does not exist, it will return an error
// If the config file is not valid, it will return an error
func GetListedPrograms() (map[string]struct{}, []string, error) {
	conf, err := themeable.LoadConfig()
	if err != nil {
		return nil, nil, err
	}
	listedPrograms := make(map[string]struct{}, len(conf))
	var listedProgramsSlice []string
	for key := range conf {
		listedPrograms[key] = struct{}{}
		listedProgramsSlice = append(listedProgramsSlice, key)
	}

	return listedPrograms, listedProgramsSlice, nil
}

// ValidateProgramFlag validates the program flag
// Given a slice of strings
// IF any of the strings are not valid programs, it will return an error
func ValidateProgramFlag(programs []string) error {
	listedPrograms, _, err := GetListedPrograms()
	if err != nil {
		return err
	}
	for _, program := range programs {
		if _, ok := listedPrograms[program]; !ok {
			return fmt.Errorf("invalid program name: %s", program)
		}
	}
	return nil
}
