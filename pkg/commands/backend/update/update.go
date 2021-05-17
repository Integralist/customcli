package update

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/pkg/arguments"
	"github.com/integralist/customcli/pkg/commands"
	"github.com/integralist/customcli/pkg/flags"
)

func New() commands.Command {
	return commands.Command{
		Children:    commands.NewChildren(), // defines `help`
		Name:        "update",
		Description: "Update a backend on a Fastly service version",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("backend update:", args)
}

func Run(args []string) {
	var helpflag bool

	fs := flag.NewFlagSet("backend update", flag.ExitOnError)
	fs.BoolVar(&helpflag, flags.Help.Long, false, flags.Help.Description)
	fs.BoolVar(&helpflag, flags.Help.Short, false, flags.Help.Description)

	// TODO: serve custom error when flag parsing fails (e.g. they pass --foobar)
	fs.Parse(args[arguments.Flags:])

	// custom parse flags for the sake of help output
	// cmdFlags := flags.Parse(fs)

	if helpflag {
		// example := "fastly backend update --service-id=SERVICE-ID --version=VERSION --name=NAME [<flags>]"
		// description := "Update a backend on a Fastly service version"
		// print.Help(example, description, "COMMAND FLAGS", cmdFlags)
		os.Exit(1)
	}

	fmt.Println("backend update logic")
}
