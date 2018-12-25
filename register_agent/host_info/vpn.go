package host_info

import (
	"register_agent/clog"
	"os"
)

type VPNInfo struct {
	VPNIp string `json:"vpn_ip"`
}

func GetVPNInfo(info *VPNInfo) error {
	ip:=os.Getenv("ifconfig_local")
	if ip=="" {
		clog.Fatal("ifconfig_local not exist.")
		os.Exit(1)
	}

	info.VPNIp=ip
	
	return nil
}
