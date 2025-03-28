package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/oss/global"
	"github.com/adnpa/IM/app/oss/utils"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"go.uber.org/zap"
)

type OssService struct {
	pb.UnimplementedOSSServer
}

// s *OssService pb.OSSServer
// var (
// 	region     string // 存储区域
// 	bucketName string // 存储空间名称
// 	objectName string // 对象名称
// )

func (s *OssService) Upload(_ context.Context, in *pb.UploadReq) (*pb.UploadResp, error) {
	if len(global.ServerConfig.OssInfo.BucketName) == 0 || len(global.ServerConfig.OssInfo.Region) == 0 || len(in.ObjectName) == 0 {
		return nil, fmt.Errorf("invalid parameters")
	}

	// tmp, _ := os.Create("tmpfile.wav")
	// defer tmp.Close()

	// tmp.Write(in.Content)
	content := bytes.NewReader(in.Content)

	request := &oss.PutObjectRequest{
		Bucket: &global.ServerConfig.OssInfo.BucketName, // 存储空间名称
		Key:    oss.Ptr(in.ObjectName),                  // 对象名称
		Body:   content,                                 // 要上传的内容
	}

	result, err := global.AliOssCli.PutObject(context.TODO(), request)
	if err != nil {
		logger.Errorf("failed to put object", zap.Error(err))
		return nil, err
	}
	logger.Info("", zap.Any("", result))

	url, err := utils.GenPresignGetUrl(in.ObjectName)

	if err != nil {
		logger.Errorf("failed to gen get url", zap.Error(err))
		return nil, err
	}
	return &pb.UploadResp{
		GetUrl: url,
	}, nil
}

func (s *OssService) Download(_ context.Context, in *pb.DownloadReq) (*pb.DownloadResp, error) {
	if len(in.BucketName) == 0 || len(in.Region) == 0 || len(in.ObjectName) == 0 {
		return nil, fmt.Errorf("invalid parameters")
	}

	// mimetype.
	// 创建获取对象的请求
	request := &oss.GetObjectRequest{
		Bucket: oss.Ptr(global.ServerConfig.OssInfo.BucketName), // 存储空间名称
		Key:    oss.Ptr(in.ObjectName),                          // 对象名称
	}
	// 执行获取对象的操作并处理结果
	result, err := global.AliOssCli.GetObject(context.TODO(), request)
	if err != nil {
		log.Fatalf("failed to get object %v", err)
	}
	defer result.Body.Close() // 确保在函数结束时关闭响应体

	log.Printf("get object result:%#v\n", result)

	// 读取对象的内容
	data, _ := io.ReadAll(result.Body)
	log.Printf("body:%s\n", data)

	return &pb.DownloadResp{
		// Content: data,
	}, nil
}
