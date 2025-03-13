package handler

import (
	"strconv"

	"github.com/adnpa/IM/internal/service/user"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	mongodb.Insert("user", &user.User{Id: uid})
	// srv := &user.UserService{}
	// utils.NowSecond()
	// params := user.RegisterReq{}
	// if err := c.BindJSON(&params); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
	// 	return
	// }

	// srv.Register(params)
	// logger.Info("end", zap.Any("args", pbData))
}

func Login(c *gin.Context) {
	// srv := &user.UserService{}

	// params := user.LoginReq{}
	// if err := c.BindJSON(&params); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
	// 	return
	// }

	// srv.Login(params)
}
