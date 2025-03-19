package handler

import (
	"context"
	"net/http"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/code"
	"github.com/adnpa/IM/app/web/global"
	"github.com/adnpa/IM/app/web/handler/forms"
	"github.com/adnpa/IM/app/web/utils"
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

func GetFriendList(c *gin.Context) {
	uid, _ := utils.GetUserId(c)
	resp, err := global.FriendCli.GetFriendsByUserId(context.Background(), &pb.GetFriendsByUserIdReq{Uid: int32(uid)})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, resp.Friends)
}

func GetUserSelfApplyList(c *gin.Context) {
	uid, _ := utils.GetUserId(c)
	resp, err := global.FriendCli.GetFriendApply(context.Background(), &pb.GetFriendApplyReq{UserId: int32(uid)})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, resp.FriendApplys)
}

func ApplyAddFriend(c *gin.Context) {
	var friendApply *pb.FriendApply
	err := c.ShouldBind(&friendApply)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrArgs))
	}

	isResp, err := global.FriendCli.IsFriend(context.Background(), &pb.IsFriendReq{Left: friendApply.From, Right: friendApply.To})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}

	if isResp.IsFriend {
		c.JSON(http.StatusOK, ErrInfo(code.ErrAlreadyFriend))
		return
	}

	resp, err := global.FriendCli.CreateFriendApply(context.Background(), &pb.CreateFriendApplyReq{FriendApply: friendApply})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, resp.Success)
}

func HandleApplyFriend(c *gin.Context) {
	var friendApply *pb.FriendApply
	err := c.ShouldBind(&friendApply)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrArgs))
	}

	resp, err := global.FriendCli.UpdateFriendApply(context.Background(), &pb.UpdateFriendApplyReq{FriendApply: friendApply})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, resp.Success)
}

func DeleteFriend(c *gin.Context) {
	var form forms.DeleteFriendForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrArgs))
	}

	uid, _ := utils.GetUserId(c)

	isResp, err := global.FriendCli.IsFriend(context.Background(), &pb.IsFriendReq{Left: int32(uid), Right: form.FriendId})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	if !isResp.IsFriend {
		c.JSON(http.StatusOK, ErrInfo(code.ErrNotFriend))
		return
	}

	resp, err := global.FriendCli.DeleteFriend(context.Background(), &pb.DeleteFriendReq{UserId: int32(uid), FriendId: form.FriendId})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, resp.Success)
}

func GetFriendDetail(c *gin.Context) {
	var form forms.GetFriendDetailForm
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrArgs))
	}

	// global.UserCli.IsFriend()

	resp, err := global.UserCli.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: form.FriendId})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, resp)
}
