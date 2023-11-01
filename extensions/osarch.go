package extensions

import (
	"runtime"

	"github.com/Shopify/go-lua"
)

func os_arch(l *lua.State) int {
	l.PushString(runtime.GOARCH)
	return 1
}
