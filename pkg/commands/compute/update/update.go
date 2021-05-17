package update

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
)

func New() commands.Command {
	return commands.Command{
		Name:        "update",
		Description: "Update a package on a Fastly Compute@Edge service version",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("compute update:", args)
}
