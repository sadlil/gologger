package gologger

import (
	"github.com/mattbaird/elastigo/lib"
	"strings"
	"time"
	"fmt"
	"os"
)

func ESPrinter(log LogInstance, packageName string, fileName string, lineNumber int, funcName string, time time.Time) {
	url, index := getParseClientOption(log.LoggerInit.Location)
	client := elastigo.NewConn()
	client.SetFromUrl(url)
	logJason := ElasticLog{packageName, fileName, funcName, lineNumber, log.LogType, log.Message, time}
	timeString := time.Format("2006-01-02")
	_, indexErr := client.Index(index, timeString, "", nil, logJason)
	if indexErr != nil {
		fmt.Println(indexErr)
		os.Exit(1)
	}
}

func getParseClientOption(url string) (string, string) {
	if url == "" {
		return "http://localhost:9200", "gologger"
	}
	pos := strings.LastIndex(url, "/")
	esUrl := url[0:pos]
	if esUrl == "" {
		esUrl = "http://localhost:9200"
	}
	index := url[pos+1:]
	return esUrl, index
}

type ElasticLog struct {
	PackageName  string
	FileName     string
	FunctionName string
	LineNumber   int
	LogType      string
	Message      string
	Time         time.Time
}
