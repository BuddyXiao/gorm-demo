package utils

import (
	"orm/internal/model/dto"
	"orm/internal/model/po"
	"reflect"
)

func UserDto2Po(user *dto.User) po.User {
	var usr po.User
	if !isEmpty(user.Username) {
		usr.Username = user.Username
	}
	if !isEmpty(user.Password) {
		usr.Password = user.Password
	}
	if !isEmpty(user.RealName) {
		usr.RealName = user.RealName
	}
	if !isEmpty(user.Province) {
		usr.Province = user.Province
	}
	if !isEmpty(user.County) {
		usr.County = user.County
	}
	if !isEmpty(user.Phone) {
		usr.Phone = user.Phone
	}
	if !isEmpty(user.UnitId) {
		usr.UnitId = user.UnitId
	}
	if !isEmpty(user.UnitName) {
		usr.UnitName = user.UnitName
	}
	if !isEmpty(user.City) {
		usr.City = user.City
	}
	return usr
}

func isEmpty(t any) bool {
	val := reflect.ValueOf(t)
	switch val.Kind() {
	case reflect.Int, reflect.Int64, reflect.Int32:
		if val.Int() == 0 {
			return true
		}
	case reflect.String:
		if val.String() == "" {
			return true
		}
	default:
		return false
	}
	return false
}
