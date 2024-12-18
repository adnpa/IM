package friend

import (
	"context"
	"github.com/adnpa/IM/internal/api/api_info"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/adnpa/IM/pkg/discovery"
	"github.com/adnpa/IM/pkg/pb/pb_friend"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetFriendClient() pb_friend.FriendClient {
	cc := discovery.GetSrvConn(config.Config.RpcRegisterName.FriendName)
	client := pb_friend.NewFriendClient(cc)
	return client
}

func AddFriend(c *gin.Context) {
	params := api_info.AddFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.AddFriendReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.AddFriend(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.AddFriendResp{}
	c.JSON(http.StatusOK, pbResp)
}

func DeleteFriend(c *gin.Context) {
	params := api_info.DeleteFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.DeleteFriendReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.DeleteFriend(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.DeleteFriendResp{}
	//c.JSON(http.StatusOK, resp)

	c.JSON(http.StatusOK, pbResp)
}

func GetFriendApplyList(c *gin.Context) {
	params := api_info.GetFriendApplyListReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.GetFriendApplyReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.GetFriendApplyList(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.GetFriendApplyListResp{}
	//c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, pbResp)

}

func GetSelfFriendApplyList(c *gin.Context) {
	params := api_info.GetSelfApplyListReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.GetFriendApplyReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.GetSelfApplyList(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.GetSelfApplyListResp{}
	//c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, pbResp)

}

func GetFriendList(c *gin.Context) {
	params := api_info.GetFriendListReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.GetFriendListReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.GetFriendList(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.GetFriendListResp{}
	//c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, pbResp)

}

func AddFriendResponse(c *gin.Context) {
	params := api_info.AddFriendResponseReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.AddFriendResponseReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.AddFriendResponse(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.AddFriendResponseResp{}
	//c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, pbResp)
}

func SetFriendRemark(c *gin.Context) {
	params := api_info.SetFriendRemarkReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.SetFriendCommentReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.SetFriendComment(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.SetFriendRemarkResp{}
	//c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, pbResp)
}

func ImportFriend(c *gin.Context) {
	params := api_info.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.ImportFriendReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.ImportFriend(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.ImportFriendResp{}
	//c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, pbResp)
}

func IsFriend(c *gin.Context) {
	params := api_info.IsFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
		return
	}

	client := GetFriendClient()
	pbParams := &pb_friend.IsFriendReq{}
	_ = utils.CopyStructFields(pbParams, &params)
	pbResp, err := client.IsFriend(context.Background(), pbParams)
	if err != nil {
		logger.L().Warn("Api AddFriend Error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, constant.ErrInfo(constant.ErrInternal))
		return
	}

	//resp := api_info.IsFriendResp{}
	//c.JSON(http.StatusOK, resp)
	c.JSON(http.StatusOK, pbResp)
}
