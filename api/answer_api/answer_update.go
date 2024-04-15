package answer_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Token    string `header:"Authorization"`
	AnswerID uint   `header:"AnswerID"`
	Content  string `header:"Content"`
}

func (AnswerApi) AnswerUpdate(c *gin.Context) {
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
	// 从数据库中获取答案
	var answer models.AnswerModel
	err = global.DB.First(&answer, "id = ?", req.AnswerID).Error
	if err != nil {
		res.FailWithCode(res.NotFoundError, c)
		return
	}
	// 检查用户是否有权限更新这篇答案
	if answer.UserID != userID.UserID {
		res.FailWithCode(res.ForbiddenError, c)
		return
	}

	// 更新答案
	answer.Content = req.Content
	global.DB.Save(&answer)
	res.OkWithMessage("更新成功", c)
}
