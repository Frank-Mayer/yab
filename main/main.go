package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	lua "github.com/yuin/gopher-lua"
	"selene.frankmayer.dev/docs"
	"selene.frankmayer.dev/extensions"
	"selene.frankmayer.dev/util"
)

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
	var err error
	util.ConfigPath, err = get_config_path()
	if err != nil {
		log.Fatal(err)
	}

	// run each file passed as an argument
	for i := 0; i < arg_separator_index; i++ {
		file := argsWithoutProg[i]
		init_file, err := getInitFile(util.ConfigPath, file)
		if err != nil {
			log.Fatal(err)
		}
		err = runLuaFile(init_file)
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
		fmt.Println(util.Version)
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

func runLuaFile(init_file string) error {
	// setup lua
	l := lua.NewState()
	defer l.Close()
	extensions.RegisterExtensions(l)

	package_path := util.GetPackagePath()
	setup_code := "package.path = '" + strings.ReplaceAll(package_path, "\\", "\\\\") + ";'"
	err := l.DoString(setup_code)
	if err != nil {
		log.Error("Error setting up lua", "error", err, "code", setup_code)
		return err
	}

	// run lua file
	err = l.DoFile(init_file)
	if err != nil {
		return err
	}

	return nil
}

func get_config_path() (string, error) {
	pathname := path.Join(".", ".selene")

	// check for current directory
	if _, err := os.Stat(pathname); !os.IsNotExist(err) {
		return pathname, nil
	}

	// check for global config
	return util.GetGlobalConfigPath()
}

func initProject() {
	util.ConfigPath = ".selene"
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
		return
	} else if err != nil {
		log.Fatal(err)
	} else if !f.IsDir() {
		log.Fatal(".selene is not a directory")
	}
	log.Info(".selene already exists")

	err = initDefinitons()
	if err != nil {
		log.Error(err)
	}
}

func initDemoConfig() error {
	filename := path.Join(".selene", "hello.lua")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("print('Hello, world!')")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Created " + filename)
	log.Info("Run with `" + util.BinName() + " hello`")

	return nil
}

func initDefinitons() error {
	config_path, err := util.GetGlobalConfigPath()
	if err != nil {
		return err
	}

	lib_path := path.Join(config_path, "lib")
	filename := path.Join(lib_path, "Selene.lua")

	// chek if exists
	_, err = os.Stat(filename)
	if err == nil {
		log.Info("Definitions already exist. Updating...")
		err = os.Remove(filename)
		if err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	// create directory
	err = os.MkdirAll(lib_path, 0775)
	if err != nil {
		return err
	}

	// create file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// write file
	_, err = f.WriteString(extensions.Definitions())
	if err != nil {
		return err
	}
	log.Info("Lua API definitions created", "location", lib_path)

	return nil
}
