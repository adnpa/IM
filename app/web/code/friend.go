package code

const (
	ErrNotFriend = iota + 200000
	ErrAlreadyFriend

	ErrChatKafkaSend      = 30001
	ErrChatUnknownMsgType = 30002
	ErrChatMsgTimeout     = 30003
)
