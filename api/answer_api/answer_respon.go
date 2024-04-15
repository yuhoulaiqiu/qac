package answer_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"github.com/gin-gonic/gin"
)

func (AnswerApi) AnswerQuest(c *gin.Context) {
	var cr models.AnswerModel
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
	//问题表中对应问题回答量+1
	var quest models.QuestionModel
	if err = global.DB.First(&quest, "id = ?", cr.QuestID).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if err = global.DB.Model(&quest).Update("answer_count", quest.AnswerCount+1).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithMessage("回答发布成功", c)
}
