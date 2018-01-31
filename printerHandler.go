package gologger

import (
	"path"
	"runtime"
	"strings"
	"time"
	"os"
)

func logPrinter(log LogInstance) {
	info := retrieveCallInfo()
	timer := time.Now()
	logPrint(log, info, timer)
}

func logPrint(log LogInstance, info *callerInfo, time time.Time) {
	Print(log, info.packageName, info.fileName, info.line, info.funcName, time)
	if log.LogType == "CRT" {
		os.Exit(1)
	}
}

type callerInfo struct {
	packageName string
	fileName    string
	funcName    string
	line        int
}

func retrieveCallInfo() *callerInfo {
	pc, file, line, _ := runtime.Caller(3)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &callerInfo{
		packageName: packageName,
		fileName:    fileName,
		funcName:    funcName,
		line:        line,
	}
}
