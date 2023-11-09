package extensions

import (
	"runtime"

	"github.com/yuin/gopher-lua"
)

func osType(l *lua.LState) int {
	l.Push(lua.LString(runtime.GOOS))
	return 1
}
