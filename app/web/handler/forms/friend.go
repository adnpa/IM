package forms

type DeleteFriendForm struct {
	FriendId int32 `json:"friend_id,omitempty"`
}

type GetFriendDetailForm struct {
	FriendId int32 `json:"friend_id,omitempty"`
}
