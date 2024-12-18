package user

import (
	"context"
	"github.com/adnpa/IM/internal/api/api_info"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetUserClient() pb_user.UserClient {
	cc := discovery.GetSrvConn(config.Config.RpcRegisterName.UserName)
	client := pb_user.NewUserClient(cc)
	return client
}

func GetUserInfo(c *gin.Context) {
	//logger.L().Info("GetUserinfo")
	logger.L().Info("GetUserinfo")
	params := api_info.GetUsersInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetUserClient()
	pbParams := &pb_user.GetUserInfoReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.GetUserInfo(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api GetUserInfo Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.GetUsersInfoResp{
	//	CommResp: api_info.CommResp{
	//		ErrCode: pbResp.ErrorCode,
	//		ErrMsg:  pbResp.ErrorMsg,
	//	},
	//	UserInfoList: ,
	//	Data: ,
	//}
	c.JSON(http.StatusOK, pbResp)
}

func GetSelfUserInfo(c *gin.Context) {
	params := api_info.GetSelfUserInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetUserClient()
	pbParams := &pb_user.GetUserInfoReq{}
	pbResp, err := client.GetUserInfo(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api GetSelfUserInfo Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}
	//resp := api_info.GetSelfUserInfoResp{}
	c.JSON(http.StatusOK, pbResp)
}

func UpdateUserInfo(c *gin.Context) {
	params := api_info.UpdateUserInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	pbParams := &pb_user.UpdateUserInfoReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	client := GetUserClient()
	pbResp, err := client.UpdateUserInfo(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api GetUserInfo Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.UpdateUserInfoResp{}
	c.JSON(http.StatusOK, pbResp)
}
