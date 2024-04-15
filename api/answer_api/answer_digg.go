package answer_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	AnswerLikeKey = "answer:like:"
	UserLikeKey   = "user:answer:like:"
)

type diggRequest struct {
	AnswerID string `json:"answer_id"`
	UserID   string `json:"user_id"`
}

// AnswerDigg 点赞
func (AnswerApi) AnswerDigg(c *gin.Context) {
	var req diggRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	userLikeKey := fmt.Sprintf("%s%s", UserLikeKey, req.UserID)
	// 判断用户是否已经点赞
	if global.Redis.SIsMember(c, userLikeKey, req.AnswerID).Val() {
		res.FailWithCode(res.RepeatDiggError, c)
		return
	}

	// 将用户ID添加到答案的点赞用户集合中
	global.Redis.SAdd(c, userLikeKey, req.AnswerID)

	// 点赞
	key := fmt.Sprintf("%s%s", AnswerLikeKey, req.AnswerID)
	err = global.Redis.Incr(c, key).Err()
	if err != nil {
		res.FailWithCode(res.RedisError, c)
		return
	}
	// 同步到数据库
	err = SyncAnswerDiggCountToDB(c, req.AnswerID)
	if err != nil {

		res.FailWithCode(res.ArgumentError, c)
		return
	}
	res.OkWithMessage("点赞成功", c)
}

// AnswerUnDigg 取消点赞
func (AnswerApi) AnswerUnDigg(c *gin.Context) {
	var req diggRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	userLikeKey := fmt.Sprintf("%s%s", UserLikeKey, req.UserID)
	// 判断用户是否已经点赞
	if global.Redis.SIsMember(c, userLikeKey, req.AnswerID).Val() {
		// 将用户ID从答案的点赞用户集合中移除
		global.Redis.SRem(c, userLikeKey, req.AnswerID)

		key := fmt.Sprintf("%s%s", AnswerLikeKey, req.AnswerID)
		err = global.Redis.Decr(c, key).Err()
		if err != nil {
			res.FailWithCode(res.RedisError, c)
			return
		}

		// 同步到数据库
		err = SyncAnswerDiggCountToDB(c, req.AnswerID)
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

// GetAnswerDiggCount 获取答案的点赞数
func GetAnswerDiggCount(c *gin.Context, answerID uint) (int, error) {
	key := AnswerLikeKey + strconv.Itoa(int(answerID))
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

// SyncAnswerDiggCountToDB 将Redis中的点赞数同步到数据库中
func SyncAnswerDiggCountToDB(c *gin.Context, answerID string) error {
	questionUID, err := strconv.ParseUint(answerID, 10, 64)
	if err != nil {
		return err
	}
	count, err := GetAnswerDiggCount(c, uint(questionUID))
	if err != nil {
		return err
	}
	// 更新数据库中的点赞数
	var answer models.AnswerModel
	err = global.DB.First(&answer, "id = ?", answerID).Error
	if err != nil {
		return err
	}
	answer.DiggCount = count
	global.DB.Save(&answer)
	return nil
}
