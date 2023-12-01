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
		"Run a script from the `package.json` file using the first javascript package manager found. Trying pnpm, bun and npm in that order.",
		[]string{"script string"},
		[]string{},
		jsRun,
		"",
		"",
	},
	{
		"js_install",
		"Install dependencies from `package.json` using the first javascript package manager found. Trying pnpm, bun and npm in that order.",
		[]string{},
		[]string{},
		jsInstall,
		"",
		"",
	},
	{
		"git_clone_or_pull",
		"Clones a git repository to a specified destination. If the repository already exists, it will pull the latest changes instead.",
		[]string{"url string", "destination string"},
		[]string{},
		gitCloneOrPull,
		"",
		"",
	},
	{
		"zip",
		"Create a zip file containing the given files.",
		[]string{"files table", "output string"},
		[]string{},
		makeZip,
		"",
		"Yab.zip(\n\t{'foo.txt', 'bar.txt', 'baz/'},\n\t'archive.zip'\n)",
	},
	{
		"watch",
		"Watch file or directory paths for changes and call a function when a change occurs. " +
			"The callback function will be called with the file path and the event type as arguments. " +
			"The event type can be one of 'create', 'write', 'remove', 'rename' or 'chmod'.",
		[]string{"paths table", "callback function(string, string)"},
		[]string{},
		watch,
		"",
		"Yab.watch('foo.txt', function(file, event)\n\tprint('foo.txt changed!')\nend)",
	},
	{
		"block",
		"Block the current thread and wait for an interrupt signal.",
		[]string{},
		[]string{},
		block,
		"",
		"Yab.block()",
	},
	{
		"find",
		"Find files matching a pattern in a directory.",
		[]string{"pattern string"},
		[]string{"A table containing the matching file paths."},
		find,
		"table",
		"Yab.find('*.txt')",
	},
	{
		"find",
		"Find files matching a pattern in a directory.",
		[]string{"root string", "pattern string"},
		[]string{"A table containing the matching file paths."},
		find,
		"table",
		"Yab.find('foo', '*.txt')",
	},
	{
		"fileinfo",
		"Get information about a file.",
		[]string{"path string"},
		[]string{"A table containing the file information (name, size, mode, modtime, isdir, sys). See https://pkg.go.dev/io/fs#FileInfo for details."},
		fileinfo,
		"table",
		"local foo_info = Yab.fileinfo('foo.txt')\nprint(foo_info.size)",
	},
}

func Definitions() string {
	sb := strings.Builder{}
	sb.WriteString("---@meta\n")
	sb.WriteString("---@class Yab\n")
	sb.WriteString("Yab = {}\n")
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
		sb.WriteString("Yab.")
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

	l.SetGlobal("Yab", table)
}
