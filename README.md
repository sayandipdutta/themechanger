# Program to change themes for certain programs that don't support Dynamic Themeing in Windows

## Requirements:
1. Windows 10 or higher
2. Auto Dark Mode >= 10.0.0
3. Golang >= 1.12.5 (If you plan to build from source)

## Supported Programs
- [WindowsTerminal](https://www.microsoft.com/en-us/p/windows-terminal/9n0dx20hk701?activetab=pivot:overviewtab)
- [OneCommander](https://www.onecommander.com/)*
- [PythonIDLE](https://www.python.org/downloads/)*
- [Spyder](https://www.spyder-ide.org/)*

**\* Program must be closed while the theme switch takes place, for the theme switch to take effect.**

## Installation
No need to install, just download the `.exe` file, and put the path to the file in Auto Dark Mode scripts file (see [Usage](#usage))

If you plan to build from source, you will need to install Go from [here](https://golang.org/doc/install).
After you are done installing. Get the repo and run the following command:
```
git clone https://github.com/sayandipdutta/themechanger.git
cd themechange
go mod tidy
go build main/themeChanger.go
```
`themeChange.exe` file will be created in the `pwd`.

## Usage
1. Open the Auto Dark Mode scripts file (see [Installation](#installation))
2. Add the following line to the end of the file:

```
  - Name: ProgramThemeSwitcher
    Command: full\path\to\themeChange.exe   # eg. D:\Programs\themechanger\themeChange.exe
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
    Command: D:\Programs\themechanger\themeChange.exe    # Example Path
    ArgsLight: [--theme, light]
    ArgsDark: [--theme, dark]
    AllowedSources: [Any]
```
Usage of `themeChange.exe` is as follows:
```
Usage of themeChange.exe:

Accepted value of theme flag: 
        light
        dark    //default value.

Accepted value of program flag:
        all                                     //default value. Theme all programs.
or,     "<program_name> <program_name2> ..."    Theme only specified programs.

program should be provided in double quotes ("") if it contains spaces.

Flags:
  -program string
        Program to be themed (default "all")
  -theme string
        Type of theme to be set (default "dark")

Example:

.\themeChange.exe                       //default value of theme == dark, program == all
.\themeChange.exe --theme=dark          //default value of program == all
.\themeChange.exe --theme=light --program="OneCommander Spyder"
```
