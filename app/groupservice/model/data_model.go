package model

import (
	"aim/commonmodel"
	"time"

	"gorm.io/gorm"
)

type GroupWithUserInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	GroupID         int64
	UserID          int64
	GroupRemarkName string
	Role            commonmodel.GroupRole

	LastReadTime time.Time
}

func (_ *GroupWithUserInfo) Data() {}

type GroupInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	GroupID   int64
	GroupName string
}

func (_ *GroupInfo) Data() {}

type SessionInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	SessionID    int64
	UserID       int64
	GoalUserID   int64
	LastReadTime time.Time
}

func (_ *SessionInfo) Data() {}

type GroupApplyInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	GoalID      int64
	ApplyUserID int64
}

func (_ *GroupApplyInfo) Data() {}

type GroupMuteInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	GroupID     int64
	UserID      int64
	MuteEndTime time.Time
	MuteReason  string
}

func (_ *GroupMuteInfo) Data() {}
