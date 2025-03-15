package constant

type ErrCode int

// 公共 1000-1999
const (
	NoError ErrCode = 0
	ErrArgs ErrCode = 1000 + iota
	ErrInternal
	ErrDatabase
	ErrSendCode
	SendCodeTooFrequent
)

// 2000 - 2099
const (
	UserExist ErrCode = 2000 + iota
	UserNotExist
	TokenGenErr
	PasswordNotMatch
)

func StatusText(code ErrCode) string {
	switch code {
	case NoError:
		return ""
	case ErrArgs:
		return "Args Input Error"
	case ErrDatabase:
		return "Database Error"
	case ErrInternal:
		return "Internal Server Error, please try later"
	case UserExist:
		return "User Exist"
	case UserNotExist:
		return "User Not Exist"
	case TokenGenErr:
		return "Gen Token Err"
	case PasswordNotMatch:
		return "Password Not Match"
	case ErrSendCode:
		return "Send Code Fail"
	case SendCodeTooFrequent:
		return "Send Code Too Frequent"
	default:
		return ""
	}
}
