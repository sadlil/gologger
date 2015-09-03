package gologger

import (
	"runtime"
	"path/filepath"
	"strings"
	"github.com/sadlil/gologger/printer"
	"github.com/sadlil/gologger/logger"
)

func logPrinter( log logger.LogInstance ) {
	_, fileName, lineNumber, _ := runtime.Caller(2)
	logPrint(log, fileName, lineNumber)
}

func logPrint(log logger.LogInstance, fileName string, lineNumber int) {
	fileName = parseFileName(fileName)
	printer.Print(log, fileName, lineNumber)
}

func parseFileName(fileName string) string {
	fileName = filepath.ToSlash(fileName)
	lastIndex := strings.LastIndex(fileName, "/")
	fileName = fileName[lastIndex +1 : len(fileName)]
	return fileName
}
