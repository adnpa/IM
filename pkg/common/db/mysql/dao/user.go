package dao

import (
	"context"
	"github.com/adnpa/IM/model"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/adnpa/IM/query"
	"go.uber.org/zap"
	"time"
)

func GetAllUserUid() []string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var res []string
	li, err := query.User.WithContext(ctx).Select(query.User.UID).Find()
	if err != nil {
		return nil
	}
	for _, i := range li {
		res = append(res, i.UID)
	}
	return res
}

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
	user.CreateTime = time.Now()
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

	//user.UpdatedAt = time.Now()
	logger.L().Info("db", zap.Any("", user.Name))
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
