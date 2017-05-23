package printer
import (
	"github.com/sadlil/gologger/printer/console"
	"github.com/sadlil/gologger/logger"
	"time"
	"github.com/sadlil/gologger/printer/file"
	"github.com/sadlil/gologger/printer/es"
)

func Print(log logger.LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	if(log.LoggerInit.PrinterType == "console") {
		console.ConsolePrinter(log, packageName, fileName, lineNumber, funcName, time)
	} else if(log.LoggerInit.PrinterType == "file") {
		file.FilePrinter(log, packageName, fileName, lineNumber, funcName, time)
	} else if(log.LoggerInit.PrinterType == "es") {
		es.ESPrinter(log, packageName, fileName, lineNumber, funcName, time)
	}
}
