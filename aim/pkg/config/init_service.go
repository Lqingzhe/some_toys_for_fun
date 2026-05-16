package commonconfig

import (
	"aim/commonmodel"
	newlog "aim/pkg/log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
)

func initNacosClient(logger *zap.Logger, timeout time.Duration, serviceAddr ...commonmodel.ServiceAddr) naming_client.INamingClient {
	sc := make([]constant.ServerConfig, len(serviceAddr))
	for i, j := range serviceAddr {
		sc[i] = *constant.NewServerConfig(j.Host, uint64(j.Port))
	}
	cc := constant.ClientConfig{
		NamespaceId: "public",
		TimeoutMs:   uint64(timeout.Seconds()),
	}
	nacosClient, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		newlog.LogInitFatal(logger, err, "Init Nacos Service Register Error")
	}
	return nacosClient
}
func ResolverService(serviceName string, Config commonmodel.GateWayConfig, logger *zap.Logger) client.Option {
	return client.WithResolver(
		resolver.NewNacosResolver(
			initNacosClient(
				logger,
				Config.ServiceInfo[serviceName].TimeOut,
				Config.ServiceInfo[serviceName].ServiceAddr...,
			),
		),
	)
}
func RegisterService(Config commonmodel.ServiceConfig, logger *zap.Logger) server.Option {
	return server.WithRegistry(
		registry.NewNacosRegistry(
			initNacosClient(
				logger,
				Config.Timeout,
				Config.ServiceAddr,
			),
		),
	)
}
