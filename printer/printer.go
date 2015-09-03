package printer
import (
	"github.com/sadlil/gologger/printer/console"
	"github.com/sadlil/gologger/logger"
	"fmt"
)

func Print(log logger.LogInstance, fileName string, lineNumber int) {
	if(log.LoggerInit.PrinterType == "console") {
		console.ConsolePrinter(log, fileName, lineNumber)
	} else if(log.LoggerInit.PrinterType == "file") {
		fmt.Println("File")
	} else if(log.LoggerInit.PrinterType == "es") {
		fmt.Println("ES")
	}
}
