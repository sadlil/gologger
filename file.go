package gologger

import (
	"fmt"
	"os"
	"path"
	"time"
)

func FilePrinter(log LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	logFileName := log.LoggerInit.Location
	if logFileName == "" {
		logFileName = "log.txt"
	}
	basePath := path.Dir(logFileName)
	filePath := path.Base(logFileName)

	if os.MkdirAll(basePath, 0777) != nil {
		fmt.Println("Unable to create directory")
		os.Exit(1)
	}

	logFileName = path.Join(basePath, filePath)
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logString := fmt.Sprintf("[%s] [%s] [%s::%s::%s] [%d] %s\n", log.LogType, time.Format("2006-01-02 15:04:05"), packageName, fileName, funcName, lineNumber, log.Message)
	_, fileWriteErr := file.WriteString(logString)
	if fileWriteErr != nil {
		fmt.Println(fileWriteErr)
		os.Exit(1)
	}
}
