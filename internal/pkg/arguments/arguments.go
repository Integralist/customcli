package arguments

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Program = iota
	Command
	Subcommand
	Flags
)

// IdentifyCommand parses the arguments provided looking for a 'command'.
//
// NOTE: the CLI is comprised of nested commands and so a 'command' in
// practical terms could be a category such as 'backend' (e.g. fastly backend)
// or an actual 'subcommand' such as 'create' (e.g. fastly backend create).
func IdentifyCommand(cmds map[string]string, args []string) (string, error) {
	commandIndex := 0
	commandSeen := false

	for _, arg := range args {
		if commandSeen {
			break
		}

		if strings.HasPrefix(arg, "-") == true {
			commandIndex++
			continue
		}

		for cmd := range cmds {
			if arg == cmd {
				commandSeen = true
				break
			}
		}
	}

	if !commandSeen {
		err := fmt.Sprintf("expected command but got %s", strings.Join(args, " "))
		return "", errors.New(err)
	}

	return args[commandIndex], nil
}
