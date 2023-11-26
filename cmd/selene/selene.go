package main

import (
	"github.com/Frank-Mayer/selene/internal/docs"
	"github.com/Frank-Mayer/selene/internal/extensions"
	"github.com/Frank-Mayer/selene/internal/util"

	"fmt"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/yuin/gopher-lua"
)

func main() {
	log.SetLevel(log.WarnLevel)

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
	passArgs := false
	var passArgsList []string
	argSeparatorIndex := len(argsWithoutProg)
	for i := 0; i < len(argsWithoutProg); i++ {
		arg := argsWithoutProg[i]
		if passArgs {
			passArgsList = append(passArgsList, arg)
		} else {
			switch arg {
			case "--debug":
				log.SetLevel(log.DebugLevel)
				log.Debug("Debug mode enabled")
			case "--silent":
				log.SetLevel(10)
			case "--":
				passArgs = true
				argSeparatorIndex = i
			}
		}
	}
	extensions.SetArgs(passArgsList)

	// find config folder
	var err error
	util.ConfigPath, err = getConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	// run each file passed as an argument
	for i := 0; i < argSeparatorIndex; i++ {
		file := argsWithoutProg[i]
		if strings.HasPrefix(file, "-") {
			continue
		}
		initFile, err := getInitFile(util.ConfigPath, file)
		if err != nil {
			log.Fatal(err)
		}
		err = runLuaFile(initFile)
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

func getInitFile(configPath string, file string) (string, error) {
	initFile := path.Join(configPath, file+".lua")
	if _, err := os.Stat(initFile); err != nil {
		return "", err
	}
	return initFile, nil
}

func runLuaFile(initFile string) error {
	// setup lua
	l := lua.NewState()
	defer l.Close()
	extensions.RegisterExtensions(l)

	packagePath := util.GetPackagePath()
	setupCode := "package.path = '" + strings.ReplaceAll(packagePath, "\\", "\\\\") + ";'"
	err := l.DoString(setupCode)
	if err != nil {
		log.Error("Error setting up lua", "error", err, "code", setupCode)
		return err
	}

	// run lua file
	err = l.DoFile(initFile)
	if err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
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
	configPath, err := util.GetGlobalConfigPath()
	if err != nil {
		return err
	}

	libPath := path.Join(configPath, "lib")
	filename := path.Join(libPath, "Selene.lua")

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
	err = os.MkdirAll(libPath, 0775)
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
	log.Info("Lua API definitions created", "location", libPath)

	return nil
}
