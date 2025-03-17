package util_test

import (
	"testing"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func TestGetConfig(t *testing.T) {
	sc := constant.ServerConfig{
		IpAddr: "127.0.0.1", // Nacos的服务地址
		Port:   8848,        // Nacos的服务端口
	}

	cc := constant.ClientConfig{
		// Endpoint:    "acm.aliyun.com:8080",
		NamespaceId: "9ce6c088-a8e9-4867-a14f-94fe0937056b",
		// RegionId:    "cn-shanghai",
		// AccessKey:   "LTAI4G8KxxxxxxxxxxxxxbwZLBr",
		// SecretKey:   "n5jTL9YxxxxxxxxxxxxaxmPLZV9",
		OpenKMS:   true,
		TimeoutMs: 5000,
		LogLevel:  "debug",
	}

	client, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: []constant.ServerConfig{sc},
		},
	)

	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: "user-srv.yaml",
		Group:  "dev",
	})

	t.Log(content)
}
