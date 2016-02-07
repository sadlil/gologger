package gologger
import "github.com/spf13/pflag"

// Flags
var l int // Leveled flag
var c string // Category flag
var forceToGlog bool
var alsoLogToGlog bool

func init() {
	pflag.IntVar(&l, "logLevel", 0, "log level in which gologger logs to")
	pflag.StringVar(&c, "logCatagory", "", "log catagory which used for logging")
	pflag.BoolVar(&forceToGlog, "forceToGlog", false, "use glog for all logging. cancel other logging.")
	pflag.BoolVar(&alsoLogToGlog, "alsoLogToGlog", false, "also use glog as side logging.")
}

func SetLogLevel(l int) {

}

func setGlogLevel(l int) {

}

func SetLogCategory(s string) {

}

func ForceToGlog() {

}

func AlsoLogToGlog() {

}

func DisableForceGlog() {

}

func DisableGlog() {

}