package friend

import (
	"github.com/adnpa/IM/internal/api/api_info"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBlack(c *gin.Context) {
	params := api_info.AddBlacklistReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
	}

	resp := api_info.AddBlacklistResp{}
	c.JSON(http.StatusOK, resp)
}

func GetBlacklist(c *gin.Context) {
	params := api_info.GetBlackListReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
	}

	resp := api_info.GetBlackListResp{}
	c.JSON(http.StatusOK, resp)
}

func RemoveBlack(c *gin.Context) {
	params := api_info.RemoveBlackListReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, constant.ErrInfo(constant.ErrArgs))
	}

	resp := api_info.RemoveBlackListResp{}
	c.JSON(http.StatusOK, resp)
}
