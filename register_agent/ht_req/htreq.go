package ht_req

import (
    "strings"
    "register_agent/clog"
    "register_agent/cfg"
    "encoding/json"
    "io/ioutil"
    "time"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/net/context"
)

type response struct {
    Code string `json:"code"`
    Message string  `json:"message"`
}

func registOnce(s string) bool {
    var url string=cfg.CONF.URL
	
	client_cfg:=&clientcredentials.Config{
        ClientID:       cfg.CONF.ClientID,
        ClientSecret:   cfg.CONF.ClientSecret,
        TokenURL:       cfg.CONF.TokenURL,
    }

	client:=client_cfg.Client(context.Background())

    res,err:=client.Post(url,"application/json",strings.NewReader(s))
    if err!=nil {
        clog.Fatal("post url:%s failed,error:%s",url,err)
        return false
    }

    defer res.Body.Close()

    if res.StatusCode!=200 {
        clog.Fatal("post url:%s response code:%d !=200",url,res.StatusCode)
        return false
    }

    body,err:=ioutil.ReadAll(res.Body)
    if err!=nil {
        clog.Fatal("read body failed.error:%s",err)
        return false
    }

    var jsonResp response     
    err=json.Unmarshal(body,&jsonResp)
    if err!=nil {
        clog.Fatal("unmarshal reponse:%s failed.error:%s",body,err)
        return false
    }

    if jsonResp.Code!="00006" {
        clog.Fatal("reponse code:%s !=00006,error:%s",jsonResp.Code,jsonResp.Message)
        return false
    }

    return true
}

func DeviceRegist(s string) {
    var ret bool
    var sec int=5

    for {
        ret=registOnce(s)
        if ret {
            clog.Info("register deivce ok")
            break
        }

        clog.Fatal("register device %d seconds later.",sec)

        time.Sleep(time.Duration(sec)*time.Second)
    }
}
