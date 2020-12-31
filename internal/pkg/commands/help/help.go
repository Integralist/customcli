package help

import (
	"fmt"
	"strings"

	"github.com/integralist/customcli/internal/pkg/commands/backend"
)

func Run(args []string) {
	// TODO: figure out a dynamic way to handle this...

	// NOTE: we pass the full args (e.g. fastly help <command>) to ensure
	// help.ErrorContext doesn't incorrectly attach extra information to
	// the help output.
	//
	// we pass nil for the third argument because we know ahead of time that we
	// don't want a contextual error message (e.g. ERROR: error parsing
	// arguments: expected command but got...)
	//
	// TODO: passing nil is a code smell

	// TODO: fill in the missing commands.
	switch strings.Join(args[2:], " ") {
	case "backend":
		// TODO: check all usage of flag.Arg()
		backend.Help(backend.Cmds, args[1:], nil)
	default:
		fmt.Println("TODO: show help for:", args[2:])
	}
}
