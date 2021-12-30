package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/sayandipdutta/themechanger/themeable"
)

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

func GetListedPrograms() (map[string]struct{}, error) {
	conf, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	t := reflect.TypeOf(conf)
	listedPrograms := make(map[string]struct{}, t.NumField())
	for key := range listedPrograms {
		listedPrograms[key] = struct{}{}
	}
	return listedPrograms, nil
}

func ValidateProgramFlag(programs map[string]struct{}) error {
	listedPrograms, err := GetListedPrograms()
	if err != nil {
		return err
	}
	for program := range programs {
		if _, ok := listedPrograms[program]; !ok {
			return fmt.Errorf("invalid program name: %s", program)
		}
	}
	return nil
}
