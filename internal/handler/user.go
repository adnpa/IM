package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/adnpa/IM/internal/constant"
	"github.com/adnpa/IM/internal/handler/forms"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/common/logger"
	"github.com/adnpa/IM/pkg/common/pb"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Register(c *gin.Context) {
	form := forms.RegisterForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, ErrInfo(constant.ErrArgs))
		return
	}

	//todo validate code

	// todo 简化--------------------------------------------------
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cli := pb.NewUserClient(conn)

	resp, err := cli.CreateUser(context.Background(), &pb.CreateUserReq{
		Nickname: form.Username,
		Mobile:   form.Mobile,
		Email:    form.Email,
		Password: form.Password,
	})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// --------------------------------------------------------

	token, expiredAt, err := utils.GenerateToken(strconv.FormatInt(resp.Uid, 10))
	if err != nil {
		logger.Error("gen token error", zap.Error(err))
		c.JSON(http.StatusOK, ErrInfo(constant.TokenGenErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     resp.Uid,
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

	// todo 验证码

	// todo 简化--------------------------------------------------
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cli := pb.NewUserClient(conn)

	uResp, err := cli.GetUserByEmail(context.Background(), &pb.GetUserByEmailReq{Email: form.Email})
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(constant.UserNotExist))
		return
	}
	resp, err := cli.CheckPassWord(context.Background(), &pb.CheckPassWordReq{
		Password:          form.Password,
		EncryptedPassword: uResp.Usr.PassWord,
		Salt:              uResp.Usr.Salt})
	if err != nil || !resp.Match {
		c.JSON(http.StatusOK, ErrInfo(constant.UserNotExist))
		return
	}
	// --------------------------------------------------------

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

// helper
