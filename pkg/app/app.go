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
	"github.com/integralist/customcli/pkg/formatting"
)

// Run bootstraps the CLI.
//
// - Initialises all subcommands.
// - Appends subcommands to a base command.
// - Sets global flag options onto the base command.
// - Identifies which subcommand should be executed.
func Run(args []string) {
	app := commands.Command{
		Name:        "fastly",
		Description: "A tool to interact with the Fastly API",
	}
	cmds := []commands.Command{
		backend.New(),
		compute.New(),
		configure.New(),
		// ...
	}

	c := commands.NewChildren()
	c.Add(cmds...)

	app.Children = c
	setGlobalFlags(app, args)

	cmd, err := commands.Find(c.Cmds, flag.Args())
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}
	cmd.Exec(args)
}

func setGlobalFlags(c commands.Command, args []string) {
	var (
		help    bool
		token   string
		verbose bool
	)

	helpFlag := commands.Flag{
		Long:        "help",
		Short:       "h",
		Description: "Show context-sensitive help",
	}
	tokenFlag := commands.Flag{
		Long:        "token",
		Short:       "t",
		Description: "Fastly API token (or via FASTLY_API_TOKEN)",
	}
	verboseFlag := commands.Flag{
		Long:        "verbose",
		Short:       "v",
		Description: "Verbose logging",
	}

	flag.BoolVar(&help, helpFlag.Long, false, helpFlag.Description)
	flag.BoolVar(&help, helpFlag.Short, false, helpFlag.Description)
	flag.StringVar(&token, tokenFlag.Long, "", tokenFlag.Description)
	flag.StringVar(&token, tokenFlag.Short, "", tokenFlag.Description)
	flag.BoolVar(&verbose, verboseFlag.Long, false, verboseFlag.Description)
	flag.BoolVar(&verbose, verboseFlag.Short, false, verboseFlag.Description)

	c.Flags = append(c.Flags, helpFlag, tokenFlag, verboseFlag)

	inputs := commands.HelpOutput(c, args)
	inputs = append(inputs, formatting.Bold("COMMANDS"), commands.GenHelp(c.Children.Cmds))
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), formatting.HelpTemplate, inputs...)
	}
	flag.Parse()
}
