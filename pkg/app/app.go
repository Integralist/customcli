package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/pkg/commands"
	"github.com/integralist/customcli/pkg/commands/backend"
	"github.com/integralist/customcli/pkg/commands/compute"
	"github.com/integralist/customcli/pkg/commands/configure"
	"github.com/integralist/customcli/pkg/flags"
)

// Run bootstraps the CLI.
//
// - Parses the global flags.
// - Identifies which subcommand should execute.
func Run(args []string) {
	cmds := []commands.Command{
		backend.New(),
		compute.New(),
		configure.New(),
		// ...
	}

	b := commands.NewBase()
	b.Add(cmds...)

	globals := flags.ParseGlobal(b, args)
	fmt.Printf("globals: %+v\n", globals) // TODO: pass globals into subcommands!

	// TODO: is the 'arguments' package constants used anywhere?
	c, err := commands.Find(b.List, flag.Args())
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	c.Exec(args)
}
