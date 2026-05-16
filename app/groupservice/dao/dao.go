package dao

import (
	"aim/app/groupservice/model"
	"aim/commonmodel"
	commonconfig "aim/pkg/config"
	newlog "aim/pkg/log"
	"context"

	"go.uber.org/zap"
)

func InitDB(MysqlConfig *commonmodel.MysqlConfig, logger *zap.Logger) *model.DBContext {
	MysqlCtx, err := commonconfig.MakeMysql(MysqlConfig)
	if err != nil {
		newlog.LogInitFatal(logger, err, "Init Mysql Failed")
	}

	return &model.DBContext{
		Mysql: MysqlCtx,
	}
}
func CloseDB(MysqlContext *commonmodel.MysqlContext) {
	commonconfig.DBClose(MysqlContext.Client)
}
func Add(ctx context.Context, info commonmodel.DBOperater, dbContext *model.DBContext) error {
	return info.AddInfo(ctx, dbContext)
}
func Update(ctx context.Context, info commonmodel.DBOperater, dbContext *model.DBContext) (bool, error) {
	return info.UpdateInfo(ctx, dbContext)
}
func Delete(ctx context.Context, info commonmodel.DBOperater, dbContext *model.DBContext) error {
	return info.DeleteInfo(ctx, dbContext)
}
func Get(ctx context.Context, info commonmodel.DBOperater, dbContext *model.DBContext) (bool, error) {
	return info.GetInfo(ctx, dbContext)
}
