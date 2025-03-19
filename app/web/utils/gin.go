package utils

import (
	"fmt"
	"strconv"

	"github.com/adnpa/IM/app/web/constant"
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (int64, error) {
	uid, ok := c.Get(constant.USER_ID_KEY)
	if !ok {
		return 0, fmt.Errorf("user id not exist")
	}
	if val, ok := uid.(int64); ok {
		return val, nil
	} else if val, ok := uid.(int); ok {
		return int64(val), nil
	} else if val, ok := uid.(string); ok {
		intUid, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		}
		return intUid, nil
	} else {
		return 0, fmt.Errorf("type convert fail")
	}
}
