package host_info

import (
	"register_agent/cfg"
    "syscall"
    "register_agent/clog"
    "os"
    "strings"
)

type HostInfo struct {
	UUID string `json:"uuid"`
	DeviceType string `json:"device_type"`
	SoftwareVersion string `json:"software_version"`
	OsVersion string `json:"os_version"`
    Hostname string `json:"hostname"`
}

func int8ToString(bs [65]int8) string {
    buf:=make([]byte,len(bs))
    
    var i int 
    var b int8

    for i,b=range bs {
        if b==0 {
            break
        }
        buf[i]=byte(b)
    }

    return string(buf[:i]) 
}

func getOsVersion() string {
    var name syscall.Utsname

    err:=syscall.Uname(&name)
    if err!=nil {
        clog.Fatal("syscall uname failed.error:%s",err)
        os.Exit(1)
    }
   
    sysname:=int8ToString(name.Sysname)
    release:=int8ToString(name.Release)
    version:=int8ToString(name.Version)

    return sysname+" "+release+" "+version 
}

func GetHostInfo(info *HostInfo) error {
	info.DeviceType=cfg.CONF.DeviceType
	info.UUID=cfg.CONF.UUID
    info.Hostname=strings.Join([]string{info.DeviceType,cfg.CONF.Province,cfg.CONF.City,cfg.CONF.District,cfg.CONF.Address,cfg.CONF.Organization},"-")

    info.OsVersion=getOsVersion()
    info.SoftwareVersion="1.0"

	return nil
}
