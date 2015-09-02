package gologger
import "fmt"

type GoLogger struct {
}


func getLogger() GoLogger {
	return GoLogger{}
}



func (log GoLogger) Log(message string) {
	fmt.Println("[LOG]", message)
}


func (log GoLogger) Message(message string) {
	fmt.Println("[MSG]", message)
}


func (log GoLogger) Warn(message string) {
	fmt.Println("[WRN]", message)
}


func (log GoLogger) Debug(message string) {
	fmt.Println("[DBG]", message)
}

func (log GoLogger) Error(message string) {
	fmt.Println("[ERR]", message)
}



func (log GoLogger) ReplaceMessage(message string) {
	fmt.Println("[RM]", message)
}
