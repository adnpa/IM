package main

import (
	"flag"
	"strconv"

	"github.com/adnpa/IM/app/web/handler"
	"github.com/adnpa/IM/app/web/initialize"
	"github.com/adnpa/IM/app/web/middlewares"
	"github.com/adnpa/IM/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitConfig()
	initialize.InitSrvConn()

	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())
	r.Use(cors.Default())

	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.POST("/register", handler.Register)
		userRouterGroup.POST("/pwd_login", handler.PasswordLogin)

		// authRouterGroup.GET("detail", middlewares.JWTAuth(), handler.GetUserDetail)
		// authRouterGroup.PATCH("update", middlewares.JWTAuth(), handler.UpdateUser)
	}

	friendGroup := r.Group("/friend")
	{
		friendGroup.GET("/friend_info_list", handler.GetFriendList)
		friendGroup.GET("/friend_detail", handler.GetFriendDetail)
		friendGroup.GET("/self_apply_list", handler.GetUserSelfApplyList)
		friendGroup.GET("/apply_list", handler.GetUserApplyList)
		friendGroup.POST("/apply_add_friend", handler.ApplyAddFriend)
		friendGroup.POST("/handle_apply", handler.HandleApplyFriend)
		friendGroup.POST("/delete_friend", handler.DeleteFriend)
	}

	// groupGroup := r.Group("/group")
	// {
	// TODO: 完成接口
	// groupGroup.GET("/group_info_list", handler.GetGroups)

	// groupGroup.GET("/self_apply_list", handler.GetGroupApplyList)
	// groupGroup.GET("/apply_list", handler.GetGroupApplyList)

	// groupGroup.POST("/create_group", handler.CreateGroup)
	// groupGroup.POST("/delete_group", handler.CreateGroup)

	// groupGroup.POST("/appoint", handler.ApplyGroup)
	// groupGroup.POST("/handle_apply", handler.ApplyGroup)
	// groupGroup.POST("/remove_member", handler.ApplyGroup)
	// groupGroup.POST("/block", handler.ApplyGroup)
	// }

	// TODO: 搜索服务
	// searchGroup := r.Group("/search")
	// {

	// }

	ginPort := flag.Int("port", 10000, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	r.Run(utils.ServerIP + ":" + strconv.Itoa(*ginPort))
}
