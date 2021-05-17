package backend

import (
	"fmt"

	"github.com/integralist/customcli/pkg/commands"
	"github.com/integralist/customcli/pkg/commands/backend/create"
	"github.com/integralist/customcli/pkg/commands/backend/delete"
	"github.com/integralist/customcli/pkg/commands/backend/describe"
	"github.com/integralist/customcli/pkg/commands/backend/list"
	"github.com/integralist/customcli/pkg/commands/backend/update"
)

// TODO: ensure all types have comments
func New() commands.Command {
	cmds := []commands.Command{
		create.New(),
		delete.New(),
		describe.New(),
		list.New(),
		update.New(),
	}

	c := commands.NewChildren()
	c.Add(cmds...)

	cmd := commands.Command{
		Children:    c,
		Name:        "backend",
		Description: "Manipulate Fastly service version backends",
		Exec:        run,
	}

	return cmd
}

func run(args []string) {
	fmt.Printf("run backend %+v", args)
}
