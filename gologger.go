package gologger

type GoLogger struct {
}

func GetLogger() GoLogger {
	return GoLogger{}
}

func (log GoLogger) Log(message string) {
	logPrinter(logInstance{ "LOG", message})
}

func (log GoLogger) Message(message string) {
	logPrinter(logInstance{ "MSG", message})
}

func (log GoLogger) Info(message string) {
	logPrinter(logInstance{ "INF", message})
}

func (log GoLogger) Warn(message string) {
	logPrinter(logInstance{ "WRN", message})
}

func (log GoLogger) Debug(message string) {
	logPrinter(logInstance{ "DBG", message})
}

func (log GoLogger) Error(message string) {
	logPrinter(logInstance{ "ERR", message})
}

func (log GoLogger) ReplaceMessage(message string) {
	logPrinter(logInstance{ "RSS", message})
}
