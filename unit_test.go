package pathhelper

import (
	"testing"
)

func TestGetCurrentExecDir(t *testing.T) {
	logger.Printf("Testing GetCurrentExecDir()...\n")
	dir, err := GetCurrentExecDir()
	if err != nil {
		t.Errorf("GetCurrentExecDir() err: %s\n", err)
		return
	}
	logger.Printf("dir = %s\n", dir)
}

func TestGetFileNameWithoutExt(t *testing.T) {
	p := "/a/b/c.apk"
	logger.Printf("Testing GetFileNameWithoutExt(%s)...\n", p)
	f := GetFileNameWithoutExt(p)
	logger.Printf(" = %s\n", f)
}

func TestPathFileExist(t *testing.T) {
	arr := []string{"/usr/", "~/go", "./", "~/xxx"}

	for _, v := range arr {
		b := PathFileExist(v)
		logger.Printf("PathFileExist(%v) = %v", v, b)
	}
}
