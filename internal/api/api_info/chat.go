package api_info

import "github.com/adnpa/IM/pkg/pb/pb_ws"

type ParamsUserNewestSeq struct {
	ReqIdentifier int    `json:"reqIdentifier" binding:"required"`
	SendID        string `json:"sendID" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
	MsgIncr       int    `json:"msgIncr" binding:"required"`
}

type ParamsUserPullMsg struct {
	ReqIdentifier *int   `json:"reqIdentifier" binding:"required"`
	SendID        string `json:"sendID" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
	Data          struct {
		SeqBegin *int64 `json:"seqBegin" binding:"required"`
		SeqEnd   *int64 `json:"seqEnd" binding:"required"`
	}
}

type ParamsUserPullMsgBySeqList struct {
	ReqIdentifier int      `json:"reqIdentifier" binding:"required"`
	SendID        string   `json:"sendID" binding:"required"`
	OperationID   string   `json:"operationID" binding:"required"`
	SeqList       []uint32 `json:"seqList"`
}

type ParamsUserSendMsg struct {
	SenderPlatformID int32  `json:"senderPlatformID" binding:"required"`
	SendID           string `json:"sendID" binding:"required"`
	SenderNickName   string `json:"senderNickName"`
	SenderFaceURL    string `json:"senderFaceUrl"`
	OperationID      string `json:"operationID" binding:"required"`
	Data             struct {
		SessionType int32                  `json:"sessionType" binding:"required"`
		MsgFrom     int32                  `json:"msgFrom" binding:"required"`
		ContentType int32                  `json:"contentType" binding:"required"`
		RecvID      string                 `json:"recvID" `
		GroupID     string                 `json:"groupID" `
		ForceList   []string               `json:"forceList"`
		Content     []byte                 `json:"content" binding:"required"`
		Options     map[string]bool        `json:"options" `
		ClientMsgID string                 `json:"clientMsgID" binding:"required"`
		CreateTime  int64                  `json:"createTime" binding:"required"`
		OffLineInfo *pb_ws.OfflinePushInfo `json:"offlineInfo" `
	}
}
