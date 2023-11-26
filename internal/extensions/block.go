package extensions

import (
	"os"
	"os/signal"

	"github.com/charmbracelet/log"
	"github.com/yuin/gopher-lua"
)

func block(_ *lua.LState) int {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	log.Debug("Received interrupt signal, stog blocking")

	return 0
}
