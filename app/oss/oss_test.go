package main

import (
	"testing"

	"github.com/adnpa/IM/app/oss/global"
	"github.com/adnpa/IM/app/oss/initialize"
	"github.com/adnpa/IM/app/oss/utils"
	"github.com/joho/godotenv"
)

func TestPresign(t *testing.T) {
	initialize.InitConfig()
	initialize.InitAliOssCli()
	err := godotenv.Load()
	if err != nil {
		t.Error("Error loading .env file")
	}

	t.Log(global.ServerConfig.OssInfo.BucketName)
	t.Log(global.ServerConfig.OssInfo.Region)
	getUrl, err := utils.GenPresignGetUrl("9fe89d61-6ff5-4517-8a43-65c3c5f98e91/auth0-step1.png")
	t.Log(getUrl)
	t.Log(err)
	// putUrl, err := utils.GenPresignPutUrl("bbb.txt", "bbb.txt")
	// t.Log(putUrl)
	// t.Log(err)
	t.Log("a")
}
