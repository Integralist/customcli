package deploy

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
)

func New() commands.Command {
	return commands.Command{
		Name:        "deploy",
		Description: "Deploy a package to a Fastly Compute@Edge service",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("compute deploy:", args)
}
