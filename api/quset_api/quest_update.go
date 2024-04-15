package quset_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Token   string `header:"Authorization"`
	QuestID uint   `header:"QuestID"`
	Title   string `header:"Title"`
	Content string `header:"Content"`
}

func (QuestApi) QuestUpdate(c *gin.Context) {
	var req UpdateRequest
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
	// 检查用户是否有权限更新这个问题
	if question.UserID != userID.UserID {
		res.FailWithCode(res.ForbiddenError, c)
		return
	}
	// 更新问题
	question.Title = req.Title
	question.Content = req.Content
	global.DB.Save(&question)
	res.OkWithMessage("更新成功", c)
}
