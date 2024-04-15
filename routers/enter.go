// router总入口，调用其他router

package routers

import (
	"QAComm/global"
	"github.com/gin-gonic/gin"
)

// RouterGroup api分组
type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// 路由分组
	apiRouterGroup := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	// 路由分层
	// 用户api
	routerGroupApp.UserRouter()
	//问题api
	routerGroupApp.QuestRouter()
	//回答api
	routerGroupApp.AnswerRouter()
	return router
}
