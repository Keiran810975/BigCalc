package utils

import (
	"bigcalc/database"
	"bigcalc/models"
	"bigcalc/status"
	"github.com/gin-gonic/gin"
)

//修改密码
func ChangePassword(c *gin.Context){
	var user models.User
	id:=status.CurrentUserId
	if err:=database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found",})
		return
	}
	
	var newpassword string
	if err:= c.ShouldBindJSON(&newpassword); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request",})
		return
	}
	user.Password = newpassword

	if err:=database.Db.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update password",})
		return
	}
}

//退出登录
func Logout(c *gin.Context){
	var user models.User
	id:=status.CurrentUserId
	if err:=database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found",})
		return
	}
    user.IsLogin = false
    if err:=database.Db.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to logout",})
		return
    }

	c.JSON(200, gin.H{"message": "Logout successfully",})
}

//上传头像
func UploadAvatar(c *gin.Context){
	var user models.User
	id:=status.CurrentUserId
	if err:=database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found",})
		return
	}

	//未完成


}