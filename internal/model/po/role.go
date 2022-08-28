package po

type Role struct {
	BaseModel
	Name string
	Code string
}

func (Role) TableName() string {
	return "t_role"
}
