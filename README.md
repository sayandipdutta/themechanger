# Program to change themes for certain programs that don't support Dynamic Themeing in Windows

## Requirements:
1. Windows 10 or higher
2. Auto Dark Mode >= 10.0.0
3. Golang >= 1.23.5 (If you plan to build from source)

## Installation
```sh
go install github.com/sayandipdutta/themechanger
```

If you plan to build from source, you will need to install Go from [here](https://golang.org/doc/install).
After you are done installing. Get the repo and run the following command:
```
git clone https://github.com/sayandipdutta/themechanger.git
cd themechanger
go mod tidy
go build
```
`themechanger.exe` file will be created in the `pwd`.

You can then add it to your `%PATH` (`$Env:PATH` from `powershell`).

## Usage
1. Open the Auto Dark Mode scripts file (see [Installation](#installation))
2. Add the following line to the end of the file:

```
  - Name: ProgramThemeSwitcher
    Command: themechanger.exe  # if added to PATH, else the full path to the exe
    ArgsLight: [-light]
    ArgsDark: []
    AllowedSources: [Any]
```
The `scripts.yml` file should look like this:
```
Enabled: true
Component:
  TimeoutMillis: 10000
  Scripts:
  - Name: ProgramThemeSwitcher
    Command: themechanger.exe  # if added to PATH, else the full path to the exe
    ArgsLight: [-light]
    ArgsDark: []
    AllowedSources: [Any]
```
Usage of `themeChange.exe` is as follows:
```
Usage of themeChange.exe:

Flags:
  -light light
        If given, light theme will be set, otherwise `dark`.

Example:
.\themeChange.exe
.\themeChange.exe -light
```

## Configuration:
`themechanger` looks for the config file in `$Env:THEMECHANGER_CONFIG`.

So populate a `json` file with the following structure:
```json5
{
  "YourAppName": {
    "light": "X:full\\path\\to\\light\\themeconfig\\for\\YourAppName",
    "dark": "X:full\\path\\to\\dark\\themeconfig\\for\\YourAppName",
    "configpath": "X:full\\path\\to\\YourApp's\\Default\\ConfigPath"
  },
  // ... other apps
}
```

And set the path of the file as `$Env:THEMECHANGER_CONFIG`
