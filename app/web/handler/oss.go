package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/global"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	logger.Info("form data", zap.Any("form field name", c.Request.FormValue("name")),
		zap.Any("form", form))
	fileHeader, err := c.FormFile("file")
	if err != nil {
		logger.Info("sdlkfja", zap.Error(err))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		logger.Error("fail open file")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		logger.Error("fail read file")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	objectName := fmt.Sprintf("%s/%s", uuid.NewString(), fileHeader.Filename)
	resp, err := global.OssCli.Upload(context.Background(), &pb.UploadReq{ObjectName: objectName, Content: data})
	if err != nil {
		logger.Error("fail upload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	logger.Info("grpc service resp", zap.Any("resp", resp))

	c.JSON(http.StatusOK, gin.H{
		"res_url": resp.GetUrl,
	})
}

func Download(c *gin.Context) {

}
