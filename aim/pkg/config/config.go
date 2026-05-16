package commonconfig

import (
	"aim/commonmodel"
	"fmt"
)

func GetCommonConfig() commonmodel.CommonConfig {

}
func GetMysqlConfig() commonmodel.MysqlConfig {
	newStruct := commonmodel.MysqlConfig{}

	newStruct.Url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local&parseTime=true&timeout=%s&readTimeout=%s&writeTimeout=%s",
		newStruct.Username,
		newStruct.Password,
		newStruct.Host,
		newStruct.Port,
		newStruct.DBName,
		newStruct.Timeout.String(),
		newStruct.ReadTimeout.String(),
		newStruct.WriteTimeout.String(),
	)
	return newStruct
}
func GetRedisConfig() commonmodel.RedisConfig {

}
func GetMongoDBConfig() commonmodel.MongoDBConfig {

}

func GetGatewayConfig() commonmodel.GateWayConfig {

}
func GetLimitersConfig() commonmodel.LimiterConfig {

}
func GetTokenConfig() commonmodel.TokenConfig {

}

func GetServiceConfig() commonmodel.ServiceConfig {

}
func GetUserConfig() commonmodel.UserConfig {
	newStruct := commonmodel.UserConfig{}

	if newStruct.MaxUsernameLength > 255 {
		newStruct.MaxUsernameLength = 255
	}
	if newStruct.MaxPasswordLength > 255 {
		newStruct.MaxPasswordLength = 255
	}
	if newStruct.MaxUserNameLength > 255 {
		newStruct.MaxUserNameLength = 255
	}
	if newStruct.MaxIntroduceLength > 255 {
		newStruct.MaxIntroduceLength = 255
	}
	if newStruct.MaxNickNameLength > 255 {
		newStruct.MaxNickNameLength = 255
	}
}
func GetGroupConfig() commonmodel.GroupConfig {

}
func GetMessageConfig() commonmodel.MessageConfig {
	newStruct := commonmodel.MessageConfig{}

	if newStruct.MaxMessageByteLength > 65535 {
		newStruct.MaxMessageByteLength = 65535
	}
}
func GetKafkaConfig() commonmodel.KafkaConfig {

}
