package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt的payload数据
type JwtPayLoad struct {
	UserName string `json:"user_name"` //用户名
	UserID   uint   `json:"user_id"`   //用户id

}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
