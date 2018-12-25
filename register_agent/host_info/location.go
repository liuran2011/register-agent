package host_info

import (
	"register_agent/cfg"
)

type LocationInfo struct {
	Province string `json:"province"`
	City string `json:"city"`
	District string `json:"district"`
	Address string `json:"address"`
	Organization string `json:"org"`
}

func GetLocationInfo(info *LocationInfo) error {
	info.Address=cfg.CONF.Address
	info.City=cfg.CONF.City
	info.District=cfg.CONF.District
	info.Organization=cfg.CONF.Organization
	info.Province=cfg.CONF.Province
	
	return nil
}