package global

import (
	"gorm.io/gorm"
	"survey_backend/config"
)

var (
	Config *config.Config
	Db     *gorm.DB
)
