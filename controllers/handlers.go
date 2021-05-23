package controllers

import (
	"net/http"

	"github.com/Lisnyk-M/go-endpoints/api"
	"github.com/Lisnyk-M/go-endpoints/models"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	var users []models.User
	api.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
