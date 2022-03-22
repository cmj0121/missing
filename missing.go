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
func (missing *Missing) Run() (exit int) {
	arg.MustParse(missing)

	found_missing, err := missing.InitPy.Execute()
	switch {
	case err != nil:
		fmt.Println(err)
		exit = 1
	case found_missing:
		exit = 2
	}

	return
}

// the version info
func (missing Missing) Version() (ver string) {
	ver = fmt.Sprintf("missing (v%d.%d.%d)", MAJOR, MINOR, MACRO)
	return
}
