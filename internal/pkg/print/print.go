package print

import (
	"flag"
	"fmt"
	"strings"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/commands"
	"github.com/integralist/customcli/internal/pkg/flags"
	"github.com/integralist/customcli/internal/pkg/formatting"
)

// Help generates help output for commands.
func Help(usage string, description string, title string, items interface{}) {
	var short string
	globalFlags := make([]flags.Flag, 0)

	// TODO: don't rely on global state access like flag.VisitAll
	// TODO: how to reuse flags.Parse which only works for FlagSet
	flag.VisitAll(func(f *flag.Flag) {
		// flags are visited in lexical order so we know that the short flag will
		// appear first followed by the long flag.
		if len(f.Name) == 1 {
			short = f.Name
			return
		}
		globalFlags = append(globalFlags, flags.Flag{
			Long:        f.Name,
			Short:       short,
			Description: f.Usage,
		})
	})

	// items can be either a collection of subcommands or flags for a command/subcommand.
	var itemsParsed string
	switch t := items.(type) {
	case map[string]string:
		itemsParsed = commands.Gen(t)
	case []flags.Flag:
		itemsParsed = flags.GenExample(t...)
	}

	inputs := []interface{}{
		formatting.Bold("USAGE"),
		usage,
		description,
		formatting.Bold("GLOBAL FLAGS"),
		flags.GenExample(globalFlags...),
		formatting.Bold(title),
		itemsParsed,
	}

	formatting.Template(flag.CommandLine.Output(), formatting.HelpTemplate, inputs...)
}

// GenUsage generates example usage for print.Help().
//
// Example: fastly foo bar --abc=ABC --xyz=XYZ
func GenUsage(fs *flag.FlagSet, args []string) string {
	cmd := args[arguments.Command : len(args)-1]
	usage := "fastly " + strings.Join(cmd, " ")

	// TODO: skip 'optional' flags and instead display [<flags>]
	fs.VisitAll(func(f *flag.Flag) {
		if f.Name != "h" && f.Name != "help" && len(f.Name) > 1 {
			usage = fmt.Sprintf("%s --%s=%s", usage, f.Name, strings.ToUpper(f.Name))
		}
	})

	return usage
}

// ErrorContext prints additional contextual information to the bottom of help
// output such as when an unexpected subcommand was passed.
func ErrorContext(args []string, err error) {
	if len(args) > 0 && args[0] != "-help" && args[0] != "help" {
		flags.ErrorContext(err)
	}
}
