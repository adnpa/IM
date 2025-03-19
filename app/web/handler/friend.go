package handler

import (
	"github.com/gin-gonic/gin"
)

type FriendInfo struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
	Comment  string `json:"comment"`
	Avatar   string `json:"avatar"`
}

type GetFriendResp struct {
	Friends []FriendInfo
}

func GetFriendList(c *gin.Context) {}

func GetUserSelfApplyList(c *gin.Context) {}

func GetUserApplyList(c *gin.Context) {}

func ApplyAddFriend(c *gin.Context) {}

func HandleApplyFriend(c *gin.Context) {}

func DeleteFriend(c *gin.Context) {}

func GetFriendDetail(c *gin.Context) {}
