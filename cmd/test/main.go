package main

import (
	"github.com/Frank-Mayer/selene/internal/extensions"

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
	test("return Selene.os_type()", runtime.GOOS)
	test("return Selene.os_arch()", runtime.GOARCH)
	test("return Selene.stdall(\"echo foo\")", "foo\n")
	test("return Selene.stdout(\"echo foo\")", "foo\n")

	pack := []byte("{\"scripts\": {\"test\": \"echo foo\"}}")
	err := os.WriteFile("./package.json", pack, 0664)
	if err != nil {
		log.Error("Writing package.json", "err", err)
	} else {
		test("return Selene.js_run(\"test\")", "true")
		err := os.Remove("./package.json")
		if err != nil {
			log.Fatal("Error removing package.json", "err", err)
		}
	}

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
