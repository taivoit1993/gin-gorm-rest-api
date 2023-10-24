package userroute

import (
	usercontroller "gin-gorm-rest-api/internal/user/controller"
	"gin-gorm-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.POST("/login", usercontroller.Login)
		user.GET("/", usercontroller.CreateUser)
		user.POST("/", usercontroller.CreateUser)
		user.PUT("/:id", usercontroller.UpdateUser)
		user.DELETE("/:id", usercontroller.DeleteUser)
		user.GET("/validate", middleware.RequireAuth, usercontroller.CheckLoggin)

	}

}
