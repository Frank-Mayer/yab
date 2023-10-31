package extensions

import "github.com/Shopify/go-lua"

type Function struct {
	Name        string
	Description string
	Parameters  []string
	Returns     []string
	Function    lua.Function
}

var Functions = []Function{
	{"os_type",
		"Returns the operating system type.",
		[]string{},
		[]string{"\"windows\" or \"unix\" on the respective system."},
		os_type,
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
		[]string{"executable"},
		[]string{"true if the executable is available, false otherwise."},
		checkExec,
	},
	{
		"stdall",
		"Call a shell command and return the full output (stdout + stderr) in one string.",
		[]string{"command"},
		[]string{"The output of the command."},
		stdall,
	},
	{
		"stdout",
		"Call a shell command and return the output (stdout) in one string.",
		[]string{"command"},
		[]string{"The output of the command."},
		stdout,
	},
	{
		"stderr",
		"Call a shell command and return the error output (stderr) in one string.",
		[]string{"command"},
		[]string{"The output of the command."},
		stderr,
	},
}

func RegisterExtensions(l *lua.State) {
	l.CreateTable(0, len(Functions))

	for _, f := range Functions {
		l.PushGoFunction(f.Function)
		l.SetField(-2, f.Name)
	}

	l.SetGlobal("selene")
}
