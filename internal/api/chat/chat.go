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

func SendMsg(c *gin.Context) {
	params := api_info.ParamsUserSendMsg{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
	}

	c.JSON(http.StatusOK, "")
}

func DelMsg(c *gin.Context) {
	params := api_info.DelMsgReq{
		OpUserID:    "",
		UserID:      "",
		SeqList:     nil,
		OperationID: "",
	}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}
