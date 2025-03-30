package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/adnpa/IM/app/web/handler"
	"github.com/adnpa/IM/app/web/initialize"
	"github.com/adnpa/IM/app/web/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitConfig()
	initialize.InitSrvConn()

	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRouterGroup := r.Group("/auth")
	{
		authRouterGroup.POST("/register", handler.Register)
		authRouterGroup.POST("/pwd_login", handler.PasswordLogin)
	}

	userRouterGroup := r.Group("/user", middlewares.JWTAuth())
	{
		userRouterGroup.GET("self_profile", handler.GetSelfProfile)
		userRouterGroup.POST("update_self_profile", handler.UpdateSelfProfile)
		// authRouterGroup.GET("detail", middlewares.JWTAuth(), handler.GetUserDetail)
		// authRouterGroup.PATCH("update", middlewares.JWTAuth(), handler.UpdateUser)
	}

	ossRouterGroup := r.Group("/oss", middlewares.JWTAuth())
	{
		ossRouterGroup.POST("/upload", handler.Upload)
		// ossRouterGroup.POST("/download")
	}

	friendGroup := r.Group("/friend", middlewares.JWTAuth())
	{
		friendGroup.GET("/info_list", handler.GetFriendList)
		friendGroup.GET("/detail", handler.GetFriendDetail)
		friendGroup.GET("/self_apply_list", handler.GetUserSelfApplyList)
		friendGroup.POST("/apply_add_friend", handler.ApplyAddFriend)
		friendGroup.POST("/handle_apply", handler.HandleApplyFriend)
		friendGroup.POST("/delete_friend", handler.DeleteFriend)
	}

	groupGroup := r.Group("/group", middlewares.JWTAuth())
	{
		// TODO: 完成接口
		groupGroup.GET("/info_list", handler.GetUserGroups)

		// groupGroup.GET("/self_apply_list", handler.GetGroupApplyList)
		// groupGroup.GET("/apply_list", handler.GetGroupApplyList)

		// groupGroup.POST("/create_group", handler.CreateGroup)
		// groupGroup.POST("/delete_group", handler.CreateGroup)

		// groupGroup.POST("/appoint", handler.ApplyGroup)
		// groupGroup.POST("/handle_apply", handler.ApplyGroup)
		// groupGroup.POST("/remove_member", handler.ApplyGroup)
		// groupGroup.POST("/block", handler.ApplyGroup)
	}

	// TODO: 搜索服务
	// /search/friend?id=101
	searchGroup := r.Group("/search")
	{
		searchGroup.GET("/friend", handler.SearchFriend)
	}

	ginPort := flag.Int("port", 10000, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	r.Run("localhost:" + strconv.Itoa(*ginPort))
}
