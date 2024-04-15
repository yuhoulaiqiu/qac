package answer_api

import (
	"QAComm/global"
	"QAComm/models"
	"QAComm/models/res"
	"QAComm/utils/page_list"
	"github.com/gin-gonic/gin"
)

type Query struct {
	models.PageInfo
	QuestId int `form:"quest_id"`
}

func (AnswerApi) AnswerList(c *gin.Context) {
	var cr Query
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	answerList, count, err := page_list.PageListByQuestId(models.AnswerModel{}, page_list.Option{
		PageInfo: cr.PageInfo,
		Debug:    false,
	}, cr.QuestId)
	if err != nil {
		global.Log.Error(err)
	}
	res.OkWithList(answerList, count, c)
}
