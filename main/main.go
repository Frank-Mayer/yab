package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Shopify/go-lua"
	"github.com/charmbracelet/log"
	"selene.frankmayer.io/docs"
	"selene.frankmayer.io/extensions"
	"selene.frankmayer.io/util"
)

const version = "0.1.0"

func main() {
	// process arguments
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 1 {
		treatSpecialArgs(argsWithoutProg[0])
	}

	// if no args, exit
	if len(argsWithoutProg) == 0 {
		docs.Help()
		log.Fatal("No arguments given")
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

	// find config folder
	config_path, err := get_config_path()
	if err != nil {
		log.Fatal(err)
	}

	// run each file passed as an argument
	for i := 0; i < arg_separator_index; i++ {
		file := argsWithoutProg[i]
		init_file, err := getInitFile(config_path, file)
		if err != nil {
			log.Fatal(err)
		}
		err = runLuaFile(config_path, init_file)
		if err != nil {
			log.Fatal("Error running file: "+file, "error", err)
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
	case "--update", "--upgrade", "-u":
		util.Update()
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
	package_path := path.Join(config_path, "?.lua")
	setup_code := "package.path = '" + package_path + ";'"
	err := lua.DoString(l, setup_code)
	if err != nil {
		log.Error("Error setting up lua", "error", err, "code", setup_code)
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
	init := path.Join(".", ".selene")

	// check for current directory
	if _, err := os.Stat(init); !os.IsNotExist(err) {
		return init, nil
	}

	// check for XDG_CONFIG_HOME
	if config_home := path.Join(os.Getenv("XDG_CONFIG_HOME"), "selene"); config_home != "" {
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
			log.Fatal(err)
		}
		log.Info("Created .selene")
		err = initDemoConfig()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else if !f.IsDir() {
		log.Fatal(".selene is not a directory")
	}
	log.Warn(".selene already exists")
	os.Exit(0)
}

func initDemoConfig() error {
	filename := path.Join(".selene", "hello.lua")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("print('Hello, world!')")
	log.Info("Created", filename)
	log.Info("Run with `" + util.BinName() + " hello`")
	return nil
}
