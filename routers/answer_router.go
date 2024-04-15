package routers

import "QAComm/api"

func (r RouterGroup) AnswerRouter() {
	answerApi := api.ApiGroupApp.AnswerApi
	r.POST("/answer", answerApi.AnswerQuest)
	r.GET("/answer", answerApi.AnswerList)
	r.DELETE("/answer", answerApi.AnswerDelete)
	r.PUT("/answer", answerApi.AnswerUpdate)
}
