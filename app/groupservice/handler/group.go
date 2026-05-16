package handler

import (
	"aim/app/groupservice/service"
	"aim/kitex_gen/kitexgroupservie"
	newlog "aim/pkg/log"
	"context"
)

func (s *GroupServiceImpl) GetGroupInfo(ctx context.Context, req *kitexgroupservie.GetGroupInfoReq) (resp *kitexgroupservie.GetGroupInfoResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)
	serviceStruct := service.NewGroup()
	return
}

// ChangeGroupInfo implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) ChangeGroupInfo(ctx context.Context, req *kitexgroupservie.ChangeGroupInfoReq) (resp *kitexgroupservie.ChangeGroupInfoResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// SearchGroup implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) SearchGroup(ctx context.Context, req *kitexgroupservie.SearchGroupReq) (resp *kitexgroupservie.SearchGroupResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// CreateGroup implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) CreateGroup(ctx context.Context, req *kitexgroupservie.CreateGroupReq) (resp *kitexgroupservie.CreateGroupResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// DeleteGroup implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) DeleteGroup(ctx context.Context, req *kitexgroupservie.DeleteGroupReq) (resp *kitexgroupservie.DeleteGroupResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// LeaveGroup implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) LeaveGroup(ctx context.Context, req *kitexgroupservie.LeaveGroupReq) (resp *kitexgroupservie.LeaveGroupResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// SetGroupApply implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) SetGroupApply(ctx context.Context, req *kitexgroupservie.SetGroupApplyReq) (resp *kitexgroupservie.SetGroupApplyResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// GetGroupApplyList implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) GetGroupApplyList(ctx context.Context, req *kitexgroupservie.GetGroupApplyListReq) (resp *kitexgroupservie.GetGroupApplyListResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// GetLastVisitTime implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) GetLastVisitTime(ctx context.Context, req *kitexgroupservie.GetLastVisitTimeReq) (resp *kitexgroupservie.GetLastVisitTimeResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// AgreeGroupApply implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) AgreeGroupApply(ctx context.Context, req *kitexgroupservie.AgreeGroupApplyReq) (resp *kitexgroupservie.AgreeGroupApplyResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// TransformGroupOwner implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) TransformGroupOwner(ctx context.Context, req *kitexgroupservie.TransformGroupOwnerReq) (resp *kitexgroupservie.TransformGroupOwnerResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// KickOutGroup implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) KickOutGroup(ctx context.Context, req *kitexgroupservie.KickOutGroupReq) (resp *kitexgroupservie.KickOutGroupResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// SetManager implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) SetManager(ctx context.Context, req *kitexgroupservie.SetManagerReq) (resp *kitexgroupservie.SetManagerResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// RevokeManager implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) RevokeManager(ctx context.Context, req *kitexgroupservie.RevokeManagerReq) (resp *kitexgroupservie.RevokeManagerResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// GetGroupInfoWithUser implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) GetGroupInfoWithUser(ctx context.Context, req *kitexgroupservie.GetGroupInfoWithUserReq) (resp *kitexgroupservie.GetGroupInfoWithUserResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}

// UpdateGroupInfoWithUser implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) UpdateGroupInfoWithUser(ctx context.Context, req *kitexgroupservie.UpdateGroupInfoWithUserReq) (resp *kitexgroupservie.UpdateGroupInfoWithUserResp, err error) {
	logger := newlog.AddTraceAndEquipID(s.Logger, req.CommonInfo.Trace, s.EquipID)

	return
}
