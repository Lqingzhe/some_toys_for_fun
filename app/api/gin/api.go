package gin

import (
	"aim/app/api/handler"
	"aim/app/api/middleware"
	"aim/app/api/model"
	"aim/commonmodel"
	newlog "aim/pkg/log"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiConfig struct {
	logger        *zap.Logger
	snowNode      *snowflake.Node
	dbContext     *model.DBContext
	limiterConfig commonmodel.LimiterConfig
	tokenConfig   commonmodel.TokenConfig
	serviceClient model.ServiceClient
	equipID       int64
	RoutTimeOut   map[string]time.Duration
}

func NewConfig(logger *zap.Logger, snowNode *snowflake.Node, dbContext *model.DBContext, limiterConfig commonmodel.LimiterConfig, tokenConfig commonmodel.TokenConfig, equipID int64, RoutTimeout map[string]time.Duration, serviceClient model.ServiceClient) *ApiConfig {
	return &ApiConfig{
		logger:        logger,
		snowNode:      snowNode,
		dbContext:     dbContext,
		limiterConfig: limiterConfig,
		tokenConfig:   tokenConfig,
		serviceClient: serviceClient,
		equipID:       equipID,
		RoutTimeOut:   RoutTimeout,
	}
}
func (A *ApiConfig) Begin(port string) {
	handlerConfig := handler.NewHandlerConfig(A.logger, A.snowNode, A.dbContext, A.equipID, A.tokenConfig, A.serviceClient)

	g := gin.New()
	gin.SetMode(gin.ReleaseMode)
	g.Use(gin.Recovery())

	g.Use(middleware.Log(A.logger, A.equipID))

	g.POST("ping", handlerConfig.Ping) //测试

	needLogin := g.Group("")
	needLogin.Use(
		middleware.AnalyseToken(A.tokenConfig, A.dbContext),
		middleware.Limiter(A.dbContext, A.limiterConfig),
	)
	{
		user := needLogin.Group("/user")
		{
			g.POST("/user/register", middleware.Limiter(A.dbContext, A.limiterConfig), middleware.SetTimeOut(A.RoutTimeOut["/user/register"]), handlerConfig.Register)

			//添加功能：登录时，先确认websocked为未登录，然后登录websocked，访问messageService获取未发送消息，发送
			g.POST("/user/login", middleware.Limiter(A.dbContext, A.limiterConfig), middleware.SetTimeOut(A.RoutTimeOut["/user/login"]), handlerConfig.Login)
			g.POST("/user/refresh-token", middleware.Limiter(A.dbContext, A.limiterConfig), middleware.SetTimeOut(A.RoutTimeOut["/user/refresh-token"]), handlerConfig.RefreshToken)

			//添加功能：登出websocked
			user.POST("/logout-all-device", middleware.SetTimeOut(A.RoutTimeOut["/user/logout-all-device"]), handlerConfig.LogoutAll)
			//添加功能：登出websocked
			user.POST("/logout-a-device", middleware.SetTimeOut(A.RoutTimeOut["/user/logout-a-device"]), handlerConfig.LogoutOne)

			//添加功能：查询websocked查看用户的登录状态
			user.POST("/get-user-info", middleware.SetTimeOut(A.RoutTimeOut["/user/get-user-info"]), handlerConfig.GetUserInfo)
			user.POST("/get-other-user-info", middleware.SetTimeOut(A.RoutTimeOut["/user/get-other-user-info"]), handlerConfig.GetOtherUserInfo)
			user.POST("/update-user-info", middleware.SetTimeOut(A.RoutTimeOut["/user/update-user-info"]), handlerConfig.UpdateUserInfo)
			user.POST("remark", middleware.SetTimeOut(A.RoutTimeOut["/user/remark"]), handlerConfig.Remark)
		}
		message := needLogin.Group("/message")
		{

		}
		group := needLogin.Group("/group")
		{

		}
	}
	g.POST("/user/register")
	g.POST("/user/login")
	//路由注册

	err := g.Run(":" + port)
	if err != nil {
		newlog.LogInitError(A.logger, err, "http begin error")
		return
	}
}
