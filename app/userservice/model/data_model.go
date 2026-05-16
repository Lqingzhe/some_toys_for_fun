package model

import (
	"time"

	"gorm.io/gorm"
)

type UserInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID        int64  `gorm:"primary_key;"`
	UserName      string `gorm:"NOT NULL"`
	Introduction  string
	BirthdayYear  int64
	BirthdayMonth int64
	BirthdayDay   int64
}

func (_ *UserInfo) Data() {}

type UserLoginInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID   int64  `gorm:"primary_key;"`
	Password string `gorm:"NOT NULL"`
	Salt     string `gorm:"NOT NULL"`
}

func (_ *UserLoginInfo) Data() {}

type RemarkInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID     int64 `gorm:"uniqueIndex:idx_user_id;NOT NULL;"`
	GoalUserID int64 `gorm:"uniqueIndex:idx_user_id;NOT NULL;"`
	NickName   string
}

func (_ *RemarkInfo) Data() {}
