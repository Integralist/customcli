package formatting

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

var Bold = color.New(color.Bold).SprintFunc()
var RedBold = color.New(color.Bold, color.FgRed).SprintFunc()

// TODO: we handle COMMANDS already but SUBCOMMANDS is used when using
// fastly help <command>. We fix this already with `-help <command>` (see
// `category.Usage()` called in backend.go)
var HelpTemplate = `%s
  %s

%s

%s
%s
%s
%s
`

// Template prints to an io.Writer an interpolated template.
func Template(w io.Writer, template string, input ...interface{}) {
	fmt.Fprintf(w, template, input...)
}
