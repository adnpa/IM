package constant

const (
	FriendAgreeFlag  = 1
	FriendRefuseFlag = 0

	SingleChatType = 1
	GroupChatType  = 2

	//消息源（选项）
	IsHistory            = "history"
	IsPersistent         = "persistent"
	IsOfflinePush        = "offlinePush"
	IsUnreadCount        = "unreadCount"
	IsConversationUpdate = "conversationUpdate"
	IsSenderSync         = "senderSync"
)
