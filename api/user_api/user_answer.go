package user_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/page_list"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserAnswer(c *gin.Context) {
	var cr Query
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	answerList, count, err := page_list.PageListByUserId(models.AnswerModel{}, page_list.Option{
		PageInfo: cr.PageInfo,
		Debug:    false,
	}, cr.UserId)
	if err != nil {
		global.Log.Error(err)
	}
	res.OkWithList(answerList, count, c)
}
