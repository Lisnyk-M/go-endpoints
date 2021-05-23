package main

import (
	"github.com/Lisnyk-M/go-endpoints/http_func"
	// "github.com/Lisnyk-M/go-endpoints/models"
	// "github.com/Lisnyk-M/go-endpoints/controllers"
	// "github.com/Lisnyk-M/go-endpoints/api"
	"github.com/gin-gonic/gin"
)

// func db_connection() *gorm.DB {
// 	if err := godotenv.Load(); err != nil {
// 		fmt.Println("No .env file found")
// 	}
// 	postgresql_string, exists := os.LookupEnv("POSTGRESQL_CONNECTION")

// 	if exists {
// 		fmt.Println(postgresql_string)
// 	} else {
// 		fmt.Println(".env not exist")
// 	}

// 	pgUrl, err := pq.ParseURL(postgresql_string)
// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	db, err := gorm.Open("postgres", pgUrl)

// 	if err != nil {
// 		panic("Failed to connect to database! 222222222222222")
// 	}
// 	return db
// }

func main() {
	db := api.db_connection()
	// db.AutoMigrate(&models.User{})
	r := gin.Default()

	// r.GET("/user/:id", func(c *gin.Context) {
	// 	var user models.User
	// 	id := c.Param("id")
	// 	res := db.First(&user, id)

	// 	if res.Error != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "Not found :)"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, user)
	// })

	// r.GET("/users", func(c *gin.Context) {
	// 	var users []models.User
	// 	db.Find(&users)
	// 	c.JSON(http.StatusOK, users)
	// })
	r.GET("/users", controllers.getUsers())

	// r.DELETE("/user/:id", func(c *gin.Context) {
	// 	var user models.User
	// 	id := c.Param("id")

	// 	res := db.First(&user, id)
	// 	if res.Error != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": "Not found :)"})
	// 		return
	// 	}

	// 	result := db.Delete(&user, "email = ?", user.Email)
	// 	if result.Error != nil {
	// 		c.JSON(http.StatusOK, res.Error)
	// 		return
	// 	}

	// 	c.JSON(200, gin.H{"message": "deleted user " + user.Email + " successfully!"})
	// })

	// r.POST("/auth/register", func(c *gin.Context) {
	// 	var user models.User
	// 	type RegisterForm struct {
	// 		User     string `json:"user" binding:"required"`
	// 		Email    string `json:"email" binding:"required"`
	// 		Password string `json:"password" binding:"required"`
	// 	}
	// 	var form RegisterForm

	// 	errr := c.ShouldBindJSON(&form)
	// 	if errr != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": errr.Error()})
	// 		return
	// 	}

	// 	res := db.First(&user, "email = ?", form.Email)
	// 	if res.Error == nil {
	// 		c.JSON(400, gin.H{"error": "user allready exists"})
	// 		return
	// 	}

	// 	result := db.Create(&models.User{Name: form.User, Email: form.Email, Password: form.Password})
	// 	if result.Error != nil {
	// 		c.JSON(400, gin.H{"error": res.Error})
	// 		return
	// 	}

	// 	c.JSON(201, gin.H{"data": res.RowsAffected})
	// })
	http_func.Send()

	r.Run()
}
