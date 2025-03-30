package constant

type Role int

const (
	RoleOwner = iota + 1
	RoleManager
	RoleMember
)
