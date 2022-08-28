package dto

type User struct {
	Username string
	Password string
	RealName string
	Phone    int
	Province string
	City     string
	County   string
	UnitId   int
	UnitName string
}

func NewUser(username, password, realname string, phone int, province, city, county string, unitId int, uintName string) *User {
	return &User{
		Username: username,
		Password: password,
		RealName: realname,
		Phone:    phone,
		Province: province,
		City:     city,
		County:   county,
		UnitId:   unitId,
		UnitName: uintName,
	}
}
