package service

import (
	"aim/app/groupservice/dao"
	"aim/app/groupservice/dao/groupwithuser"
	"aim/app/groupservice/dao/sessioninfo"
	"aim/app/groupservice/model"
	newerror "aim/pkg/error"
	"context"
	"time"
)

type UserOfAll struct {
	dbContext *model.DBContext
}

func NewUserInfoOfGroup(dbContext *model.DBContext) *UserOfAll {
	return &UserOfAll{
		dbContext: dbContext,
	}
}
func (u *UserOfAll) GetGroupAndSessionID(ctx context.Context, userID int64) (groupID []int64, sessionID []int64, err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("user_of_all:GetGroupAndSessionInfo")
	groupStruct := groupwithuser.NewStruct(0, userID, "", "", time.Time{}, groupwithuser.WithUserID)
	sessionStruct := sessioninfo.NewStruct(0, userID, 0, time.Time{}, sessioninfo.WithUserID)
	exist, err := dao.Get(ctx, groupStruct, u.dbContext)
	if err != nil {
		return nil, nil, err
	}
	if !exist {
		groupID = nil
	} else {
		groupID = make([]int64, len(groupStruct.Info))
		for i, v := range groupStruct.Info {
			groupID[i] = v.GroupID
		}
	}
	exist, err = dao.Get(ctx, sessionStruct, u.dbContext)
	if err != nil {
		return nil, nil, err
	}
	if !exist {
		sessionID = nil
	} else {
		sessionID = make([]int64, len(sessionStruct.Info))
		for i, v := range sessionStruct.Info {
			sessionID[i] = v.SessionID
		}
	}
	return groupID, sessionID, nil
} //未获取到的Info空值为nil
