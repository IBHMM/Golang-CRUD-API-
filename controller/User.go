package controller

import (
	"app/config"
	"app/helper"
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleGetAllUsers(c *gin.Context) {

	defer config.Disconnect()
	err := config.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var database *gorm.DB = config.DB
	var User []model.User
	result := database.Find(&User)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, User)
}

func HandleAddUser(c *gin.Context) {
	defer config.Disconnect()
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Body"})
		return
	}

	if len(user.Email) < 8 && len(user.Password) < 3 && len(user.Username) < 3 {
		c.JSON(400, gin.H{"error": "Invalid Cridentials", "message": user})
		return
	}

	if err := config.Connect(); err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	InsertedUser := config.DB.Create(&user)

	if InsertedUser.Error != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(201, gin.H{"error": nil, "message": user})
}

func HandleGetUserById(c *gin.Context) {
	defer config.Disconnect()

	UserID, err := helper.ParseInt(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var user model.User

	if err := config.Connect(); err != nil {
		c.JSON(500, gin.H{"error": "Problem During Db connection"})
		return
	}

	if result := config.DB.Where("id=?", UserID).First(&user); result.Error != nil {
		c.JSON(500, gin.H{"error": "Error Occurred"})
		return
	}

	if len(user.Username) < 3 {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}

	c.JSON(200, gin.H{"error": nil, "message": user})
}

func HandleChangeUser(c *gin.Context) {
	if err := config.Connect(); err != nil {
		c.JSON(500, gin.H{"error": "Problem During Db connection"})
		return
	}
	defer config.Disconnect()

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err, "message": "Invalid request body"})
		return
	}

	var existingUser model.User
	if result := config.DB.First(&existingUser, user.ID); result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	config.DB.Save(&existingUser)

	c.JSON(200, gin.H{"error": nil, "message": existingUser})
}

func HandleDeleteUser(c *gin.Context) {
	if err := config.Connect(); err != nil {
		c.JSON(500, gin.H{"message": "Connection Error", "error": err})
		return
	}

	defer config.Disconnect()

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request Body", "error": err})
		return
	}

	if result := config.DB.First(&model.User{}, user.ID); result.Error != nil {
		c.JSON(404, gin.H{"messasge": "There is no such User", "error": result.Error})
		return
	}

	config.DB.Delete(&model.User{}, user.ID)
	c.JSON(200, gin.H{"error": nil, "message": "User Deleted"})
}
