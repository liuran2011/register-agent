package host_info

import (
    "os/exec"
    "register_agent/clog"
    "strings"
    "bytes"
    "fmt"
)

type DiskInfo struct {
	Number int `json:"disk_number"`
	Manufacture []string `json:"disk_manufacture"`
}

func getDiskDeviceList() []string {
    cmd:=exec.Command("/bin/sh","-c",`lsblk -P |grep "TYPE=\"disk\"" | awk '{print $1}'`)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Error("get disk list failed.error:%s",err)
        return nil
    }

    list:=strings.Split(out.String(),"\n")
    for i,v:=range list {
        list[i]=strings.Trim(strings.TrimLeft(v,"NAME=\""),"\"")
    }

    return list
}

func getVendorAndModel(d []string) []string {
    var result []string

    for _,v:=range d {
        str:=fmt.Sprintf("udisks --show-info /dev/%s|grep vendor -A1",v)
        cmd:=exec.Command("/bin/sh","-c",str)

        var out bytes.Buffer
        cmd.Stdout=&out

        err:=cmd.Run()
        if err!=nil {
            clog.Error("get disk:%s info failed.error:%s",v,err)
            continue
        }

        li:=strings.Split(out.String(),"\n")
        li[0]=strings.Trim(strings.TrimLeft(li[0],"    vendor:")," ")
        li[1]=strings.Trim(strings.TrimLeft(li[1],"    model:")," ")

        if len(li[0])==0 && len(li[1])==0 {
            result=append(result,"Unkown vendor, /dev/"+v)
        } else {
            result=append(result,li[0]+","+li[1])
        }

    }

    return result
}

func GetDiskInfo(info *DiskInfo) error {
    disk:=getDiskDeviceList()
    list:=getVendorAndModel(disk)

    if len(list)>0 {
        info.Number=len(list)
        info.Manufacture=list
    } else {
        info.Number=1
        info.Manufacture=[]string{"Unkown vendor"}
    }

	return nil
}
