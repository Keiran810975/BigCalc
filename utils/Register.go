package utils

import (
	"bigcalc/database"
	"bigcalc/models"
)

func IsRegistered(email string) bool {
	if err := database.Db.Where("email = ?", email).First(&models.User{}).Error; err != nil {
		return false
	}
	return true
}

// 判断两个密码是不是一样
func Confirm(password, passwordConfirm string) bool {
	return password == passwordConfirm
}
