package controllers

import (
	"bigcalc/database"
	"bigcalc/models"
	"bigcalc/status"
	"fmt"
	"bigcalc/utils"

	"github.com/gin-gonic/gin"
)

var History struct {
	CanvasData string `json:"canvasData"`
}

//获取每个用户自己的表名
func getCanvasTableName(userID uint) string {
	return fmt.Sprintf("canvas_user_%d", userID)
}

// 检查表是否存在
func isTableExist(tableName string) bool {
	var count int64
	database.CanvasDb.Table("bigcalc." + tableName).Count(&count)
	return count > 0
}

// 创建表
func createTable(tableName string) error {
	// 根据模型动态创建表（示例是 CanvasData 模型，你可以根据需要调整）
	err := database.CanvasDb.Table("bigcalc." + tableName).AutoMigrate(&models.CanvasData{})
	return err
}

func GetHistory(c *gin.Context) {
	fmt.Println("GetHistory")
	// requestBody.CanvasData = c.PostForm("canvasData")
	if err := c.ShouldBindJSON(&History); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
		return
	}
	fmt.Println(History.CanvasData)
	var newCanvasData models.CanvasData
	newCanvasData.CanvasData = History.CanvasData

	if(status.CurrentUserId == 0) {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		fmt.Println("未登录")
		fmt.Println(status.CurrentUserId)
		return
	}

	newCanvasData.UserID=status.CurrentUserId
	newCanvasData.CreatedAt=utils.GetCurrentTime()

	tableName:=getCanvasTableName(status.CurrentUserId)

	// 检查表是否存在
	if !isTableExist(tableName) {
		// 如果表不存在，创建表
		if err := createTable(tableName); err != nil {
			c.JSON(500, gin.H{"error": "Failed to create table"})
			return
		}
	}

	if err := database.CanvasDb.Table(tableName).Create(&newCanvasData).Error; err != nil {
		fmt.Println("保存画板失败", err)
		c.JSON(500, gin.H{"error": "Failed to create canvas data"})
		return
	}
	fmt.Println("保存画板成功")

	c.JSON(200, gin.H{
		"message": "Canvas data received successfully",
	})

}

func GetCanvas(c *gin.Context) {
	// 检查用户是否已登录
	if status.CurrentUserId == 0 {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		fmt.Println("未登录")
		return
	}

	tableName:=getCanvasTableName(status.CurrentUserId)

	if !isTableExist(tableName) {
		// 如果表不存在，直接返回空数据
		c.JSON(200, gin.H{
			"canvasData": nil,
			"message":    "No canvas data found for the user",
		})
		return
	}

	// 查询最新的画板数据
	var canvasData models.CanvasData
	err := database.CanvasDb.Table(tableName).First(&canvasData).Error
	if err != nil {
		fmt.Println("获取画板数据失败", err)
		c.JSON(500, gin.H{
			"error": "Failed to retrieve canvas data",
		})
		return
	}
	fmt.Println(canvasData.CanvasData)

	c.JSON(200, gin.H{
		"canvasData": canvasData.CanvasData,
		"message":    "Canvas data retrieved successfully",
	})
}
