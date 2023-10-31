package extensions

import (
	"regexp"

	"github.com/Shopify/go-lua"
)

// string.gmatch (s, pattern)
func gmatch(l *lua.State) int {
	s := lua.CheckString(l, 1)
	pattern := lua.CheckString(l, 2)
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(s, -1)
	var i int = 0

	iter := func(l *lua.State) int {
		if i >= len(matches) {
			return 0
		}

		l.PushString(matches[i])
		i++
		return 1
	}

	l.PushGoFunction(iter)
	return 1
}

func addMissingStringFunctions(l *lua.State) {
	// get the string table
	l.Global("string")

	// add the gmatch
	l.PushGoFunction(gmatch)
	l.SetField(-2, "gmatch")

	// pop the string table
	l.SetGlobal("string")
}
