package gologger

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
	//elastigo "github.com/mattbaird/elastigo/lib"
)


func TestSimpleConsole(t *testing.T) {
	log := GetLogger(CONSOLE, SimpleLog)
	assert.NotNil(t, log)
	assert.Equal(t, "console", log.Logger.PrinterType)
	assert.Equal(t, "simple", log.Logger.Location)
	log.Message("Message Log")
	log.Warn("Warring Log")
	log.Debug("Debug Log")
	log.Info("Info Log")
	log.Error("Error Log")
	log.Log("Log Log")
	fmt.Println()
}

func TestColorConsole(t *testing.T) {
	log := GetLogger(CONSOLE, ColoredLog)
	assert.NotNil(t, log)
	assert.Equal(t, "console", log.Logger.PrinterType)
	assert.Equal(t, "color", log.Logger.Location)
	log.Message("Message Log")
	log.Warn("Warring Log")
	log.Debug("Debug Log")
	log.Info("Info Log")
	log.Error("Error Log")
	log.Log("Log Log")
	fmt.Println()
}


func TestDefaultFileLog(t *testing.T) {
	defaultFileLog := GetLogger(FILE, "")
	assert.NotNil(t, defaultFileLog)
	defaultFileLog.Log("Test File Log")
	file, err := os.Stat("log.txt")
	assert.Nil(t, err)
	assert.Equal(t, file.Name(), "log.txt")
	initFileSize := file.Size()
	defaultFileLog.Log("Second Line")
	file, err = os.Stat("log.txt")
	afterFileSize := file.Size()
	assert.NotEqual(t, initFileSize, afterFileSize)
	os.Remove("log.txt")
}

func TestFileLog(t *testing.T) {
	fileLog := GetLogger(FILE, "customfolder/customlog.txt")
	assert.NotNil(t, fileLog)
	fileLog.Log("Test File Log")
	file, err := os.Stat("customfolder/customlog.txt")
	assert.Nil(t, err)
	assert.Equal(t, file.Name(), "customlog.txt")
	initFileSize := file.Size()
	fileLog.Log("Second Line")
	file, err = os.Stat("customfolder/customlog.txt")
	afterFileSize := file.Size()
	assert.NotEqual(t, initFileSize, afterFileSize)
	os.RemoveAll("customfolder/customlog.txt")
}
/*
func TestEsLog(t *testing.T) {
	esLog := GetLogger(ELASTICSEARCH, "http://localhost:9200/customlogindex")
	assert.NotNil(t, esLog)
	assert.Equal(t, esLog.Logger.PrinterType, "es")
	esLog.Log("Log Into ES")
	client := elastigo.NewConn()
	client.SetFromUrl("http://localhost:9200")
	hits, err := client.Search("customlogindex", "", nil, nil)
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, hits.Hits.Hits, 0)
}*/
