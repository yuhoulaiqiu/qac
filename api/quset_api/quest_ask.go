package quset_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"github.com/gin-gonic/gin"
)

func (QuestApi) PutQuestion(c *gin.Context) {
	var cr models.QuestionModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//入库
	if err = global.DB.Create(&cr).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("问题发布成功", c)
}
