# Program to change themes for certain programs that don't support Dynamic Themeing in Windows

## Requirements:
1. Windows 10 or higher
2. Auto Dark Mode >= 10.0.0
3. Golang >= 1.12.5 (If you plan to build from source)

## Supported Programs
- [WindowsTerminal](https://www.windowsterminal.com/)
- [OneCommander](https://www.onecmd.com/)*
- [PythonIDLE](https://www.python.org/downloads/)*
- [Spyder](https://www.spyder-ide.org/)*

**\* Program must be closed while the theme switch takes place, for the theme switch to take effect.**

## Installation
No need to install, just download the `.exe` file, and put the path to the file in Auto Dark Mode scripts file (see [Usage](#usage))

If you plan to build from source, you will need to install Go from [here](https://golang.org/doc/install).
After you are done installing. Get the repo and run the following command:
```
cd themeChange
go build themeChanger.go
```
`themeChange.exe` file will be created in the same directory as the repo.

## Usage
1. Open the Auto Dark Mode scripts file (see [Installation](#installation))
2. Add the following line to the end of the file:

```
  - Name: ProgramThemeSwitcher
    Command: full\path\to\themeChange.exe   # eg. D:\Programs\themeChange\themeChange.exe
    ArgsLight: [--theme, light]
    ArgsDark: [--theme, dark]
    AllowedSources: [Any]
```
The `scripts.yml` file should look like this:
```
Enabled: true
Component:
  TimeoutMillis: 10000
  Scripts:
  - Name: ProgramThemeSwitcher
    Command: D:\Programs\themeChange\themeChange.exe    # Example Path
    ArgsLight: [--theme, light]
    ArgsDark: [--theme, dark]
    AllowedSources: [Any]
```
Usage of `themeChange.exe` is as follows:
```
Usage of themeChange.exe:

Accepted value of theme flag: 
        light
        dark

Flags:
  -theme string
        Type of theme to be set (default "dark")
```
