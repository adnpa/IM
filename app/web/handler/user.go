package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/global"
	"github.com/adnpa/IM/app/web/handler/forms"
	"github.com/adnpa/IM/internal/constant"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	form := forms.RegisterForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, ErrInfo(constant.ErrArgs))
		return
	}

	//TODO: validate code

	resp, err := global.UserCli.CreateUser(context.Background(), &pb.CreateUserReq{
		Nickname: form.Username,
		Mobile:   form.Mobile,
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	token, expiredAt, err := utils.GenerateToken(strconv.FormatInt(int64(resp.Uid), 10))
	if err != nil {
		logger.Error("gen token error", zap.Error(err))
		c.JSON(http.StatusOK, ErrInfo(constant.TokenGenErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":     resp.Uid,
		"token":  token,
		"expire": expiredAt,
	})
}

func PasswordLogin(c *gin.Context) {
	form := forms.PwdLoginForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, ErrInfo(constant.ErrArgs))
		return
	}

	// TODO: 验证码

	uResp, err := global.UserCli.GetUserByEmail(context.Background(), &pb.GetUserByEmailReq{Email: form.Email})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(constant.UserNotExist))
		return
	}
	resp, err := global.UserCli.CheckPassWord(context.Background(), &pb.CheckPassWordReq{
		Password:          form.Password,
		EncryptedPassword: uResp.Usr.PassWord,
		Salt:              uResp.Usr.Salt})
	if err != nil || !resp.Match {
		c.JSON(http.StatusOK, ErrInfo(constant.UserNotExist))
		return
	}

	token, expiredAt, err := utils.GenerateToken(strconv.FormatInt(int64(uResp.Usr.Id), 10))
	if err != nil {
		logger.Error("gen token error", zap.Error(err))
		c.JSON(http.StatusOK, ErrInfo(constant.TokenGenErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     uResp.Usr.Id,
		"token":  token,
		"expire": expiredAt,
	})
}

