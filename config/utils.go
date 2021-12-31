package config

import (
	"fmt"
	"log"
	"os"

	"github.com/sayandipdutta/themechanger/setup"
	"github.com/sayandipdutta/themechanger/themeable"
)

func SetLogger() *log.Logger {
	file, err := os.OpenFile(setup.GetParentDir().LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

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
