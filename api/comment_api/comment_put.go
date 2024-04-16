package comment_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentRequest struct {
	AnswerID        uint   `json:"answer_id"`
	UserID          uint   `json:"user_id"`
	Content         string `json:"content"`
	ParentCommentID *uint  `json:"parent_comment_id,omitempty"` // 可选的父评论ID
}

// Comment 评论
func (CommentApi) Comment(c *gin.Context) {
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
	//如果有父评论，将父评论的评论数加1
	if parentCommentID != nil {
		err = global.DB.Model(&models.CommentModel{}).Where("id = ?", parentCommentID).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			res.FailWithCode(res.DBError, c)
			return
		}
	}

	err = global.DB.Create(&comment).Error
	if err != nil {
		res.FailWithCode(res.DBError, c)
		return
	}
	res.OkWithMessage("评论成功", c)
}
