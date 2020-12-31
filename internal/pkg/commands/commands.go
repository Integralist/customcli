package commands

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// List represents all available commands
var List = map[string]string{
	"backend":         "Manipulate Fastly service version backends",
	"compute":         "Manage Compute@Edge packages",
	"configure":       "Configure the Fastly CLI",
	"dictionary":      "Manipulate Fastly edge dictionaries",
	"dictionaryitem":  "Manipulate Fastly edge dictionary items",
	"domain":          "Manipulate Fastly service version domains",
	"healthcheck":     "Manipulate Fastly service version healthchecks",
	"help":            "Show help.",
	"logging":         "Manipulate Fastly service version logging endpoints",
	"service":         "Manipulate Fastly services",
	"service-version": "Manipulate Fastly service versions",
	"stats":           "View statistics (historical and realtime) for a Fastly service",
	"update":          "Update the CLI to the latest version",
	"version":         "Display version information for the Fastly CLI",
	"whoami":          "Get information about the currently authenticated account",
}

// Pad adds spacing so command descriptions in help output align.
func Pad(name string, longest int) string {
	return strings.Repeat(" ", longest-len(name))
}

// Gen generates a string consisting of the given commands/subcommands.
func Gen(list map[string]string) string {
	var b bytes.Buffer

	// the following iteration over list is required for two reasons:
	// 1. sort the map for the sake of consistent output.
	// 2. identify the longest command for more structured output.
	keys := make([]string, 0, len(list))
	longest := 0
	for name := range list {
		keys = append(keys, name)
		if len(name) > longest {
			longest = len(name)
		}
	}
	sort.Strings(keys)

	for _, name := range keys {
		description := list[name]
		pad := Pad(name, longest)
		line := fmt.Sprintf("  %s"+pad+"  %s\n", name, description)
		b.Write([]byte(line))
	}
	return b.String()
}
