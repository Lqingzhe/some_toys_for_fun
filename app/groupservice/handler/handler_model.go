package handler

import (
	"aim/app/groupservice/model"
	"aim/commonmodel"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

type GroupServiceImpl struct {
	SnowNode    *snowflake.Node
	DBContext   *model.DBContext
	Logger      *zap.Logger
	GroupConfig commonmodel.GroupConfig
	EquipID     int64
}

func NewGroupServiceImpl(SnowNode *snowflake.Node, DBContext *model.DBContext, Logger *zap.Logger, GroupConfig commonmodel.GroupConfig, EquipID int64) *GroupServiceImpl {
	return &GroupServiceImpl{
		SnowNode:    SnowNode,
		DBContext:   DBContext,
		Logger:      Logger,
		GroupConfig: GroupConfig,
		EquipID:     EquipID,
	}
}
