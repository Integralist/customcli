package main

import (
	"os"

	"github.com/integralist/customcli/pkg/app"
)

func main() {
	app.Run(os.Args)
}

// TODO: check remaining help variations
//
// fastly help backend
// fastly help backend create
//
// NOTE:
//
// 1. help subcommand prints 'verbose' help (same sections but expanded commands)
//   	should support custom footer (e.g. For help on a specific command, try...)
// 2. --json flag appears to be used for help, but what else?
// 		https://github.com/fastly/cli/blob/master/pkg/app/run.go#L736-L743
