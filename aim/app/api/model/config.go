package model

import (
	"aim/commonmodel"
)

type Config struct {
	commonmodel.CommonConfig
	commonmodel.GateWayConfig
	DBConfig
	commonmodel.LimiterConfig
	commonmodel.TokenConfig
}
