package create

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
		Name:        "create",
		Description: "Create a backend on a Fastly service version",
		Exec:        run,
	}
}

func run(args []string) {
	var (
		helpflag bool

		// Required.
		serviceID      string
		serviceVersion int
		backendName    string
		backendAddress string

		// Optional.
		backendComment string

		// TODO: implement all other flags.
	)

	fs := flag.NewFlagSet("backend create", flag.ExitOnError)
	fs.BoolVar(&helpflag, flags.Help.Long, false, flags.Help.Description)
	fs.BoolVar(&helpflag, flags.Help.Short, false, flags.Help.Description)
	fs.StringVar(&serviceID, flags.ServiceID.Long, "", flags.ServiceID.Description)
	fs.StringVar(&serviceID, flags.ServiceID.Short, "", flags.ServiceID.Description)
	fs.IntVar(&serviceVersion, flags.ServiceVersion.Long, 0, flags.ServiceVersion.Description)
	fs.StringVar(&backendName, "name", "", "Backend name")
	fs.StringVar(&backendName, "n", "", "Backend name")
	fs.StringVar(&backendAddress, "address", "", "A hostname, IPv4, or IPv6 address for the backend")
	fs.StringVar(&backendComment, "comment", "", "A descriptive note")

	// TODO: serve custom error when flag parsing fails (e.g. they pass --foobar)
	fs.Parse(args[arguments.Flags:])

	// custom parse flags for the sake of help output
	// cmdFlags := flags.Parse(fs)

	// TODO: handle required flags (e.g. error if no value provided).

	if helpflag {
		// usage := commands.GenUsage(fs, args)
		// description := "Create a backend on a Fastly service version"
		// title := "COMMAND FLAGS"
		// print.Help(usage, description, title, cmdFlags)
		os.Exit(1)
	}

	fmt.Println("backend create:", args)
}
