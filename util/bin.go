package util

import (
	"errors"
	"os"
	"path"
)

var (
	ConfigPath string
)

func BinName() string {
	bin_name := os.Args[0]
	if len(bin_name) > 24 {
		return "selene"
	}
	return bin_name
}

func GetPackagePath() string {
	return path.Join(ConfigPath, "?.lua") + ";" +
		path.Join(ConfigPath, "?", "init.lua") + ";" +
		path.Join(".", "?.lua") + ";" +
		path.Join("?", "init.lua")

}

func GetGlobalConfigPath() (string, error) {
    pathname := path.Join(os.Getenv("HOME"), ".config", "selene")
    if _, err := os.Stat(pathname); !os.IsNotExist(err) {
        return pathname, nil
    }

    pathname = path.Join(os.Getenv("XDG_CONFIG_HOME"), "selene")
    if _, err := os.Stat(pathname); !os.IsNotExist(err) {
        return pathname, nil
    }

    pathname = path.Join(os.Getenv("APPDATA"), "selene")
    if _, err := os.Stat(pathname); !os.IsNotExist(err) {
        return pathname, nil
    }

    return "", errors.New("Could not find config path")
}
