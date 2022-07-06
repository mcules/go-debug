package go_debug

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type Color struct {
	Log      string
	Notice   string
	Warning  string
	Error    string
	Fatal    string
	Panic    string
	Reset    string
	Function string
}

var colors Color

// init the debug package
func init() {
	colors.Reset = "\033[0m"
	colors.Function = "\033[35m"
	colors.Log = "\033[37m"
	colors.Warning = "\033[33m"
	colors.Error = "\033[91m"
	colors.Fatal = "\033[31m"
	colors.Panic = "\033[45m"
	colors.Notice = "\033[96m"
}

// Log writes message to console
func Log(message ...interface{}) {
	writeToConsole("Log", message, colors.Log)
}

// Notice writes message to console
func Notice(message ...interface{}) {
	writeToConsole("Notice", message, colors.Notice)
}

// Warning writes message to console
func Warning(message ...interface{}) {
	writeToConsole("Warning", message, colors.Warning)
}

// Error writes message to console
func Error(message ...interface{}) {
	writeToConsole("Error", message, colors.Error)
}

// Fatal writes message to console
func Fatal(message ...interface{}) {
	log.Fatal(colors.Fatal, colors.Function, logFormat(), colors.Fatal, message, colors.Reset)
}

// Panic writes message to console
func Panic(message ...interface{}) {
	log.Panic(colors.Panic, colors.Function, logFormat(), colors.Panic, message, colors.Reset)
}

// writeToConsole writes message to console
func writeToConsole(logType string, message interface{}, color string) {
	log.Println(color, logType, colors.Function, logFormat(), color, fmt.Sprint(message), colors.Reset)
}

// logFormat defines the log format
func logFormat() string {
	counter, _, _, success := runtime.Caller(3)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}

	packagename := runtime.FuncForPC(counter).Name()
	lastIndex := strings.LastIndex(packagename, "/")

	return packagename[lastIndex+1:]
}
