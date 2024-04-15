package user_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/jwts"
	"QAComm/utils/pwd"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.StandardClaims
}

func (UserApi) Login(c *gin.Context) {
	var cr LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ?", cr.UserName).Error
	if err != nil {
		// 即未找到
		global.Log.Warn("用户名错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	//校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	//生成token
	token, err := jwts.GetToken(jwts.JwtPayLoad{
		UserName: userModel.UserName,
		UserID:   userModel.ID,
	})
	//生成refreshToken
	rToken, err := jwts.GetRefreshToken(jwts.JwtPayLoad{
		UserName: userModel.UserName,
		UserID:   userModel.ID,
	})
	tx := global.DB.Begin()
	// 更新 token 和 refreshToken
	userModel.Token = token
	userModel.RefreshToken = rToken
	if err := tx.Save(&userModel).Error; err != nil {
		tx.Rollback()
		// 处理错误
		return
	}
	tx.Commit()
	res.OkWithData(map[string]string{
		"token":        token,
		"refreshToken": rToken,
	}, c)
}
