package handler

import (
	"aim/kitex_gen/kitexgroupservie"
	newlog "aim/pkg/log"
	"context"
)

func (s *GroupServiceImpl) GetGroupAndSessionID(ctx context.Context, req *kitexgroupservie.GetGroupAndSessionIDReq) (resp *kitexgroupservie.GetGroupAndSessionIDResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}
