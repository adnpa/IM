package dao

import (
	"context"
	"github.com/adnpa/IM/common/db/mysql/model"
	"github.com/adnpa/IM/common/db/mysql/query"
	"time"
)

func GetUserByUid(uid string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	first, err := query.User.WithContext(ctx).Where(query.User.UID.Eq(uid)).First()
	if err != nil {
		return nil, err
	}
	return first, err
}

func GetUsersByUidL(uidL []string) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return query.User.WithContext(ctx).Where(query.User.UID.In(uidL...)).Find()
}

func CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//todo password
	//user.Password = utils.EncryptPassword([]byte(user))
	err := query.User.WithContext(ctx).Create(user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	user.UpdatedAt = time.Now()
	_, err := query.User.WithContext(ctx).Where(query.User.UID.Eq(user.UID)).Updates(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := query.User.WithContext(ctx).Where(query.User.UID.Eq(user.UID)).Delete(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUsers(uidL []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := query.User.WithContext(ctx).Where(query.User.UID.In(uidL...)).Delete()
	if err != nil {
		return err
	}
	return nil
}
