package utils

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/adnpa/IM/app/oss/global"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

func TypeDisposition(typ string) string {
	if strings.HasPrefix(typ, "image/") {
		return "inline"
	} else {
		return "attachment"
	}
}

func GenPresignGetUrl(objectName string) (string, error) {
	if len(objectName) == 0 {
		return "", fmt.Errorf("invalid parameters, object name required")
	}
	req := &oss.GetObjectRequest{
		Bucket: oss.Ptr(global.ServerConfig.OssInfo.BucketName),
		Key:    oss.Ptr(objectName),
		// ResponseContentType:        oss.Ptr(typ),
		ResponseContentDisposition: oss.Ptr(fmt.Sprintf("%s; filename=%s", TypeDisposition("image/png"), url.QueryEscape("tt.png"))),
		// ResponseContentDisposition: oss.Ptr("inline; filename=\"%s\""),
	}

	return genPresignUrl(req)
}

func GenPresignPutUrl(objectName, downloadName, typ string) (string, error) {
	if len(objectName) == 0 {
		return "", fmt.Errorf("invalid parameters, object name required")
	}
	req := &oss.PutObjectRequest{
		Bucket:             oss.Ptr(global.ServerConfig.OssInfo.BucketName),
		Key:                oss.Ptr(objectName),
		ContentType:        oss.Ptr(typ),
		ContentDisposition: oss.Ptr(fmt.Sprintf("%s; filename=%s", TypeDisposition(typ), downloadName)),
	}
	return genPresignUrl(req)
}

// 生成文件临时访问url
func genPresignUrl(request any) (string, error) {
	if len(global.ServerConfig.OssInfo.BucketName) == 0 {
		return "", fmt.Errorf("invalid parameters, bucket name required")
	}
	if len(global.ServerConfig.OssInfo.Region) == 0 {
		return "", fmt.Errorf("invalid parameters, region required")
	}

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(global.ServerConfig.OssInfo.Region)

	// 创建OSS客户端
	client := oss.NewClient(cfg)

	// 生成GetObject的预签名URL
	result, err := client.Presign(context.TODO(), request, oss.PresignExpires(10*time.Minute))
	if err != nil {
		return "", fmt.Errorf("failed to presign %v", err)
	}
	// log.Printf("request method:%v\n", result.Method)
	// log.Printf("request expiration:%v\n", result.Expiration)
	// log.Printf("request url:%v\n", result.URL)

	if len(result.SignedHeaders) > 0 {
		//当返回结果包含签名头时，使用预签名URL发送GET请求时也包含相应的请求头，以免出现不一致，导致请求失败和预签名错误
		log.Printf("signed headers:\n")
		for k, v := range result.SignedHeaders {
			log.Printf("%v: %v\n", k, v)
		}
	}

	return result.URL, nil
}
