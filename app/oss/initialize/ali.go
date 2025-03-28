package initialize

import (
	"github.com/adnpa/IM/app/oss/global"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

func InitAliOssCli() {
	// ali := global.ServerConfig.OssInfo

	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion("cn-guangzhou")

	client := oss.NewClient(cfg)
	global.AliOssCli = client
}
