package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/web/code"
	"github.com/adnpa/IM/app/web/constant"
	"github.com/adnpa/IM/app/web/global"
	"github.com/adnpa/IM/app/web/handler/forms"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	form := forms.RegisterForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, ErrInfo(code.ErrArgs))
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
		c.JSON(http.StatusOK, ErrInfo(code.TokenGenErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":    resp.Uid,
		"token":  token,
		"expire": expiredAt,
	})
}

func PasswordLogin(c *gin.Context) {
	form := forms.PwdLoginForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, ErrInfo(code.ErrArgs))
		return
	}

	// TODO: 验证码

	uResp, err := global.UserCli.GetUserByEmail(context.Background(), &pb.GetUserByEmailReq{Email: form.Email})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.UserNotExist))
		return
	}
	resp, err := global.UserCli.CheckPassWord(context.Background(), &pb.CheckPassWordReq{
		Password:          form.Password,
		EncryptedPassword: uResp.Usr.PassWord,
		Salt:              uResp.Usr.Salt})
	if err != nil || !resp.Match {
		c.JSON(http.StatusOK, ErrInfo(code.UserNotExist))
		return
	}

	token, expiredAt, err := utils.GenerateToken(strconv.FormatInt(int64(uResp.Usr.Id), 10))
	if err != nil {
		logger.Error("gen token error", zap.Error(err))
		c.JSON(http.StatusOK, ErrInfo(code.TokenGenErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     uResp.Usr.Id,
		"token":  token,
		"expire": expiredAt,
	})
}

func GetSelfProfile(c *gin.Context) {
	uid, _ := c.Get(constant.USER_ID_KEY)
	intUid, _ := strconv.ParseInt(uid.(string), 10, 32)
	resp, err := global.UserCli.GetUserById(context.Background(), &pb.GetUserByIdReq{Id: int32(intUid)})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrUnauthorized))
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"nickname": resp.Usr.Nickname,
			"gender":   resp.Usr.Sex,
			"avatar":   resp.Usr.Avatar,
			"birthday": resp.Usr.Birthday,
			"memo":     resp.Usr.Memo,
		})
}

func UpdateSelfProfile(c *gin.Context) {
	uid, _ := c.Get(constant.USER_ID_KEY)
	intUid, _ := strconv.ParseInt(uid.(string), 10, 32)

	form := forms.UpdateSelfProfileForm{}
	if err := c.ShouldBind(&form); err != nil {
		logger.Info("get arg fail", zap.Any("args", form))
		c.JSON(http.StatusOK, ErrInfo(code.ErrArgs))
		return
	}

	pbArgs := &pb.UpdateUserReq{Id: int32(intUid)}
	copier.Copy(pbArgs, form)

	logger.Info("", zap.Any("pb", pbArgs))
	_, err := global.UserCli.UpdateUser(context.Background(), pbArgs)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(code.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, "ok")
}
