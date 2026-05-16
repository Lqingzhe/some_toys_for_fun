package config

import (
	"aim/app/api/model"
	"aim/pkg/config"
)

func initDBConfig() model.DBConfig {
	newConfig := model.DBConfig{}

	newConfig.Redis = commonconfig.GetRedisConfig()

	return newConfig
}
