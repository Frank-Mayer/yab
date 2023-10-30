package extensions

import "github.com/Shopify/go-lua"

var lua_args []string

func SetArgs(args []string) {
	lua_args = args
}

func args(l *lua.State) int {
	l.CreateTable(len(lua_args), 0)
	for i := 0; i < len(lua_args); i++ {
		l.PushString(lua_args[i])
		l.RawSetInt(-2, i+1)
	}
	return 1
}
