package database
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"bigcalc/models"
)

var Db *gorm.DB

func InitDB(){
	dsn := "root:keiran114514@tcp(127.0.0.1:3306)/bigcalc?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&models.User{})
}