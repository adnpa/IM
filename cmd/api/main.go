package main

import (
	"flag"
	"strconv"

	"github.com/adnpa/IM/internal/handler"
	"github.com/adnpa/IM/internal/middleware"
	"github.com/adnpa/IM/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(cors.Default())

	authRouterGroup := r.Group("/auth")
	{
		authRouterGroup.POST("/register", handler.Register)
		authRouterGroup.POST("/token", handler.Login)
	}

	friendGroup := r.Group("/friend")
	{
		friendGroup.GET("/get_friends", handler.GetFriendList)
		friendGroup.POST("/add_friend", handler.AddFriend)
	}

	groupGroup := r.Group("/group")
	{
		groupGroup.GET("/get_groups", handler.GetGroups)
		groupGroup.POST("/create_group", handler.CreateGroup)
		groupGroup.GET("/apply_group", handler.ApplyGroup)
	}

	ginPort := flag.Int("port", 10000, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	r.Run(utils.ServerIP + ":" + strconv.Itoa(*ginPort))
}
