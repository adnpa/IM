package api_info

import "github.com/adnpa/IM/pkg/pb/pb_ws"

type ParamsCommFriend struct {
	OperationID string `json:"operationID" binding:"required"`
	ToUserID    string `json:"toUserID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"`
}

//===================================

//type GetFriendsInfoReq struct {
//	ParamsCommFriend
//}
//type GetFriendsInfoResp struct {
//	CommResp
//	FriendInfoList []*pb_ws.FriendInfo      `json:"-"`
//	Data           []map[string]interface{} `json:"data"`
//}

type GetFriendListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"`
}
type GetFriendListResp struct {
	CommResp
	FriendInfoList []*pb_ws.FriendInfo      `json:"-"`
	Data           []map[string]interface{} `json:"data"`
}

type GetFriendApplyListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"`
}
type GetFriendApplyListResp struct {
	CommResp
	FriendRequestList []*pb_ws.FriendRequest   `json:"-"`
	Data              []map[string]interface{} `json:"data"`
}

type GetSelfApplyListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"`
}
type GetSelfApplyListResp struct {
	CommResp
	FriendRequestList []*pb_ws.FriendRequest   `json:"-"`
	Data              []map[string]interface{} `json:"data"`
}

type AddFriendReq struct {
	ParamsCommFriend
	ReqMsg string `json:"reqMsg"`
}
type AddFriendResp struct {
	CommResp
}

type AddFriendResponseReq struct {
	ParamsCommFriend
	Flag      int32  `json:"flag" binding:"required,oneof=-1 0 1"`
	HandleMsg string `json:"handleMsg"`
}
type AddFriendResponseResp struct {
	CommResp
}

type DeleteFriendReq struct {
	ParamsCommFriend
}
type DeleteFriendResp struct {
	CommResp
}

type SetFriendRemarkReq struct {
	ParamsCommFriend
	Remark string `json:"remark" binding:"required"`
}
type SetFriendRemarkResp struct {
	CommResp
}

type GetBlackListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"`
}
type GetBlackListResp struct {
	CommResp
	BlackUserInfoList []*pb_ws.PublicUserInfo  `json:"-"`
	Data              []map[string]interface{} `json:"data"`
}

type AddBlacklistReq struct {
	ParamsCommFriend
}
type AddBlacklistResp struct {
	CommResp
}

type RemoveBlackListReq struct {
	ParamsCommFriend
}
type RemoveBlackListResp struct {
	CommResp
}

type ImportFriendReq struct {
	FriendUserIDList []string `json:"friendUserIDList" binding:"required"`
	OperationID      string   `json:"operationID" binding:"required"`
	FromUserID       string   `json:"fromUserID" binding:"required"`
}
type ImportFriendResp struct {
	CommResp
	UserIDResultList []UserIDResult `json:"data"`
}
type UserIDResult struct {
	UserID string `json:"userID""`
	Result int32  `json:"result"`
}

//type PublicUserInfo struct {
//	UserID   string `json:"userID"`
//	Nickname string `json:"nickname"`
//	FaceUrl  string `json:"faceUrl"`
//	Gender   int32  `json:"gender"`
//}

type IsFriendReq struct {
	ParamsCommFriend
}
type IsFriendResp struct {
	CommResp
	Response Response `json:"data"`
}
type Response struct {
	Friend bool `json:"isFriend"`
}
