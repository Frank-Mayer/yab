package extensions

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
	"github.com/yuin/gopher-lua"
)

type jsPackageManager uint8

const (
	none jsPackageManager = iota
	npm
	pnpm
	bun
)

func findJsPackageManager() jsPackageManager {
	if _, err := exec.LookPath("pnpm"); err == nil {
		log.Debug("Found pnpm")
		return pnpm
	}

	if _, err := exec.LookPath("bun"); err == nil {
		log.Debug("Found bun")
		return bun
	}

	if _, err := exec.LookPath("npm"); err == nil {
		log.Debug("Found npm")
		return npm
	}

	log.Error("No JS package manager found")
	return none
}

func jsRun(l *lua.LState) int {
	script := l.CheckString(1)
	var command string
	switch findJsPackageManager() {
	case pnpm:
		command = "pnpm run " + script
	case bun:
		command = "bun run " + script
	case npm:
		command = "npm run " + script
	default:
		l.Push(lua.LFalse)
		return 1
	}

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error("Error running script", "error", err)
		l.Push(lua.LFalse)
		return 1
	}

	l.Push(lua.LTrue)
	return 1
}

func jsInstall(l *lua.LState) int {
	var command string
	switch findJsPackageManager() {
	case pnpm:
		command = "pnpm install"
	case bun:
		command = "bun install"
	case npm:
		command = "npm install"
	default:
		l.Push(lua.LFalse)
		return 1
	}
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error("Error running script", "error", err)
		l.Push(lua.LFalse)
		return 1
	}

	l.Push(lua.LTrue)
	return 1
}
