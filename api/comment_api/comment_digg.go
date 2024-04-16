package comment_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	CommentLikeKey = "answer:like:"
	UserLikeKey    = "user:comment:like:"
)

type diggRequest struct {
	CommentID string `json:"comment_id"`
	UserID    string `json:"user_id"`
}

// CommentDigg 评论点赞
func (CommentApi) CommentDigg(c *gin.Context) {
	var req commentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	commentID := req.AnswerID
	userID := req.UserID
	userLikeKey := UserLikeKey + strconv.Itoa(int(userID))

	// 判断用户是否已经点赞
	if global.Redis.SIsMember(c, userLikeKey, commentID).Val() {
		res.FailWithCode(res.RepeatDiggError, c)
		return
	}

	// 将用户ID添加到评论的点赞用户集合中
	global.Redis.SAdd(c, userLikeKey, commentID)

	// 点赞
	key := CommentLikeKey + strconv.Itoa(int(commentID))
	err = global.Redis.Incr(c, key).Err()
	if err != nil {
		res.FailWithCode(res.RedisError, c)
		return
	}
	// 同步到数据库
	err = SyncCommentDiggCountToDB(c, strconv.Itoa(int(commentID)))
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	res.OkWithMessage("点赞成功", c)
}

// CommentUnDigg 取消点赞
func (CommentApi) CommentUnDigg(c *gin.Context) {
	var req diggRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	userLikeKey := fmt.Sprintf("%s%s", UserLikeKey, req.UserID)
	// 判断用户是否已经点赞
	if global.Redis.SIsMember(c, userLikeKey, req.CommentID).Val() {
		// 将用户ID从答案的点赞用户集合中移除
		global.Redis.SRem(c, userLikeKey, req.CommentID)

		key := fmt.Sprintf("%s%s", CommentLikeKey, req.CommentID)
		err = global.Redis.Decr(c, key).Err()
		if err != nil {
			res.FailWithCode(res.RedisError, c)
			return
		}

		// 同步到数据库
		err = SyncCommentDiggCountToDB(c, req.CommentID)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		res.OkWithMessage("取消点赞成功", c)

	} else {
		res.FailWithCode(res.InvalidOperationError, c)
		return
	}
}

// GetCommentDiggCount 获取答案的点赞数
func GetCommentDiggCount(c *gin.Context, commentID uint) (int, error) {
	key := CommentLikeKey + strconv.Itoa(int(commentID))
	countStr, err := global.Redis.Get(c, key).Result()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// SyncCommentDiggCountToDB 将Redis中的点赞数同步到数据库中
func SyncCommentDiggCountToDB(c *gin.Context, answerID string) error {
	questionUID, err := strconv.ParseUint(answerID, 10, 64)
	if err != nil {
		return err
	}
	count, err := GetCommentDiggCount(c, uint(questionUID))
	if err != nil {
		return err
	}
	// 更新数据库中的点赞数
	var comment models.CommentModel
	err = global.DB.First(&comment, "id = ?", answerID).Error
	if err != nil {
		return err
	}
	comment.DiggCount = count
	global.DB.Save(&comment)
	return nil
}
