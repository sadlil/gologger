package gologger

type Printer interface {
	String() string
	Format(...interface{}) string
	Formatf(string, ...interface{}) string
	Print(Log) bool
}

