package dictionary

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"create":   "Create a Fastly edge dictionary on a Fastly service version",
	"list":     "List all dictionaries on a Fastly service version",
	"describe": "Show detailed information about a Fastly edge dictionary",
	"update":   "Update name of dictionary on a Fastly service version",
	"delete":   "Delete a Fastly edge dictionary from a Fastly service version",
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("dictionary cmd:", cmd)
}
