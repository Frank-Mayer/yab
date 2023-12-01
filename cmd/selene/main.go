package main

import (
	"github.com/Frank-Mayer/selene/internal/cli"
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

	var err error
	if util.ConfigPath, err = getConfigPath(); err != nil {
		log.Fatal(err)
	}

	cli := cli.NewCli()
	cli.Parse()

	if cli.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")
	} else if cli.Silent {
		log.SetLevel(10)
	}

	if cli.Version {
		fmt.Println(util.Version)
	}

	if cli.Help {
		docs.Help()
	}

	if cli.Def {
		initDefinitons()
	}

	if cli.Update {
		util.Update()
	}

	for _, file := range cli.Configs {
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

func initDefinitons() {
	configPath, err := util.GetGlobalConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	libPath := path.Join(configPath, "lib")
	filename := path.Join(libPath, "Selene.lua")

	// chek if exists
	_, err = os.Stat(filename)
	if err == nil {
		log.Info("Definitions already exist. Updating...")
		err = os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
	} else if !os.IsNotExist(err) {
		log.Fatal(err)
	}

	// create directory
	err = os.MkdirAll(libPath, 0775)
	if err != nil {
		log.Fatal(err)
	}

	// create file
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write file
	_, err = f.WriteString(extensions.Definitions())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Lua API definitions created", "location", libPath)
}
