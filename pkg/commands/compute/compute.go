package compute

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
	"github.com/integralist/customcli/pkg/commands/compute/build"
	"github.com/integralist/customcli/pkg/commands/compute/deploy"
	"github.com/integralist/customcli/pkg/commands/compute/initialization"
	"github.com/integralist/customcli/pkg/commands/compute/update"
	"github.com/integralist/customcli/pkg/commands/compute/validate"
)

// Cmds represents all available subcommands.
var Cmds = map[string]string{
	"init":     "Initialize a new Compute@Edge package locally",
	"build":    "Build a Compute@Edge package locally",
	"deploy":   "Deploy a package to a Fastly Compute@Edge service",
	"update":   "Update a package on a Fastly Compute@Edge service version",
	"validate": "Validate a Compute@Edge package",
}

func New() commands.Command {
	cmds := []commands.Command{
		build.New(),
		deploy.New(),
		initialization.New(),
		update.New(),
		validate.New(),
	}

	b := commands.NewBase()
	b.Add(cmds...)

	return commands.Command{
		Base:        b,
		Name:        "compute",
		Description: "Manage Compute@Edge packages",
		Exec:        run,
	}
}

func run(args []string) {
	fmt.Println("compute:", args)
}
