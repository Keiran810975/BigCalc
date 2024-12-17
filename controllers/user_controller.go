package controllers

import (
	"bigcalc/database"
	"bigcalc/models"
	"bigcalc/status"
	"bigcalc/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	//将客户端请求中的 JSON 数据绑定到 Go 中的 user 结构体实例
	fmt.Println("66666666")
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("hahahahaha")
	//如果数据库中已经存在相同的UserID
	if utils.IsRegistered(user.Email) {
		fmt.Println("Register Falied: User already exists")
		return
	}

	//俩密码不一样
	if !utils.Confirm(user.Password, user.ConfirmPassword) {
		fmt.Println("Register Falied: Passwords do not match")
		fmt.Println(user.Password)
		fmt.Println(user.ConfirmPassword)
		return
	}

	//保存用户信息到数据库
	if err := database.Db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	fmt.Println("保存成功")

	//注册成功
	fmt.Println("xixixixixixi")
	c.JSON(200, gin.H{"message": "User registered successfully"})

}

func Login(c *gin.Context) {
	var user models.User
	//将客户端请求中的 JSON 数据绑定到 Go 中的 user 结构体实例
	fmt.Println("dsdasd")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//未注册
	if !utils.IsRegistered(user.Email) {
		fmt.Println("用户未注册")
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	//从数据库里找这个用户
	var dbUser models.User
	if err := database.Db.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		fmt.Println("没找到用户")
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	//密码不对
	if !utils.PasswordCorrect(user.Password, dbUser.Password) {
		fmt.Println("密码不对")
		c.JSON(401, gin.H{"error": "Incorrect password"})
		return
	}

	//登录成功
	dbUser.IsLogin = true

	//更新登录状态
	if err := database.Db.Save(&dbUser).Error; err != nil {
		fmt.Println("保存用户失败")
		c.JSON(500, gin.H{"error": "Failed to save user"})
		return
	}

	//更新当前用户信息
	status.CurrentUserId = dbUser.UserId
	fmt.Println(status.CurrentUserId)
	fmt.Println("登录成功！")
	c.JSON(200, gin.H{"message": "User logged in successfully"})

}
