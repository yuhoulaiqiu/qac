package main

import (
	"QAComm/core"
	"QAComm/flag"
	"QAComm/global"
	"QAComm/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	//global.Redis = core.InitRedis()
	// 调用接口
	r := routers.InitRouter()
	// 运行
	addr := global.Config.System.Addr()
	global.Log.Infof("问答论坛运行在%v中", addr)
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
