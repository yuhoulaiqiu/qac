package answer_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"github.com/gin-gonic/gin"
)

type commentRequest struct {
	AnswerID        uint   `json:"answer_id"`
	UserID          uint   `json:"user_id"`
	Content         string `json:"content"`
	ParentCommentID *uint  `json:"parent_comment_id,omitempty"` // 可选的父评论ID
}

// CommentOnAnswer 评论回答
func (AnswerApi) CommentOnAnswer(c *gin.Context) {
	var req commentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	answerID := req.AnswerID
	userID := req.UserID
	content := req.Content
	parentCommentID := req.ParentCommentID

	// 创建新的评论
	comment := models.CommentModel{
		UserID:          userID,
		AnswerID:        answerID,
		Content:         content,
		ParentCommentID: parentCommentID, // 设置父评论ID
	}

	err = global.DB.Create(&comment).Error
	if err != nil {
		res.FailWithCode(res.DBError, c)
		return
	}
	res.OkWithMessage("评论成功", c)
}
