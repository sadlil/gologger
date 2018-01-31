package gologger

type Logger struct {
	PrinterType string
	Location    string
}

type LogInstance struct {
	LogType    string
	Message    string
	LoggerInit Logger
}
