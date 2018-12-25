package host_info

import (
	"encoding/json"
)

type hostinfo struct {
	HostInfo
	CpuInfo
	DiskInfo
	LocationInfo
	NetworkInfo
	VPNInfo
}

func MarshalHostInfo() ([]byte,error) {
	var info *hostinfo=new(hostinfo)
	
	GetCpuInfo(&(info.CpuInfo))
	GetDiskInfo(&(info.DiskInfo))
	GetHostInfo(&(info.HostInfo))
	GetLocationInfo(&(info.LocationInfo))
	GetNetworkInfo(&(info.NetworkInfo))
	GetVPNInfo(&(info.VPNInfo))
	
	return json.Marshal(*info)
}
