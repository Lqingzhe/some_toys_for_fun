package model

import "aim/commonmodel"

type DBContext struct {
	Mysql   *commonmodel.MysqlContext
	Redis   *commonmodel.RedisContext
	MongoDB *commonmodel.MongoDBContext
}
