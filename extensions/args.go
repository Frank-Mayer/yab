package extensions

import "github.com/yuin/gopher-lua"

var lua_args []string

func SetArgs(args []string) {
	lua_args = args
}

func args(l *lua.LState) int {
	table := l.NewTable()
	for i := 0; i < len(lua_args); i++ {
		l.SetTable(table, lua.LNumber(i+1), lua.LString(lua_args[i]))
	}
	l.Push(table)
	return 1
}
