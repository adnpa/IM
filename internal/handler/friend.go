package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/adnpa/IM/internal/service/friend"
	"github.com/adnpa/IM/internal/service/user"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func AddFriend(c *gin.Context) {
	srv := &friend.FriendService{}
	ownId, _ := strconv.ParseInt(c.Query("own"), 10, 64)
	friId, _ := strconv.ParseInt(c.Query("fri"), 10, 64)
	params := friend.AddFriendReq{OwnId: ownId, FriId: friId}
	// if err := c.BindJSON(&params); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
	// 	return
	// }

	srv.FriendApply(params)
	c.ProtoBuf()
	c.JSON(http.StatusOK, gin.H{"errCode": 0, "msg": "ok"})
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
	var friendL []FriendInfo
	var resp GetFriendResp
	id := c.Query("uid")
	idNum, _ := strconv.ParseInt(id, 10, 64)
	cur, err := mongodb.GetAll("friend", bson.M{"ownerid": idNum})
	if err != nil {
		logger.Error("", zap.Error(err))
		return
	}

	for cur.Next(context.TODO()) {
		var result friend.Friend
		var u *user.User

		cur.Decode(&result)
		logger.Info("res", zap.Any("friend", result))
		mongodb.GetDecode("user", bson.M{"id": result.FriendID}, &u)
		logger.Info("user", zap.Any("", u))
		tmp := FriendInfo{
			Uid:      u.Id,
			Nickname: u.Nickname,
		}
		friendL = append(friendL, tmp)
	}
	resp.Friends = friendL
	logger.Info("resp", zap.Any("friL", friendL))
	c.JSON(http.StatusOK, resp)
}

func GetFriendDetail(c *gin.Context) {}
