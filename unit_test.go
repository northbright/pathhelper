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
