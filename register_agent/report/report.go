package report

import (
    "register_agent/cfg"
    "os/exec"
    "os"
    "fmt"
    "register_agent/clog"
    "bytes"
    "time"
    "strings"
)

func modifyHostname() {
    file:=strings.Join([]string{cfg.CONF.FalconDir,"agent/config/cfg.json"},"/")

    hostname:=strings.Join([]string{cfg.CONF.DeviceType,cfg.CONF.Province,cfg.CONF.City,cfg.CONF.District,cfg.CONF.Address,cfg.CONF.Organization},"-")
   
    str:=fmt.Sprintf(`sed -i '/hostname/d' %s; sed -i '/"ip": "",/i\    "hostname":"%s",' %s`,
        file,hostname,file)
    cmd:=exec.Command("/bin/sh","-c",str)

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("write hostname:%s to file:%s failed.error:%s",
            hostname,file,err)
        os.Exit(1)
    }
}

func startAgent() bool {
    modifyHostname()

    cstr:=fmt.Sprintf("export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin; cd %s; ./open-falcon start agent",cfg.CONF.FalconDir)
    cmd:=exec.Command("/bin/sh","-c",cstr)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("start falcon agent failed.error:%s",err)
        return false
    }

    return true
}

func stopAgent() bool {
    cstr:=fmt.Sprintf("export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin; cd %s; ./open-falcon stop agent",cfg.CONF.FalconDir)
    cmd:=exec.Command("/bin/sh","-c",cstr)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("stop falcon agent failed. error:%s",err)
        return false
    }

    return true
}

func checkAgent() bool {
    cstr:=fmt.Sprintf("export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin; cd %s; ./open-falcon check agent",cfg.CONF.FalconDir)
    cmd:=exec.Command("/bin/sh","-c",cstr)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("run check falcon agent failed.error:%s",err)
        return false
    }

    var name, status,pid string
    n,err:=fmt.Sscanf(out.String(),"%s%s%s",&name,&status,&pid)
    if err!=nil || n!=3 {
        clog.Fatal("parse check falcon agent cmd failed.error:%s",err)
        return false
    }

    if status!="UP" {
        clog.Fatal("falcon agent status:%s !=UP",status)
        return false
    }

    return true
}

func StopReport() {
    flag:=checkAgent()
    if flag {
        stopAgent()
    }
}

func StartReport() {
    flag:=checkAgent()
    if flag {
        clog.Error("falcon-agent already running. stop it first...")
        stopAgent()
    }

    for {
        flag=startAgent()
        if flag {
            clog.Info("start falcon-agent ok!")
            flag=checkAgent()
        }

        if flag {
            clog.Info("check falcon-agent ok!")
            break
        }

        clog.Info("start falcon agent 5 seconds later...")

        time.Sleep(time.Duration(5)*time.Second)
    }
}
