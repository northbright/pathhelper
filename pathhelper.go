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
	// DEBUG is debug mode.
	// Set this flag to true to output debug messages from this package.
	DEBUG = false
)

// GetCurrentExecDir gets the current executable path.
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

// GetAbsPath returns absolute path of input path. It'll join the directory of current executable and input path if it's relative.
func GetAbsPath(p string) (absPath string, err error) {
	if filepath.IsAbs(p) {
		return p, nil
	}

	dir := ""
	if dir, err = GetCurrentExecDir(); err != nil {
		return "", err
	}

	return path.Join(dir, p), nil
}

// GetAbsPaths returns absolute paths of input paths.
//
//   Params:
//     pathMap: key: path name, value: path.
//       It'll join the directory of current executable and input path if it's relative.
//       Ex: m := map[string]string{"uploadDir": "./uploads", "photoDir": "./photos",}
//   Return:
//     absPathMap: key: path name, value: absolute path.
func GetAbsPaths(pathMap map[string]string) (absPathMap map[string]string, err error) {
	absPathMap = map[string]string{}
	for k, v := range pathMap {
		if absPathMap[k], err = GetAbsPath(v); err != nil {
			return absPathMap, err
		}
	}
	return absPathMap, nil
}

// CreateDirs creates directories with given paths and permission.
//
//   Params:
//     dirMap: key: dir name, value: dir path.
//       It'll join the directory of current executable and input path if it's relative.
//       Ex: m := map[string]string{"uploadDir": "./uploads", "photoDir": "/tmp/photos",}
//     perm: permission bits.
func CreateDirs(dirMap map[string]string, perm os.FileMode) (err error) {
	for _, dir := range dirMap {
		absDir := ""
		if absDir, err = GetAbsPath(dir); err != nil {
			return err
		}

		if err = os.MkdirAll(absDir, perm); err != nil {
			return err
		}
	}
	return nil
}

// GetFileNameWithoutExt gets the file name without extended name of given file path.
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
