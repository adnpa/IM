package handler

import (
	"github.com/gin-gonic/gin"
)

func AddFriend(c *gin.Context) {
	// srv := &friend.FriendService{}
	// ownId, _ := strconv.ParseInt(c.Query("own"), 10, 64)
	// friId, _ := strconv.ParseInt(c.Query("fri"), 10, 64)
	// params := friend.AddFriendReq{OwnId: ownId, FriId: friId}
	// if err := c.BindJSON(&params); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
	// 	return
	// }

	// srv.FriendApply(params)
	// c.ProtoBuf()
	// c.JSON(http.StatusOK, gin.H{"errCode": 0, "msg": "ok"})
}

type FriendInfo struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
	Comment  string `json:"comment"`
	Avatar   string `json:"avatar"`
}

type GetFriendResp struct {
	Friends []FriendInfo
}

func GetFriendList(c *gin.Context) {

	// logger.Info("resp", zap.Any("friL", friendL))
	// c.JSON(http.StatusOK, resp)
}

func GetFriendDetail(c *gin.Context) {}
