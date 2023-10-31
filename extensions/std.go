package extensions

import (
	"os/exec"
	"strings"

	"github.com/Shopify/go-lua"
)

// call a shell command and return the full output (stdout + stderr) in one string
func stdall(l *lua.State) int {
	command := lua.CheckString(l, 1)
	parts := strings.Fields(command)

	// Check if there are any parts
	if len(parts) == 0 {
		l.PushString("")
		return 1
	}

	// Assign the first part to the 'name' variable
	name := parts[0]

	// Assign the rest of the parts to the 'arg' slice
	args := parts[1:]

	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		l.PushString(err.Error())
		return 1
	}
	l.PushString(string(out))
	return 1
}

// call a shell command and return the output (stdout) in one string
func stdout(l *lua.State) int {
	command := lua.CheckString(l, 1)
	parts := strings.Fields(command)

	// Check if there are any parts
	if len(parts) == 0 {
		l.PushString("")
		return 1
	}

	// Assign the first part to the 'name' variable
	name := parts[0]

	// Assign the rest of the parts to the 'arg' slice
	args := parts[1:]

	out, err := exec.Command(name, args...).Output()
	if err != nil {
		l.PushString(err.Error())
		return 1
	}
	l.PushString(string(out))
	return 1
}

// call a shell command and return the error (stderr) in one string
func stderr(l *lua.State) int {
	command := lua.CheckString(l, 1)
	parts := strings.Fields(command)

	// Check if there are any parts
	if len(parts) == 0 {
		l.PushString("")
		return 1
	}

	// Assign the first part to the 'name' variable
	name := parts[0]

	// Assign the rest of the parts to the 'arg' slice
	args := parts[1:]

	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		l.PushString(err.Error())
		return 1
	}
	l.PushString(string(out))
	return 1
}
