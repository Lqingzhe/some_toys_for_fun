package config

import (
	"aim/app/userservice/model"
	"aim/pkg/config"
)

func initDBConfig() model.DBConfig {
	newConfig := model.DBConfig{}

	newConfig.Mysql = commonconfig.GetMysqlConfig()

	return newConfig
}
