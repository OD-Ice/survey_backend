package main

import (
	"survey_backend/core"
	"survey_backend/flag"
	"survey_backend/global"
	"survey_backend/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化mysql连接
	global.Db = core.InitGorm()
	// 命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Error("服务启动失败：", err.Error())
	}
}
