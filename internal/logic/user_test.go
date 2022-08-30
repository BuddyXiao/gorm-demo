package logic

import (
	"fmt"
	"gorm.io/gorm"
	"orm/internal/dao"
	"orm/internal/model/dto"
	"orm/internal/model/po"
	"orm/internal/utils"
	"testing"
)

func TestCreateUser(t *testing.T) {
	dao.Init()
	user := dto.NewUser("xiaoming", "11111111",
		"xiaoming", 155555555, "湖北", "武汉", "关谷", 1, "关谷一号")
	var usrPo = utils.UserDto2Po(user)
	result := dao.DB.Session(&gorm.Session{DryRun: true}).Create(&usrPo).Statement

	fmt.Println(result.SQL.String())
	//assertCreate(t, result)
}

func TestCreateSimple(t *testing.T) {
	dao.Init()
	user := po.User{
		Username: "Jinzhu",
		Phone:    18,
	}
	result := dao.DB.Session(&gorm.Session{DryRun: false}).Create(&user)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
}

func TestUpdateSimple(t *testing.T) {
	dao.Init()
	var user po.User
	user.ID = 1563001523807162372
	user.Province = "湖北省"
	result := dao.DB.Model(&user).Updates(&user)
	fmt.Println(result.Error)
}

func TestDeleteSimple(t *testing.T) {
	dao.Init()
	id := 1563001523807162373
	var user po.User
	user.ID = id
	result := dao.DB.Session(&gorm.Session{DryRun: false}).Delete(&user)
	fmt.Println(result.Statement.SQL.String())
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
}

func TestQuerySimple(t *testing.T) {
	dao.Init()
	var user = po.User{}
	user.ID = 1563001523807162372
	result := dao.DB.Session(&gorm.Session{DryRun: false}).Find(&user).Statement
	fmt.Println(result.SQL.String())
}

func TestQueryList(t *testing.T) {
	dao.Init()
	var users []po.User
	result := dao.DB.Session(&gorm.Session{DryRun: true}).Find(&users).Statement
	fmt.Println(result.SQL.String())
}

func TestAssociation(t *testing.T) {
	dao.Init()
	var user = po.User{}
	user.ID = 1563001523807162370
	var orleIds []int
	dao.DB.Model(&po.UserRole{}).Where("user_id = ?", user.ID).Select("role_id").Find(&orleIds)
	fmt.Println("orleIds:", orleIds)
	var orles []po.Role
	dao.DB.Where("deleted = 0").Find(&orles, "id in ?", orleIds)
	fmt.Println("roles:", orles)
	dao.DB.First(&user)
	user.Roles = orles
	//dao.DB.Preload("Roles").Where("deleted > 0").Find(&user) 无法使用预加载
	fmt.Println(user)
}

func TestAssociationDryRun(t *testing.T) {
	dao.Init()
	db := dao.DB.Session(&gorm.Session{DryRun: false})
	var user = po.User{}
	user.ID = 1563001523807162370
	var orleIds []int
	s1 := db.Model(&po.UserRole{}).Where("user_id = ?", user.ID).Select("role_id").Find(&orleIds).Statement
	fmt.Println("orleIds:", orleIds, " , sql:", s1.SQL.String())
	var roles []po.Role
	s2 := db.Where("deleted = 0").Where("id in ?", orleIds).Find(&roles).Statement
	fmt.Println("roles:", roles, ",sql:", s2.SQL.String())
	s3 := db.First(&user).Statement
	user.Roles = roles
	fmt.Println("users", user, "\nsql:", s3.SQL.String())
	//dao.DB.Preload("Roles").Where("deleted > 0").Find(&user) 无法使用预加载
}

func TestJoins(t *testing.T) {
	dao.Init()
	db := dao.DB.Session(&gorm.Session{DryRun: true})
	var user = po.User{}
	user.ID = 1563001523807162370
	db.Joins("")
}
func assertCreate(t *testing.T, result *gorm.DB) {
	t.Helper()
	if result.Error != nil {
		t.Errorf("create user err: %s", result.Error.Error())
	}
	fmt.Println(result.RowsAffected)

}
