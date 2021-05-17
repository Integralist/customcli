package build

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
)

func New() commands.Command {
	return commands.Command{
		Name:        "build",
		Description: "Build a Compute@Edge package locally",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("compute build:", args)
}
