package gologger

import "github.com/sadlil/gologger/logger"

const (
	CONSOLE string = "console"
	FILE string = "file"
	ELASTICSEARCH string = "es"
	ConsoleAndFile = "cf"
	SimpleLog string = "simple"
	ColoredLog string = "color"
)

type GoLogger struct  {
	Logger logger.GoLogger
}

func GetLogger(printerType string, location string) GoLogger {
	return GoLogger{logger.GoLogger{printerType, location}}
}

func (log GoLogger) Log(message string) {
	logPrinter(logger.LogInstance{ "LOG", message, log.Logger})
}

func (log GoLogger) Message(message string) {
	logPrinter(logger.LogInstance{ "MSG", message, log.Logger})
}

func (log GoLogger) Info(message string) {
	logPrinter(logger.LogInstance{ "INF", message, log.Logger})
}

func (log GoLogger) Warn(message string) {
	logPrinter(logger.LogInstance{ "WRN", message, log.Logger})
}

func (log GoLogger) Debug(message string) {
	logPrinter(logger.LogInstance{ "DBG", message, log.Logger})
}

func (log GoLogger) Error(message string) {
	logPrinter(logger.LogInstance{ "ERR", message, log.Logger})
}

func (log GoLogger) ReplaceMessage(message string) {
	logPrinter(logger.LogInstance{ "RSS", message, log.Logger})
}
