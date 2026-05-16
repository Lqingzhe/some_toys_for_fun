package groupmuteinfo

import (
	"aim/app/groupservice/model"
	newerror "aim/pkg/error"
	"aim/tool"
	"context"
	"time"
)

type MuteInfo struct {
	model.GroupMuteInfo
	Info            []*model.GroupMuteInfo
	whereWithUserID bool
}
type operate func(*MuteInfo)

func NewStruct(groupID int64, userID int64, muteEndTime time.Time, muteReason string, Operates ...operate) *MuteInfo {
	newStruct := &MuteInfo{
		GroupMuteInfo: model.GroupMuteInfo{
			GroupID:     groupID,
			UserID:      userID,
			MuteEndTime: muteEndTime,
			MuteReason:  muteReason,
		},
	}
	for _, Operate := range Operates {
		Operate(newStruct)
	}
	return newStruct
}
func WithWhereUserID(muteInfo *MuteInfo) {
	muteInfo.whereWithUserID = true
}
func (m *MuteInfo) AddInfo(ctx context.Context, dbContext any) (err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("db_manage:AddInfo")
	DB, err := tool.TypeAssert[model.DBContext](dbContext)
	if err != nil {
		return err
	}

}
func (m *MuteInfo) GetInfo(ctx context.Context, dbContext any) (exist bool, err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("db_manage:GetInfo")
	DB, err := tool.TypeAssert[model.DBContext](dbContext)
	if err != nil {
		return false, err
	}

}
func (m *MuteInfo) UpdateInfo(ctx context.Context, dbContext any) (exist bool, err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("db_manage:UpdateInfo")
	DB, err := tool.TypeAssert[model.DBContext](dbContext)
	if err != nil {
		return false, err
	}

}
func (m *MuteInfo) DeleteInfo(ctx context.Context, dbContext any) (err error) {
	defer func(trace string) {
		err = newerror.TranslateError(err).AddErrorTrace(trace)
	}("db_manage:DeleteInfo")
	DB, err := tool.TypeAssert[model.DBContext](dbContext)
	if err != nil {
		return err
	}

}
