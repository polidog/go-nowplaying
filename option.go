package nowplaying

import (
	"runtime"
	"os"
)

type Option struct {
	Version bool `long:"version" description:"Show version"`
	Filename string `short:"f" long:"filename" description:"Config file name."`
}

func (o Option) GetFilename() string {
	if len(o.Filename) > 0 {
		return o.Filename
	}
	return homeDir() + "/.nowplaying.toml"
}

func homeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
