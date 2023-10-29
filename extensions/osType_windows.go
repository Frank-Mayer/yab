//go:build windows
// +build windows

package extensions

import "github.com/Shopify/go-lua"

func os_type(l *lua.State) int {
	l.PushString("windows")
	return 1
}
