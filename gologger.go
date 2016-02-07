package gologger

import (
	"runtime"
	"path"
	"time"
)

// LoggerType is used to indicate the the logger type.
// Where or How to Log.
type LoggerType string

const (
	// Glog will use "github.com/golang/glog" to log this will use the leveled
	// handled log for logging like glog.V().Infoln().
	// This will not apply the category filter provided by gologger. By Providing
	// `gologger.EnableC` as options, can apply the category filter before log
	// to glog.
	Glog LoggerType = "glog"

	// Console Logger. Simply Will use the fmt package. Two modes available for
	// Logging. Colored Log or Plain text log to Console. Provides Options as
	// `gologger.Plain` and `gologger.Colored`
	Console = "console"

	// All logs to filesystem. The file path as options Needs to be provided.
	// if path is not provided logs to the default package location.
	File = "file"

	// Stores logs in elasticsearch in logstash format. Kibana is able
	// to use those logs. The elasticsearch location needs to be set via options.
	// If no options is set then will use `localhost:9200` to communicate with
	// elasticsearch.
	ElasticSearch = "elasticsearch"

	// Custom Type. Defined by user.
	Custom = "custom"
)

const (
	EnableC = iota
	Plain
	Colored
)

type Logger struct {
	log     *Log
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
		log:     &Log{},
		printer: newConsolePrinter(false),
	}
}

func NewCustom(name string, p Printer) *Logger {
	return &Logger{
		log:     &Log{},
		printer: p,
		logType: Custom,
	}
}

func (l *Logger) C(c string) *Logger {
	l.log.category = c
	return l
}

func (l *Logger) L(level int) *Logger {
	l.log.level = level
	return l
}

func (l *Logger) Info(args ...interface{}) {
	m := l.printer.Format(args...)
	l.log.message = m
	l.callerInfo()
	l.printer.Print(l.log)
}

func (l *Logger) reset() {
	l.log = &Log{}
}

func (l *Logger) process() {

}

func (l *Logger) callerInfo() {
	pc, file, line, _ := runtime.Caller(2)
	path.Split(file)
	pkg := runtime.FuncForPC(pc).Name()

	l.log.caller = &caller{
		line: line,
		method: stripPackage(pkg),
		fileName: getFileName(file),
		pkg: stripFileName(pkg),
		pkgPath: stripGOPATH(file),
		timestamp: time.Now(),
	}
}
