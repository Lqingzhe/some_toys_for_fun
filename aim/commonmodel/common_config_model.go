package commonmodel

import "time"

type CommonConfig struct {
	Service string
	EquipID int
}
type GateWayConfig struct {
	Port        string
	ServiceInfo map[string]ServiceInfo
	RoutTimeOut map[string]time.Duration
}
type ServiceInfo struct {
	TimeOut     time.Duration
	ServiceAddr []ServiceAddr
}
type ServiceConfig struct {
	Timeout time.Duration
	ServiceAddr
}
type ServiceAddr struct {
	Host string
	Port int64
}
type LimiterConfig struct {
	MaxToken        int64
	GenerateToken   int64
	RedisExpireTime time.Duration
}

type TokenConfig struct {
	TokenPassword          string
	SaltByteLen            int64
	RefreshTokenExpireTime time.Duration
	AccessTokenExpireTime  time.Duration
}
type UserConfig struct {
	SaltByteLen       int64
	MaxPasswordLength int64
	MinPasswordLength int64
	MaxUsernameLength int64

	MaxUserNameLength  int64
	MaxIntroduceLength int64

	MaxNickNameLength int64
}
type GroupConfig struct {
	MaxGroupNameLength     int64
	MaxGroupNickNameLength int64
}
type MessageConfig struct {
	MaxMessageByteLength int64
}
type KafkaConfig struct {
	Host string
	Port int
}
