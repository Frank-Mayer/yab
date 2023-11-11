package docs

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"selene.frankmayer.dev/extensions"
	"selene.frankmayer.dev/util"
)

func Help() {
	width := util.TermWidth()

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
	sb.WriteString("# Documentation v")
	sb.WriteString(util.Version)
	sb.WriteString("\n\n")

	bin_name := util.BinName()

	sb.WriteString("## Usage:\n\n")
	sb.WriteString(bin_name + " [configs ...]\n\n")
	sb.WriteString(bin_name + " [configs ...] -- [args ...]\n\n")

	sb.WriteString("## Command Line Arguments\n\n")
	sb.WriteString("**" + bin_name + " [--version, -v]**\n\n")
	sb.WriteString("Prints the version of the program.\n\n")
	sb.WriteString("**" + bin_name + " [--help, -h]**\n\n")
	sb.WriteString("Prints this help.\n\n")
	sb.WriteString("**" + bin_name + " --init**\n\n")
	sb.WriteString("Initializes a new Selene project.\n\n")
	sb.WriteString("**" + bin_name + " [--update, --upgrade, -u]**\n\n")
	sb.WriteString("Updates the Selene binary to the latest version.\n\n")

	sb.WriteString("## Lua API Functions (in the `Selene` global table)\n\n")
	for _, f := range extensions.Functions {
		sb.WriteString(addFunction(&f))
		sb.WriteString("\n")
	}
	return sb.String()
}

func addFunction(f *extensions.Function) string {
	var sb strings.Builder
	sb.WriteString("### 𝑓 " + f.Name + "\n\n")
	sb.WriteString("*" + f.Description + "*\n\n")
	sb.WriteString("**Parameters:** ")
	if len(f.Parameters) > 0 {
		sb.WriteString("\n")
		for _, p := range f.Parameters {
			sb.WriteString("* ")
			param_words := strings.Split(p, " ")
			sb.WriteString(param_words[0])
			sb.WriteString(" `")
			sb.WriteString(strings.Join(param_words[1:], " "))
			sb.WriteString("`\n")
		}
	} else {
		sb.WriteString("None\n")
	}
	sb.WriteString("\n")
	sb.WriteString("**Returns:** ")
	switch len(f.Returns) {
	case 0:
		sb.WriteString("None\n")
	case 1:
		sb.WriteString(f.Returns[0] + "\n")
	default:
		sb.WriteString("\n")
		for _, r := range f.Returns {
			sb.WriteString("* " + r + "\n")
		}
	}

	if f.Example != "" {
		sb.WriteString("\n**Example:**\n\n```lua\n")
		sb.WriteString(f.Example)
		sb.WriteString("\n```\n")
	}

	return sb.String()
}
