package handler

import (
	"context"
	"net/http"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/code"
	"github.com/adnpa/IM/app/web/global"
	"github.com/adnpa/IM/app/web/utils"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GroupInfoView struct {
	GroupId int64  `json:"group_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
}

func GetUserGroups(c *gin.Context) {
	var results []GroupInfoView
	uid, _ := utils.GetUserId(c)
	logger.Info("uid", zap.Any("", uid))
	resp, err := global.GroupCli.GetUserGroups(context.Background(), &pb.GetUserGroupsReq{UserId: int32(uid)})
	if err != nil {
		logger.Error("get group list", zap.Error(err))
		c.JSON(http.StatusOK, ErrInfo(code.ErrUserNotFound))
		return
	}

	for _, f := range resp.GroupIds {
		tmp := GroupInfoView{
			GroupId: f,
		}

		g, err := global.GroupCli.GetGroupInfoById(context.Background(), &pb.GetGroupInfoByIdReq{GroupId: tmp.GroupId})
		if err != nil {
			logger.Info("", zap.Any("groupid", f), zap.Error(err))
			c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
			return
		}
		tmp.Name = g.GroupInfo.GroupName
		tmp.Avatar = g.GroupInfo.AvatarUrl
		results = append(results, tmp)
	}

	c.JSON(http.StatusOK, results)
}

// func GetUserGroups(c *gin.Context){}
