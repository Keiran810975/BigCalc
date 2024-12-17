package database
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"bigcalc/models"
	"fmt"
)

var Db *gorm.DB
var CanvasDb *gorm.DB

func InitUserDB(){
	dsn := "root:keiran114514@tcp(127.0.0.1:3306)/bigcalc?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic("failed to connect database")
	}

	//设置所有用户的 is_login=0
	if err := Db.Model(&models.User{}).Where("1 = 1").Update("is_login", 0).Error; err != nil {
		fmt.Println("Failed to update is_login:", err)
	}

	Db.AutoMigrate(&models.User{})
}

func InitCanvasDB() {
	dsn := "root:keiran114514@tcp(127.0.0.1:3306)/bigcalc?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	CanvasDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect to canvas database")
	}
	CanvasDb.AutoMigrate(&models.CanvasData{})
}