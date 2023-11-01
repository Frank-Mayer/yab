package extensions

import (
	"runtime"

	"github.com/Shopify/go-lua"
)

func osType(l *lua.State) int {
	l.PushString(runtime.GOOS)
	return 1
}
