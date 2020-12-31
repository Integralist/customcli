package healthcheck

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"create":   "Create a healthcheck on a Fastly service version",
	"list":     "List healthchecks on a Fastly service version",
	"describe": "Show detailed information about a healthcheck on a Fastly service version",
	"update":   "Update a healthcheck on a Fastly service version",
	"delete":   "Delete a healthcheck on a Fastly service version",
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("healthcheck cmd:", cmd)
}
