package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PullHistoryMsgReq struct{}

func PullHistoryMsg(c *gin.Context) {
	// id, _ := strconv.ParseInt(c.Query("uid"), 10, 64)
	// chat.HistoryMsgQueue.PullUsrOffLineMsg(id)
	c.JSON(http.StatusOK, gin.H{})
}

type GetMsgByIdsReq struct {
	Ids []int64
}

func GetMsgByIds(c *gin.Context) {

}
