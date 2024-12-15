package auth

import (
	"context"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb"

	"github.com/gin-gonic/gin"
	"net/http"
)

type paramRegisterReq struct {
	Secret   string `json:"secret" binding:"required,max=32"`
	Platform int32  `json:"platform" binding:"required,min=1,max=7"`
	UID      string `json:"uid" binding:"required,min=1,max=64"`
	Name     string `json:"name" binding:"required,min=1,max=64"`
	Icon     string `json:"icon" binding:"omitempty,max=1024"`
	Gender   int32  `json:"gender" binding:"omitempty,oneof=0 1 2"`
	Mobile   string `json:"mobile" binding:"omitempty,max=32"`
	Birth    string `json:"birth" binding:"omitempty,max=16"`
	Email    string `json:"email" binding:"omitempty,max=64"`
	Ex       string `json:"ex" binding:"omitempty,max=1024"`
}

func registerParams2PB(params *paramRegisterReq) *pb.RegisterReq {
	pbData := pb.RegisterReq{
		UID:    params.UID,
		Name:   params.Name,
		Icon:   params.Icon,
		Gender: params.Gender,
		Mobile: params.Mobile,
		Birth:  params.Birth,
		Email:  params.Email,
		Ex:     params.Ex,
	}
	return &pbData
}

func UserRegister(c *gin.Context) {
	grpcConn := discovery.GetSrvConn(config.Config.RpcRegisterName.AuthName)
	cli := pb.NewAuthClient(grpcConn)

	params := paramRegisterReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	pbData := registerParams2PB(&params)

	registerResp, err := cli.Register(context.Background(), pbData)
	if err != nil || !registerResp.Success {
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err})
		return
	}

}
