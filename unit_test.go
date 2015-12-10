package pathhelper_test

import (
	"fmt"
	"github.com/northbright/pathhelper"
	"os"
)

func ExampleGetCurrentExecDir() {
	dir, err := pathhelper.GetCurrentExecDir()
	if err != nil {
		return
	}

	// temp test folder will be changed each time, output to os.Stderr.
	// Ex: /tmp/go-build662839082/github.com/northbright/pathhelper/_test
	fmt.Fprintf(os.Stderr, "current dir: %v\n", dir)
	// Output:
}

func ExampleGetAbsPath() {
	// Relative path will be joined to a absolute path with current test executable in under /tmp.
	// Ex: /tmp/go-build662839082/github.com/northbright/pathhelper/_test/config.json
	pathArr := []string{"/", "/var/log/boot.log", "config.json", "./example/test.go"}

	for _, v := range pathArr {
		if p, err := pathhelper.GetAbsPath(v); err != nil {
			fmt.Fprintf(os.Stderr, "GetAbsPath(%v) error: %v\n", v, err)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", p)
		}
	}
	// Output:
}

func ExampleRelativePathsToAbsPaths() {
	// Init absolute path map:
	// Join all relative paths with current executable dir.
	m := map[string]string{
		"uploadDir": "./uploads",
		"photoDir":  "./photos",
	}

	if err := pathhelper.RelativePathsToAbsPaths(m); err != nil {
		fmt.Fprintf(os.Stderr, "RelativePathsToAbsPaths() error: %v\n", err)
	}

	fmt.Fprintf(os.Stderr, "RelativePathsToAbsPaths():\n")
	for k, v := range m {
		fmt.Fprintf(os.Stderr, "k: %v, v: %v\n", k, v)
	}
	// Output:
}

func ExampleGetFileNameWithoutExt() {
	p := "/a/b/c.apk"
	f := pathhelper.GetFileNameWithoutExt(p)
	fmt.Println(f)
	// Output:
	// c
}

func ExamplePathFileExist() {
	arr := []string{"/usr/", "~/go", "./", "~/xxx"}

	for _, v := range arr {
		b := pathhelper.PathFileExist(v)
		fmt.Printf("%v: %v\n", v, b)
	}

	// Output:
	// /usr/: true
	// ~/go: false
	// ./: true
	// ~/xxx: false
}
