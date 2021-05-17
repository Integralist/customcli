package commands

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type Run func(args []string)

type Base struct {
	List map[string]Command
}

func (b *Base) Add(c ...Command) {
	for _, cmd := range c {
		b.List[cmd.Name] = cmd
	}
}

func NewBase() *Base {
	return &Base{
		List: make(map[string]Command),
	}
}

func Find(cmds map[string]Command, args []string) (Command, error) {
	err := fmt.Errorf("couldn't find command: %s", strings.Join(args, " "))

	for i, arg := range args {
		for k, v := range cmds {
			if arg == k {
				if v.Base != nil {
					// Ensure a subcommand followed only by flags doesn't recurse.
					var (
						cmdCount  int
						flagCount int
					)
					subargs := args[i+1:]
					for _, arg := range subargs {
						if strings.HasPrefix(arg, "--") {
							flagCount++
						} else {
							cmdCount++
						}
					}
					if flagCount == 0 && cmdCount == 0 || flagCount == len(subargs) {
						return v, nil
					}
					return Find(v.List, subargs)
				}

				// Ensure a subcommand, with a nil Base, followed by another command
				// doesn't return the currently matched subcommand.
				subargs := args[i+1:]
				for _, arg := range subargs {
					if !strings.HasPrefix(arg, "--") {
						return Command{}, err
					}
				}
				return v, nil
			}
		}
	}

	return Command{}, err
}

type Command struct {
	*Base

	Name        string
	Description string
	Exec        Run
}

// Pad adds spacing so command descriptions in help output align.
func Pad(name string, longest int) string {
	return strings.Repeat(" ", longest-len(name))
}

// Gen generates a string consisting of the given commands/subcommands.
func Gen(cmds map[string]Command) string {
	var b bytes.Buffer

	// the following iteration over list is required for two reasons:
	// 1. sort the map for the sake of consistent output.
	// 2. identify the longest command for more structured output.
	keys := make([]string, 0, len(cmds))
	longest := 0
	for name := range cmds {
		keys = append(keys, name)
		if len(name) > longest {
			longest = len(name)
		}
	}
	sort.Strings(keys)

	for _, name := range keys {
		description := cmds[name].Description
		pad := Pad(name, longest)
		line := fmt.Sprintf("  %s"+pad+"  %s\n", name, description)
		b.Write([]byte(line))
	}
	return b.String()
}
