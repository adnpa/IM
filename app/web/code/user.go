package code

const (
	ErrUserNotFound = iota + 10000
	UserExist
	UserNotExist
	TokenGenErr
	PasswordNotMatch
)
