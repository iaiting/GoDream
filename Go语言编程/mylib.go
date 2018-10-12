package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func simple_log(format string) {

	funcptr, file_fullname, line, ok := runtime.Caller(1)

	file_name := filepath.Base(file_fullname)

	if ok {
		func_name := runtime.FuncForPC(funcptr).Name()

		fmt.Printf("[%s][%s][%d] %s\n", file_name, func_name, line, format)
	}
}
