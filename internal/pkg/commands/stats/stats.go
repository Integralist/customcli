package stats

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"regions":    "List stats regions",
	"historical": "View historical stats for a Fastly service",
	"realtime":   "View realtime stats for a Fastly service",
}

func Run(args []string) {
	cmd, err := arguments.IdentifyCommand(Cmds, args[2:])
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	fmt.Println("stats cmd:", cmd)
}
