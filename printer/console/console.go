package console

import (
	"fmt"
	"github.com/sadlil/gologger/logger"
	"time"
)

func ConsolePrinter(log logger.LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	color := getColor(log)
	color.Set()
	fmt.Printf("[%s] [%s] [%s::%s::%s] [%d] %s\n", log.LogType, time.Format("2006-01-02 15:04:05"), packageName, fileName, funcName, lineNumber, log.Message)
	Unset()
}


func getColor(log logger.LogInstance) *Color {
	var color *Color

	if(log.LoggerInit.Location == "simple") {
		color = New(Reset)
		return color
	}

	switch log.LogType {
	case "LOG" :
		color = New(Reset)
		break
	case "MSG" :
		color = New(FgBlue)
		break
	case "INF" :
		color = New(FgGreen)
		break
	case "WRN" :
		color = New(FgMagenta)
		break
	case "DBG" :
		color = New(FgYellow)
		break
	case "ERR" :
		color = New(FgRed)
		break
	case "RSS" :
		color = New(Reset)
		break
	default:
		color = New(Reset)
		break
	}
	return color
}
