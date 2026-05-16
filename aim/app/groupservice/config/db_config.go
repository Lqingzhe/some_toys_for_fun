package config

import (
	"aim/app/groupservice/model"
	commonconfig "aim/pkg/config"
)

func initDBConfig() model.DBConfig {
	newConfig := model.DBConfig{}

	newConfig.Mysql = commonconfig.GetMysqlConfig()
	newConfig.Redis = commonconfig.GetRedisConfig()

	return newConfig
}
