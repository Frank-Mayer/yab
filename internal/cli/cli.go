package cli

import (
	"flag"

	"github.com/Frank-Mayer/yab/internal/extensions"
	"github.com/charmbracelet/log"
)

type Cli struct {
	Version bool
	Help    bool
	Def     bool
	Update  bool
	Debug   bool
	Silent  bool
	Configs []string
}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) Parse() {
	flag.BoolVar(&c.Version, "version", false, "Prints the version of the program.")
	flag.BoolVar(&c.Version, "v", false, "Prints the version of the program.")
	flag.BoolVar(&c.Help, "help", false, "Prints this help.")
	flag.BoolVar(&c.Help, "h", false, "Prints this help.")
	flag.BoolVar(&c.Def, "def", false, "Creates definitions file in global config.")
	flag.BoolVar(&c.Update, "update", false, "Updates the Yab binary to the latest version.")
	flag.BoolVar(&c.Update, "upgrade", false, "Updates the Yab binary to the latest version.")
	flag.BoolVar(&c.Update, "u", false, "Updates the Yab binary to the latest version.")
	flag.BoolVar(&c.Debug, "debug", false, "Enables debug logging.")
	flag.BoolVar(&c.Silent, "silent", false, "Disables logging.")
	flag.Parse()

	args := flag.Args()
	sepIndex := -1
	for i, arg := range args {
		if arg == "--" {
			sepIndex = i
			break
		}
	}
	if sepIndex == -1 {
		c.Configs = args
	} else {
		c.Configs = args[:sepIndex]
		extensions.SetArgs(args[sepIndex+1:])
	}

	if c.Silent && c.Debug {
		log.Fatal("Cannot use both --silent and --debug")
	}
}
