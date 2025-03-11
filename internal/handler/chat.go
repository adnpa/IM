package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PullHistoryMsgReq struct{}

func PullHistoryMsg(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}

type GetMsgByIdsReq struct {
	Ids []int64
}

func GetMsgByIds(c *gin.Context) {
	
}
