package missing

// find the missing but necessary __init__.py
type InitPy struct {
	// the project's base dir
	BaseDir string `help:"the python project's basedir"`

	// search the __init__.py used for unittest
	Unittest        bool   `arg:"-U" default:"true" help:"check the test-case used on unittest"`
	UnittestPattern string `default:"test*" help:"the file pattern"`
}
