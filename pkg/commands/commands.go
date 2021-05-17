package commands

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/integralist/customcli/pkg/formatting"
)

type Run func(args []string)

// Children represents a collection of subcommands.
type Children struct {
	Cmds  map[string]Command
	Flags []Flag
}

func (c *Children) Add(cmds ...Command) {
	for _, cmd := range cmds {
		c.Cmds[cmd.Name] = cmd
	}
}

// TODO: where do we implement...
//
// if len(args) == 1 {
// 	flag.Usage()
// 	ErrorContext(errors.New("command not specified"))
// 	os.Exit(1)
// }

// if help {
// 	flag.Usage()
// 	os.Exit(1)
// }
func (c *Children) Help(args []string) {
	inputs := c.helpOutput()
	inputs = append(inputs, formatting.Bold("COMMANDS"), GenHelp(c.Cmds))
	fmt.Fprintf(os.Stdout, formatting.HelpTemplate, inputs...)
}

func (c *Children) helpOutput() []interface{} {
	inputs := []interface{}{
		formatting.Bold("\nUSAGE"),
		c.genUsage(),
		"GENERATE >> A tool to interact with the Fastly API",
		formatting.Bold("GLOBAL FLAGS"),
		c.genFlags(),
	}

	// The check for c.Parent.Children means we don't try to generate command flags
	// for the top-level group that was initialised in app.Run()
	//
	// TODO: rethink this...
	// if len(c.Flags) > 0 && c.Parent.Children != nil {
	if len(c.Flags) > 0 {
		inputs = append(inputs, formatting.Bold("COMMAND FLAGS"), c.genFlags())
	}

	return inputs
}

func HelpOutput(c Command, args []string) []interface{} {
	inputs := []interface{}{
		formatting.Bold("\nUSAGE"),
		GenUsage(c, args),
		c.Description,
		formatting.Bold("GLOBAL FLAGS"),
		GenFlags(c),
	}

	// Avoid generating command flags that are a duplicate of the global flags
	// when rendering help output for the root app command.
	if len(c.Flags) > 0 && c.Name != "fastly" {
		inputs = append(inputs, formatting.Bold("COMMAND FLAGS"), GenFlags(c))
	}

	return inputs
}

// genUsage generates example usage for help output.
//
// Example: fastly foo bar --abc=ABC --xyz=XYZ
//
// TODO: implement usage generation
func (c *Children) genUsage() string {
	return "fastly [<flags>] <command> [<args> ...]"

	// cmd := args[arguments.Command : len(args)-1]
	// usage := "fastly " + strings.Join(cmd, " ")

	// // TODO: skip 'optional' flags and instead display [<flags>]
	// fs.VisitAll(func(f *flag.Flag) {
	// 	if f.Name != "h" && f.Name != "help" && len(f.Name) > 1 {
	// 		usage = fmt.Sprintf("%s --%s=%s", usage, f.Name, strings.ToUpper(f.Name))
	// 	}
	// })

	// return usage
}

func GenUsage(c Command, args []string) string {
	fmt.Printf("\n\ncommand %+v\n\n", c)
	fmt.Printf("\n\nargs %+v\n\n", args)
	return "fastly [<flags>] <command> [<args> ...]"
}

// genFlags generates example flag usage for help output.
//
// It is expected to be a custom replacement for flag.PrintDefaults so that
// long and short flags can be displayed on one line rather than being
// unnecessarily duplicated across multiple lines.
func (c *Children) genFlags() string {
	if len(c.Flags) == 0 {
		fmt.Printf("\n\ngenFlags c.Flags %+v\n\n", c)
		// fs = b.Children.Flags
	}
	fmt.Printf("\n\ngenFlags c.Flags len %+v\n\n", len(c.Flags))

	var buf bytes.Buffer

	longest := 0
	for _, f := range c.Flags {
		if len(f.Long) > longest {
			longest = len(f.Long)
		}
	}

	for _, f := range c.Flags {
		pad := formatting.Pad(f.Long, longest)

		// TODO: in real fastly tool it also shows assignment value, and it uses
		// double dashes not single.
		// e.g. --token=TOKEN
		var line string
		if len(f.Short) > 0 {
			line = fmt.Sprintf("  -%s, -%s"+pad+"  %s\n", f.Short, f.Long, f.Description)
		} else {
			line = fmt.Sprintf("      -%s"+pad+"  %s\n", f.Long, f.Description)
		}
		buf.Write([]byte(line))
	}
	return buf.String()
}

