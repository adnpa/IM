package friend

import "github.com/adnpa/IM/pkg/common/db/mongodb"

type FriendService struct{}

type AddFriendReq struct {
	OwnId int64 `json:"own_id"`
	FriId int64 `json:"fri_id"`
}

func (s *FriendService) FriendApply(req AddFriendReq) bool {
	err1 := mongodb.Insert("friend", &Friend{
		OwnerID:  req.OwnId,
		FriendID: req.FriId,
	})
	err2 := mongodb.Insert("friend", &Friend{
		OwnerID:  req.FriId,
		FriendID: req.OwnId,
	})
	return err1 == nil && err2 == nil
}

func (s *FriendService) GetFriend() {

}
