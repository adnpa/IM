package dao

import (
	"context"
	"errors"
	"github.com/adnpa/IM/model"
	"github.com/adnpa/IM/pkg/common/constant"
	"github.com/adnpa/IM/query"
	"time"
)

func GetFriendsByUserUid(uid string) ([]*model.Friend, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return query.Friend.WithContext(ctx).Where(query.Friend.OwnerID.Eq(uid)).Find()
}

func IsFriend(uid1, uid2 string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	f, err := query.Friend.WithContext(ctx).
		Where(query.Friend.OwnerID.Eq(uid1), query.Friend.FriendID.Eq(uid2)).
		Or(query.Friend.OwnerID.Eq(uid2), query.Friend.FriendID.Eq(uid1)).First()
	if f != nil && err == nil {
		return true
	}

	return false
}

func AddFriend(fq *model.FriendRequest) error {
	if fq.Flag == constant.FriendRefuseFlag {
		return errors.New("add friend refuse")
	}
	return query.Q.Transaction(func(tx *query.Query) error {
		var err error
		_, err = tx.FriendRequest.Where(query.FriendRequest.UserID.Eq(fq.UserID), query.FriendRequest.ReqID.Eq(fq.ReqID)).Update(query.FriendRequest.Flag, fq.Flag)
		if err != nil {
			return err
		}

		f1 := &model.Friend{
			OwnerID:  fq.UserID,
			FriendID: fq.ReqID,
		}
		f2 := &model.Friend{
			OwnerID:  fq.ReqID,
			FriendID: fq.UserID,
		}
		f1.CreateTime = time.Now()
		f2.CreateTime = time.Now()
		err = tx.Friend.Create(f1)
		if err != nil {
			return err
		}
		err = tx.Friend.Create(f2)
		if err != nil {
			return err
		}

		return nil
	})
}

func DeleteFriend(userUid, friendUid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return query.Q.Transaction(func(tx *query.Query) error {
		var err error

		f, err := tx.Friend.
			Where(tx.Friend.OwnerID.Eq(userUid), tx.Friend.FriendID.Eq(friendUid)).
			Or(tx.Friend.OwnerID.Eq(friendUid), tx.Friend.FriendID.Eq(userUid)).Find()
		if err != nil {
			return err
		}
		_, err = tx.Friend.WithContext(ctx).Delete(f...)
		if err != nil {
			return err
		}

		fq, err := tx.FriendRequest.
			Where(tx.FriendRequest.UserID.Eq(userUid), tx.FriendRequest.ReqID.Eq(friendUid)).
			Or(tx.FriendRequest.UserID.Eq(friendUid), tx.FriendRequest.ReqID.Eq(userUid)).Find()
		if err != nil {
			return err
		}
		_, err = tx.FriendRequest.WithContext(ctx).Delete(fq...)
		return err
	})
}

func SetComment(uid, friendUid, comment string) error {
	friend := query.Friend
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := query.Friend.WithContext(ctx).
		Where(friend.OwnerID.Eq(uid), friend.FriendID.Eq(friendUid)).
		Or(friend.OwnerID.Eq(friendUid), friend.FriendID.Eq(uid)).
		Update(friend.Comment, comment)
	return err
}