func GenFlags(c Command) string {
	var buf bytes.Buffer

	longest := 0
	for _, f := range c.Flags {
		if len(f.Long) > longest {
			longest = len(f.Long)
		}
	}

	for _, f := range c.Flags {
		pad := formatting.Pad(f.Long, longest)

		// TODO: in real fastly tool it also shows assignment value, and it uses
		// double dashes not single.
		// e.g. --token=TOKEN
		var line string
		if len(f.Short) > 0 {
			line = fmt.Sprintf("  -%s, -%s"+pad+"  %s\n", f.Short, f.Long, f.Description)
		} else {
			line = fmt.Sprintf("      -%s"+pad+"  %s\n", f.Long, f.Description)
		}
		buf.Write([]byte(line))
	}
	return buf.String()
}

// TODO: think about pulling Service ID from manifest file for demonstration
// purposes so it's closer to how fastly/cli works (this will mean needing to
// implement a global --service-id flag).
func (c *Children) SetGlobalFlags() {
	var (
		help    bool
		token   string
		verbose bool
	)

	helpFlag := Flag{"help", "h", "Show context-sensitive help"}
	tokenFlag := Flag{"token", "t", "Fastly API token (or via FASTLY_API_TOKEN)"}
	verboseFlag := Flag{"verbose", "v", "Verbose logging"}

	flag.BoolVar(&help, helpFlag.Long, false, helpFlag.Description)
	flag.BoolVar(&help, helpFlag.Short, false, helpFlag.Description)
	flag.StringVar(&token, tokenFlag.Long, "", tokenFlag.Description)
	flag.StringVar(&token, tokenFlag.Short, "", tokenFlag.Description)
	flag.BoolVar(&verbose, verboseFlag.Long, false, verboseFlag.Description)
	flag.BoolVar(&verbose, verboseFlag.Short, false, verboseFlag.Description)

	c.Flags = append(c.Flags, helpFlag, tokenFlag, verboseFlag)
	c.setGlobalFlagUsage()
}

// flag.Usage is called when there is an error parsing flags.
// So we need to modify the flag.Usage because we need access to the defined
// global flags, and if there's an error parsing those flags then the usage
// output needs to be updated so it can display our custom output.
func (c *Children) setGlobalFlagUsage() {
	inputs := c.helpOutput()
	inputs = append(inputs, GenHelp(c.Cmds))
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), formatting.HelpTemplate, inputs...)
	}
	flag.Parse()
}

func NewChildren() *Children {
	c := &Children{
		Cmds: make(map[string]Command),
	}
	help := Command{
		Children:    c,
		Name:        "help",
		Description: "default help for the current command",
		Exec:        c.Help,
	}
	c.Add(help)
	return c
}

// Flag represents flag meta data.
type Flag struct {
	Long        string
	Short       string
	Description string
}

func Find(cmds map[string]Command, args []string) (Command, error) {
	err := fmt.Errorf("couldn't find command: %s", strings.Join(args, " "))

	for i, arg := range args {
		for name, cmd := range cmds {
			if arg == name {
				if cmd.Children != nil {
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
						return cmd, nil
					}
					return Find(cmd.Children.Cmds, subargs)
				}

				// Ensure a subcommand, with a nil Children, followed by another command
				// doesn't return the currently matched subcommand.
				subargs := args[i+1:]
				for _, arg := range subargs {
					if !strings.HasPrefix(arg, "--") {
						return Command{}, err
					}
				}
				return cmd, nil
			}
		}
	}

	return Command{}, err
}

type Command struct {
	*Children

	Name        string
	Description string
	Exec        Run
	Flags       []Flag
}

// GenHelp generates a string consisting of the given commands/subcommands.
func GenHelp(cmds map[string]Command) string {
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
		pad := formatting.Pad(name, longest)
		line := fmt.Sprintf("  %s"+pad+"  %s\n", name, description)
		b.Write([]byte(line))
	}
	return b.String()
}
