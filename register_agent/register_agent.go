package main

import (
	"register_agent/clog"
	"os"
	hostinfo "register_agent/host_info"
    "register_agent/ht_req"
    "register_agent/report"
    "register_agent/cfg"
    "register_agent/utils"
    "os/signal"
    "syscall"
)

func main() {
    c:=make(chan os.Signal)
    signal.Notify(c,syscall.SIGINT,syscall.SIGTERM)
    utils.EnsureSingleInstance()

	cfg.Init()
	
	r,err:=hostinfo.MarshalHostInfo()
	if err!=nil {
		clog.Fatal("get hostinfo failed. error:%s",err)
		os.Exit(1)
	}

    go func() {
        for s:= range c {
            switch s {
                case syscall.SIGINT,syscall.SIGTERM:
                    clog.Info("got signal, stop falcon agent...")
                    report.StopReport()
                    os.Exit(0)
                default:
                    clog.Info("non sigint or sigterm signal receive.")
            }
        }
    }()

    ht_req.DeviceRegist(string(r))
    
    report.StartReport()
}
