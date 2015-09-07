# gologger

A Simple Easy to use go logger library.  Displays Colored log into console in any unix or windows platform.
You can even store your logs in file, ElasticSearch or MySQL Database.

Still in Development Phase. 


Developed:

    - Colored and Simple Log into Console
    - File Logging
Developing:

    - Elasticsearch logging support
    - MySQL logging support
    
How to Get:

    go get github.com/sadlil/gologger
How To Use:
    
    import "github.com/sadlil/gologger"
    
    logger = gologger.GetLogger(gologger.CONSOLE, gologger.SimpleLog) 
    // Displays Simple plain log in console
    
    logger = gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog) 
    // Displays Colorful log in console
    
    logger = gologger.GetLogger(gologger.FILE, fileName) 
    // Log all the message in the given file.
    If file is not presents then creates it. if filename is "" creates 
    a default file named logs.txt in ur project directory.
    
    
    
    logger.Log(Message)
Displays

    [Log] [Time] [Package Name::File Name::Function Name] [Line Number] Message
    
You Can also use those functions now --

    logger.Log(message string)
    logger.Message(message string)
    logger.Info(message string)
    logger.Warn(message string)
    logger.Debug(message string)
    logger.Error(message string)

You can use two or more same type or diffrent type logger in same application. 

    logger1 = gologger.GetLogger(gologger.CONSOLE, gologger.SimpleLog)
    logger2 = gologger.GetLogger(gologger.FILE, "filelog.log")
    
    logger1.log("Hello Console") // loges into console.
    logger2.log("Hello File") // loges into file.


