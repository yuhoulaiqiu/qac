package jwts

import (
	"QAComm/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// GetToken 创建token
func GetToken(user JwtPayLoad) (string, error) {
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(global.Config.Jwt.Secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func GetRefreshToken(user JwtPayLoad) (string, error) {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(global.Config.Jwt.Secret))
	if err != nil {
		return "", err
	}
	return rt, err
}
