package logger

type GoLogger struct {
	PrinterType string
	Location string
}

type LogInstance struct {
	LogType string
	Message string
	LoggerInit GoLogger
}
