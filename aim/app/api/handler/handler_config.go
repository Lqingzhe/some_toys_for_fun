package handler

import (
	"aim/app/api/model"
	"aim/commonmodel"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type HandlerConfig struct {
	logger        *zap.Logger
	snowNode      *snowflake.Node
	dbContext     *model.DBContext
	equipID       int64
	tokenConfig   commonmodel.TokenConfig
	serviceClient model.ServiceClient
}

func NewHandlerConfig(logger *zap.Logger, snowNode *snowflake.Node, dbContext *model.DBContext, equipID int64, tokenConfig commonmodel.TokenConfig, serviceClient model.ServiceClient) *HandlerConfig {
	return &HandlerConfig{
		logger:        logger,
		snowNode:      snowNode,
		dbContext:     dbContext,
		equipID:       equipID,
		tokenConfig:   tokenConfig,
		serviceClient: serviceClient,
	}
}
