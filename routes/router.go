package routes

import (
	"gin_jwt_demo/controller"
	"gin_jwt_demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 获取 token
	// /login?username=username&passwor=password
	r.POST("/login", controller.Login)

	api := r.Group("/api")
	api.Use(middleware.JWT())
	{
		// 查询
		// /api/user?username=username
		api.GET("/user", controller.GetUser)
		// 增加
		// /api/user?username=username&password=password
		api.POST("/user", controller.NewUser)
		//修改
		// /api/user/:username/edit?password=password
		api.POST("/user/:username/edit", controller.ChangePassword)
		// 删除
		// /api/user?username=username
		api.DELETE("/user", controller.DeleteUser)
	}

	return r
}
