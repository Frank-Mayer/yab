package extensions

import (
	"runtime"

	"github.com/yuin/gopher-lua"
)

func osArch(l *lua.LState) int {
	l.Push(lua.LString(runtime.GOARCH))
	return 1
}
