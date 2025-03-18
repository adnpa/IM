工作流
数据库设计 -> 生成结构 -> protobuf结构 -> 生成接口




  rpc GetFriendsByUserId(GetFriendsByUserIdReq) returns (GetFriendsByUserIdResp);
  rpc CreateFriend(CreateFriendReq) returns (CreateFriendResp);
  rpc DeleteFriend(DeleteFriendReq) returns (DeleteFriendResp);
  rpc UpdateFriend(UpdateFriendReq) returns (UpdateFriendResp);

  rpc GetFriendApply(GetFriendApplyReq) returns (GetFriendApplyResp);

  rpc GetFriendApplyByFromId(GetFriendApplyByFromIdReq) returns (GetFriendApplyByFromIdResp);
  rpc GetFriendApplyByToId(GetFriendApplyByToIdReq)
      returns (GetFriendApplyByToIdResp);

  rpc CreateFriendApply(CreateFriendApplyReq) returns (CreateFriendApplyResp);
  rpc UpdateFriendApply(UpdateFriendApplyReq) returns (UpdateFriendApplyResp);
  rpc DeleteFriendApply(DeleteFriendApplyReq) returns (DeleteFriendApplyResp);


go开发，使用gorm，protobuf，实现微服务

model
type Friendship struct {
	UserID    int32          `gorm:"column:user_id;primaryKey;comment:用户ID" json:"user_id"`                      // 用户ID
	FriendID  int32          `gorm:"column:friend_id;primaryKey;comment:好友ID" json:"friend_id"`                  // 好友ID
	Comment   string         `gorm:"column:comment;comment:备注" json:"comment"`                                   // 备注
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

type FriendApply struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true;comment:申请ID" json:"id"`             // 申请ID
	FromID      int32          `gorm:"column:from_id;not null;comment:申请者ID" json:"from_id"`                       // 申请者ID
	ToID        int32          `gorm:"column:to_id;not null;comment:被申请者ID" json:"to_id"`                          // 被申请者ID
	Status      int32          `gorm:"column:status;comment:申请状态" json:"status"`                                   // 申请状态
	ApplyReason string         `gorm:"column:apply_reason;comment:申请理由" json:"apply_reason"`                       // 申请理由
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

另一个微服务user的示例如下，用gorm结构增删改查
func (s *UserService) GetUserByMobile(_ context.Context, in *pb.GetUserByMobileReq) (*pb.GetUserByMobileResp, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: in.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return &pb.GetUserByMobileResp{
		Usr: Model2PB(user),
	}, nil
}
