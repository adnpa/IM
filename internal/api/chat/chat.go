package chat

import (
	"github.com/adnpa/IM/internal/api/api_info"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSeq(c *gin.Context) {
	params := api_info.ParamsUserNewestSeq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
	}

	resp := api_info.UpdateUserInfoResp{}
	c.JSON(http.StatusOK, resp)
}

func PullMsgBySeqList(c *gin.Context) {
	params := api_info.ParamsUserPullMsgBySeqList{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
	}

	c.JSON(http.StatusOK, "")
}

//func newUserSendMsgReq(token string, params *paramsUserSendMsg) *pb_chat.SendMsgReq {
//
//}

func SendMsg(c *gin.Context) {
	params := api_info.UpdateUserInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
	}

	c.JSON(http.StatusOK, "")
}
