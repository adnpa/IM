package auth

import (
	"context"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type paramsToken struct {
	Secret   string `json:"secret" binding:"required,max=32"`
	Platform int32  `json:"platform" binding:"required,min=1,max=8"`
	UID      string `json:"uid" binding:"required,min=1,max=64"`
}

func tokenParams2Pb(params *paramsToken) *pb.TokenReq {
	pbData := pb.TokenReq{
		Platform: params.Platform,
		UID:      params.UID,
	}
	return &pbData
}

func UserToken(c *gin.Context) {
	conn := discovery.GetSrvConn(config.Config.RpcRegisterName.AuthName)
	cli := pb.NewAuthClient(conn)

	params := &paramsToken{}
	if err := c.ShouldBind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pbData := tokenParams2Pb(params)

	reply, err := cli.Token(context.Background(), pbData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if reply.ErrCode == 0 {
		c.JSON(http.StatusOK, gin.H{
			"errCode": reply.ErrCode,
			"errMsg":  reply.ErrMsg,
			"data": gin.H{
				"uid":         pbData.UID,
				"token":       reply.Token,
				"expiredTime": reply.ExpiredTime,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"errCode": reply.ErrCode,
			"errMsg":  reply.ErrMsg,
		})
	}
}
