package main

import (
	"flag"
	"github.com/adnpa/IM/internal/api/auth"
	"github.com/adnpa/IM/internal/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	authRouterGroup := r.Group("/auth")
	{
		authRouterGroup.POST("/register", auth.UserRegister)
		authRouterGroup.POST("/token", auth.UserToken)
	}
	ginPort := flag.Int("port", 10000, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	r.Run(utils.ServerIP + ":" + strconv.Itoa(*ginPort))
}
