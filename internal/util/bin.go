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
	binName := os.Args[0]
	if len(binName) > 24 {
		return "selene"
	}
	return binName
}

func GetPackagePath() string {
	return path.Join(ConfigPath, "?.lua") + ";" +
		path.Join(ConfigPath, "?", "init.lua") + ";" +
		path.Join(".", "?.lua") + ";" +
		path.Join("?", "init.lua")

}

func GetGlobalConfigPath() (string, error) {
	if xdgConfigHome, exists := os.LookupEnv("XDG_CONFIG_HOME"); exists {
		pathname := path.Join(xdgConfigHome, "selene")
		// check if pathname exists
		_, err := os.Stat(pathname)
		if err == nil {
			return pathname, nil
		}
		if os.IsNotExist(err) {
			// create directory
			err := os.MkdirAll(pathname, 0755)
			if err != nil {
				return "", err
			}
			return pathname, nil
		}
	}

	if home, exists := os.LookupEnv("APPDATA"); exists {
		pathname := path.Join(home, "selene")
		// check if pathname exists
		_, err := os.Stat(pathname)
		if err == nil {
			return pathname, nil
		}
		if os.IsNotExist(err) {
			// create directory
			err := os.MkdirAll(pathname, 0755)
			if err != nil {
				return "", err
			}
			return pathname, nil
		}
	}

	if home, exists := os.LookupEnv("HOME"); exists {
		pathname := path.Join(home, ".config", "selene")
		// check if pathname exists
		_, err := os.Stat(pathname)
		if err == nil {
			return pathname, nil
		}
		if os.IsNotExist(err) {
			// create directory
			err := os.MkdirAll(pathname, 0755)
			if err != nil {
				return "", err
			}
			return pathname, nil
		}
	}

	return "", errors.New("could not find or create global config path")
}
