package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-282/hngstage2task/initializers"
	"github.com/pro-282/hngstage2task/models"
)

func PostCreate(c *gin.Context) {
	// Get data off req body
	var body struct{ Name string }
	c.Bind(&body)

	// Generate a new UUID
	id := uuid.New()

	// Create a User
	user := models.User{ID: id, Name: body.Name}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, user)
}

func FetchAllUsers(c *gin.Context) {
	// Get the Users
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, users)
}

func FetchUserById(c *gin.Context) {
	// Get Id from request
	id := c.Param("id")

	// Get the user
	var user models.User
	result := initializers.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		log.Fatal(result.Error)
		c.Status(400)
		return
	}

	// Respond with result
	c.JSON(200, user)
}

func UpdateUserName(c *gin.Context) {
	// Get Id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct{ Name string }
	c.Bind(&body)

	// find the user we are updating
	var user models.User
	result := initializers.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		log.Fatal(result.Error)
		c.Status(400)
	}

	// update it
	user.Name = body.Name
	initializers.DB.Save(&user)

	// Respond with it
	c.JSON(200, user)
}

func UserDelete(c *gin.Context) {
	// Get the id of the post
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.User{}, "id = ?", id)

	// Respond
	c.Status(200)
}
