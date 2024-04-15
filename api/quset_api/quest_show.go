package quset_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/page_list"
	"github.com/gin-gonic/gin"
)

func (QuestApi) QuestionList(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	questionList, count, err := page_list.PageList(models.QuestionModel{}, page_list.Option{
		PageInfo: cr,
		Debug:    false,
	})
	if err != nil {
		global.Log.Error(err)
	}
	res.OkWithList(questionList, count, c)
}
