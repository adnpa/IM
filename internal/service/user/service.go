package user

import (
	"fmt"
	"math/rand"

	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct{}

type RegisterReq struct {
	Nickname string
	Mobile   string
	Email    string
	Password string
}

func (s *UserService) Register(req RegisterReq) error {

	if mongodb.Exist("user", bson.M{"mobile": req.Mobile}) {
		return fmt.Errorf("registed already")
	}
	newUser := &User{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Passwd:   req.Password,
		Token:    fmt.Sprintf("%08d", rand.Int31()),
	}
	return mongodb.Insert("user", newUser)
}

type LoginReq struct {
	Mobile   string
	Password string
}
type LoginResp struct {
	Token      string
	ServiceUrl string
}

func (s *UserService) Login(req LoginReq) (*LoginResp, error) {
	user := &User{}
	err := mongodb.GetDecode("user", bson.M{"mobile": req.Mobile}, user)
	if err != nil {
		return nil, err
	}

	if req.Password != user.Passwd {
		return nil, fmt.Errorf("password not match")
	}

	// 刷新token,安全
	user.Token = fmt.Sprintf("%08d", rand.Int31())
	mongodb.Update("user", user)

	// srvUrl := discovery.GetService("chat")
	srvUrl := ""
	return &LoginResp{
		Token:      user.Token,
		ServiceUrl: srvUrl,
	}, nil
}
