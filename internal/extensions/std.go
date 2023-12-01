package extensions

import (
	"bufio"
	"strings"
	"sync"

	"github.com/Frank-Mayer/yab/internal/util"
	"github.com/charmbracelet/log"
	"github.com/yuin/gopher-lua"
)

// call a shell command and return the full output (stdout + stderr) in one string
func stdall(l *lua.LState) int {
	command := l.CheckString(1)

	cmd := util.System(command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating stdout pipe", "error", err)
		return 0
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Error creating stderr pipe", "error", err)
		return 0
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting command", "error", err)
		return 0
	}

	sb := strings.Builder{}
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			sb.WriteString(scanner.Text() + "\n")
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			sb.WriteString(scanner.Text() + "\n")
		}
	}()
	wg.Wait()

	if err := cmd.Wait(); err != nil {
		log.Error("Error executing command", "command", command, "error", err)
	}

	l.Push(lua.LString(sb.String()))
	return 1
}

// call a shell command and return the output (stdout) in one string
func stdout(l *lua.LState) int {
	command := l.CheckString(1)

	cmd := util.System(command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating stdout pipe", "error", err)
		return 0
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting command", "error", err)
		return 0
	}

	sb := strings.Builder{}
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			sb.WriteString(scanner.Text() + "\n")
		}
	}()

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		log.Error("Error executing command", "command", command, "error", err)
	}

	l.Push(lua.LString(sb.String()))
	return 1
}

// call a shell command and return the error (stderr) in one string
func stderr(l *lua.LState) int {
	command := l.CheckString(1)

	cmd := util.System(command)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Error creating stderr pipe", "error", err)
		return 0
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting command", "error", err)
		return 0
	}

	sb := strings.Builder{}
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			sb.WriteString(scanner.Text() + "\n")
		}
	}()
	wg.Wait()

	if err := cmd.Wait(); err != nil {
		log.Error("Error executing command", "command", command, "error", err)
	}

	l.Push(lua.LString(sb.String()))
	return 1
}
