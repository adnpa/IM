package util_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/adnpa/IM/app/oss/global"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

func TestGetConfig(t *testing.T) {
	p, _ := os.Getwd()
	configFileName := "user-srv.yaml"

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("err:%w config file:%s/%s, ", err, p, configFileName))
	}
	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic(err)
	}
	t.Log(global.NacosConfig)
	sc := constant.ServerConfig{
		IpAddr: global.NacosConfig.Host, // Nacos的服务地址
		Port:   global.NacosConfig.Port, // Nacos的服务端口
	}

	cc := constant.ClientConfig{
		// Endpoint:    "acm.aliyun.com:8080",
		NamespaceId: global.NacosConfig.Namespace,
		// RegionId:    "cn-shanghai",
		// AccessKey:   "LTAI4G8KxxxxxxxxxxxxxbwZLBr",
		// SecretKey:   "n5jTL9YxxxxxxxxxxxxaxmPLZV9",
		OpenKMS:   true,
		TimeoutMs: 5000,
		LogLevel:  "debug",
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: []constant.ServerConfig{sc},
		},
	)
	if err != nil {
		t.Error(err)
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(content)
	t.Log("111")
}
