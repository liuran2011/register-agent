package host_info

import (
    "io/ioutil"
    "os"
    "strings"
    "fmt"
    "os/exec"
    "bytes"
    "register_agent/clog"
)

type NetworkInfo struct {
	Number int `json:"interface_number"`
	Manufacture []string `json:"interfaces"`
}

const sysfsDir string ="/sys/class/net"

func trimInvalidInterface(fs []os.FileInfo) []os.FileInfo {
    if len(fs)==0 {
        return fs
    }

    validFlag:=make([]int,len(fs))

    for i,v:=range fs {
        if v.Mode().IsRegular() {
            clog.Info("entry:%s is regular file, ignore...",v.Name())
            continue
        }

        link,err:=os.Readlink(sysfsDir+"/"+v.Name())
        if err!=nil {
            clog.Error("readlink %s failed.error:%s",v.Name(),err)
            continue
        }

        if strings.Index(link,"/devices/virtual") != -1 {
            clog.Info("virtual interface:%s, ignore...",v.Name())
            continue
        }

        validFlag[i]=1
    }

    count:=0
    for _,v:=range validFlag {
        if v==1 {
            count++
        }
    }

    result:=make([]os.FileInfo,count)

    count=0
    for i,v:=range fs {
        if validFlag[i]==1 {
            result[count]=v
            count++
        }
    }

    return result
}

func getPci(dev string) (string,error) {
    cmd:=fmt.Sprintf("ethtool -i %s | grep bus-info",dev)
    command:=exec.Command("/bin/sh","-c",cmd)

    var out bytes.Buffer
    command.Stdout=&out
    err:=command.Run()
    if err!=nil {
        clog.Error("ethtool get bus info failed for:%s",dev)
        return "",err
    }

    pci:=strings.Trim(strings.TrimLeft(out.String(),"bus-info: "),"\n")
    
    return pci,nil
}

func getInterfaceDesc(pci string) (string,error) {
    cstr:=fmt.Sprintf("lspci -s %s",pci)
    cmd:=exec.Command("/bin/sh","-c",cstr)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Error("exec command %s failed. error:%s",cstr,err)
        return "",err
    }

    index:=strings.LastIndex(out.String(),":")
    if index!=-1 {
        return strings.TrimLeft(out.String()[index+1:len(out.String())-1]," "),nil
    } else {
        return out.String(),nil
    }
}

func getInterfaceInfo(fs []os.FileInfo) []string {
    info:=make([]string,len(fs))
    count:=0

    for _,v:=range fs {
        pci,err:=getPci(v.Name())
        if err!=nil {
            continue
        }

        desc,err:=getInterfaceDesc(pci)
        if err!=nil {
            continue
        }

        info[count]=desc
        count++
   }

   return info
}

func GetNetworkInfo(info *NetworkInfo) error {
    fileInfo,err:=ioutil.ReadDir(sysfsDir)
    if err!=nil {
        clog.Fatal("read dir:%s failed.error:%s",sysfsDir,err)
        os.Exit(1)
    }

    fileInfo=trimInvalidInterface(fileInfo)
    info.Number=len(fileInfo)
    info.Manufacture=getInterfaceInfo(fileInfo)

	return nil	
}
