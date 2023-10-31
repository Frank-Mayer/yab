package docs

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/glamour"
	"selene.frankmayer.io/extensions"
)

type Function struct {
	Name        string
	Description string
	Parameters  []string
	Returns     []string
}

var functions = []Function{
	{"os_type",
		"Returns the operating system type.",
		[]string{},
		[]string{"\"windows\" or \"unix\" on the respective system."},
	},
	{
		"args",
		"Returns the command line arguments passed to the program.",
		[]string{},
		[]string{"A table containing the command line arguments."},
	},
}

func Help() {
	term_width := os.Getenv("COLUMNS")

	width, err := strconv.Atoi(term_width)
	if err != nil {
		width = 80
	}

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)

	out, err := r.Render(Markdown())
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}

func Markdown() string {
	var sb strings.Builder
	sb.WriteString("# Documentation\n\n")

	bin_name := os.Args[0]
	if len(bin_name) > 24 {
		bin_name = "selene"
	}
	sb.WriteString("## Usage:\n\n")
	sb.WriteString(bin_name + " [configs ...]\n\n")
	sb.WriteString(bin_name + " [configs ...] -- [args ...]\n\n")

	sb.WriteString("## Command Line Arguments\n\n")
	sb.WriteString(bin_name + " [--version, -v]\n\n")
	sb.WriteString("Prints the version of the program.\n\n")
	sb.WriteString(bin_name + " [--help, -h]\n\n")
	sb.WriteString("Prints this help.\n\n")

	sb.WriteString("## Lua API Functions\n\n")
	for _, f := range functions {
		sb.WriteString(addFunction(&f))
		sb.WriteString("\n")
	}
	return sb.String()
}

func addFunction(f *Function) string {
	var sb strings.Builder
	sb.WriteString("### " + extensions.Name(f.Name) + "\n")
	sb.WriteString(f.Description + "\n\n")
	sb.WriteString("#### Parameters\n\n")
	if len(f.Parameters) > 0 {
		for _, p := range f.Parameters {
			sb.WriteString("* " + p + "\n")
		}
	} else {
		sb.WriteString("None\n")
	}
	sb.WriteString("\n")
	sb.WriteString("#### Returns\n\n")
	switch len(f.Returns) {
	case 0:
		sb.WriteString("None\n")
	case 1:
		sb.WriteString(f.Returns[0] + "\n")
	default:
		for _, r := range f.Returns {
			sb.WriteString("* " + r + "\n")
		}
	}
	return sb.String()
}
