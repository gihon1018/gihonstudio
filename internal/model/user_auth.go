package model

import (
	"time"

	"gorm.io/gorm"
)

type UserAuth struct {
	Uid          uint64         `gorm:"primaryKey;autoIncrement" json:"uid"`
	Username     string         `gorm:"size:16;not null;uniqueIndex:idx_auths_username" json:"username"`
	PasswordHash string         `gorm:"size:256;not null" json:"password_hash"`
	AccessToken  string         `gorm:"size:512" json:"access_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
}
