package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Shopify/go-lua"
	"selene.frankmayer.io/extensions"
)

const version = "0.1.0"

func main() {
	// find config folder
	config_path, err := get_config_path()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 1 &&
		(argsWithoutProg[0] == "--version" || argsWithoutProg[0] == "-v") {
		fmt.Println(version)
		os.Exit(0)
	}

	// if no args, run init.lua
	if len(argsWithoutProg) == 0 {
		file := "init"
		init_file, err := get_init_file(config_path, file, true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = run_lua_file(config_path, init_file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		// run each file passed as an argument
		for _, file := range argsWithoutProg {
			init_file, err := get_init_file(config_path, file, false)
			if err != nil {
				fmt.Println("Could not find file: ", file)
				fmt.Println(err)
				os.Exit(1)
			}
			err = run_lua_file(config_path, init_file)
			if err != nil {
				fmt.Println("Error running file: ", file)
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

func get_init_file(config_path string, file string, ensure_exist bool) (string, error) {
	init_file := path.Join(config_path, file+".lua")
	if _, err := os.Stat(init_file); os.IsNotExist(err) {
		if ensure_exist {
			if _, err := os.Create(init_file); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return init_file, nil
}

func run_lua_file(config_path string, init_file string) error {
	// setup lua
	l := lua.NewState()
	lua.OpenLibraries(l)
	extensions.RegisterExtensions(l)

	// set package.path to config folder
	err := lua.DoString(l, fmt.Sprintf("package.path = '%s/?.lua;'", config_path))
	if err != nil {
		return err
	}

	// run lua file
	err = lua.DoFile(l, init_file)
	if err != nil {
		return err
	}

	return nil
}

func get_config_path() (string, error) {
	init := "./.selene"

	// check for current directory
	if _, err := os.Stat(init); !os.IsNotExist(err) {
		return init, nil
	}

	init = "./selene"
	// check for XDG_CONFIG_HOME
	if config_home := path.Join(os.Getenv("XDG_CONFIG_HOME"), init); config_home != "" {
		if _, err := os.Stat(config_home); !os.IsNotExist(err) {
			return config_home, nil
		}
	}

	// check for appdata
	if appdata := path.Join(os.Getenv("APPDATA"), init); appdata != "" {
		if _, err := os.Stat(appdata); !os.IsNotExist(err) {
			return appdata, nil
		}
	}

	return "", errors.New("Could not find config file")
}
