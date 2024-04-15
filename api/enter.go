// Api入口，通过ApiGroup结构体绑定所有种类的api，并进行实体化

package api

import (
	"QAComm/api/answer_api"
	"QAComm/api/quset_api"
	"QAComm/api/user_api"
)

type ApiGroup struct {
	UserApi   user_api.UserApi
	QuestApi  quset_api.QuestApi
	AnswerApi answer_api.AnswerApi
}

var ApiGroupApp = new(ApiGroup)
