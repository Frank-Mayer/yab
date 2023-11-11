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
	if xdg_config_home, exists := os.LookupEnv("XDG_CONFIG_HOME"); exists {
		pathname := path.Join(xdg_config_home, "selene")
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

	return "", errors.New("Could not find or create global config path")
}
