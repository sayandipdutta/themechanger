package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// LoadConfig loads the config file and returns the ThemeConfig
// If the config file does not exist, it will return an error
// If the config file is not valid, it will return an error
func loadConfig() (map[string]themeConfig, error) {
	// read config file (JSON) and set theme
	confPath := os.Getenv("THEMECHANGER_CONFIG")
	if confPath == "" {
		return nil, fmt.Errorf("envvar THEMECHANGER_CONFIG is not defined!")
	}
	jsonFile, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	result := map[string]themeConfig{}
	if err = json.Unmarshal([]byte(byteValue), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// light and dark theme names, and config file path
type themeConfig struct {
	Light      string `json:"light"`      // prefered theme name for light mode
	Dark       string `json:"dark"`       // prefered theme name for dark mode
	ConfigPath string `json:"configpath"` // config file path
}

func (programTheme *themeConfig) setTheme(isLight bool) error {
	// Open the source file
	src := programTheme.Dark
	if isLight {
		src = programTheme.Light
	}
	dst := programTheme.ConfigPath
	bkp_filename, err := backup(dst)
	if err != nil {
		return err
	}
	if err = copyFile(src, dst); err != nil {
		return err
	}
	if bkp_filename != "" {
		os.Remove(bkp_filename)
	}
	return nil
}

func backup(src string) (string, error) {
	_, err := os.Stat(src)
	if err != nil {
		return "", nil
	}
	bkp := src + ".bkp"
	err = copyFile(src, bkp)
	if err != nil {
		return "", fmt.Errorf("Failed to backup %s", src)
	}
	return bkp, nil
}

func copyFile(src string, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close() // Ensure the source file is closed

	// Create the destination file
	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destinationFile.Close() // Ensure the destination file is closed

	// Copy the contents from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %w", err)
	}

	return nil
}
