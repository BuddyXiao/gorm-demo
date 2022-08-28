package po

type User struct {
	BaseModel
	Username string `gorm:"column:user_name"`
	Password string
	RealName string `gorm:"column:real_name"`
	Phone    int
	Province string
	City     string
	County   string
	UnitId   int    `gorm:"column:unit_id"`
	UnitName string `gorm:"column:unit_name"`
}

func (User) TableName() string {
	return "t_user"
}
