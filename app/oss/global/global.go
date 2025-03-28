package global

import (
	"github.com/adnpa/IM/app/oss/config"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
)

var (
	AliOssCli    *oss.Client
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
