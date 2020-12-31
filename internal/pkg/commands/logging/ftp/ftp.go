package ftp

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"create":   "Create an FTP logging endpoint on a Fastly service version",
	"list":     "List FTP endpoints on a Fastly service version",
	"describe": "Show detailed information about an FTP logging endpoint on a Fastly service version",
	"update":   "Update an FTP logging endpoint on a Fastly service version",
	"delete":   "Delete an FTP logging endpoint on a Fastly service version",
}

func Run(args []string) {
	args = args[3:]
	cmd, err := arguments.IdentifyCommand(Cmds, args)
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("ftp cmd:", cmd)
}
