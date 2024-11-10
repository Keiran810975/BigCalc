package bigcalc

import (
	"bigcalc/config"
	"bigcalc/controllers"
	"bigcalc/database"
	"bigcalc/utils"
	"github.com/gin-gonic/gin"
)


func main(){
	config.InitConfig()
	database.InitDB()
	r:=gin.Default()
	r.POST("/register",controllers.Register)
	r.POST("/login",controllers.Login)
	r.PUT("/changepassword",utils.ChangePassword)
	r.POST("/logout",utils.Logout)
	r.Run()
}