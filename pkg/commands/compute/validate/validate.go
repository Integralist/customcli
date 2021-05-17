package validate

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
)

func New() commands.Command {
	return commands.Command{
		Children:    commands.NewChildren(), // defines `help`
		Name:        "validate",
		Description: "Validate a Compute@Edge package",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("compute validate", args)
}
