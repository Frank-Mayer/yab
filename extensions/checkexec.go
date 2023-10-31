package extensions

import (
	"os/exec"

	"github.com/Shopify/go-lua"
)

func checkExec(l *lua.State) int {
	pkg := lua.CheckString(l, 1)
	_, err := exec.LookPath(pkg)
	if err != nil {
		l.PushBoolean(false)
	} else {
		l.PushBoolean(true)
	}
	return 1
}
