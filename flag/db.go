package flag

import (
	"survey_backend/global"
	"survey_backend/models"
)

func Makemigrations() {
	// 生成表结构
	err := global.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.AnswerModel{},
		&models.RoleModel{},
		&models.OptionModel{},
		&models.QuestionModel{},
		&models.QuestionnaireModel{},
	)
	if err != nil {
		global.Log.Errorf("表结构生成失败：%s", err.Error())
		return
	}
	global.Log.Info("表结构生成成功")
}
