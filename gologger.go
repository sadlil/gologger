package gologger

type Logger struct {
	log *Log
	printer Printer
	logType logType
	configs string
}


func Get(name string, configs ...interface{}) *Logger {
	if v, ok := activeLogger[name]; ok {
		return v
	}
	if len(configs) > 0 {
		return New(name, configs[0].(LoggerType), configs[1:])
	}
	return New(name, Console)
}

func New(name string, t LoggerType, config ...interface{}) *Logger {
	return &Logger{
		log: &Log{},

	}
}

func NewCustom(name string, p Printer) *Logger {
	return &Logger{}
}

func (l *Logger) C(c string) *Logger {
	l.log.category = c
	return l
}

func (l *Logger) L(level int) *Logger {
	l.log.level = level
	return l
}

func (l *Logger) Info() {

}
