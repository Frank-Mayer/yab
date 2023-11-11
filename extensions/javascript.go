package extensions

import (
	"os"
	"os/exec"

	"github.com/yuin/gopher-lua"
)

type jsPackageManager uint8

const (
	none jsPackageManager = iota
	npm
	yarn
	pnpm
	bun
)

func findJsPackageManager() jsPackageManager {
	if _, err := exec.LookPath("pnpm"); err == nil {
		return pnpm
	}

	if _, err := exec.LookPath("yarn"); err == nil {
		return yarn
	}

	if _, err := exec.LookPath("bun"); err == nil {
		return bun
	}

	if _, err := exec.LookPath("npm"); err == nil {
		return npm
	}

	return none
}

func jsRun(l *lua.LState) int {
	script := l.CheckString(1)
	var command string
	switch findJsPackageManager() {
	case yarn:
		command = "yarn run " + script
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
	cmd.Run()
	l.Push(lua.LTrue)
	return 1
}

func jsInstall(l *lua.LState) int {
	var command string
	switch findJsPackageManager() {
	case yarn:
		command = "yarn install"
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
	cmd.Run()
	l.Push(lua.LTrue)
	return 1
}
