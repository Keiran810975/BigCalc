package utils
import (
	"bigcalc/database"
	"bigcalc/models"
)

func IsRegistered(id uint) bool {
    if err:= database.Db.Where("id = ?", id).First(&models.User{}).Error; err != nil {
        return false
    }
    return true
}

//判断两个密码是不是一样
func Confirm(password, passwordConfirm string) bool {
    return password == passwordConfirm
}