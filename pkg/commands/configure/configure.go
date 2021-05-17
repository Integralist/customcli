package configure

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
)

func New() commands.Command {
	return commands.Command{
		Children:    commands.NewChildren(), // defines `help`
		Name:        "configure",
		Description: "Configure the Fastly CLI",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("configure:", args)
}
