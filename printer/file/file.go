package file
import (
	"github.com/sadlil/gologger/logger"
	"time"
	"os"
	"fmt"
)

func FilePrinter(log logger.LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	logFileName := log.LoggerInit.Location
	if(logFileName == "") {
		logFileName = "log.txt"
	}
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Creating File")
		os.Create(logFileName)
	}
	defer file.Close()
	logString := fmt.Sprintf("[%s] [%s] [%s::%s::%s] [%d] %s\n", log.LogType, time.Format("2006-01-02 15:04:05"), packageName, fileName, funcName, lineNumber, log.Message)
	_, fileWriteErr := file.WriteString(logString)
	if(fileWriteErr != nil) {
		panic(fileWriteErr)
	}
}
