package po

type UserRole struct {
	BaseModel
	UserID int `gorm:"column:user_id"`
	RoleID int `gorm:"column:role_id"`
}

func (UserRole) TableName() string {
	return "t_user_role"
}
