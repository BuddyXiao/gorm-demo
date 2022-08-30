package po

type Role struct {
	BaseModel
	Name    string `gorm:"column:name"`
	Code    string `gorm:"column:code"`
	Deleted int    `gorm:"column:deleted"`
	// 对对多关系
	Users []User `gorm:"many2many:t_user_role"`
}

func (Role) TableName() string {
	return "t_role"
}
