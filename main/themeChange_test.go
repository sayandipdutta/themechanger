package main

import (
	"testing"

	"github.com/sayandipdutta/themechanger/themeable"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestSetTheme(t *testing.T) {
	programs := []struct {
		prog  themeable.Themeable
		theme string
	}{
		{
			themeable.WindowsTerminal{
				ThemeConfig: themeable.ThemeConfig{
					Light:      "One Half Light",
					Dark:       "Nord",
					ConfigPath: `C:\Users\sayan\AppData\Local\Packages\Microsoft.WindowsTerminalPreview_8wekyb3d8bbwe\LocalState\settings.json`,
				},
			},
			"lights",
		},
		{
			themeable.WindowsTerminal{
				ThemeConfig: themeable.ThemeConfig{
					Light:      "One Half Light",
					Dark:       "Nord",
					ConfigPath: `D:\Users\sayan\AppData\Local\Packages\Microsoft.WindowsTerminalPreview_8wekyb3d8bbwe\LocalState\settings.json`,
				},
			},
			"light",
		},
	}
	for ix, program := range programs {
		err := themeable.SetTheme(program.prog, program.theme)
		if err == nil {
			t.Errorf("%d. Expected error: %v, got: %v", ix, nil, err)
		}
	}
}
