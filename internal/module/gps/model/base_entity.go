package model

import "time"

const (
	Deleted = -1
	EXIST   = 0
)

const (
	Enabled  = "enabled"
	Disabled = "disabled"
)

type BaseEntity struct {
	Id       int64     `gorm:"primaryKey"`
	CreateAt time.Time `gorm:"create_at"`
	UpdateAt time.Time `gorm:"update_at"`
	CreateBy time.Time `gorm:"create_by"`
	UpdateBy time.Time `gorm:"update_by"`
	Version  int32     `gorm:"version"`
	Deleted  int32     `gorm:"deleted"`
}

func PrePersist() BaseEntity {
	return BaseEntity{
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Version:  0,
		Deleted:  EXIST,
	}
}

func PreUpdate(entity BaseEntity) BaseEntity {
	return BaseEntity{
		Id:       entity.Id,
		UpdateAt: time.Now(),
		CreateBy: entity.CreateBy,
		CreateAt: entity.CreateAt,
		Version:  entity.Version + 1,
		Deleted:  entity.Deleted,
	}
}
