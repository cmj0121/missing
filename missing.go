package missing

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

// the CLI interface
type Missing struct {
	*InitPy `arg:"subcommand:init-py"`
}

// create the missing with the default settings
func New() *Missing {
	return &Missing{}
}

// run on the command-line mode with passed argument
func (missing *Missing) Run() {
	arg.MustParse(missing)
}

// the version info
func (missing Missing) Version() (ver string) {
	ver = fmt.Sprintf("missing (v%d.%d.%d)", MAJOR, MINOR, MACRO)
	return
}
