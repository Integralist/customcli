package dictionaryitem

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"create":      "Create a new item on a Fastly edge dictionary",
	"list":        "List items in a Fastly edge dictionary",
	"describe":    "Show detailed information about a Fastly edge dictionary item",
	"update":      "Update or insert an item on a Fastly edge dictionary",
	"delete":      "Delete an item from a Fastly edge dictionary",
	"batchmodify": "Update multiple items in a Fastly edge dictionary",
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("dictionaryitem cmd:", cmd)
}
