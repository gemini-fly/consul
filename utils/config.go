package utils

import (
	"encoding/json"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

type Config struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func App_Package(addr string) {
	DefaultConfig := consulapi.DefaultConfig()
	DefaultConfig.Address = "10.111.176.1:8500" // 设置Consul的地址

	// 创建Consul客户端
	client, err := consulapi.NewClient(DefaultConfig)
	if err != nil {
		panic(err)
	}

	// 获取配置文件内容
	pair, _, err := client.KV().Get(addr, nil)
	if err != nil {
		panic(err)
	}
	if pair != nil {
		var C Config
		err = json.Unmarshal(pair.Value, &C)
		if err != nil {
			log.Fatalf(err.Error())
		}
	} else {
		//键不存在的情况
		panic("键不存在")
	}
}
