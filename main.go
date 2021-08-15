package main

import (
	"golang-api-example/database"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//Configuration
	database.InitDatabase()

	apiName := os.Getenv("API_NAME")
	apiVersion := os.Getenv("API_VERSION")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api":     apiName,
			"version": apiVersion,
		})
	})

	r.POST("/users", func(c *gin.Context) {

		db := database.GetDatabase()

		var user database.User

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "cannot bind json: " + err.Error(),
			})

			return
		}

		err = db.Create(&user).Error
		if err != nil {
			c.JSON(400, gin.H{
				"error": "cannot create Client: " + err.Error(),
			})

			return
		}

		c.JSON(201, user)

	})

	r.GET("/users", func(c *gin.Context) {
		db := database.GetDatabase()

		var users []database.User
		err := db.Find(&users).Error
		if err != nil {
			c.JSON(400, gin.H{
				"error": "cannot list Clients: " + err.Error(),
			})

			return
		}

		c.JSON(200, users)
	})

	r.Run()
}
