package describe

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
	"github.com/integralist/customcli/internal/pkg/print"
)

func Run(args []string) {
	var helpflag bool

	fs := flag.NewFlagSet("backend describe", flag.ExitOnError)
	fs.BoolVar(&helpflag, flags.Help.Long, false, flags.Help.Description)
	fs.BoolVar(&helpflag, flags.Help.Short, false, flags.Help.Description)

	// TODO: serve custom error when flag parsing fails (e.g. they pass --foobar)
	fs.Parse(args[arguments.Flags:])

	// custom parse flags for the sake of help output
	cmdFlags := flags.Parse(fs)

	if helpflag {
		example := "fastly backend describe --service-id=SERVICE-ID --version=VERSION --name=NAME"
		description := "Show detailed information about a backend on a Fastly service version"
		// TODO: code smell passing nil
		print.Help(example, description, "COMMAND FLAGS", cmdFlags)
		os.Exit(1)
	}

	fmt.Println("backend describe logic")
}
