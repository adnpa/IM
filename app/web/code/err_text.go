package code

type Err map[string]any

func ErrInfo(code int) Err {
	return Err{"errCode": code, "errMsg": StatusText(code)}
}

func StatusText(code int) string {
	switch code {
	case NoError:
		return ""
	case ErrArgs:
		return "Args Input Error"
	case ErrInternal:
		return "Internal Server Error, please try later"
	case ErrUnauthorized:
		return ""
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
