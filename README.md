# gologger

A Simple Easy to use go logger library.  Displays Colored log into console in any unix or windows platform.
You can even store ur logs in any file or elasticsearch.

Still in Development Phase. 


Developed:

    - Colored Log into Console
Developing:

    - File Logging support
    - Elasticsearch logging support
    
How to Get:

    go get github.com/sadlil/gologger
How To Use:
    
    import "github.com/sadlil/gologger"
    
    logger = gologger.GetLogger()
    logger.Log(Message)
Displays

    [Log] [Calling File Name] [Line Number] Message
    
You Can also use those functions now --

    logger.Log(message string)
    logger.Message(message string)
    logger.Info(message string)
    logger.Warn(message string)
    logger.Debug(message string)
    logger.Error(message string)