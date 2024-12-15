package constant

const (
	WSGetNewestSeq     = 1001 //获取最新消息
	WSPullMsg          = 1002 //获取消息
	WSSendMsg          = 1003 //发送消息
	WSPullMsgBySeqList = 1004 //通过请求列表拉取消息
	WSPushMsg          = 2001 // 发送消息
	WSDataError        = 3001
)
