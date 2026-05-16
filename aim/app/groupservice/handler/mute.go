package handler

import (
	"aim/kitex_gen/kitexgroupservie"
	newlog "aim/pkg/log"
	"context"
)

func (s *GroupServiceImpl) SetMute(ctx context.Context, req *kitexgroupservie.SetMuteReq) (resp *kitexgroupservie.SetMuteResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// ReleaseMute implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) ReleaseMute(ctx context.Context, req *kitexgroupservie.ReleaseMuteReq) (resp *kitexgroupservie.ReleaseMuteResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// GetMuteStatus implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) GetMuteStatus(ctx context.Context, req *kitexgroupservie.GetMuteStatusReq) (resp *kitexgroupservie.GetMuteStatusResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}
