package host_info

import (
    "os"
    "bytes"
    "os/exec"
    "strconv"
    "strings"
    "register_agent/clog"
)

type CpuInfo struct {
	Number int `json:"cpu_number"`
	Freq string `json:"cpu_freq"`
	Model string `json:"cpu_model"`
	Manufacture string `json:"cpu_manufacture"`
}

func getModelName() string {
    cmd:=exec.Command("/bin/sh", "-c",`cat /proc/cpuinfo |grep "model name" | awk -F ':' '{print $2}' | head -n1`)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("get cpu model name failed.error:%s",err)
        os.Exit(1)
    }   
    
    return strings.TrimLeft(strings.Trim(out.String(),"\n")," ")
}

func getNumber() int {
    cmd:=exec.Command("/bin/sh", "-c",`cat /proc/cpuinfo  | grep processor  | wc -l`)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("get cpu number failed.error:%s",err)
        os.Exit(1)
    }   
  
    c,err:=strconv.Atoi(strings.Trim(out.String(),"\n"))
    if err!=nil {
        clog.Fatal("convert string:%s to int failed.error:%s",out.String(),err)
        os.Exit(1)
    }

    return c
}

func getManufacture() string {
    cmd:=exec.Command("/bin/sh", "-c",`cat /proc/cpuinfo  |grep "vendor_id" |awk -F ':' '{print $2}' | head  -n1`)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("get cpu manufacture failed.error:%s",err)
        os.Exit(1)
    }   
    
    return strings.TrimLeft(strings.Trim(out.String(),"\n")," ")
}

func getFreq() string {
    cmd:=exec.Command("/bin/sh", "-c",`cat /proc/cpuinfo  |grep "cpu MHz" |awk -F ':' '{print $2}' | head  -n1`)

    var out bytes.Buffer
    cmd.Stdout=&out

    err:=cmd.Run()
    if err!=nil {
        clog.Fatal("get cpu freq failed.error:%s",err)
        os.Exit(1)
    }   
    
    return strings.TrimLeft(strings.Trim(out.String(),"\n")," ")
}

func GetCpuInfo(info *CpuInfo) error {
    info.Model=getModelName()
    info.Number=getNumber()
    info.Manufacture=getManufacture()
    info.Freq=getFreq()

	return nil
}
