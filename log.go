package gologger

import "time"

type logType string

const (
	INF logType = "INF"
	LOG         = "LOG"
	DBG         = "DBG"
	STK			= "STK"
	ERR         = "ERR"
	WRN         = "WRN"
	CTL         = "CTL"
	FAT         = "FAT"
	PAN         = "PAN"
)

func (l logType) String() string {
	return string(l)
}

type Log struct {
	level    int
	category string

	caller  *caller
	logType logType
	message string
}

func (l *Log) Level() int {
	return l.level
}

func (l *Log) Category() string {
	return l.category
}

func (l *Log) Caller() *caller {
	return l.caller
}

func (l *Log) Type() logType {
	return l.logType
}

func (l *Log) Message() string {
	return l.message
}

type caller struct {
	line     int32
	method   string
	fileName string
	pkg      string
	pkgPath  string

	timestamp time.Time
}

func (c *caller) Line() int32 {
	return c.line
}


func (c *caller) Method() string {
	return c.method
}

func (c *caller) File() string {
	return c.fileName
}

func (c *caller) Package() string {
	return c.pkg
}

func (c *caller) PackagePath() string {
	return c.pkgPath
}

func (c *caller) TimeStamp() time.Time {
	return c.timestamp
}