package extensions

import (
	"runtime"

	"github.com/Shopify/go-lua"
)

func osArch(l *lua.State) int {
	l.PushString(runtime.GOARCH)
	return 1
}
