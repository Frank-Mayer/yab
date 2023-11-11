package extensions

import (
	"strings"

	"github.com/yuin/gopher-lua"
)

type Function struct {
	Name        string
	Description string
	Parameters  []string
	Returns     []string
	Function    func(l *lua.LState) int
	Ret         string
    Example     string
}

var Functions = []Function{
	{
		"os_type",
		"Returns the operating system type.",
		[]string{},
		[]string{"\"windows\", \"linux\" or \"darwin\" on the respective system."},
		osType,
		"'windows'|'linux'|'darwin'",
        "",
	},
	{
		"os_arch",
		"Returns the operating system architecture.",
		[]string{},
		[]string{"\"amd64\" or \"arm64\" on the respective system."},
		osArch,
		"'amd64'|'arm64'",
        "",
	},
	{
		"args",
		"Returns the command line arguments passed to the program.",
		[]string{},
		[]string{"A table containing the command line arguments."},
		args,
		"table",
        "",
	},
	{
		"check_exec",
		"Checks if an executable is available in the system's PATH.",
		[]string{"executable string"},
		[]string{"true if the executable is available, false otherwise."},
		checkExec,
		"boolean",
        "",
	},
	{
		"stdall",
		"Call a shell command and return the full output (stdout + stderr) in one string.",
		[]string{"command string"},
		[]string{"The output of the command."},
		stdall,
		"string",
        "",
	},
	{
		"stdout",
		"Call a shell command and return the output (stdout) in one string.",
		[]string{"command string"},
		[]string{"The output of the command."},
		stdout,
		"string",
        "",
	},
	{
		"stderr",
		"Call a shell command and return the error output (stderr) in one string.",
		[]string{"command string"},
		[]string{"The output of the command."},
		stderr,
		"string",
        "",
	},
	{
		"js_run",
		"Run a script from the `package.json` file using the first javascript package manager found. Trying pnpm, yarn, bun and npm in that order.",
		[]string{"script string"},
		[]string{"true if a javascript package manager was found, false otherwise."},
		jsRun,
		"boolean",
        "",
	},
	{
		"js_install",
		"Install dependencies from `package.json` using the first javascript package manager found. Trying pnpm, yarn, bun and npm in that order.",
		[]string{},
		[]string{"true if a javascript package manager was found, false otherwise."},
		jsInstall,
		"boolean",
        "",
	},
	{
		"git_clone_or_pull",
		"Clones a git repository to a specified destination. If the repository already exists, it will pull the latest changes instead.",
		[]string{"url string", "destination string"},
		[]string{"true if the repository was cloned or pulled successfully, false otherwise."},
		gitCloneOrPull,
		"boolean",
        "",
	},
    {
        "zip",
        "Create a zip file containing the given files.",
        []string{"files table", "output string"},
        []string{"true if the zip file was created successfully, false otherwise."},
        makeZip,
        "boolean",
        "Selene.zip(\n\t{'foo.txt', 'bar.txt', 'baz/'},\n\t'archive.zip'\n)",
    },
}

func Definitions() string {
	sb := strings.Builder{}
	sb.WriteString("---@meta\n")
	sb.WriteString("---@class Selene\n")
	sb.WriteString("Selene = {}\n")
	for _, f := range Functions {
		sb.WriteString("\n")
		for _, p := range f.Parameters {
			sb.WriteString("---@param ")
			sb.WriteString(p)
			sb.WriteString("\n")
		}
		sb.WriteString("---@return ")
		sb.WriteString(f.Ret)
		sb.WriteString("\n")
		sb.WriteString("---")
		sb.WriteString(f.Description)
		sb.WriteString("\n")
		sb.WriteString("Selene.")
		sb.WriteString(f.Name)
		sb.WriteString(" = function(")
		for i, p := range f.Parameters {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(strings.Split(p, " ")[0])
		}
		sb.WriteString(")\nend\n")
	}
	return sb.String()
}

func RegisterExtensions(l *lua.LState) {
	table := l.NewTable()
	for _, f := range Functions {
		l.SetTable(table, lua.LString(f.Name), l.NewFunction(f.Function))
	}

	l.SetGlobal("Selene", table)
}
