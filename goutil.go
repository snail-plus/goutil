// Package goutil Useful utils for the Go: int, string, array/slice, map, format, cli, env, filesystem, testing and more.
package goutil

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/snail-plus/goutil/stdutil"
)

// PanicIfErr if error is not empty
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Panicf format panic message use fmt.Sprintf
func Panicf(format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}

// FuncName get func name
func FuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// PkgName get current package name
//
// Usage:
//	funcName := goutil.FuncName(fn)
//	pgkName := goutil.PkgName(funcName)
func PkgName(funcName string) string {
	for {
		lastPeriod := strings.LastIndex(funcName, ".")
		lastSlash := strings.LastIndex(funcName, "/")
		if lastPeriod > lastSlash {
			funcName = funcName[:lastPeriod]
		} else {
			break
		}
	}

	return funcName
}

// GetCallStacks stacks is a wrapper for runtime.
// Stack that attempts to recover the data for all goroutines.
// from glog package
func GetCallStacks(all bool) []byte {
	return stdutil.GetCallStacks(all)
}

// GetCallersInfo returns an array of strings containing the file and line number
// of each stack frame leading
func GetCallersInfo(skip, max int) (callers []string) {
	return stdutil.GetCallersInfo(skip, max)
}
