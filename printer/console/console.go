package console
import "fmt"

func ConsolePrinter(logType string, message string, fileName string, lineNumber int) {
	color := getColor(logType)
	color.Set()
	fmt.Printf("[%s] [%s] [%d] %s\n", logType, fileName, lineNumber, message)
	Unset()
}


func getColor(logType string) *Color {
	var color *Color
	switch logType {
	case "LOG" :
		color = New(Reset)
		break
	case "MSG" :
		color = New(Reset)
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
