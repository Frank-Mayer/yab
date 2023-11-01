package extensions

import (
	"runtime"

	"github.com/Shopify/go-lua"
)

func os_type(l *lua.State) int {
	l.PushString(runtime.GOOS)
	return 1
}
