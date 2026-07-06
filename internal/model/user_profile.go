package model

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	Uid       uint64         `gorm:"primaryKey" json:"uid"`
	Username  string         `gorm:"size:16;not null;uniqueIndex:idx_profiles_username" json:"username"`
	Nickname  string         `gorm:"size:16;not null" json:"nickname"`
	Signature string         `gorm:"size:32" json:"signature"`
	Email     string         `gorm:"size:64" json:"email"`
	Exp       int32          `gorm:"default:0;index:idx_profiles_exp,sort:desc" json:"exp"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
