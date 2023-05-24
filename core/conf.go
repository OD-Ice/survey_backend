package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"survey_backend/config"
	"survey_backend/global"
)

// InitConf 读取配置文件
func InitConf() {
	const ConfigFile = "settings.yaml"
	structConfig := &config.Config{}
	yamlConfig, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取yaml文件出错: %s", err))
	}
	err = yaml.Unmarshal(yamlConfig, structConfig)
	if err != nil {
		log.Fatalf("解析yaml文件出错: %s", err)
	}
	log.Println("配置初始化成功")
	// 赋值到全局变量
	global.Config = structConfig
}
