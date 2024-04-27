package main

import (
	"app/controller"

	"app/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {

	r := gin.Default()

	r.GET("/", controller.HandleHome)
	r.GET("/api/users", controller.HandleGetAllUsers)
	r.GET("/api/user", controller.HandleGetUserById)

	r.POST("/api/adduser", controller.HandleAddUser)
	r.PUT("/api/changeuser", controller.HandleChangeUser)

	r.DELETE("/api/deletuser", controller.HandleDeleteUser)
	r.Run()
}
