package handler

import (
	"net/http"
	"strconv"

	"github.com/adnpa/IM/internal/constant"
	"github.com/adnpa/IM/internal/handler/forms"
	"github.com/adnpa/IM/internal/service/user"
	"github.com/adnpa/IM/internal/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	form := forms.RegisterForm{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, ErrInfo(constant.ErrArgs))
		return
	}

	//todo validate code

	srv := &user.UserService{}
	user, err := srv.CreateUser(form)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(constant.UserExist))
		return
	}

	token, expiredAt, err := utils.GenerateToken(strconv.FormatInt(user.Id, 10))
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(constant.TokenGenErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":     user.Id,
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

	srv := &user.UserService{}
	user, err := srv.GetUserByMobile(form)
	if err != nil {
		c.JSON(http.StatusOK, ErrInfo(constant.UserNotExist))
		return
	}

	if utils.DoPasswordsMatch(user.Passwd, form.Password, []byte(user.Salt)) {
		token, expiredAt, err := utils.GenerateToken(strconv.FormatInt(user.Id, 10))
		if err != nil {
			c.JSON(http.StatusOK, ErrInfo(constant.TokenGenErr))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":     user.Id,
			"token":  token,
			"expire": expiredAt,
		})
	}
	c.JSON(http.StatusOK, ErrInfo(constant.PasswordNotMatch))
}

// helper
