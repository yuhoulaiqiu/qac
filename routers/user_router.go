// 调用user api，绑定在对应的RouterGroup中

package routers

import (
	"QAComm/api"
)

func (r RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.POST("/login", userApi.Login)
	r.POST("/register", userApi.Register)
	r.GET("/user/quest", userApi.UserQuest)
	r.GET("/user/answer", userApi.UserAnswer)
}
