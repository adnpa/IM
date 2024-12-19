package dao

import (
	"context"
	"github.com/adnpa/IM/model"
	"github.com/adnpa/IM/query"
	"time"
)

func AddBlacklist(userUid, friendUid string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	b := &model.UserBlackList{OwnerID: userUid, BlockID: friendUid}
	err := query.UserBlackList.WithContext(ctx).Create(b)
	return err
}

func RemoveBlacklist(userUid, friendUid string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	_, err := query.UserBlackList.WithContext(ctx).Where(query.UserBlackList.OwnerID.Eq(userUid), query.UserBlackList.BlockID.Eq(friendUid)).Delete()
	return err
}

func IsInBlackList(userUid, friendUid string) bool {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	res, _ := query.UserBlackList.WithContext(ctx).Where(query.UserBlackList.OwnerID.Eq(userUid), query.UserBlackList.BlockID.Eq(friendUid)).First()
	return res == nil
}

func GetBlacklist(userUid string) ([]*model.UserBlackList, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	res, err := query.UserBlackList.WithContext(ctx).Where(query.UserBlackList.OwnerID.Eq(userUid)).Find()
	return res, err

}
