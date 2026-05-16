package main

import (
	"aim/app/api/config"
	"aim/app/api/dao"
	api "aim/app/api/gin"
	"aim/app/api/model"
	kitexuserservice "aim/kitex_gen/kitexuserservice/userservice"

	commonconfig "aim/pkg/config"
	"aim/pkg/id"
	"aim/pkg/log"
)

func main() {
	Config := config.InitConfig()

	logger := newlog.InitLog(Config.Service, Config.EquipID)
	defer logger.Sync()

	snowNode := id.InitSnowNode(Config.EquipID, logger)

	dbContext := dao.InitDB(&Config.DBConfig.Redis, logger)
	defer dao.CloseDB(dbContext)

	UserClient := kitexuserservice.MustNewClient(
		"user_service",
		commonconfig.ResolverService("user_service", Config.GateWayConfig, logger),
	)

	httpStruct := api.NewConfig(
		logger,
		snowNode,
		dbContext,
		Config.LimiterConfig,
		Config.TokenConfig,
		int64(Config.EquipID),
		Config.RoutTimeOut,
		model.ServiceClient{
			UserClient: UserClient,
		})

	httpStruct.Begin(Config.Port)
}
