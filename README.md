# gologger

A Simple Easy to use go logger library.  Displays Colored log into console in any unix or windows platform.
You can even store your logs in file, ElasticSearch or MySQL Database.


#### Developed:
    - Colored and Simple Log into Console
    - File Logging
    - ElasticSearch Logging


#### How to Get:
``` sh
$ go get github.com/sadlil/gologger
```
#### How To Use:
``` go
import "github.com/sadlil/gologger"

Console::
logger = gologger.GetLogger(gologger.CONSOLE, gologger.SimpleLog)
// Displays Simple plain log in console

logger = gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
// Displays Colorful log in console


File::
logger = gologger.GetLogger(gologger.FILE, fileName)
// Log all the message in the given file.
// If file is not presents then creates it. if filename is "" creates
// a default file named logs.txt in ur project directory.


ElasticSearch::
logger = gologger.GetLogger(gologger.ELASTICSEARCH, location)
// Logs everything into elasticsearch. if location is "" then it
// looks for elasticsearch by default in http://localhost:9200
// and logs stored in index 'gologger'.
// If you want to provide custom location and index for your log
// you must provide the location in this format
// "http://Your_ES_Url:ES_Port/Index",
// If you only want to change the default index name you can do that
// by sending "/YourIndex", it will use default localhost for ES.
```
#### Display:
``` go
logger.Log(Message) -> displays in console.
```
``` sh
[Log] [Time] [Package Name::File Name::Function Name] [Line Number] Message
```
#### Logger Functions Available:
``` go
logger.Log(message string)
logger.Message(message string)
logger.Info(message string)
logger.Warn(message string)
logger.Debug(message string)
logger.Error(message string)
```
#### You can use two or more same type or diffrent type logger in same application.
``` go
logger1 = gologger.GetLogger(gologger.CONSOLE, gologger.SimpleLog)
logger2 = gologger.GetLogger(gologger.FILE, "filelog.log")

logger1.log("Hello Console") // loges into console.
logger2.log("Hello File") // loges into file.
```


  **Any suggestion and Issue would be highly appreciated**. You can create an issue or pull request
  if you think i need to modify my library.


#### Used Libraries
1. Color Console - [faith/color](https://github.com/fatih/color)
2. ElasticSearch - [olivere/elastic](https://github.com/olivere/elastic)
