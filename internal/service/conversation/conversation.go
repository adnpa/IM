package conversation

type Conversation struct {
	Id          int64
	LeftUserId  int64
	RightUserId int64
	MsgIds      []int64
	// UnreadNum int //未读数
}
