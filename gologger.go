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
	logPrinter(logInstance{ "MESSAGE", message})
}

func (log GoLogger) Info(message string) {
	logPrinter(logInstance{ "INFO", message})
}

func (log GoLogger) Warn(message string) {
	logPrinter(logInstance{ "WARN", message})
}

func (log GoLogger) Debug(message string) {
	logPrinter(logInstance{ "DEBUG", message})
}

func (log GoLogger) Error(message string) {
	logPrinter(logInstance{ "EROR", message})
}

func (log GoLogger) ReplaceMessage(message string) {
	logPrinter(logInstance{ "RS", message})
}
