package main

import (
	"github.com/Lisnyk-M/go-endpoints/http_func"
	// "github.com/Lisnyk-M/go-endpoints/models"
	"github.com/Lisnyk-M/go-endpoints/api"
	"github.com/Lisnyk-M/go-endpoints/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	api.Db_connection()
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUserById)
	r.DELETE("/user/:id", controllers.DeleteUserById)
	r.POST("/auth/register", controllers.Register)
	r.GET("/auth/get-email", controllers.GetEmail)

	http_func.Send()
	r.Run()
}
