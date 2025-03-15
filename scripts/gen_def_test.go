package scripts

import (
	"fmt"
	"testing"
)

func GenCRUD(name string) []string {
	return []string{
		fmt.Sprintf("Get%sById", name),
		fmt.Sprintf("Get%sByIds", name),
		fmt.Sprintf("Create%s", name),
		fmt.Sprintf("Update%s", name),
		fmt.Sprintf("Delete%s", name),
	}
}

func GenDef(names []string) {
	rpcTemp := "rpc %s(%s) returns (%s); "
	msgTemp := "message %s{\n}"
	rpcFuncList := []string{}
	msgList := []string{}

	for _, name := range names {
		rpcFuncList = append(rpcFuncList, fmt.Sprintf(rpcTemp, name, name+"Req", name+"Resp"))
		msgList = append(msgList, fmt.Sprintf(msgTemp, name+"Req"))
		msgList = append(msgList, fmt.Sprintf(msgTemp, name+"Resp"))
	}

	for _, fn := range rpcFuncList {
		fmt.Println(fn)
	}

	for _, msg := range msgList {
		fmt.Println(msg)
		fmt.Println()
	}
}

func TestGenDef(t *testing.T) {
	// names := []string{
	// 	"GetUserByPage",
	// 	"GetUserByMobile",
	// 	"GetUserByEmail",
	// 	"GetUserById",
	// 	"GetUserByIds",
	// 	"CreateUser",
	// 	"UpdateUser",
	// 	"DeleteUser",
	// 	"CheckPassWord",
	// }
	GenDef(GenCRUD("GroupApply"))
}
