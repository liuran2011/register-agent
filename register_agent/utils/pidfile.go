package utils

import (
    "os"
    "io/ioutil"
    "fmt"
    "register_agent/clog"
    "strconv"
    "strings"
)

const pidfile string="/var/run/register-agent.pid"

func stopInstance() {
    p,err:=ioutil.ReadFile(pidfile)

    if err!=nil {
        clog.Error("read pidfile:%s failed.error:%s",pidfile,err)
        return
    }

    pidStr:=fmt.Sprintf("%s",p)
    pid,err:=strconv.Atoi(strings.Trim(pidStr,"\n"))
    if err!=nil {
        clog.Error("convert pid:%s failed,error:%s",p,err)
        return 
    }

    proc,err:=os.FindProcess(pid)
    if err!=nil {
        clog.Error("find pid:%d failed.error:%s",pid,err)
        return
    }

    err=proc.Kill()
    if err!=nil {
        clog.Error("kill process failed. error:%s",err)
        return
    }

    clog.Info("stop already running instance, pid:%d",pid)
}

func saveInstancePid() {
    f,err:=os.Create(pidfile)
    if err!=nil {
        clog.Fatal("create pidfile:%s failed.error:%s",pidfile,err)
        os.Exit(1)
    }

    defer f.Close()

    pidstr:=fmt.Sprintf("%d",os.Getpid())
    f.WriteString(pidstr)

    clog.Info("save instance pid:%d to pidfile:%s",os.Getpid(),pidfile)
}

func EnsureSingleInstance() {
    stopInstance()
    saveInstancePid()
}
