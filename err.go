package errs

import (
	"log"
	"runtime"
	"strings"
)

var Separator = "/"

//Log will log the file path and line number. It will split the path after global Separator like "gihub.com". Change the separator using errs.Separator
func Log(err error) {
	if err == nil {
		return
	}
	_, fl, line, _ := runtime.Caller(1)

	file := strings.SplitAfterN(fl, Separator, 2)
	var f *string
	if len(file) > 1 {
		f = &file[1]
	} else {
		f = &file[0]
	}
	log.Printf("error at file: %v line: %v => %v\n", *f, line, err)
}
