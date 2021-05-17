package formatting

import (
	"strings"

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

// Pad adds spacing so command descriptions in help output align.
func Pad(name string, longest int) string {
	return strings.Repeat(" ", longest-len(name))
}
