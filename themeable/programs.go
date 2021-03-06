package themeable

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sayandipdutta/themechanger/setup"
	"gopkg.in/ini.v1"
)

// LoadConfig loads the config file and returns the ThemeConfig
// If the config file does not exist, it will return an error
// If the config file is not valid, it will return an error
func LoadConfig() (map[string]ThemeConfig, error) {
	// read config file (JSON) and set theme
	confpath := setup.Setup().ConfPath
	jsonFile, err := os.Open(confpath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	result := map[string]ThemeConfig{}
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// light and dark theme names, and config file path
type ThemeConfig struct {
	Light      string `json:"light"`      // prefered theme name for light mode
	Dark       string `json:"dark"`       // prefered theme name for dark mode
	ConfigPath string `json:"configpath"` // config file path
}

// GetTheme returns the theme name based on the theme flag
// If the theme flag is not valid, it will return the default theme
func (programTheme ThemeConfig) GetTheme(theme string) string {
	if theme == "light" {
		return programTheme.Light
	}
	return programTheme.Dark
}

// Interface for program that can be themed.
type Themeable interface {
	SetTheme(string) error
}

// SetTheme sets the theme for the program
// Given a themeable program and a theme name, it will set the theme
// If the theme name is not valid, it will return an error
func SetTheme(program Themeable, theme string) error {
	fmt.Println("Setting theme to: ", theme, program)
	err := program.SetTheme(theme)
	return err
}

// Spyder is a program that can be themed
type Spyder struct {
	ThemeConfig // theme config
}

// PythonIDLE is a program that can be themed
type PythonIDLE struct {
	ThemeConfig // theme config
}

// OneCommander is a program that can be themed
type OneCommander struct {
	ThemeConfig // theme config
}

// WindowsTerminal is a program that can be themed
type WindowsTerminal struct {
	ThemeConfig // theme config
}

// OneCommander.SetTheme sets the theme for OneCommander
// Given a theme name, it will set the theme
// If the theme name is not valid, it will return an error
func (programTheme OneCommander) SetTheme(theme string) error {
	// read config file (JSON) and set theme
	jsonFile, err := os.Open(programTheme.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	result := map[string]interface{}{}
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return err
	}

	var NewTheme string = programTheme.GetTheme(theme)
	result["userSettings"].(map[string]interface{})["roaming"].(map[string]interface{})["Rapidrive.Properties.Settings"].(map[string]interface{})["ThemeName"] = NewTheme

	file, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(programTheme.ConfigPath, file, 0644); err != nil {
		return err
	}
	return nil
}

// PythonIDLE.SetTheme sets the theme for PythonIDLE
// Given a theme name, it will set the theme
// If the theme name is not valid, it will return an error
func (programTheme PythonIDLE) SetTheme(theme string) error {
	cfg, err := ini.Load(programTheme.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}

	var NewTheme string = programTheme.GetTheme(theme)
	cfg.Section("Theme").Key("name").SetValue(NewTheme)

	if err = cfg.SaveTo(programTheme.ConfigPath); err != nil {
		return err
	}
	return nil
}

// Spyder.SetTheme sets the theme for Spyder
// Given a theme name, it will set the theme
// If the theme name is not valid, it will return an error
func (programTheme Spyder) SetTheme(theme string) error {
	cfg, err := os.Open(programTheme.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)

	}
	defer cfg.Close()
	scanner := bufio.NewScanner(cfg)
	scanner.Split(bufio.ScanLines)
	scanner.Split(bufio.ScanLines)
	var text []string
	var flag bool = false

	for scanner.Scan() {
		t := scanner.Text()
		if flag {
			if strings.HasPrefix(t, "[") {
				flag = false
			}
			if strings.HasPrefix(t, "selected =") {
				t = "selected = " + programTheme.GetTheme(theme)
			}
		}
		if !flag && strings.Contains(t, "[appearance]") {
			flag = true
		}
		text = append(text, t)
	}

	data := []byte(strings.Join(text, "\n"))
	if err = ioutil.WriteFile(programTheme.ConfigPath, data, 0644); err != nil {
		return err
	}
	return nil
}

// WindowsTerminal.SetTheme sets the theme for WindowsTerminal
// Given a theme name, it will set the theme
// If the theme name is not valid, it will return an error
func (programTheme WindowsTerminal) SetTheme(theme string) error {
	jsonFile, err := os.Open(programTheme.ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	result := map[string]interface{}{}
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return err
	}
	for ix := range result["profiles"].(map[string]interface{})["list"].([]interface{}) {
		var NewTheme string = programTheme.GetTheme(theme)
		result["profiles"].(map[string]interface{})["list"].([]interface{})[ix].(map[string]interface{})["colorScheme"] = NewTheme
	}

	file, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(programTheme.ConfigPath, file, 0644); err != nil {
		return err
	}
	return nil
}

type NewThemeable func(ThemeConfig) Themeable

// newOneCommander returns a new OneCommande given a theme config
func newOneCommander(config ThemeConfig) Themeable {
	return OneCommander{
		config,
	}
}

// newSpyder returns a new Spyder given a theme config
func newSpyder(config ThemeConfig) Themeable {
	return Spyder{
		config,
	}
}

// newPythonIDLE returns a new PythonIDLE given a theme config
func newPythonIDLE(config ThemeConfig) Themeable {
	return PythonIDLE{
		config,
	}
}

// newWindowsTerminal returns a new WindowsTerminal given a theme config
func newWindowsTerminal(config ThemeConfig) Themeable {
	return WindowsTerminal{
		config,
	}
}

// Registry of themeable program creators
var Registry = map[string]NewThemeable{
	"OneCommander":    newOneCommander,
	"PythonIDLE":      newPythonIDLE,
	"Spyder":          newSpyder,
	"WindowsTerminal": newWindowsTerminal,
}
