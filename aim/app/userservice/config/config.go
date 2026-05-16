package config

import (
	"aim/app/userservice/model"
	"aim/pkg/config"
)

func InitConfig() model.Config {
	newConfig := model.Config{}

	newConfig.CommonConfig = commonconfig.GetCommonConfig()
	newConfig.ServiceConfig = commonconfig.GetServiceConfig()
	newConfig.UserConfig = commonconfig.GetUserConfig()
	newConfig.DBConfig = initDBConfig()

	return newConfig
}
