// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRegister = "register"

// Register mapped from table <register>
type Register struct {
	Account  string `gorm:"column:account;primaryKey" json:"account"`
	Password string `gorm:"column:password;not null" json:"password"`
}

// TableName Register's table name
func (*Register) TableName() string {
	return TableNameRegister
}