package config

import (
	"aim/app/groupservice/model"
	commonconfig "aim/pkg/config"
)

func InitConfig() model.Config {
	newConfig := model.Config{}

	newConfig.CommonConfig = commonconfig.GetCommonConfig()
	newConfig.ServiceConfig = commonconfig.GetServiceConfig()
	newConfig.GroupConfig = commonconfig.GetGroupConfig()
	newConfig.DBConfig = initDBConfig()

	return newConfig
}
