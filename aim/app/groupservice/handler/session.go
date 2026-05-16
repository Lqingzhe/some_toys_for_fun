package handler

import (
	"aim/kitex_gen/kitexgroupservie"
	newlog "aim/pkg/log"
	"context"
)

func (s *GroupServiceImpl) CreatSession(ctx context.Context, req *kitexgroupservie.CreatSessionReq) (resp *kitexgroupservie.CreatSessionResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteSession implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) DeleteSession(ctx context.Context, req *kitexgroupservie.DeleteSessionReq) (resp *kitexgroupservie.DeleteSessionResp, err error) {
	// TODO: Your code here...
	return
}

// GetFriendLastVisitTime implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) GetFriendLastVisitTime(ctx context.Context, req *kitexgroupservie.GetFriendLastVisitTimeReq) (resp *kitexgroupservie.GetFriendLastVisitTimeResp, err error) {
	// TODO: Your code here...
	return
}

// ApplyForFriend implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) ApplyForFriend(ctx context.Context, req *kitexgroupservie.ApplyForFriendReq) (resp *kitexgroupservie.ApplyForFriendResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// GetFriendApplyList implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) GetFriendApplyList(ctx context.Context, req *kitexgroupservie.GetFriendApplyListReq) (resp *kitexgroupservie.GetFriendApplyListResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// RefuseFriendApply implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) RefuseFriendApply(ctx context.Context, req *kitexgroupservie.RefuseFriendApplyReq) (resp *kitexgroupservie.RefuseFriendApplyResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}
