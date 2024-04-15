package user_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/pwd"
	"github.com/gin-gonic/gin"
)

func (UserApi) Register(c *gin.Context) {
	var cr models.UserModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ?", cr.UserName).Error
	if err == nil {
		res.FailWithMessage("用户名重复", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	//入库
	err = global.DB.Create(&models.UserModel{
		UserName: cr.UserName,
		Password: hashPwd,
	}).Error
	if err != nil {
		res.FailWithMessage("注册失败", c)
		return
	}
	res.OkWithMessage("注册成功", c)
}
