package es
import (
	"github.com/sadlil/gologger/logger"
	"time"
	"strings"
	elastigo "github.com/mattbaird/elastigo/lib"
)

func ESPrinter(log logger.LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	url, index := getParseClientOption(log.LoggerInit.Location)
	client := elastigo.NewConn()
	client.SetFromUrl(url)
	logJason := ElasticLog{packageName, fileName, funcName, lineNumber,log.LogType, log.Message, time}
	timeString := time.Format("2006-01-02")
	_, indexErr := client.Index(index, timeString, "", nil, logJason)
	if indexErr != nil {
		panic(indexErr)
	}
}


func getParseClientOption(url string) (string, string) {
	if(url == "") {
		return "http://localhost:9200", "gologger"
	}
	pos := strings.LastIndex(url, "/")
	esUrl := url[ 0 : pos ]
	if esUrl == "" {
		esUrl = "http://localhost:9200"
	}
	index := url[ pos+1 : len(url)]
	return esUrl, index
}

type ElasticLog struct {
	PackageName string
	FileName string
	FunctionName string
	LineNumber int
	LogType string
	Message string
	Time time.Time
}
