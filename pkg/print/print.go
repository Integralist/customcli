package print

import (
	"github.com/integralist/customcli/pkg/flags"
)

// // Help generates help output for commands.
// func Help(usage string, description string, title string, items interface{}) {
// 	var short string
// 	globalFlags := make([]flags.Flag, 0)

// 	// TODO: don't rely on global state access like flag.VisitAll
// 	// TODO: how to reuse flags.Parse which only works for FlagSet
// 	flag.VisitAll(func(f *flag.Flag) {
// 		// flags are visited in lexical order so we know that the short flag will
// 		// appear first followed by the long flag.
// 		if len(f.Name) == 1 {
// 			short = f.Name
// 			return
// 		}
// 		globalFlags = append(globalFlags, flags.Flag{
// 			Long:        f.Name,
// 			Short:       short,
// 			Description: f.Usage,
// 		})
// 	})

// 	// items can be either a collection of subcommands or flags for a command/subcommand.
// 	var itemsParsed string
// 	switch t := items.(type) {
// 	case map[string]string:
// 		itemsParsed = commands.Gen(t)
// 	case []flags.Flag:
// 		itemsParsed = flags.GenExample(t...)
// 	}

// 	inputs := []interface{}{
// 		formatting.Bold("USAGE"),
// 		usage,
// 		description,
// 		formatting.Bold("GLOBAL FLAGS"),
// 		flags.GenExample(globalFlags...),
// 		formatting.Bold(title),
// 		itemsParsed,
// 	}

// 	formatting.Template(flag.CommandLine.Output(), formatting.HelpTemplate, inputs...)
// }

// ErrorContext prints additional contextual information to the bottom of help
// output such as when an unexpected subcommand was passed.
func ErrorContext(args []string, err error) {
	if len(args) > 0 && args[0] != "-help" && args[0] != "help" {
		flags.ErrorContext(err)
	}
}
