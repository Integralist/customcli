package sftp

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"create":   "Create an SFTP logging endpoint on a Fastly service version",
	"list":     "List SFTP endpoints on a Fastly service versio",
	"describe": "Show detailed information about an SFTP logging endpoint on a Fastly service version",
	"update":   "Update an SFTP logging endpoint on a Fastly service version",
	"delete":   "Delete an SFTP logging endpoint on a Fastly service version",
}

func Run(args []string) {
	args = args[3:]
	cmd, err := arguments.IdentifyCommand(Cmds, args)
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("sftp cmd:", cmd)
}
