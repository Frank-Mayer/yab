package extensions

import (
	"os/exec"
	"strings"

	"github.com/yuin/gopher-lua"
)

// call a shell command and return the full output (stdout + stderr) in one string
func stdall(l *lua.LState) int {
	command := l.CheckString(1)
	parts := strings.Fields(command)

	// Check if there are any parts
	if len(parts) == 0 {
		l.Push(lua.LString(""))
		return 1
	}

	// Assign the first part to the 'name' variable
	name := parts[0]

	// Assign the rest of the parts to the 'arg' slice
	args := parts[1:]

	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	l.Push(lua.LString(string(out)))
	return 1
}

// call a shell command and return the output (stdout) in one string
func stdout(l *lua.LState) int {
	command := l.CheckString(1)
	parts := strings.Fields(command)

	// Check if there are any parts
	if len(parts) == 0 {
		l.Push(lua.LString(""))
		return 1
	}

	// Assign the first part to the 'name' variable
	name := parts[0]

	// Assign the rest of the parts to the 'arg' slice
	args := parts[1:]

	out, err := exec.Command(name, args...).Output()
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	l.Push(lua.LString(string(out)))
	return 1
}

// call a shell command and return the error (stderr) in one string
func stderr(l *lua.LState) int {
	command := l.CheckString(1)
	parts := strings.Fields(command)

	// Check if there are any parts
	if len(parts) == 0 {
		l.Push(lua.LString(""))
		return 1
	}

	// Assign the first part to the 'name' variable
	name := parts[0]

	// Assign the rest of the parts to the 'arg' slice
	args := parts[1:]

	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}
	l.Push(lua.LString(string(out)))
	return 1
}
