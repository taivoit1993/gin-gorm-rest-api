package main

import (
	"gin-gorm-rest-api/config"
	userroute "gin-gorm-rest-api/internal/user/routes"
	"github.com/gin-gonic/gin"
)

func RootRoute(route *gin.Engine) {
	userroute.UserRoute(route)
}

func main() {
	router := gin.Default()
	//connect db
	config.Connect()
	//load route
	RootRoute(router)
	//start server
	router.Run(":9000")
}
