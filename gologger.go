package gologger

import "github.com/sadlil/gologger/logger"

const (
	CONSOLE string = "console"
	FILE string = "file"
	ELASTICSEARCH string = "es"
	SimpleLog string = "simple"
	ColoredLog string = "color"
)

type GoLogger struct  {
	Logger logger.GoLogger
}

func GetLogger(selector ...string) GoLogger {
	if len(selector) == 0  {
		return GoLogger{logger.GoLogger{CONSOLE, ColoredLog}}
	}
	return GoLogger{logger.GoLogger{selector[0], selector[1]}}
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
