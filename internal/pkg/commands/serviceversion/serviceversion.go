package serviceversion

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"clone":      "Clone a Fastly service version",
	"list":       "List Fastly service versions",
	"update":     "Update a Fastly service version",
	"activate":   "Activate a Fastly service version",
	"deactivate": "Deactivate a Fastly service version",
	"lock":       "Lock a Fastly service version",
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("service-version cmd:", cmd)
}
