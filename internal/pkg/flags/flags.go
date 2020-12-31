package flags

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/commands"
	"github.com/integralist/customcli/internal/pkg/formatting"
)

// Help is a common flag that is better shared than repeating the definition.
var Help = Flag{
	Long:        "help",
	Short:       "h",
	Description: "Show context-sensitive help",
}

// ServiceID is a common flag that is better shared than repeating the definition.
var ServiceID = Flag{
	Long:        "service-id",
	Short:       "s",
	Description: "Service ID",
}

// ServiceVersion is a common flag that is better shared than repeating the definition.
var ServiceVersion = Flag{
	Long:        "version",
	Description: "Number of service version",
}

// Global represents globally available flags.
type Global struct {
	Help    bool
	Token   string
	Verbose bool
}

// Flag represents flag meta data.
type Flag struct {
	Long        string
	Short       string
	Description string
}

// Parse loops over given flagset and populates given the []Flag.
//
// The result of this function is typically passed to GenExample.
func Parse(fs *flag.FlagSet) []Flag {
	var short string
	cmdFlags := make([]Flag, 0)
	fs.VisitAll(func(f *flag.Flag) {
		// -h, -help, --help is a global flag so we can filter it from the list of
		// command flags otherwise it becomes overloaded in the help output.
		if f.Name == "help" || f.Name == "h" {
			return
		}

		// TODO: check what happens when global flags are provided! seems like our
		// argument parsing breaks down.

		// flags are visited in lexical order so we know that the short flag will
		// appear first followed by the long flag.
		if len(f.Name) == 1 {
			short = f.Name
			return
		}
		cmdFlags = append(cmdFlags, Flag{
			Long:        f.Name,
			Short:       short,
			Description: f.Usage,
		})
	})
	return cmdFlags
}

// GenExample generates example usage for print.Help().
//
// It is expected to be a custom replacement for flag.PrintDefaults so that
// long and short flags can be displayed on one line rather than being
// unnecessarily duplicated across multiple lines.
func GenExample(flags ...Flag) string {
	var b bytes.Buffer

	longest := 0
	for _, f := range flags {
		if len(f.Long) > longest {
			longest = len(f.Long)
		}
	}

	for _, f := range flags {
		// TODO: Pad should be in a different package as it pads commands AND flags
		pad := commands.Pad(f.Long, longest)

		// TODO: in real fastly tool it also shows assignment value, and it uses
		// double dashes not single.
		// e.g. --token=TOKEN
		var line string
		if len(f.Short) > 0 {
			line = fmt.Sprintf("  -%s, -%s"+pad+"  %s\n", f.Short, f.Long, f.Description)
		} else {
			line = fmt.Sprintf("      -%s"+pad+"  %s\n", f.Long, f.Description)
		}
		b.Write([]byte(line))
	}
	return b.String()
}

// ParseGlobal parses the expected global flags and sets flag.Usage to a custom
// implementation.
//
// TODO: think about pulling Service ID from manifest file for demonstration
// purposes so it's closer to how fastly/cli works.
func ParseGlobal(args []string) Global {
	var (
		help    bool
		token   string
		verbose bool
	)

	helpFlag := Flag{"help", "h", "Show context-sensitive help"}
	tokenFlag := Flag{"token", "t", "Fastly API token (or via FASTLY_API_TOKEN)"}
	verboseFlag := Flag{"verbose", "v", "Verbose logging"}

	flag.BoolVar(&help, helpFlag.Long, false, helpFlag.Description)
	flag.BoolVar(&help, helpFlag.Short, false, helpFlag.Description)
	flag.StringVar(&token, tokenFlag.Long, "", tokenFlag.Description)
	flag.StringVar(&token, tokenFlag.Short, "", tokenFlag.Description)
	flag.BoolVar(&verbose, verboseFlag.Long, false, verboseFlag.Description)
	flag.BoolVar(&verbose, verboseFlag.Short, false, verboseFlag.Description)

	// flag.Usage is called when there is an error parsing flags.
	inputs := []interface{}{
		formatting.Bold("\nUSAGE"),
		"fastly [<flags>] <command> [<args> ...]",
		"A tool to interact with the Fastly API",
		formatting.Bold("GLOBAL FLAGS"),
		GenExample(helpFlag, tokenFlag, verboseFlag),
		formatting.Bold("COMMANDS"),
		commands.Gen(commands.List),
	}

	// we need to also modify the flag.Usage here because we need access to the
	// defined global flags, and if there's an error parsing those flags then the
	// usage output needs to be updated so it can display our custom output.
	flag.Usage = func() {
		formatting.Template(flag.CommandLine.Output(), formatting.HelpTemplate, inputs...)
	}

	flag.Parse()

	if len(args) == 1 {
		flag.Usage()
		ErrorContext(errors.New("command not specified"))
		os.Exit(1)
	}

	if help {
		flag.Usage()
		os.Exit(1)
	}

	return Global{
		Help:    help,
		Token:   token,
		Verbose: verbose,
	}
}

// ErrorContext prints to stdout a common error followed by a context specific
// message.
func ErrorContext(err error) {
	fmt.Printf("\n%s: error parsing arguments: %v.\n", formatting.RedBold("ERROR"), err)
}
