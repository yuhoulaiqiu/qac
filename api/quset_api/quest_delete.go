package quset_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/jwts"
	"github.com/gin-gonic/gin"
)

type QuestDeleteRequest struct {
	Token   string `header:"Authorization"`
	QuestID uint   `header:"QuestID"`
}

func (QuestApi) QuestDelete(c *gin.Context) {
	var req QuestDeleteRequest
	if err := c.ShouldBindHeader(&req); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 解析 JWT token 来获取用户ID
	userID, err := jwts.ParseToken(req.Token)
	if err != nil {
		res.FailWithCode(res.UnauthorizedError, c)
		return
	}

	// 从数据库中获取问题
	var question models.QuestionModel
	err = global.DB.First(&question, "id = ?", req.QuestID).Error
	if err != nil {
		res.FailWithCode(res.NotFoundError, c)
		return
	}

	// 检查用户是否有权限删除这个问题
	if question.UserID != userID.UserID {
		res.FailWithCode(res.ForbiddenError, c)
		return
	}
	// 删除问题
	global.DB.Delete(&question)

	//删除问题的所有答案
	var answers []models.AnswerModel
	global.DB.Where("quest_id = ?", req.QuestID).Find(&answers)
	for _, answer := range answers {
		global.DB.Delete(&answer)
	}

	res.OkWithMessage("删除成功", c)
}
