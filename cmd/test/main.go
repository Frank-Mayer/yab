package main

import (
	"github.com/Frank-Mayer/yab/internal/extensions"

	"os"
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/yuin/gopher-lua"
)

var (
	hasError = false
)

func main() {
	args := os.Args[1:]
	extensions.SetArgs(args)

	// test that the lua code works
	test("return 1", "1")
	test("return 1 + 1", "2")
	test("return Yab.os_type()", runtime.GOOS)
	test("return Yab.os_arch()", runtime.GOARCH)
	test("return Yab.stdall(\"echo foo\")", "foo\n")
	test("return Yab.stdout(\"echo foo\")", "foo\n")

	if hasError {
		os.Exit(1)
	}
}

func test(code string, expectedOutput string) {
	l := lua.NewState()
	defer l.Close()

	extensions.RegisterExtensions(l)

	err := l.DoString(code)
	if err != nil {
		log.Error(
			"Running test",
			"code", code,
			"err", err)
		hasError = true
		return
	}

	// check value on stack
	if l.GetTop() != 1 {
		log.Error(
			"Running test",
			"code", code,
			"err", "no value on stack")
		hasError = true
		return
	}
	ret := l.Get(-1) // returned value
	l.Pop(1)         // remove received value

	receivedOutput := ret.String()

	if receivedOutput != expectedOutput {
		log.Error(
			"Test failed",
			"code", code,
			"expected_output", expectedOutput,
			"received_output", receivedOutput)
		hasError = true
		return
	}

	// check that the stack is now empty
	if l.GetTop() != 0 {
		log.Error(
			"Running test",
			"code", code,
			"expected_output", expectedOutput,
			"received_output", receivedOutput,
			"err", "stack not empty")
		hasError = true
		return
	}

	log.Info("Test passed: " + code + " -> " + expectedOutput)
}
