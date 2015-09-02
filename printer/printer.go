package printer
import "github.com/sadlil/gologger/printer/console"

func Print(logType string, message string, fileName string, lineNumber int) {
	console.ConsolePrinter(logType, message, fileName, lineNumber)
}