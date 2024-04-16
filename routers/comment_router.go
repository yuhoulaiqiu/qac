package routers

import "QAComm/api"

func (r RouterGroup) CommentRouter() {
	commentApi := api.ApiGroupApp.CommentApi
	r.POST("/comment", commentApi.Comment)
	r.POST("/comment/digg", commentApi.CommentDigg)
	r.POST("/comment/undigg", commentApi.CommentUnDigg)
}
