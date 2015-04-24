# pathhelper

`pathhelp` is a helper for golang `path` package

#### Features:

* Get current path of executable is running

        dir, err := pathhelper.GetCurrentExecDir()

* Get file name without extend name

        f := pathhelper.GetFileNameWithoutExt(p)

#### License

MIT License