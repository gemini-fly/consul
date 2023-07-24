package utils

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
)

func AppPackage(addr string) (string, error) {
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
		return "", err
	}

	if pair == nil {
		return "", fmt.Errorf("No Value")
	}

	return string(pair.Value), nil

}
