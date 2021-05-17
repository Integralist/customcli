package flags

import (
	"flag"
	"fmt"

	"github.com/integralist/customcli/pkg/formatting"
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

// Flag represents flag meta data.
// TODO: how to remove this type as it's not defined inside of commands.go ?
// Ideally we'd remove the flags package altogether.
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

// ErrorContext prints to stdout a common error followed by a context specific
// message.
func ErrorContext(err error) {
	fmt.Printf("\n%s: error parsing arguments: %v\n", formatting.RedBold("ERROR"), err)
}
