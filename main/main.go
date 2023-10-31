package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Shopify/go-lua"
	"selene.frankmayer.io/docs"
	"selene.frankmayer.io/extensions"
	"selene.frankmayer.io/util"
)

const version = "0.1.0"

func main() {
	// find config folder
	config_path, err := get_config_path()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// process arguments
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 1 {
		treatSpecialArgs(argsWithoutProg[0])
	}

	// if no args, exit
	if len(argsWithoutProg) == 0 {
		fmt.Println("No arguments given")
		docs.Help()
		os.Exit(1)
	}

	// check for -- to pass args to lua
	pass_args := false
	pass_args_list := []string{}
	arg_separator_index := len(argsWithoutProg)
	for i := 0; i < len(argsWithoutProg); i++ {
		arg := argsWithoutProg[i]
		if pass_args {
			pass_args_list = append(pass_args_list, arg)
		} else if arg == "--" {
			pass_args = true
			arg_separator_index = i
		}
	}
	extensions.SetArgs(pass_args_list)

	// run each file passed as an argument
	for i := 0; i < arg_separator_index; i++ {
		file := argsWithoutProg[i]
		init_file, err := getInitFile(config_path, file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = runLuaFile(config_path, init_file)
		if err != nil {
			fmt.Println("Error running file: ", file)
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func treatSpecialArgs(arg0 string) {
	switch arg0 {
	case "--help", "-h":
		docs.Help()
		os.Exit(0)
	case "--version", "-v":
		fmt.Println(version)
		os.Exit(0)
	case "--init":
		initProject()
		os.Exit(0)
	}
}

func getInitFile(config_path string, file string) (string, error) {
	init_file := path.Join(config_path, file+".lua")
	if _, err := os.Stat(init_file); err != nil {
		return "", err
	}
	return init_file, nil
}

func runLuaFile(config_path string, init_file string) error {
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

func initProject() {
	f, err := os.Stat(".selene")
	if os.IsNotExist(err) {
		err = os.Mkdir(".selene", 0775)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Created .selene")
		err = initDemoConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else if !f.IsDir() {
		fmt.Println(".selene is not a directory")
		os.Exit(1)
	}
	fmt.Println(".selene already exists")
	os.Exit(0)
}

func initDemoConfig() error {
	filename := ".selene/hello.lua"
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("print('Hello, world!')")
	fmt.Println("Created", filename)
	fmt.Println("Run with `" + util.BinName() + " hello`")
	return nil
}
