package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/customcli/internal/pkg/arguments"
	"github.com/integralist/customcli/internal/pkg/commands"
	"github.com/integralist/customcli/internal/pkg/commands/backend"
	"github.com/integralist/customcli/internal/pkg/commands/compute"
	"github.com/integralist/customcli/internal/pkg/commands/configure"
	"github.com/integralist/customcli/internal/pkg/commands/dictionary"
	"github.com/integralist/customcli/internal/pkg/commands/dictionaryitem"
	"github.com/integralist/customcli/internal/pkg/commands/domain"
	"github.com/integralist/customcli/internal/pkg/commands/healthcheck"
	"github.com/integralist/customcli/internal/pkg/commands/help"
	"github.com/integralist/customcli/internal/pkg/commands/logging"
	"github.com/integralist/customcli/internal/pkg/commands/service"
	"github.com/integralist/customcli/internal/pkg/commands/serviceversion"
	"github.com/integralist/customcli/internal/pkg/commands/stats"
	"github.com/integralist/customcli/internal/pkg/commands/update"
	"github.com/integralist/customcli/internal/pkg/commands/version"
	"github.com/integralist/customcli/internal/pkg/commands/whoami"
	"github.com/integralist/customcli/internal/pkg/flags"
)

// Run bootstraps the CLI.
//
// - Parses the global flags.
// - Identifies which subcommand should execute.
func Run(args []string) {
	globals := flags.ParseGlobal(args)
	fmt.Printf("globals: %+v\n", globals)

	cmd, err := arguments.IdentifyCommand(commands.List, flag.Args())
	if err != nil {
		flag.Usage()
		flags.ErrorContext(err)
		os.Exit(1)
	}

	// TODO: try and avoid a switch statement.
	//
	// if this was an inheritance chain of Command objects could we link-list
	// through the objects to find a match?
	switch cmd {
	case "backend":
		backend.Run(args)
	case "compute":
		compute.Run(args)
	case "configure":
		configure.Run(args)
	case "dictionary":
		dictionary.Run(args)
	case "dictionaryitem":
		dictionaryitem.Run(args)
	case "domain":
		domain.Run(args)
	case "healthcheck":
		healthcheck.Run(args)
	case "help":
		help.Run(args)
	case "logging":
		logging.Run(args)
	case "service":
		service.Run(args)
	case "service-version":
		serviceversion.Run(args)
	case "stats":
		stats.Run(args)
	case "update":
		update.Run(args)
	case "version":
		version.Run(args)
	case "whoami":
		whoami.Run(args)
	}
}
