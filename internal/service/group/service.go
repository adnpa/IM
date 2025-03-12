package group

import (
	"context"

	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type GroupService struct{}

func (s *GroupService) CreateGroup(info *GroupInfo) {
	mongodb.Insert("group_info", info)
}

func (s *GroupService) Disband(gid int64) {
	mongodb.Delete("group_info", bson.M{"id": gid})
}

func (s *GroupService) GroupApply() {}

type ApplyFlag int64

const (
	Rejected ApplyFlag = iota
	Accepted
)

func (s *GroupService) HandleApply(app *GroupApply) {
	switch app.Flag {
	case Rejected:
	case Accepted:
		mongodb.Insert("gorup_member", &GroupMember{app.GroupId, app.ApplyUserId, 0})
	}
}

func GetAllGrouUser(gid int64) []GroupMember {
	var results []GroupMember
	cur, _ := mongodb.GetAll("group_member", bson.M{"groupid": gid})
	cur.All(context.Background(), &results)
	return results
}
