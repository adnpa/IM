package constant

// 业务错误码

type Err map[string]any

func ErrInfo(code int) Err {
	return Err{"errCode": code, "errMsg": StatusText(code)}
}

const (
	NoError     = 0
	ErrArgs     = 1001
	ErrInternal = 1002

	//user
	ErrUserArgs = 10001

	//Friend

	//chat
	ErrChatKafkaSend      = 30001
	ErrChatUnknownMsgType = 30002
	ErrChatMsgTimeout     = 30003
)

func StatusText(code int) string {
	switch code {
	case NoError:
		return ""
	case ErrArgs:
		return "Args Input Error"
	case ErrInternal:
		return "Internal Server Error, please try later"
	case ErrChatKafkaSend:
		return "Server Error: kafka send msg error"
	case ErrChatUnknownMsgType:
		return "Unknown Msg Type"
	case ErrChatMsgTimeout:
		return "Msg has expired, pull fail"
	default:
		return ""
	}
}
