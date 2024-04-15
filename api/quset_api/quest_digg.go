package quset_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	QuestionLikeKey = "question:like:"
	UserLikeKey     = "user:quest:like:"
)

type diggRequest struct {
	QuestID string `json:"quest_id"`
	UserID  string `json:"user_id"`
}

func (QuestApi) QuestDigg(c *gin.Context) {
	var req diggRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return

	}
	questionID := req.QuestID
	userID := req.UserID
	userLikeKey := UserLikeKey + userID

	// 判断用户是否已经点赞
	if global.Redis.SIsMember(c, userLikeKey, questionID).Val() {
		res.FailWithCode(res.RepeatDiggError, c)
		return
	}

	// 将用户ID添加到问题的点赞用户集合中
	global.Redis.SAdd(c, userLikeKey, questionID)

	// 点赞
	key := QuestionLikeKey + questionID
	err = global.Redis.Incr(c, key).Err()
	if err != nil {
		res.FailWithCode(res.RedisError, c)
		return
	}
	// 同步到数据库
	err = SyncQuestionLikeCountToDB(c, questionID)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	res.OkWithMessage("点赞成功", c)
}

func (QuestApi) QuestUnDigg(c *gin.Context) {
	var req diggRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return

	}
	questionID := req.QuestID
	userID := req.UserID

	userLikeKey := UserLikeKey + userID
	// 判断用户是否已经点赞
	if global.Redis.SIsMember(c, userLikeKey, questionID).Val() {
		// 将用户ID从问题的点赞用户集合中移除
		global.Redis.SRem(c, userLikeKey, questionID)

		key := QuestionLikeKey + questionID
		err := global.Redis.Decr(c, key).Err()
		if err != nil {
			res.FailWithCode(res.RedisError, c)
			return
		}
		// 同步到数据库
		err = SyncQuestionLikeCountToDB(c, questionID)
		if err != nil {
			res.FailWithCode(res.RedisError, c)
			return
		}
		res.OkWithMessage("取消点赞成功", c)
	} else {
		res.FailWithCode(res.InvalidOperationError, c)
		return
	}

}

// GetQuestionLikeCount 获取问题的点赞数
func GetQuestionLikeCount(c *gin.Context, questionID uint) (int, error) {
	key := QuestionLikeKey + strconv.Itoa(int(questionID))
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

// SyncQuestionLikeCountToDB 将Redis中的点赞数同步到数据库中
func SyncQuestionLikeCountToDB(c *gin.Context, questionID string) error {
	questionUID, err := strconv.ParseUint(questionID, 10, 64)
	if err != nil {
		return err
	}
	count, err := GetQuestionLikeCount(c, uint(questionUID))
	if err != nil {
		return err
	}
	// 更新数据库中的点赞数
	var question models.QuestionModel
	err = global.DB.First(&question, "id = ?", questionID).Error
	if err != nil {
		return err
	}
	question.DiggCount = count
	global.DB.Save(&question)
	return nil
}
