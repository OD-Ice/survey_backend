package main

import (
	"fmt"
	"survey_backend/core"
	"survey_backend/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化mysql连接
	global.Db = core.InitGorm()
	fmt.Println(global.Db)
}
