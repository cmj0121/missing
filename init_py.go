package missing

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
)

var (
	DEFAULT_PATTERN = []string{
		"test*.py",
	}
)

// find the missing but necessary __init__.py
type InitPy struct {
	// the project's base dir
	BaseDir string `arg:"-b" default:"." help:"the python project's basedir"`

	// keep check the root folder
	KeepRoot bool `arg:"-r,--keep-root" help:"keep root __init__.py"`

	// search the __init__.py used for unittest
	Pattern []string `arg:"-p" help:"the file regexp pattern"`

	// the exclude pattern which exactly match the pattern
	Exclude []string `arg:"-e" help:"exclude the folder that match the search path"`
}

// find the missing __init__.py
func (cmd *InitPy) Execute() (ok bool, err error) {
	var res []*regexp.Regexp

	if len(cmd.Pattern) == 0 {
		// append the default pattern
		cmd.Pattern = append(cmd.Pattern, DEFAULT_PATTERN...)
	}

	for _, pattern := range cmd.Pattern {
		// save the file pattern
		res = append(res, regexp.MustCompile(pattern))
	}

	// sort the exclude pattern
	sort.Strings(cmd.Exclude)

	ok, err = cmd.search_missing_init_py(cmd.BaseDir, res)
	return
}

// find the missing based on the basedir and recursive if path is folder
func (cmd *InitPy) search_missing_init_py(base string, res []*regexp.Regexp) (missing bool, err error) {
	var files []fs.FileInfo

	idx := sort.SearchStrings(cmd.Exclude, base)
	if idx >= 0 && idx < len(cmd.Exclude) {
		if path.Clean(base) == path.Clean(cmd.Exclude[idx]) {
			// found in the exclude path, skip
			return
		}
	}

	files, err = ioutil.ReadDir(base)
	if err != nil {
		// cannot open dir
		err = fmt.Errorf("cannot ReadDir on %#v: %v", base, err)
		return
	}

	for _, file := range files {
		switch {
		case file.IsDir():
			var sub_missing bool
			subdir := fmt.Sprintf("%v/%v", base, file.Name())

			sub_missing, err = cmd.search_missing_init_py(subdir, res)
			if err != nil {
				// cannot open sub-dir
				return
			}

			if sub_missing {
				// found the missing __init__.py
				cmd.check_init_py_in_path(base)
				missing = true
			}
		default:
			name := file.Name()
			for _, re := range res {
				if re.MatchString(name) {
					if cmd.check_init_py_in_path(base) {
						// found the missing __init__.py
						missing = true
					}
				}
			}
		}
	}

	return
}

func (cmd *InitPy) check_init_py_in_path(base string) (missing bool) {
	if base == cmd.BaseDir && !cmd.KeepRoot {
		// need not process the root folder
		return
	}

	init_py_path := fmt.Sprintf("%v/__init__.py", base)
	if _, err := os.Stat(init_py_path); errors.Is(err, os.ErrNotExist) {
		missing = true

		fmt.Printf("create %v\n", init_py_path)
		os.Create(init_py_path) // nolint
	}

	return
}
