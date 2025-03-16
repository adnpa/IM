package group

import (
	"context"

	"github.com/adnpa/IM/pkg/common/pb"
)

// s *GroupService pb.GroupServer

type GroupService struct{}

// 群聊基础信息管理
func (s *GroupService) GetGroupInfoById(_ context.Context, _ *pb.GetGroupInfoByIdReq) (*pb.GetGroupInfoByIdResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) GetGroupInfoByIds(_ context.Context, _ *pb.GetGroupInfoByIdsReq) (*pb.GetGroupInfoByIdsResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) CreateGroupInfo(_ context.Context, _ *pb.CreateGroupInfoReq) (*pb.CreateGroupInfoResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) UpdateGroupInfo(_ context.Context, _ *pb.UpdateGroupInfoReq) (*pb.UpdateGroupInfoResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) DeleteGroupInfo(_ context.Context, _ *pb.DeleteGroupInfoReq) (*pb.DeleteGroupInfoResp, error) {
	panic("not implemented") // TODO: Implement
}

// 成员管理
func (s *GroupService) GetGroupMemberById(_ context.Context, _ *pb.GetGroupMemberByIdReq) (*pb.GetGroupMemberByIdResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) GetGroupMemberByIds(_ context.Context, _ *pb.GetGroupMemberByIdsReq) (*pb.GetGroupMemberByIdsResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) CreateGroupMember(_ context.Context, _ *pb.CreateGroupMemberReq) (*pb.CreateGroupMemberResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) UpdateGroupMember(_ context.Context, _ *pb.UpdateGroupMemberReq) (*pb.UpdateGroupMemberResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) DeleteGroupMember(_ context.Context, _ *pb.DeleteGroupMemberReq) (*pb.DeleteGroupMemberResp, error) {
	panic("not implemented") // TODO: Implement
}

// 申请管理
func (s *GroupService) GetGroupApplyById(_ context.Context, _ *pb.GetGroupApplyByIdReq) (*pb.GetGroupApplyByIdResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) GetGroupApplyByIds(_ context.Context, _ *pb.GetGroupApplyByIdsReq) (*pb.GetGroupApplyByIdsResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) CreateGroupApply(_ context.Context, _ *pb.CreateGroupApplyReq) (*pb.CreateGroupApplyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) UpdateGroupApply(_ context.Context, _ *pb.UpdateGroupApplyReq) (*pb.UpdateGroupApplyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) DeleteGroupApply(_ context.Context, _ *pb.DeleteGroupApplyReq) (*pb.DeleteGroupApplyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GroupService) mustEmbedUnimplementedGroupServer() {
	panic("not implemented") // TODO: Implement
}

// func (s *GroupService) CreateGroup(info *GroupInfo) {
// 	mongodb.Insert("group_info", info)
// }

// func (s *GroupService) Disband(gid int64) {
// 	mongodb.Delete("group_info", bson.M{"id": gid})
// }

// func (s *GroupService) GroupApply() {}

// type ApplyFlag int64

// const (
// 	Rejected ApplyFlag = iota
// 	Accepted
// )

// func (s *GroupService) HandleApply(app *GroupApply) {
// 	switch app.Flag {
// 	case Rejected:
// 	case Accepted:
// 		mongodb.Insert("gorup_member", &GroupMember{app.GroupId, app.ApplyUserId, 0})
// 	}
// }

// func GetAllGrouUser(gid int64) []GroupMember {
// 	var results []GroupMember
// 	cur, _ := mongodb.GetAll("group_member", bson.M{"groupid": gid})
// 	cur.All(context.Background(), &results)
// 	return results
// }
