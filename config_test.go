package main

import "testing"

func Test_ThemeConfig_GetTheme(t *testing.T) {
	args := []struct {
		theme ThemeConfig
		setTo string
		want  string
	}{
		{
			ThemeConfig{
				Light:      "light theme",
				Dark:       "dark theme",
				ConfigPath: "configpath",
			},
			"light",
			"light theme",
		},
		{
			ThemeConfig{
				Light:      "light theme",
				Dark:       "dark theme",
				ConfigPath: "configpath",
			},
			"dark",
			"dark theme",
		},
		{
			ThemeConfig{
				Light:      "light theme",
				Dark:       "dark theme",
				ConfigPath: "configpath",
			},
			"",
			"dark theme",
		},
	}
	for ix, arg := range args {
		got := arg.theme.GetTheme(arg.setTo)
		if got != arg.want {
			t.Errorf("%d. Expected: %v, got: %v", ix, arg.want, got)
		}
	}
}

func TestSetTheme(t *testing.T) {
	programs := []struct {
		prog  Themeable
		theme string
	}{
		{
			WindowsTerminal{
				ThemeConfig: ThemeConfig{
					Light:      "One Half Light",
					Dark:       "Nord",
					ConfigPath: `D:\Users\sayan\AppData\Local\Packages\Microsoft.WindowsTerminalPreview_8wekyb3d8bbwe\LocalState\settings.json`,
				},
			},
			"light",
		},
	}
	for ix, program := range programs {
		err := SetTheme(program.prog, program.theme)
		if err == nil {
			t.Errorf("%d. got: %v", ix, err)
		}
	}
}
