package user

import (
	"fmt"
	"math/rand"

	"github.com/adnpa/IM/internal/handler/forms"
	"github.com/adnpa/IM/internal/model"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct{}

// todo 后面参数和返回都是proto定义的

func (s *UserService) CreateUser(req forms.RegisterForm) (*model.User, error) {
	if mongodb.Exist("user", bson.M{"mobile": req.Mobile}) {
		return nil, fmt.Errorf("registed already")
	}

	salt := utils.RandomSalt()
	pwd := utils.HashPassword(req.Password, salt)

	newUser := &model.User{
		Mobile: req.Mobile,
		Passwd: pwd,
		Token:  fmt.Sprintf("%08d", rand.Int31()),
	}

	err := mongodb.Insert("user", newUser)
	return newUser, err
}

func (s *UserService) GetUserByMobile(req forms.PwdLoginForm) (*model.User, error) {
	user := &model.User{}
	err := mongodb.GetDecode("user", bson.M{"mobile": req.Mobile}, user)
	if err != nil {
		return nil, err
	}

	if req.Password != user.Passwd {
		return nil, fmt.Errorf("password not match")
	}

	// 刷新token,安全
	// user.Token = fmt.Sprintf("%08d", rand.Int31())
	// mongodb.Update("user", user)

	// srvUrl := discovery.GetService("chat")
	srvUrl := ""
	return &LoginResp{
		Token:      user.Token,
		ServiceUrl: srvUrl,
	}, nil
}
