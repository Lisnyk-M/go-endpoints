package controllers

import (
	"net/http"

	"github.com/Lisnyk-M/go-endpoints/api"
	"github.com/Lisnyk-M/go-endpoints/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	api.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	res := api.DB.First(&user, id)

	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found :)"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	res := api.DB.First(&user, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found :)"})
		return
	}

	result := api.DB.Delete(&user, "email = ?", user.Email)
	if result.Error != nil {
		c.JSON(http.StatusOK, res.Error)
		return
	}

	c.JSON(200, gin.H{"message": "deleted user " + user.Email + " successfully!"})
}

func Register(c *gin.Context) {
	var user models.User
	var form models.RegisterForm

	errr := c.ShouldBindJSON(&form)
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errr.Error()})
		return
	}

	res := api.DB.First(&user, "email = ?", form.Email)
	if res.Error == nil {
		c.JSON(400, gin.H{"error": "user allready exists"})
		return
	}

	result := api.DB.Create(&models.User{Name: form.User, Email: form.Email, Password: form.Password})
	if result.Error != nil {
		c.JSON(400, gin.H{"error": res.Error})
		return
	}

	c.JSON(201, gin.H{"data": res.RowsAffected})
}

func GetEmail(c *gin.Context) {
	var user models.User
	email := c.Param("email")

	res, err := api.DB.First(&user, "email=?")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
