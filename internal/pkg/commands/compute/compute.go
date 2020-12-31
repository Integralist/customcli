package compute

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"init":     "Initialize a new Compute@Edge package locally",
	"build":    "Build a Compute@Edge package locally",
	"deploy":   "Deploy a package to a Fastly Compute@Edge service",
	"update":   "Update a package on a Fastly Compute@Edge service version",
	"validate": "Validate a Compute@Edge package",
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("compute cmd:", cmd)
}
