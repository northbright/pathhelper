// Package pathhelper provides path related helper functions.
package pathhelper

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	DEBUG bool = false // Set this flag to true to output debug messages from this package.
)

// GetCurrentExecDir() gets the current executable path.
func GetCurrentExecDir() (dir string, err error) {
	p, err := exec.LookPath(os.Args[0])
	if err != nil {
		if DEBUG {
			fmt.Printf("exec.LookPath(%s), err: %s\n", os.Args[0], err)
		}
		return "", err
	}

	absPath, err := filepath.Abs(p)
	if err != nil {
		if DEBUG {
			fmt.Printf("filepath.Abs(%s), err: %s\n", p, err)
		}
		return "", err
	}

	dir = filepath.Dir(absPath)
	return dir, nil
}

// GetFileNameWithoutExt() gets the file name without extended name of given file path.
func GetFileNameWithoutExt(filePath string) (fileNameWithoutExt string) {
	base := path.Base(filePath)
	ext := path.Ext(base)
	return base[0 : len(base)-len(ext)]
}

// PathFileExist checks if a file or dir exists.
func PathFileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
