package service

import (
	"aim/app/groupservice/dao"
	"aim/app/groupservice/dao/groupmuteinfo"
	"aim/app/groupservice/dao/groupwithuser"
	"aim/app/groupservice/model"
	"aim/commonmodel"
	newerror "aim/pkg/error"
	"aim/tool"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
)

type Mute struct {
	groupNoticeTopic sarama.SyncProducer
	dbContext        *model.DBContext
}

func NewMute(groupNoticeTopic sarama.SyncProducer, dbContext *model.DBContext) *Mute {
	return &Mute{
		groupNoticeTopic: groupNoticeTopic,
		dbContext:        dbContext,
	}
}

func (m *Mute) SetMute(ctx context.Context, userID int64, groupID int64, goalUserID int64, MuteTimeSecond int64, MuteReason string) (err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("group_mute:SetMute")
	groupWithUserStruct := groupwithuser.NewStruct(groupID, userID, "", "", time.Time{}, groupwithuser.WithGroupID, groupwithuser.WithUserID)
	exist, err := dao.Get(ctx, groupWithUserStruct, m.dbContext)
	if err != nil {
		return err
	}
	if !exist {
		return newerror.MakeError(http.StatusNotFound, newerror.CodeResourceNotFound, "The Group Is Not Exist", fmt.Errorf("Try To Set Mute By Uninjion User"), newerror.LevelInfo)
	}
	userRole := groupWithUserStruct.Info[0].Role
	groupWithUserStruct = groupwithuser.NewStruct(groupID, goalUserID, "", "", time.Time{}, groupwithuser.WithGroupID, groupwithuser.WithUserID)
	exist, err = dao.Get(ctx, groupWithUserStruct, m.dbContext)
	if err != nil {
		return err
	}
	if !exist {
		return newerror.MakeError(http.StatusNotFound, newerror.CodeUserNotFound, "The User Did Not Join The Group", fmt.Errorf("Try To Set Mute To Uninjion User"), newerror.LevelInfo)
	}
	goalUserRole := groupWithUserStruct.Info[0].Role
	if userRole == commonmodel.Member || (userRole == commonmodel.Manager && goalUserRole != commonmodel.Member) {
		return newerror.MakeError(http.StatusForbidden, newerror.CodePermissionDenied, "You Are Not The Manager", fmt.Errorf("Try To Set Mute Without Enough Permission"), newerror.LevelInfo)
	}
	groupMuteStruct := groupmuteinfo.NewStruct(groupID, goalUserID, time.Now().Add(time.Duration(MuteTimeSecond)*time.Second), MuteReason, groupmuteinfo.WithWhereUserID)
	exist, err = dao.Update(ctx, groupMuteStruct, m.dbContext)
	if err != nil {
		return err
	}
	if !exist {
		err = dao.Add(ctx, groupMuteStruct, m.dbContext)
		if err != nil {
			return err
		}
	}
	groupWithUserStruct = groupwithuser.NewStruct(groupID, 0, "", "", time.Time{}, groupwithuser.WithGroupID)
	_, err = dao.Get(ctx, groupWithUserStruct, m.dbContext)
	if err != nil {
		return err
	}
	memberList := make([]int64, 0, len(groupWithUserStruct.Info))
	for _, info := range groupWithUserStruct.Info {
		memberList = append(memberList, info.UserID)
	}
	groupNoticeMessage := commonmodel.KafkaGroupNotice{
		GoalUserID:  memberList,
		SessionID:   groupID,
		Data:        map[string]any{"user_id": userID, "goal_user_id": goalUserID, "mute_time": MuteTimeSecond, "mute_reason": MuteReason},
		MessageCode: commonmodel.MessageCode_GroupSetMute,
	}
	_, _, err = tool.SendKafkaGroupNotice(m.groupNoticeTopic, groupNoticeMessage)
	if err != nil {
		return err
	}
	return nil
}
func (m *Mute) ReleaseMute(ctx context.Context, userID int64, groupID int64, goalUserID int64) (err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("group_mute:ReleaseMute")
	groupWithUserStruct := groupwithuser.NewStruct(groupID, userID, "", "", time.Time{}, groupwithuser.WithGroupID, groupwithuser.WithUserID)
	exist, err := dao.Get(ctx, groupWithUserStruct, m.dbContext)
	if err != nil {
		return err
	}
	if !exist {
		return newerror.MakeError(http.StatusNotFound, newerror.CodeResourceNotFound, "The Group Is Not Exist", fmt.Errorf("Try To Set Mute By Uninjion User"), newerror.LevelInfo)
	}
	if groupWithUserStruct.Info[0].Role == commonmodel.Member {
		return newerror.MakeError(http.StatusForbidden, newerror.CodePermissionDenied, "You Are Not Manager", fmt.Errorf("Try To Realse Mute Without Enough Permission"), newerror.LevelInfo)
	}
	groupMuteStruct := groupmuteinfo.NewStruct(groupID, goalUserID, time.Time{}, "")
	err = dao.Delete(ctx, groupMuteStruct, m.dbContext)
	if err != nil {
		return err
	}
	groupWithUserStruct = groupwithuser.NewStruct(groupID, 0, "", "", time.Time{}, groupwithuser.WithGroupID)
	_, err = dao.Get(ctx, groupWithUserStruct, m.dbContext)
	if err != nil {
		return err
	}
	memberList := make([]int64, 0, len(groupWithUserStruct.Info))
	for _, info := range groupWithUserStruct.Info {
		memberList = append(memberList, info.UserID)
	}
	groupNoticeMessage := commonmodel.KafkaGroupNotice{
		GoalUserID:  memberList,
		SessionID:   groupID,
		Data:        map[string]any{"user_id": userID, "goal_user_id": goalUserID},
		MessageCode: commonmodel.MessageCode_GroupReleaseMute,
	}
	_, _, err = tool.SendKafkaGroupNotice(m.groupNoticeTopic, groupNoticeMessage)
	if err != nil {
		return err
	}
	return nil
}
func (m *Mute) GetMuteStatus(ctx context.Context, userID int64, groupID int64) (isMute bool, err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("group_mute:GetMuteStatus")
	groupMuteStruct := groupmuteinfo.NewStruct(groupID, userID, time.Time{}, "", groupmuteinfo.WithWhereUserID)
	exist, err := dao.Get(ctx, groupMuteStruct, m.dbContext)
	if err != nil {
		return true, err
	}
	if !exist {
		return false, nil
	}
	if groupMuteStruct.Info[0].MuteEndTime.Unix() > time.Now().Unix() {
		return true, nil
	} else {

		return false, nil
	}
}
