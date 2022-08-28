package logic

import (
	"fmt"
	"orm/internal/model/dto"
	"orm/internal/utils"
)

func CreateUser(user *dto.User) {
	var usrPo = utils.UserDto2Po(user)
	//result := dao.DB.Create(&usrPo)
	fmt.Println(usrPo)
}
