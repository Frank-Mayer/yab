package extensions

import (
	"os/exec"

	"github.com/yuin/gopher-lua"
)

func checkExec(l *lua.LState) int {
	pkg := l.CheckString(1)
	_, err := exec.LookPath(pkg)
	if err != nil {
		l.Push(lua.LBool(false))
	} else {
		l.Push(lua.LBool(true))
	}
	return 1
}
