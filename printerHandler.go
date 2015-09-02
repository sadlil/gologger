package gologger

import (
	"runtime"
	"path/filepath"
	"strings"
	"github.com/sadlil/gologger/printer"
)


type logInstance struct {
	logType string
	message string
}


func logPrinter( log logInstance ) {
	_, fileName, lineNumber, _ := runtime.Caller(2)
	logPrint(log, fileName, lineNumber)
}

func logPrint(log logInstance, fileName string, lineNumber int) {
	fileName = parseFileName(fileName)
	printer.Print(log.logType, log.message, fileName, lineNumber)
}

func parseFileName(fileName string) string {
	fileName = filepath.ToSlash(fileName)
	lastIndex := strings.LastIndex(fileName, "/")
	fileName = fileName[lastIndex +1 : len(fileName)]
	return fileName
}


