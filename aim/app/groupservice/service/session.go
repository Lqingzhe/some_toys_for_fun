package service

import (
	"aim/app/groupservice/dao"
	"aim/app/groupservice/dao/groupapply"
	"aim/app/groupservice/dao/sessioninfo"
	"aim/app/groupservice/model"
	"aim/commonmodel"
	newerror "aim/pkg/error"
	"aim/tool"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/bwmarrin/snowflake"
)

type ServiceSession struct {
	systemTopic sarama.SyncProducer
	dbContext   *model.DBContext
	snowFlack   *snowflake.Node
}

func NewSession(systemTopic sarama.SyncProducer, dbContext *model.DBContext, snowFlack *snowflake.Node) *ServiceSession {
	return &ServiceSession{
		systemTopic: systemTopic,
		dbContext:   dbContext,
		snowFlack:   snowFlack,
	}
}
func (s *ServiceSession) CreatSession(ctx context.Context, userID int64, goalUserID int64) (SessionID int64, err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("session:CreatSession")
	SessionID = s.snowFlack.Generate().Int64()
	sessionStruct1 := sessioninfo.NewStruct(SessionID, userID, goalUserID, time.Time{})
	sessionStruct2 := sessioninfo.NewStruct(SessionID, goalUserID, userID, time.Time{})
	DB := &model.DBContext{
		Mysql: tool.BeginMysqlTransaction(s.dbContext.Mysql),
	}
	err = dao.Add(ctx, sessionStruct1, DB)
	if err != nil {
		DB.Mysql.Client.Rollback()
		return 0, err
	}
	err = dao.Add(ctx, sessionStruct2, DB)
	if err != nil {
		DB.Mysql.Client.Rollback()
		return 0, err
	}
	result := DB.Mysql.Client.Commit()
	if err2 := newerror.IsMysqlError(result); err2 != nil {
		return 0, err2
	}
	//向s.GoalUserID发送消息
	messageStruct := commonmodel.KafkaSystemMessage{
		GoalUserID:  []int64{goalUserID},
		Data:        map[string]any{"user_id": userID},
		MessageCode: commonmodel.MessageCode_FriendRequest_Success,
	}
	_, _, err = tool.SendKafkaSystemMessage(s.systemTopic, messageStruct)
	if err != nil {
		return 0, err
	}
	return SessionID, nil
}

func (s *ServiceSession) DeleteSession(ctx context.Context, sessionID int64, userID int64) (err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("session:DeleteSession")
	sessionStruct := sessioninfo.NewStruct(sessionID, userID, 0, time.Time{}, sessioninfo.WithSessionID, sessioninfo.WithUserID)
	exist, err := dao.Get(ctx, sessionStruct, s.dbContext)
	if err != nil {
		return err
	}
	if !exist {
		return newerror.MakeError(http.StatusBadRequest, newerror.CodeResourceNotFound, "He IS Not Your Friend", fmt.Errorf("Delete Unexist SessionID"), newerror.LevelInfo)
	}
	goalUserID := sessionStruct.Info[0].GoalUserID
	sessionStruct = sessioninfo.NewStruct(sessionID, 0, 0, time.Time{}, sessioninfo.WithSessionID)
	err = dao.Delete(ctx, sessionStruct, s.dbContext)
	if err != nil {
		return err
	}
	messageStruct := commonmodel.KafkaSystemMessage{
		GoalUserID:  []int64{goalUserID},
		Data:        map[string]any{"user_id": userID},
		MessageCode: commonmodel.MessageCode_FriendDelete,
	}
	_, _, err = tool.SendKafkaSystemMessage(s.systemTopic, messageStruct)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceSession) GetFriendLastVisitTime(ctx context.Context, sessionID int64, goalUserID int64) (lastVisitTime string, err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("session:GetFriendLastVisitTime")
	sessionStruct := sessioninfo.NewStruct(sessionID, goalUserID, 0, time.Time{}, sessioninfo.WithSessionID, sessioninfo.WithUserID)
	exist, err := dao.Get(ctx, sessionStruct, s.dbContext)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", newerror.MakeError(http.StatusBadRequest, newerror.CodeResourceNotFound, "He Is Not Your Friend", fmt.Errorf("Do Not Get The Info Whith SessionID and UserID"), newerror.LevelInfo)
	}
	return sessionStruct.Info[0].LastReadTime.String(), nil
}
func (s *ServiceSession) ApplyForFriend(ctx context.Context, userID int64, goalUserID int64) (err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("session:ApplyForFriend")
	groupApplyStruct := groupapply.NewStruct(userID, goalUserID)
	err = dao.Add(ctx, groupApplyStruct, s.dbContext)
	if err != nil {
		return err
	}
	messageStruct := commonmodel.KafkaSystemMessage{
		GoalUserID:  []int64{goalUserID},
		Data:        map[string]any{"user_id": userID},
		MessageCode: commonmodel.MessageCode_FriendRequest,
	}
	_, _, err = tool.SendKafkaSystemMessage(s.systemTopic, messageStruct)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceSession) GetFriendApplyList(ctx context.Context, userID int64) (applyUserID []int64, err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("session:GetApplyList")
	groupApplyStruct := groupapply.NewStruct(userID, 0, groupapply.WithGoalID)
	exist, err := dao.Get(ctx, groupApplyStruct, s.dbContext)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	applyUserID = make([]int64, len(groupApplyStruct.Info))
	for i, v := range groupApplyStruct.Info {
		applyUserID[i] = v.ApplyUserID
	}
	return applyUserID, nil
}
func (s *ServiceSession) RefuseFriendApply(ctx context.Context, userID int64, goalUserID int64) (err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("session:RefuseFriendApply")
	groupApplyStruct := groupapply.NewStruct(userID, goalUserID, groupapply.WithGoalID, groupapply.WithApplyUserID)
	err = dao.Delete(ctx, groupApplyStruct, s.dbContext)
	if err != nil {
		return err
	}
	messageStruct := commonmodel.KafkaSystemMessage{
		GoalUserID:  []int64{goalUserID},
		Data:        map[string]any{"user_id": userID},
		MessageCode: commonmodel.MessageCode_FriendRequest_Refuse,
	}
	_, _, err = tool.SendKafkaSystemMessage(s.systemTopic, messageStruct)
	if err != nil {
		return err
	}
	return nil
}
