package initialization

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
)

func New() commands.Command {
	return commands.Command{
		Name:        "init",
		Description: "Initialize a new Compute@Edge package locally",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("compute init:", args)
}
