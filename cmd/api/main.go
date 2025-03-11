package main

import (
	"flag"
	"strconv"

	"github.com/adnpa/IM/internal/handler"
	"github.com/adnpa/IM/internal/middleware"
	"github.com/adnpa/IM/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	authRouterGroup := r.Group("/auth")
	{
		authRouterGroup.POST("/register", handler.Register)
		authRouterGroup.POST("/token", handler.Login)
	}

	friendGroup := r.Group("/friend")
	{
		friendGroup.POST("/add_friend", handler.AddFriend)
		friendGroup.GET("/get_friends", handler.GetFriendList)
	}

	ginPort := flag.Int("port", 10000, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	r.Run(utils.ServerIP + ":" + strconv.Itoa(*ginPort))
}
