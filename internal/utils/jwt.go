package utils

import (
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//https://oauth.net/2/

type Claims struct {
	UID      string
	Platform int32
	jwt.RegisteredClaims
}

// 获取密钥的回调函数
func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) { return []byte(config.Config.Jwt.Secret), nil }
}

func BuildClaims(uid string, platform int32, ttl int64) Claims {
	claims := Claims{
		UID:      uid,
		Platform: platform,
	}

	if ttl > 0 {
		claims.RegisteredClaims = jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl) * time.Second)),
		}
	}

	return claims
}

func GenerateToken(uid string, platform int32) (string, int64, error) {
	claims := BuildClaims(uid, platform, config.Config.Jwt.Expire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.Config.Jwt.Secret)
	return tokenString, config.Config.Jwt.Expire, err
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, secret(), jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	//  1 check uid and platform

	//	2 multi login policy

	//仅一端登录
	//允许多端登录

	return nil, nil
}

func GetUserId(token string) (string, error) {
	c, err := ParseToken(token)
	if err != nil {
		return "", err
	}
	return c.UID, nil
}

func VerifyToken(tokenString, uid string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, secret(), jwt.WithLeeway(5*time.Second))
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(*Claims); ok && claims.UID == uid {
		return true
	} else {
		return false
	}
}
