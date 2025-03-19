package initialize

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adnpa/IM/app/web/global"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
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

	sc := constant.ServerConfig{
		IpAddr: global.NacosConfig.Host,
		Port:   global.NacosConfig.Port,
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
		panic(err)
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		logger.Panic("读取nacos失败", zap.Error(err), zap.Any("config", global.ServerConfig))
	}

}
