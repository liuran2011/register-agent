package cfg

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"register_agent/clog"
    "github.com/satori/go.uuid"
)

const cfgFilePath="/etc/register_agent/register_agent.conf"

type Conf struct {
	URL string `json:"url"`
	DeviceType string `json:"device_type"`
	UUID string `json:"uuid"`
	Province string `json:"province"`
	City string	`json:"city"`
	District string	`json:"district"`
	Address string	`json:"address"`
	Organization string	`json:"organization"`
    FalconDir string `json:"falcon_dir"`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	TokenURL string `json:"token_url"`
}

var CONF  Conf

func loadJson() {
	data,err:=ioutil.ReadFile(cfgFilePath)
	if err!=nil {
		clog.Fatal("open config file: %s failed.error:%s",cfgFilePath,err)
		os.Exit(1)
	}
	
	err=json.Unmarshal(data, &CONF)
	if err!=nil {
		clog.Fatal("parse config file:%s failed.error:%s",cfgFilePath,err)
		os.Exit(1)
	}
}

func saveJson(cfg * Conf) {
    bytes,err:=json.Marshal(cfg)
    if err!=nil {
        clog.Fatal("marshal CONF:%v to json failed. error:%s",cfg,err)
        os.Exit(1)
    }

    err=ioutil.WriteFile(cfgFilePath,bytes,0644)
    if err!=nil {
        clog.Fatal("write file %s failed. error:%s",cfgFilePath,err)
        os.Exit(1)
    }
}

func sanityCheck() {	
	if CONF.Address=="" {
		clog.Fatal("Address is null!")
		os.Exit(1)
	}
	
	if CONF.District=="" {
		clog.Fatal("District is null!")
		os.Exit(1)
	}
	
	if CONF.City=="" {
		clog.Fatal("City is null!")
		os.Exit(1)
	}
	
	if CONF.Province=="" {
		clog.Fatal("Province is null!")
		os.Exit(1)
	}
	
	if CONF.Organization=="" {
		clog.Fatal("Organization is null!")
		os.Exit(1)
	}
	
	if CONF.DeviceType=="" {
		clog.Fatal("DeviceType is null!")
		os.Exit(1)
	}
	
	if CONF.URL=="" {
		clog.Fatal("URL is null!")
		os.Exit(1)
	}

	if CONF.ClientID=="" {
		clog.Fatal("OAuth2.0 client id is null!")
		os.Exit(1)
	}

	if CONF.ClientSecret=="" {
		clog.Fatal("OAuth2.0 client secret is null!")
		os.Exit(1)
	}

	if CONF.TokenURL=="" {
		clog.Fatal("OAuth2.0 token url is null!")
		os.Exit(1)
	}
}

func Init() {
	loadJson()
	sanityCheck()
    if CONF.UUID=="" {
        uuid,err:= uuid.NewV4()
        if err!=nil {
            clog.Fatal("generator uuid failed. error:%s",err)
            os.Exit(1)
        }

        CONF.UUID=uuid.String()
        saveJson(&CONF)
    }
}
