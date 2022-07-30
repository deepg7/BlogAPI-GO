package controllers

import (
	"BlogAPI-GO/config"
	"BlogAPI-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Validate input
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	// Create book
	user := models.User{Name: input.Name, Email: input.Email, Number:input.Number}
	config.DB.Create(&user)
  
	c.JSON(http.StatusOK, gin.H{"data": user})
  }

  func FindUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
  
	c.JSON(http.StatusOK, gin.H{"data": users})
  }