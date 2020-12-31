package backend

import (
	"os"
	"strings"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/commands/backend/create"
	"github.com/integralist/customcli/internal/pkg/commands/backend/delete"
	"github.com/integralist/customcli/internal/pkg/commands/backend/describe"
	"github.com/integralist/customcli/internal/pkg/commands/backend/list"
	"github.com/integralist/customcli/internal/pkg/commands/backend/update"
	"github.com/integralist/customcli/internal/pkg/print"
)

// Cmds represents all available subcommands.
//
// NOTE: this is a public function because the `fastly help command` requires
// knowledge of the available commands to build the appropriate help output.
//
// TODO: add this docstring to all Cmds references.
var Cmds = map[string]string{
	"create":   "Create a backend on a Fastly service version",
	"list":     "List backends on a Fastly service version",
	"describe": "Show detailed information about a backend on a Fastly service version",
	"update":   "Update a backend on a Fastly service version",
	"delete":   "Delete a backend on a Fastly service version",
}

// Help prints the help info for the command, then exits.
//
// NOTE: we extract this logic into an exported function because we want to
// reuse most of the logic from the top-level 'help' command.
// e.g. fastly help <command>
func Help(cmds map[string]string, args []string, err error) {
	example := "fastly backend <command> [<args> ...]"
	description := "Manipulate Fastly service version backends"
	// TODO: originally was category.Usage and passed in a map!
	// but we need a way to generate SUBCOMMANDS and not COMMANDS too.
	// we need to update all Help examples as this function is copied/pasted in
	// each top level command.
	print.Help(example, description, "COMMANDS", cmds)
	if err != nil {
		print.ErrorContext(args, err)
	}
	os.Exit(1)
}

// Run executes the command logic.
// TODO: add this docstring to all instances of Run().
func Run(args []string) {
	a := args[arguments.Subcommand:]
	cmd, err := arguments.IdentifyCommand(Cmds, a)

	// TODO: code smell passing nil
	// TODO: this needs abstracting!
	// We want to catch whenever --help is called like:
	// fastly <command> --help <command>
	// ...so we can display the help output of the parent command.
	// But we also want to avoid displaying that particular help if the help flag
	// appears at the end of the arguments list.
	if len(args) >= arguments.Subcommand+1 && strings.Contains(args[arguments.Subcommand], "help") {
		Help(Cmds, a, nil)
	}

	if err != nil {
		// TODO: ensure all categories use this pattern instead of flag.Usage
		// TODO: the Help function needs to know if it's printing 'commands' or
		// 'command flags'! so we might have to pass that in initially and then
		// dynamically determine it when moving to a struct/object pattern.
		Help(Cmds, a, err)
	}

	switch cmd {
	case "create":
		create.Run(args)
	case "list":
		list.Run(args)
	case "describe":
		describe.Run(args)
	case "update":
		update.Run(args)
	case "delete":
		delete.Run(args)
	}
}
