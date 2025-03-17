package errors

// 统一封装错误
// https://github.com/pkg/errors



// type withCode struct {
// 	err   error
// 	code  int
// 	cause error
// 	// *stack
// }

// func WithCode(code int, format string, args ...interface{}) error {
// 	err := errors.New("fjas")
// 	errors.Parse(err)
// return &withCode{
// }
// }

// func StatusText(code int) string {
// 	switch code {
// 	case NoError:
// 		return ""
// 	case ErrArgs:
// 		return "Args Input Error"
// 	case ErrInternal:
// 		return "Internal Server Error, please try later"
// 	case ErrChatKafkaSend:
// 		return "Server Error: kafka send msg error"
// 	case ErrChatUnknownMsgType:
// 		return "Unknown Msg Type"
// 	case ErrChatMsgTimeout:
// 		return "Msg has expired, pull fail"
// 	default:
// 		return ""
// 	}
// }
