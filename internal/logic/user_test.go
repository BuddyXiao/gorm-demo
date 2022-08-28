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
	result := dao.DB.Session(&gorm.Session{DryRun: true}).Create(&user).Statement

	fmt.Println(result.SQL.String())
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
	id := 1563001523807162371
	result := dao.DB.Delete(&po.User{}, id)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
}

func assertCreate(t *testing.T, result *gorm.DB) {
	t.Helper()
	if result.Error != nil {
		t.Errorf("create user err: %s", result.Error.Error())
	}
	fmt.Println(result.RowsAffected)

}
