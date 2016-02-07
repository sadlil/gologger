package gologger
import "fmt"

type Printer interface {
	String() string
	Format(...interface{}) string
	Formatf(string, ...interface{}) string
	Print(*Log)
}

type consolePrinter struct {
	colored bool
}

func newConsolePrinter(c bool) Printer {
	if c {
		return newColorConsolePrinter()
	}
	return newSimpleConsolePrinter()
}

func newSimpleConsolePrinter() Printer {
	return &consolePrinter{false}
}

func newColorConsolePrinter() Printer {
	return &consolePrinter{true}
}

func (c consolePrinter) String() string {
	if c.colored {
		return "colored console logger"
	}
	return "plain console logger"
}

func (c consolePrinter) Format(args ...interface{}) string {
	return fmt.Sprintln(args...)
}


func (c consolePrinter) Formatf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func (c consolePrinter) Print(log *Log) {
	if c.colored {
		fmt.Println("colored", log.caller)
		return
	}
	fmt.Println(log.caller)
}