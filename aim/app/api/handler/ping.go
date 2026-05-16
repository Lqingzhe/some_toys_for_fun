package handler

import (
	newerror "aim/pkg/error"
	newlog "aim/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *HandlerConfig) Ping(c *gin.Context) {
	userID := c.GetInt64("user_id")
	ip := c.ClientIP()
	logger := c.MustGet("logger").(*zap.Logger)
	logger = newlog.AddGateWayInfo(logger, http.StatusOK, userID, ip, c.FullPath())
	newlog.SetGinLog(c, logger, "Ping", newerror.LevelInfo)
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "pong",
	})
}
