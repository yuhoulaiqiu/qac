package answer_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/jwts"
	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	Token    string `header:"Authorization"`
	AnswerID uint   `header:"AnswerID"`
}

func (AnswerApi) AnswerDelete(c *gin.Context) {
	var req DeleteRequest
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

	// 从数据库中获取答案
	var answer models.AnswerModel
	err = global.DB.First(&answer, "id = ?", req.AnswerID).Error
	if err != nil {
		res.FailWithCode(res.NotFoundError, c)
		return
	}

	// 检查用户是否有权限删除这篇答案
	if answer.UserID != userID.UserID {
		res.FailWithCode(res.ForbiddenError, c)
		return
	}
	// 删除答案
	global.DB.Delete(&answer)
	// 对应问题回答数减一
	var question models.QuestionModel
	err = global.DB.First(&question, "id = ?", answer.QuestID).Error
	if err != nil {
		res.FailWithCode(res.NotFoundError, c)
		return
	}
	question.AnswerCount -= 1
	global.DB.Save(&question)
	res.OkWithMessage("删除成功", c)
}
