package routers

import (
	"LoveDiary/api"
	"LoveDiary/middleware"
	"LoveDiary/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 设置静态资源目录
	r.Static("/uploads", "./uploads")

	auth := r.Group("api")
	auth.Use(middleware.JwtToken())
	{
		auth.GET("parse/token", api.GetTokenInfo)
		auth.GET("user/info", api.GetUserInfo)
		auth.PUT("user/changepw", api.ChangePassword)
		auth.PUT("user/edit", api.EditUser)
	}

	router := r.Group("api")
	{
		//router.POST("upload", api.UploadImage)
		router.GET("users", api.GetUsers)
		router.POST("login", api.Login)
		router.GET("test", api.Test)
		router.POST("user/add", api.CreateUser)
		router.PUT("user/:id", api.EditUserByAdmin)
		router.DELETE("user/:id", api.DeleteUser)
		router.GET("user/:id", api.GetUserInfoByAdmin)
		router.GET("user/search", api.SearchUsers)
		router.PUT("user/changepw/:id", api.ChangePasswordByAdmin)
		router.GET("user", api.GetUser)
		router.PUT("user/update/:id", api.UpdateUserInfo)
	}

	_ = r.Run(utils.HttpPort)
}
