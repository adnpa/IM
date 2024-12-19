package main

import (
	"flag"
	"github.com/adnpa/IM/internal/api/auth"
	"github.com/adnpa/IM/internal/api/chat"
	"github.com/adnpa/IM/internal/api/friend"
	"github.com/adnpa/IM/internal/api/user"
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

	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.POST("/get_self_user_info", user.GetSelfUserInfo)
		userRouterGroup.POST("/update_user_info", user.UpdateUserInfo)
		userRouterGroup.POST("/get_user_info", user.GetUserInfo)
	}

	friendRouterGroup := r.Group("/friend")
	{
		friendRouterGroup.POST("/add_friend", friend.AddFriend)
		friendRouterGroup.POST("/delete_friend", friend.DeleteFriend)
		friendRouterGroup.POST("/get_friend_apply_list", friend.GetFriendApplyList)
		friendRouterGroup.POST("/get_self_friend_apply_list", friend.GetSelfFriendApplyList)
		friendRouterGroup.POST("/get_friend_list", friend.GetFriendList)
		friendRouterGroup.POST("/add_friend_response", friend.AddFriendResponse)
		friendRouterGroup.POST("/set_friend_remark", friend.SetFriendRemark)

		friendRouterGroup.POST("/add_black", friend.AddBlack)
		friendRouterGroup.POST("/get_black_list", friend.GetBlacklist)
		friendRouterGroup.POST("/remove_black", friend.RemoveBlack)

		friendRouterGroup.POST("/import_friend", friend.ImportFriend)
		friendRouterGroup.POST("/is_friend", friend.IsFriend)
	}

	chatGroup := r.Group("/msg")
	{
		chatGroup.POST("/newest_seq", chat.GetSeq)
		chatGroup.POST("/send_msg", chat.SendMsg)
		chatGroup.POST("/pull_msg_by_seq", chat.PullMsgBySeqList)
		chatGroup.POST("/del_msg", chat.DelMsg)
	}

	//groupRouterGroup := r.Group("/group")
	//{
	//	groupRouterGroup.POST("/create_group", group.CreateGroup)                                   //1
	//	groupRouterGroup.POST("/set_group_info", group.SetGroupInfo)                                //1
	//	groupRouterGroup.POST("join_group", group.JoinGroup)                                        //1
	//	groupRouterGroup.POST("/quit_group", group.QuitGroup)                                       //1
	//	groupRouterGroup.POST("/group_application_response", group.ApplicationGroupResponse)        //1
	//	groupRouterGroup.POST("/transfer_group", group.TransferGroupOwner)                          //1
	//	groupRouterGroup.POST("/get_recv_group_applicationList", group.GetRecvGroupApplicationList) //1
	//	groupRouterGroup.POST("/get_user_req_group_applicationList", group.GetUserReqGroupApplicationList)
	//	groupRouterGroup.POST("/get_groups_info", group.GetGroupsInfo)                   //1
	//	groupRouterGroup.POST("/kick_group", group.KickGroupMember)                      //1
	//	groupRouterGroup.POST("/get_group_member_list", group.GetGroupMemberList)        //no use
	//	groupRouterGroup.POST("/get_group_all_member_list", group.GetGroupAllMemberList) //1
	//	groupRouterGroup.POST("/get_group_members_info", group.GetGroupMembersInfo)      //1
	//	groupRouterGroup.POST("/invite_user_to_group", group.InviteUserToGroup)          //1
	//	groupRouterGroup.POST("/get_joined_group_list", group.GetJoinedGroupList)        //1
	//}

	ginPort := flag.Int("port", 10006, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	r.Run(utils.ServerIP + ":" + strconv.Itoa(*ginPort))
}
