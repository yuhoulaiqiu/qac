package routers

import "QAComm/api"

func (r RouterGroup) QuestRouter() {
	questApi := api.ApiGroupApp.QuestApi
	r.POST("/quest", questApi.PutQuestion)
	r.GET("/quest", questApi.QuestionList)
	r.DELETE("/quest", questApi.QuestDelete)
	r.PUT("/quest", questApi.QuestUpdate)
}
