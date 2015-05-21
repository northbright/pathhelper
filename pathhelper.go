package pathhelper

import (
	"github.com/northbright/fnlog"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	logger *log.Logger
)

func GetCurrentExecDir() (dir string, err error) {
	p, err := exec.LookPath(os.Args[0])
	if err != nil {
		logger.Printf("exec.LookPath(%s), err: %s\n", os.Args[0], err)
		return "", err
	}

	absPath, err := filepath.Abs(p)
	if err != nil {
		logger.Printf("filepath.Abs(%s), err: %s\n", p, err)
		return "", err
	}

	dir = filepath.Dir(absPath)
	return dir, nil
}

func GetFileNameWithoutExt(filePath string) (fileNameWithoutExt string) {
	base := path.Base(filePath)
	ext := path.Ext(base)
	return base[0 : len(base)-len(ext)]
}

// Check if a file or dir exists.
func PathFileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func init() {
	// init funlog's logger
	logger = fnlog.New("")
}
