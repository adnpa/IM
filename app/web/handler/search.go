package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/code"
	"github.com/adnpa/IM/app/web/global"
	"github.com/gin-gonic/gin"
)

func SearchFriend(c *gin.Context) {
	fmt.Println("search friend")
	friendId := c.Query("id")
	intFriendId, err := strconv.ParseInt(friendId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrArgs))
		return
	}

	resp, err := global.UserCli.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: int32(intFriendId)})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}

	c.JSON(http.StatusOK, resp.Usr)
}
