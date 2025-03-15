package handler

import (
	"github.com/adnpa/IM/internal/constant"
	"github.com/gin-gonic/gin"
)

func ErrInfo(code constant.ErrCode) gin.H {
	return gin.H{"errCode": code, "errMsg": constant.StatusText(code)}
}
