package printer
import (
	"github.com/sadlil/gologger/printer/console"
	"github.com/sadlil/gologger/logger"
	"fmt"
	"time"
	"github.com/sadlil/gologger/printer/file"
)

func Print(log logger.LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	if(log.LoggerInit.PrinterType == "console") {
		console.ConsolePrinter(log, packageName, fileName, lineNumber, funcName, time)
	} else if(log.LoggerInit.PrinterType == "file") {
		file.FilePrinter(log, packageName, fileName, lineNumber, funcName, time)
	} else if(log.LoggerInit.PrinterType == "es") {
		fmt.Println("ES")
	}
}
