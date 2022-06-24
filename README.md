# pathhelper

[![Build Status](https://travis-ci.org/northbright/pathhelper.svg?branch=master)](https://travis-ci.org/northbright/pathhelper)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/pathhelper)](https://goreportcard.com/report/github.com/northbright/pathhelper)
[![GoDoc](https://godoc.org/github.com/northbright/pathhelper?status.svg)](https://godoc.org/github.com/northbright/pathhelper)

Update: this repo is no longer maintained. Use [pathelper](https://github.com/northbright/pathelper) instead.

pathhelper is a [Golang](http://golang.org) package which provides path related helper functions.

#### Get absolute path of current executalbe directory.

    dir, err := pathhelper.GetCurrentExecDir()
    if err != nil {
        return
    }

#### Get absolute path of given relative path.  
It joins absolute path of current executable directory and given relative path into a single absolute path.

    p := "./config"
    absPath := ""
    if absPath, err = pathhelper.GetAbsPath(p); err != nil {
        fmt.Fprintf(os.Stderr, "GetAbsPath(%v) error: %v\n", p, err)
    }

#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/pathhelper)

#### License
* [MIT License](./LICENSE)
