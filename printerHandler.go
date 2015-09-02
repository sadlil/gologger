package gologger

import (
	"runtime"
	"fmt"
	"path/filepath"
	"strings"
)


type logInstance struct {
	message string
	logType string
}


func logPrinter( log logInstance ) {
	_, fileName, lineNumber, _ := runtime.Caller(2)
	logPrint(log, fileName, lineNumber)
}

func logPrint(log logInstance, fileName string, lineNumber int) {
	fileName = parseFileName(fileName)
	fmt.Printf("[%s] [%s] [%d] : %s\n", log.logType, fileName, lineNumber, log.message)
}

func parseFileName(fileName string) string {
	fileName = filepath.ToSlash(fileName)
	lastIndex := strings.LastIndex(fileName, "/")
	fileName = fileName[lastIndex +1 : len(fileName)-1]
	return fileName
}
