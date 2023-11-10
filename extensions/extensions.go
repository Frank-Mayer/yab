package extensions

import "github.com/yuin/gopher-lua"

type Function struct {
	Name        string
	Description string
	Parameters  []string
	Returns     []string
	Function    func(l *lua.LState) int
}

var Functions = []Function{
	{
		"os_type",
		"Returns the operating system type.",
		[]string{},
		[]string{"\"windows\", \"linux\" or \"darwin\" on the respective system."},
		osType,
	},
	{
		"os_arch",
		"Returns the operating system architecture.",
		[]string{},
		[]string{"\"amd64\" or \"arm64\" on the respective system."},
		osArch,
	},
	{
		"args",
		"Returns the command line arguments passed to the program.",
		[]string{},
		[]string{"A table containing the command line arguments."},
		args,
	},
	{
		"check_exec",
		"Checks if an executable is available in the system's PATH.",
		[]string{"executable: *string*"},
		[]string{"true if the executable is available, false otherwise."},
		checkExec,
	},
	{
		"stdall",
		"Call a shell command and return the full output (stdout + stderr) in one string.",
		[]string{"command: *string*"},
		[]string{"The output of the command."},
		stdall,
	},
	{
		"stdout",
		"Call a shell command and return the output (stdout) in one string.",
		[]string{"command: *string*"},
		[]string{"The output of the command."},
		stdout,
	},
	{
		"stderr",
		"Call a shell command and return the error output (stderr) in one string.",
		[]string{"command: *string*"},
		[]string{"The output of the command."},
		stderr,
	},
	{
		"js_run",
		"Run a script from the `package.json` file using the first javascript package manager found. Trying pnpm, yarn, bun and npm in that order.",
		[]string{"script: *string*"},
		[]string{"true if a javascript package manager was found, false otherwise."},
		jsRun,
	},
	{
		"js_install",
		"Install dependencies from `package.json` using the first javascript package manager found. Trying pnpm, yarn, bun and npm in that order.",
		[]string{},
		[]string{"true if a javascript package manager was found, false otherwise."},
		jsInstall,
	},
}

func RegisterExtensions(l *lua.LState) {
	table := l.NewTable()
	for _, f := range Functions {
		l.SetTable(table, lua.LString(f.Name), l.NewFunction(f.Function))
	}

	l.SetGlobal("Selene", table)
}
