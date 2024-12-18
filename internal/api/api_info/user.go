package api_info

import "github.com/adnpa/IM/pkg/pb/pb_ws"

type GetUsersInfoReq struct {
	OperationID string   `json:"operationID" binding:"required"`
	UserIDList  []string `json:"userIDList" binding:"required"`
}
type GetUsersInfoResp struct {
	CommResp
	UserInfoList []*pb_ws.PublicUserInfo  `json:"-"`
	Data         []map[string]interface{} `json:"data"`
}

type GetSelfUserInfoReq struct {
	OperationID string `json:"operationID" binding:"required"`
	UserID      string `json:"userID" binding:"required"`
}
type GetSelfUserInfoResp struct {
	CommResp
	UserInfo *pb_ws.UserInfo        `json:"-"`
	Data     map[string]interface{} `json:"data"`
}

type UpdateUserInfoReq struct {
	UserInfo
	OperationID string `json:"operationID" binding:"required"`
}

type UpdateUserInfoResp struct {
	CommResp
}
