package api_info

type CommResp struct {
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

type UserInfo struct {
	Uid    string `json:"uid" binding:"required,min=1,max=64"`
	Name   string `json:"name" binding:"omitempty,min=1,max=64"`
	Icon   string `json:"icon" binding:"omitempty,max=1024"`
	Gender int32  `json:"gender" binding:"omitempty,oneof=0 1 2"`
	Mobile string `json:"mobile" binding:"omitempty,max=32"`
	Birth  uint32 `json:"birth" binding:"omitempty"`
	Email  string `json:"email" binding:"omitempty,max=64"`
	Ex     string `json:"ex" binding:"omitempty,max=1024"`
}

type GroupAddMemberInfo struct {
	UserID    string `json:"userID" binding:"required"`
	RoleLevel int32  `json:"roleLevel" binding:"required"`
}
