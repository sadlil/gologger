package gologger

const (
	CONSOLE       string = "console"
	FILE          string = "file"
	ELASTICSEARCH string = "es"
	SimpleLog     string = "simple"
	ColoredLog    string = "color"
)

type GoLogger struct {
	Logger
}

func GetLogger(selector ...string) GoLogger {
	if len(selector) == 0 {
		return GoLogger{Logger{CONSOLE, ColoredLog}}
	}
	return GoLogger{Logger{selector[0], selector[1]}}
}

func (log GoLogger) Log(message string) {
	logPrinter(LogInstance{LogType: "LOG", Message: message, LoggerInit: log.Logger})
}

func (log GoLogger) Message(message string) {
	logPrinter(LogInstance{LogType: "MSG", Message: message, LoggerInit: log.Logger})
}

func (log GoLogger) Info(message string) {
	logPrinter(LogInstance{LogType: "INF", Message: message, LoggerInit: log.Logger})
}

func (log GoLogger) Warn(message string) {
	logPrinter(LogInstance{LogType: "WRN", Message: message, LoggerInit: log.Logger})
}

func (log GoLogger) Debug(message string) {
	logPrinter(LogInstance{LogType: "DBG", Message: message, LoggerInit: log.Logger})
}

func (log GoLogger) Error(message string) {
	logPrinter(LogInstance{LogType: "ERR", Message: message, LoggerInit: log.Logger})
}

func (log GoLogger) Fatal(message string) {
	logPrinter(LogInstance{LogType: "CRT", Message: message, LoggerInit: log.Logger})
}

func (log GoLogger) ReplaceMessage(message string) {
	logPrinter(LogInstance{LogType: "RSS", Message: message, LoggerInit: log.Logger})
}
