package dao

import (
	"context"
	"github.com/adnpa/IM/common/db/mysql/model"
	"github.com/adnpa/IM/common/db/mysql/query"
	"time"
)

func GetSelfApplyList(uid string) ([]*model.FriendRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return query.FriendRequest.WithContext(ctx).Where(query.FriendRequest.ReqID.Eq(uid)).Find()
}

func GetApplyList(uid string) ([]*model.FriendRequest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return query.FriendRequest.WithContext(ctx).Where(query.FriendRequest.UserID.Eq(uid)).Find()
}

func GetFriendReq(reqId, userId string) (*model.FriendRequest, error) {
	fq := query.FriendRequest
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return query.FriendRequest.WithContext(ctx).Where(fq.ReqID.Eq(reqId), fq.UserID.Eq(userId)).First()
}

func AddFriendRequest(fq *model.FriendRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := query.FriendRequest.WithContext(ctx).Create(fq)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFriendRequest(fq *model.FriendRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fq.UpdatedAt = time.Now()
	_, err := query.FriendRequest.WithContext(ctx).Where(query.FriendRequest.UserID.Eq(fq.UserID)).Updates(fq)
	if err != nil {
		return err
	}
	return nil
}
