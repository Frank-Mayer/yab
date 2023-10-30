package extensions

import "github.com/Shopify/go-lua"

func register(name string, l *lua.State, f lua.Function) {
	l.Register("selene_"+name, f)
}

func RegisterExtensions(l *lua.State) {
	register("os_type", l, os_type)
	register("args", l, args)
}
