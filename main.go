package main

import (
	"bigcalc/config"
	"bigcalc/controllers"
	"bigcalc/database"
	"bigcalc/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	database.InitUserDB()
	database.InitCanvasDB()
	// CORS配置
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowOrigins = []string{"http://localhost:3000"}                   // 允许来自前端的请求
	//corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}            // 允许的请求方法
	//corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // 允许的请求头
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}                   // 允许的来源
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的 HTTP 方法，包括 OPTIONS
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // 允许的请求头
	corsConfig.AllowCredentials = true                                            // 如果需要处理带凭证的请求，设置为 true

	r := gin.Default()

	r.Use(cors.New(corsConfig))
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/history", controllers.GetHistory)
	r.GET("/history/last/:id", controllers.GetCanvas)
	r.PUT("/changepassword", utils.ChangePassword)
	r.GET("/history/list", controllers.GetHistoryList)
	r.POST("/logout", utils.Logout)
	r.Run(":8080")
}
