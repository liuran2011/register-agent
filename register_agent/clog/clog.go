package clog

import (
	"fmt"
    "log"
    "log/syslog"
    "os"
)

var infoLogger *log.Logger
var errorLogger *log.Logger
var debugLogger *log.Logger
var infologger *log.Logger
var fatalLogger *log.Logger

const logPrefix="register-agent"

func init() {
    info,err:=syslog.Dial("","",syslog.LOG_USER|syslog.LOG_INFO,logPrefix)
    if err!=nil {
        fmt.Fprintf(os.Stderr,"create info logger failed.error:%s",err)
        os.Exit(1)
    }
    infoLogger=log.New(info,"",0)

    debug,err:=syslog.Dial("","",syslog.LOG_USER|syslog.LOG_DEBUG,logPrefix)
    if err!=nil {
        fmt.Fprintf(os.Stderr,"create debug logger failed.error:%s",err)
        os.Exit(1)
    }
    debugLogger=log.New(debug,"",0)

    errorl,err:=syslog.Dial("","",syslog.LOG_USER|syslog.LOG_ERR,logPrefix)
    if err!=nil {
        fmt.Fprintf(os.Stderr,"create error logger failed.error:%s",err)
        os.Exit(1)
    }
    errorLogger=log.New(errorl,"",0)

    fatal,err:=syslog.Dial("","",syslog.LOG_USER|syslog.LOG_CRIT,logPrefix)
    if err!=nil {
        fmt.Fprintf(os.Stderr,"create fatal logger failed.error:%s",err)
        os.Exit(1)
    }
    fatalLogger=log.New(fatal,"",0)
}

func Error(format string, args ...interface{}) {
    errorLogger.Printf(format,args...) 
}

func Debug(format string, args ...interface{}) {
	debugLogger.Printf(format,args...)
}

func Info(format string, args ...interface{}) {
	infoLogger.Printf(format,args...)
}

func Fatal(format string, args ...interface{}) {
	fatalLogger.Printf(format,args...)
}
