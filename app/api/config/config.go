package config

import (
	"aim/app/api/model"
	"aim/pkg/config"
)

func InitConfig() model.Config {
	newConfig := model.Config{}

	newConfig.CommonConfig = commonconfig.GetCommonConfig()
	newConfig.GateWayConfig = commonconfig.GetGatewayConfig()
	newConfig.DBConfig = initDBConfig()
	newConfig.LimiterConfig = commonconfig.GetLimitersConfig()
	newConfig.TokenConfig = commonconfig.GetTokenConfig()

	return newConfig
}
