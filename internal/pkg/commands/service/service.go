package service

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"create":   "Create a Fastly service",
	"list":     "List Fastly services",
	"describe": "Show detailed information about a Fastly service",
	"update":   "Update a Fastly service",
	"delete":   "Delete a Fastly service",
	"search":   "Search for a Fastly service by name",
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("service cmd:", cmd)
}
