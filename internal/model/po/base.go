package po

import "time"

type BaseModel struct {
	ID       int       `gorm:"column:id;primaryKey;autoIncrement"`
	CreateAt time.Time `gorm:"column:gmt_create;autoCreateTime:milli"`
	UpdateAt time.Time `gorm:"column:gmt_modified;autoUpdateTime:milli"`
}
