package setup

import (
	"path"
	"path/filepath"
	"runtime"
)

type SetupFiles struct {
	LogPath  string
	ConfPath string
}

func GetParentDir() SetupFiles {
	_, filename, _, _ := runtime.Caller(1)
	c := filepath.Join(path.Dir(path.Dir(filename)), `config\config.json`)
	l := filepath.Join(path.Dir(path.Dir(filename)), "logs.txt")
	return SetupFiles{
		LogPath:  l,
		ConfPath: c,
	}
}
