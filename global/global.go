package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"survey_backend/config"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Log    *logrus.Logger
	Redis  *redis.Client
)
