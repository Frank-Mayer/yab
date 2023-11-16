package extensions

import (
	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	lua "github.com/yuin/gopher-lua"
)

func watch(l *lua.LState) int {
	paths := l.CheckTable(1)
	callback := l.CheckFunction(2)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
	}

	go func() {
        defer watcher.Close()
		for {
			select {
			case event := <-watcher.Events:
				if event.Op != 0 {
                    var op string
                    switch event.Op {
                    case fsnotify.Create:
                        op = "create"
                    case fsnotify.Write:
                        op = "write"
                    case fsnotify.Remove:
                        op = "remove"
                    case fsnotify.Rename:
                        op = "rename"
                    case fsnotify.Chmod:
                        op = "chmod"
                    default:
                        op = "unknown"
                    }
                    log.Debug("Event received", "event", event)
					err := l.CallByParam(lua.P{
						Fn:      callback,
						NRet:    0,
						Protect: true,
					}, lua.LString(event.Name), lua.LString(op))
					if err != nil {
						log.Error("Error calling callback", "error", err)
					}
				}
			case err := <-watcher.Errors:
				if err != nil {
					log.Error("Error watching file", "error", err)
				}
			}
		}
	}()

	for i := 1; i <= paths.Len(); i++ {
		path := paths.RawGetInt(i).String()
		log.Debug("Adding path to watcher", "path", path)
		err = watcher.Add(path)
		if err != nil {
			log.Error("Error adding path to watcher", "path", path, "error", err)
		}
	}

	return 0
}
