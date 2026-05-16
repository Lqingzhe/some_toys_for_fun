package handler

import (
	"aim/app/api/service"
	"aim/kitex_gen/kitexcommonmodel"
	"aim/kitex_gen/kitexuserservice"
	newerror "aim/pkg/error"
	newlog "aim/pkg/log"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

//

func (h *HandlerConfig) Register(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	var req struct {
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    newerror.CodeInvalidJSON,
			"message": "JSON unmarshal error",
		})
		logger := newlog.AddError(h.logger, err, newerror.CodeInvalidJSON)
		logger = newlog.AddGateWayInfo(logger, http.StatusBadRequest, 0, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Register", newerror.LevelInfo)
		return
	}
	kitexReq := &kitexuserservice.RegisterReq{Password: req.Password}
	kitexResp, err := h.serviceClient.UserClient.Register(ctx, kitexReq)
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, 0, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Register", err2.LogLevel)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "success",
		"data": gin.H{
			"user_info": gin.H{
				"user_id": kitexResp.UserId,
			},
		},
	})
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, kitexResp.UserId, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "Register", newerror.LevelInfo)
	return
}
func (h *HandlerConfig) Login(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	var req struct {
		UserID   int64  `json:"user_id"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    newerror.CodeInvalidJSON,
			"message": "JSON unmarshal error",
		})
		logger := newlog.AddError(h.logger, err, newerror.CodeInvalidJSON)
		logger = newlog.AddGateWayInfo(logger, http.StatusBadRequest, 0, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Login", newerror.LevelInfo)
		return
	}
	kitexReq := &kitexuserservice.LoginReq{
		CommonInfo: &kitexcommonmodel.CommonInfo{Trace: c.GetString("trace")},
		UserId:     req.UserID,
		Password:   req.Password,
	}
	_, err := h.serviceClient.UserClient.Login(ctx, kitexReq)
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, req.UserID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Login", err2.LogLevel)
		return
	}

	serviceStruct := service.NewToken(h.dbContext, h.tokenConfig)
	accessToken, refreshToken, err := serviceStruct.MakeTokens(ctx, req.UserID, c.GetHeader("X-Device-ID"))
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, req.UserID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Login", err2.LogLevel)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "success",
		"data": gin.H{
			"token_info": gin.H{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		},
	})
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, req.UserID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "Login", newerror.LevelInfo)
	return
}

func (h *HandlerConfig) LogoutAll(c *gin.Context) {
	userID := c.GetInt64("user_id")
	tokenStruct := service.NewToken(h.dbContext, h.tokenConfig)
	err := tokenStruct.ReleaseAllToken(c.Request.Context(), userID)
	if err != nil {
		err2 := newerror.TranslateError(err)
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Logout", err2.LogLevel)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		return
	}
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, userID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "Logout", newerror.LevelInfo)
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "Logout Success",
	})
}
func (h *HandlerConfig) LogoutOne(c *gin.Context) {
	userID := c.GetInt64("user_id")
	tokenStruct := service.NewToken(h.dbContext, h.tokenConfig)
	err := tokenStruct.ReleaseOneTokenWithDeviceID(c.Request.Context(), userID, c.GetHeader("X-Device-ID"))
	if err != nil {
		err2 := newerror.TranslateError(err)
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Logout", err2.LogLevel)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		return
	}
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, userID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "Logout", newerror.LevelInfo)
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "Logout Success",
	})
}

func (h *HandlerConfig) RefreshToken(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    newerror.CodeInvalidJSON,
			"message": "JSON unmarshal error",
		})
		logger := newlog.AddError(h.logger, err, newerror.CodeInvalidJSON)
		logger = newlog.AddGateWayInfo(logger, http.StatusBadRequest, 0, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "RefreshToken", newerror.LevelInfo)
		return
	}
	tokenStruct := service.NewToken(h.dbContext, h.tokenConfig)
	userID, accessToken, refreshToken, err := tokenStruct.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, 0, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "RefreshToken", err2.LogLevel)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "success",
		"data": gin.H{
			"token_info": gin.H{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		},
	})
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, userID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "RefreshToken", newerror.LevelInfo)
	return
}

//

func (h *HandlerConfig) GetUserInfo(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	userID := c.GetInt64("user_id")
	kitexReq := kitexuserservice.GetUserInfoReq{
		CommonInfo: &kitexcommonmodel.CommonInfo{
			Trace: c.GetString("trace"),
		},
		UserId: userID,
	}
	kitexResp, err := h.serviceClient.UserClient.GetUserInfo(ctx, &kitexReq)
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, userID, c.ClientIP(), c.FullPath())
	}
	type userInfo struct {
		UserID        int64  `json:"user_id"`
		UserName      string `json:"user_name"`
		Introduction  string `json:"introduction"`
		BirthdayYear  int64  `json:"birthday_year"`
		BirthdayMonth int64  `json:"birthday_month"`
		BirthdayDay   int64  `json:"birthday_day"`
	}

	type remarkInfo struct {
		GoalUserID int64  `json:"goal_user_name"`
		NickName   string `json:"nick_name"`
	}
	RemarkInfoList := make([]remarkInfo, len(kitexResp.RemarkInfo))
	for i, j := range kitexResp.RemarkInfo {
		RemarkInfoList[i] = remarkInfo{
			GoalUserID: j.GoalUserID,
			NickName:   j.NickName,
		}
	}
	var resp = struct {
		UserInfo    userInfo
		RemarkInfos []remarkInfo
	}{
		UserInfo: userInfo{
			UserID:        userID,
			UserName:      kitexResp.UserInfo.UserName,
			Introduction:  kitexResp.UserInfo.Introduction,
			BirthdayYear:  kitexResp.UserInfo.BirthdayYear,
			BirthdayMonth: kitexResp.UserInfo.BirthdayMonth,
			BirthdayDay:   kitexResp.UserInfo.BirthdayDay,
		},
		RemarkInfos: RemarkInfoList,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "success",
		"data":    resp,
	})
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, userID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "GetUserInfo", newerror.LevelInfo)
	return
}
func (h *HandlerConfig) GetOtherUserInfo(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	userID := c.GetInt64("user_id")
	var req struct {
		GoalUserID int64 `json:"goal_user_id"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    newerror.CodeInvalidJSON,
			"message": "JSON unmarshal error",
		})
		logger := newlog.AddError(h.logger, err, newerror.CodeInvalidJSON)
		logger = newlog.AddGateWayInfo(logger, http.StatusBadRequest, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "GetOtherUserInfo", newerror.LevelInfo)
		return
	}
	kitexReq := kitexuserservice.GetOtherUserInfoReq{
		CommonInfo: &kitexcommonmodel.CommonInfo{
			Trace: c.GetString("trace"),
		},
		GoalUserId: req.GoalUserID,
	}
	kitexResp, err := h.serviceClient.UserClient.GetOtherUserInfo(ctx, &kitexReq)
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "GetOtherUserInfo", newerror.LevelInfo)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "success",
		"data": gin.H{
			"user_info": gin.H{
				"user_name":      kitexResp.UserInfo.UserName,
				"introduction":   kitexResp.UserInfo.Introduction,
				"birthday_year":  kitexResp.UserInfo.BirthdayYear,
				"birthday_month": kitexResp.UserInfo.BirthdayMonth,
				"birthday_day":   kitexResp.UserInfo.BirthdayDay,
			},
		},
	})
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, userID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "GetOtherUserInfo", newerror.LevelInfo)
	return
}
func (h *HandlerConfig) UpdateUserInfo(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	userID := c.GetInt64("user_id")
	var req struct {
		UserName      string `json:"user_name"`
		Introduction  string `json:"introduction"`
		BirthdayYear  int64  `json:"birthday_year"`
		BirthdayMonth int64  `json:"birthday_month"`
		BirthdayDay   int64  `json:"birthday_day"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    newerror.CodeInvalidJSON,
			"message": "JSON unmarshal error",
		})
		logger := newlog.AddError(h.logger, err, newerror.CodeInvalidJSON)
		logger = newlog.AddGateWayInfo(logger, http.StatusBadRequest, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "UpdateUserInfo", newerror.LevelInfo)
		return
	}
	kitexReq := kitexuserservice.UpdateUserInfoReq{
		CommonInfo: &kitexcommonmodel.CommonInfo{
			Trace: c.GetString("trace"),
		},
		UserInfo: &kitexcommonmodel.UserInfo{
			UserID:        userID,
			UserName:      req.UserName,
			Introduction:  req.Introduction,
			BirthdayYear:  req.BirthdayYear,
			BirthdayMonth: req.BirthdayMonth,
			BirthdayDay:   req.BirthdayDay,
		},
	}
	_, err := h.serviceClient.UserClient.UpdateUserInfo(ctx, &kitexReq)
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "UpdateUserInfo", err2.LogLevel)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "success",
	})
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, userID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "UpdateUserInfo", newerror.LevelInfo)
	return
}
func (h *HandlerConfig) Remark(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)
	userID := c.GetInt64("user_id")
	var req struct {
		GoalUserID int64  `json:"goal_user_id"`
		NickName   string `json:"nick_name"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    newerror.CodeInvalidJSON,
			"message": "JSON unmarshal error",
		})
		logger := newlog.AddError(h.logger, err, newerror.CodeInvalidJSON)
		logger = newlog.AddGateWayInfo(logger, http.StatusBadRequest, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Remark", newerror.LevelInfo)
		return
	}
	kitexReq := kitexuserservice.RemarkReq{
		CommonInfo: &kitexcommonmodel.CommonInfo{
			Trace: c.GetString("trace"),
		},
		RemarkInfo: &kitexcommonmodel.RemarkInfo{
			UserID:     userID,
			GoalUserID: req.GoalUserID,
			NickName:   req.NickName,
		},
	}
	_, err := h.serviceClient.UserClient.Remark(ctx, &kitexReq)
	if err != nil {
		err2 := newerror.TranslateError(err)
		c.AbortWithStatusJSON(err2.HttpCode, gin.H{
			"code":    err2.StatueCode,
			"message": err2.HttpMessage,
		})
		logger := newlog.AddError(h.logger, err, err2.StatueCode)
		logger = newlog.AddGateWayInfo(logger, err2.HttpCode, userID, c.ClientIP(), c.FullPath())
		newlog.SetGinLog(c, logger, "Remark", err2.LogLevel)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    newerror.CodeSuccess,
		"message": "success",
	})
	logger := newlog.AddGateWayInfo(h.logger, http.StatusOK, userID, c.ClientIP(), c.FullPath())
	newlog.SetGinLog(c, logger, "Remark", newerror.LevelInfo)
	return
}
