package themeable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sayandipdutta/themechanger/setup"
	"gopkg.in/ini.v1"
)

func LoadConfig() (map[string]ThemeConfig, error) {
	// read config file (JSON) and set theme
	confpath := setup.GetParentDir().ConfPath
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
func SetTheme(program Themeable, theme string) error {
	fmt.Println("Setting theme to: ", program)
	err := program.SetTheme(theme)
	return err
}

type Spyder struct {
	ThemeConfig
}
type PythonIDLE struct {
	ThemeConfig
}
type OneCommander struct {
	ThemeConfig
}
type WindowsTerminal struct {
	ThemeConfig
}

func (programTheme OneCommander) SetTheme(theme string) error {
	// read config file (JSON) and set theme
	jsonFile, err := os.Open(programTheme.ConfigPath)
	if err != nil {
		return err
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

// read config file (CFG) and set theme
func (programTheme PythonIDLE) SetTheme(theme string) error {
	cfg, err := ini.Load(programTheme.ConfigPath)
	if err != nil {
		return err
	}

	var NewTheme string = programTheme.GetTheme(theme)
	cfg.Section("Theme").Key("name").SetValue(NewTheme)

	if err = cfg.SaveTo(programTheme.ConfigPath); err != nil {
		return err
	}
	return nil
}

// read config file (INI) and set theme
func (programTheme Spyder) SetTheme(theme string) error {
	cfg, err := ini.Load(programTheme.ConfigPath)
	if err != nil {
		return err
	}

	var NewTheme string = programTheme.GetTheme(theme)
	cfg.Section("appearance").Key("selected").SetValue(NewTheme)

	if err = cfg.SaveTo(programTheme.ConfigPath); err != nil {
		return err
	}
	return nil
}

// read config file (JSON) and set theme
func (programTheme WindowsTerminal) SetTheme(theme string) error {
	jsonFile, err := os.Open(programTheme.ConfigPath)
	if err != nil {
		return err
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

func newOneCommander(config ThemeConfig) Themeable {
	return OneCommander{
		config,
	}
}
func newSpyder(config ThemeConfig) Themeable {
	return Spyder{
		config,
	}
}
func newPythonIDLE(config ThemeConfig) Themeable {
	return PythonIDLE{
		config,
	}
}
func newWindowsTerminal(config ThemeConfig) Themeable {
	return WindowsTerminal{
		config,
	}
}

var Registry = map[string]NewThemeable{
	"OneCommander":    newOneCommander,
	"PythonIDLE":      newPythonIDLE,
	"Spyder":          newSpyder,
	"WindowsTerminal": newWindowsTerminal,
}
