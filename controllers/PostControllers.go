package controllers

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-282/hngstage2task/initializers"
	"github.com/pro-282/hngstage2task/models"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	// Get data off req body
	var body struct{ Name string }
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body, Name must be a string"})
		return
	}

	// Generate a new UUID
	id := uuid.New()

	// Create a User
	user := models.User{ID: id, Name: body.Name}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create user"})
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
	param := c.Param("id")

	// Check if the param is a valid UUID
	if isValidUUID(param) {
		// Fetch user by ID
		var user models.User
		result := initializers.DB.First(&user, "id = ?", param)

		if result.Error != nil {
			c.JSON(400, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	} else {
		// Fetch user by Username
		var user models.User
		result := initializers.DB.First(&user, "name = ?", param)

		if result.Error != nil {
			c.JSON(400, gin.H{"error": "User not found"})
			return
		}

		c.JSON(200, user)
	}
}

func UpdateUserName(c *gin.Context) {
	// Get Id or Username from request
	param := c.Param("id")

	// Check if the param is a valid UUID
	if isValidUUID(param) {
		// Update user by ID
		var user models.User
		result := initializers.DB.First(&user, "id = ?", param)

		if result.Error != nil {
			c.JSON(400, gin.H{"error": "User not found"})
			return
		}

		// Get the data off req body
		var body struct{ Name string }
		if err := c.Bind(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Update user's name
		user.Name = body.Name
		initializers.DB.Save(&user)

		c.JSON(200, user)
	} else {
		// Update user by Username
		var user models.User
		result := initializers.DB.First(&user, "name = ?", param)

		if result.Error != nil {
			c.JSON(400, gin.H{"error": "User not found"})
			return
		}

		// Get the data off req body
		var body struct{ Name string }
		if err := c.Bind(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Update user's name
		user.Name = body.Name
		initializers.DB.Save(&user)

		c.JSON(200, user)
	}
}

func UserDelete(c *gin.Context) {
	// Get the id or username
	param := c.Param("id")

	var result *gorm.DB

	if isValidUUID(param) {
		// Delete the User by id
		result = initializers.DB.Delete(&models.User{}, "id = ?", param)
	} else {
		// Delete the User by name
		result = initializers.DB.Delete(&models.User{}, "name = ?", param)
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to delete user"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Respond
	c.Status(200)
}

// Function to check if a string is a valid UUID
func isValidUUID(s string) bool {
	// This regex pattern checks for a valid UUID format
	// Example UUID format: 123e4567-e89b-12d3-a456-426655440000
	uuidPattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

	// Use regular expression to match the pattern
	match, _ := regexp.MatchString(uuidPattern, s)
	return match
}
