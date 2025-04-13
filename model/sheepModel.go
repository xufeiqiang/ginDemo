package model

import "time"

type SheepInfos struct {
	// gorm 会默认根据结构体名来推导表名sheep_infos
	UUID       string `gorm:"primaryKey"` // 主键
	Code       string `gorm:"size:36;not null"`
	Category   string `gorm:"size:36;not null"`
	Age        uint   `gorm:"not null"`
	Birth      time.Time
	Avatar     []byte // 头像
	CreateTime time.Time
	CreateUser string
	UpdateTime time.Time
	UpdateUser string
}
